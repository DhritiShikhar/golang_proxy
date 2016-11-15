package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	// create a server
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Webserver creation unsuccessful")
	} else {
		for {
			connection, _ := server.Accept() // accept connection
			reader := bufio.NewReader(connection)
			req, _ := http.ReadRequest(reader)
			backend, _ := net.Dial("tcp", "127.0.0.1:8081")
			backendReader := bufio.NewReader(backend)
			req.Write(backend)
			resp, _ := http.ReadResponse(backendReader, req)
			resp.Close = true
			resp.Write(connection)
			log.Printf("proxied %s: got %d", req.URL.Path, resp.StatusCode)
			connection.Close()
		}
	}
}
