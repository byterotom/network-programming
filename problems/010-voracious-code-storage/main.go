package main

import "net"

const (
	PROTOCOL = "tcp"
	ADDRESS  = ":6969"
)

func main() {
	listener, err := net.Listen(PROTOCOL, ADDRESS)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

}
