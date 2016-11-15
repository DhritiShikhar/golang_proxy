package main

import (
	"fmt"
	"net"
)

func main() {

	// create a server
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Webserver creation unsuccessful")
	}
}
