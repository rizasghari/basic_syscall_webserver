package handler

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Handle(fd int) {
	message := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"Content-Length: 25\r\n" +
		"\r\n" +
		"Server with syscall"

	for {
		cSock, cAddr, err := syscall.Accept(fd)
		if err != nil {
			log.Fatal("error (accept): ", err)
		} else {
			log.Printf("Client Accepted - socket: %v, address: %v", cSock, cAddr)
		}

		go func(clientSocket int, clientAddress syscall.Sockaddr) {

			buf := make([]byte, 1024)
			n, err := syscall.Read(cSock, buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Read error: %s\n", err)
				return
			}
			fmt.Printf("Received request: %s\n", string(buf[:n]))

			err = syscall.Sendmsg(clientSocket, []byte(message), []byte{}, clientAddress, 0)
			if err != nil {
				log.Fatal("error (send): ", err)
			}
			syscall.Close(clientSocket)
		}(cSock, cAddr)
	}
}