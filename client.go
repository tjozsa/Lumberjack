package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//open connection
	conn, err := net.Dial("tcp", ":50000")
	if err != nil {
		// No connection could be made becasue the target machine actively refused it.
		fmt.Println("Error dialing", err.Error())
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n"))
		if err != nil {
			fmt.Println("Error while sending.", err.Error())
		}
	}
}
