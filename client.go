package main

import "net"

// Responsible for keepign user information
// name , current tcp connection , current room

type client struct {
	conn net.Addr			   // current tcp connection
	nick string 			   // name
	room *room                 // point to a room
	commands chan<- command    // command send by channel 
}