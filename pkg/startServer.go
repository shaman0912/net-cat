package pkg

import (
	"fmt"
	"log"
	"net"
	"os"
)

// create port and type by port
const (
	CONN_PORT = ":8989"
	CONN_TYPE = "tcp"
)

// start server by default port
func StartServerDefPort() {
	l, err := net.Listen(CONN_TYPE, CONN_PORT)
	if err != nil {
		log.Fatalln("Wrong func StartServerDefPort")
		return
	}
	defer l.Close()
	log.Printf("Listening on the port :%s", CONN_PORT)
	go BroadCaster(&mutex)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("unable to acept connection: %s", err.Error())
			conn.Close()
			continue
		}
		go HandleConnection(conn, &mutex)
	}
}

func StartServerMyPort() {
	arg := os.Args[1]
	port := fmt.Sprintf(":%s", arg)
	l, err := net.Listen(CONN_TYPE, port)
	if err != nil {
		log.Fatalf("unable to start server:%s", err.Error())
	}
	defer l.Close()
	fmt.Printf("Listening on the port %s/n", port)
	go BroadCaster(&mutex)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("unaable to acept connetction: %s", err.Error())
			conn.Close()
			continue

		}
		go HandleConnection(conn, &mutex)
	}
}
