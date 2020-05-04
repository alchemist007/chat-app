package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start listens on each channel and updates the pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Pool:Size of connection pool: ", len(pool.Clients))
			fmt.Println("Pool:Current pool of clients=")
			for client := range pool.Clients {
				fmt.Println("Pool:", *client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Pool:Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected"})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Pool:Sending message to all clients")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
