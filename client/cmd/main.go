package main

import (
	"flag"

	"github.com/itocode21/broadcast-system/client/internal/client"
)

func main() {
	serverAddr := flag.String("addr", "localhost:8080", "Server address to connect to")
	flag.Parse()

	client.Connect(*serverAddr)
}
