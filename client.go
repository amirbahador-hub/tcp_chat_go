package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Responsible for keepign user information
// name , current tcp connection , current room

type client struct {
	conn net.Conn			   // current tcp connection
	nick string 			   // name
	room *room                 // point to a room
	commands chan<- command    // command send by channel 
}


func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, "")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
			case "/nick":
			case "/join":
			case "/rooms":
			case "/msg":
			case "/quit":
			default:
				c.err(fmt.Errorf("unknown command: %s", cmd))
		
		}
	}
}


func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}