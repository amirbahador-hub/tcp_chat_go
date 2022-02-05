package main

import (
	"net"
	"log"
)

func main() {
	s := newServer()   // initial server

	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server : %s", err.Error())
	}

	defer listener.Close()
	log.Printf("start server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}

}