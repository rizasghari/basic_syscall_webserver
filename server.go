package main

import (
	"log"
	"net"
	"syscall"
)

type Server struct {
	Host string
	Port int
}

func NewServer(port int, host string) *Server {
	return &Server{
		Port: port,
		Host: host,
	}
}

func (s Server) start() (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("Error starting socket server:", err) 
	}
	srv := &syscall.SockaddrInet4{Port: s.Port}
	addrs, err := net.LookupHost(s.Host)
	if err != nil {
		log.Fatal("Error lookup host:", err)
	}
	for _, addr := range addrs {
		ip := net.ParseIP(addr).To4()
		copy(srv.Addr[:], ip)
		if err = syscall.Bind(fd, srv); err != nil {
			log.Fatal("error (bind): ", err)
		}
	}
	if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatal("error (listening): ", err)
	} else {
		log.Println("Listening on ", s.Host, ":", s.Port)
	}
	if err != nil {
		log.Fatal("error (port listening): ", err)
	}

	return fd, nil
}