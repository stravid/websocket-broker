package main

import (
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(request *http.Request) bool { return true },
}

type WebSocketHandler struct {
    broker *Broker
}

func (webSocketHandler WebSocketHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    webSocket, err := upgrader.Upgrade(writer, request, nil)

    if err != nil {
        return
    }

    webSocketConnection := &WebSocketConnection{webSocket: webSocket}
    webSocketHandler.broker.register <- webSocketConnection

    defer func() { webSocketHandler.broker.unregister <- webSocketConnection }()

    for {
        _, _, err := webSocket.ReadMessage()

        if err != nil {
            return
        }
    }
}
