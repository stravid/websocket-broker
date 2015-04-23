package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/soveran/redisurl"
    "net/http"
)

func main() {
    redisURL := flag.String("redis-url", "redis://127.0.0.1:6379/0", "connect to this Redis server")
    channel := flag.String("channel", "", "subscribe to this channel")

    flag.Parse()

    fmt.Println("DEBUG  redisURL:", *redisURL)
    fmt.Println("DEBUG  channel:", *channel)

    // Setup collector and Redis connection so we are receiving messages
    connection, error := redisurl.ConnectToURL(*redisURL)
    if error != nil {
        fmt.Println(error)
        os.Exit(1)
    }
    defer connection.Close()

    receivedMesssages := make(chan string)
    go Collector(connection, receivedMesssages, *channel)

    broker := &Broker{
        register: make(chan *WebSocketConnection),
        unregister: make(chan *WebSocketConnection),
        webSocketConnections: make(map[*WebSocketConnection]bool),
        receivedMessages: receivedMesssages,
    }

    go broker.run()

    http.Handle("/", WebSocketHandler{broker: broker})

    if err := http.ListenAndServe("localhost:3000", nil); err != nil {
        fmt.Println("ListenAndServe:", err)
    }
}
