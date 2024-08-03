package main

import (
	"log"
)

func main() {
	server := NewServer(8080, "127.0.0.1",)

	fd, err := server.start()
	if err != nil {
		log.Fatal("error (start server): ", err)
	}

	handler := NewHandler()
	handler.handle(fd)
}