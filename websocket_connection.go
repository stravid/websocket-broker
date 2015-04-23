package main

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
    webSocket *websocket.Conn
}

func (webSocketConnection *WebSocketConnection) send(message string) {
    webSocketConnection.webSocket.WriteMessage(websocket.TextMessage, []byte(message))
}
