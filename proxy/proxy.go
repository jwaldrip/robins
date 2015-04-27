package proxy

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
)

// Proxy constructs methods to proxy TCP traffic
type Proxy struct {
	hosts []string
	port  string
}

// New creates a new Proxy
func New(hosts []string, port string) *Proxy {
	if port[0] != ':' {
		port = ":" + port
	}
	return &Proxy{
		hosts: hosts,
		port:  port,
	}
}

// Listen listens for new incoming connections
func (p *Proxy) Listen() error {
	ln, err := net.Listen("tcp", p.port)
	if err != nil {
		return err
	}
	fmt.Println("listening on", p.port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go p.handle(conn)
	}
}

func (p *Proxy) handle(clientConn net.Conn) {
	var serverConn net.Conn
	defer clientConn.Close()

	shuffledHosts := make([]string, len(p.hosts))
	perm := rand.Perm(len(p.hosts))
	for i, v := range perm {
		shuffledHosts[v] = p.hosts[i]
	}

	var connErr error
	for _, host := range shuffledHosts {
		serverConn, connErr = net.Dial("tcp", host+p.port)
		if connErr == nil {
			break
		}
	}

	if connErr != nil {
		fmt.Fprintln(os.Stderr, connErr)
		return
	}

	defer serverConn.Close()

	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
}
