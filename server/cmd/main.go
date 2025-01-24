package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/itocode21/broadcast-system/server/internal/server"
)

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	srv := server.NewBroadcastServer()
	go srv.Run()

	http.HandleFunc("/ws", srv.HandleWebSocket)

	log.Println("Server started on port", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
