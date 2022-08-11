package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	chatMessages    = make(chan string)
)
var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Client1 -> Server -> HandleConnection(Client1)
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	clientMessages := make(chan string)
	go MessageWrite(conn, clientMessages)
	// Client1:2560 Platzi.com, 38
	// platzi.com:38
	clientName := conn.RemoteAddr().String()

	clientMessages <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	chatMessages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incomingClients <- clientMessages

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		chatMessages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- clientMessages
	chatMessages <- fmt.Sprintf("%s said goodbye!", clientName)
}
func MessageWrite(conn net.Conn, clientMessages <-chan string) {
	for messsage := range clientMessages {
		fmt.Fprintln(conn, messsage)
	}
}
