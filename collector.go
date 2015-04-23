package main

import "github.com/garyburd/redigo/redis"
import "fmt"

func Collector(connection redis.Conn, receivedMesssages chan string, channel string) {
    pubSubConnection := redis.PubSubConn{Conn: connection}
    pubSubConnection.Subscribe(channel)

    for {
        switch message := pubSubConnection.Receive().(type) {
        case redis.Message:
            fmt.Println("Collector received message from Redis.")
            receivedMesssages <- string(message.Data)
        }
    }
}
