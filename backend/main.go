package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alchemist007/chat-app/pkg/websocket"
)

var id int = 0

func serveWebsocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("main:Websocket endpoint hit by new client")

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	// go websocket.Writer(ws)
	// websocket.Reader(ws)
	id = id + 1
	client := &websocket.Client{
		ID:   strconv.Itoa(id),
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client

	fmt.Println("main: New Client ID = ", client.ID)
	client.Read()
}

func setupRoutes() {

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(pool, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8081", nil)
}
