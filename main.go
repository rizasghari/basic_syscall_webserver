package main

import (
	"log"

	"github.com/rizasghari/syscall_webserver/handler"
	"github.com/rizasghari/syscall_webserver/server"
)

func main() {
	server := server.NewServer(8080, "127.0.0.1")

	fd, err := server.Start()
	if err != nil {
		log.Fatal("error (start server): ", err)
	}

	handler := handler.NewHandler()
	handler.Handle(fd)
}
