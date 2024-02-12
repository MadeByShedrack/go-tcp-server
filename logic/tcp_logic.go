package logic

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "TCP"
)

func TcpServer() {
	listening, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listening.Close()
	for {
		conn, err := listening.Accept()
		if err != nil {
			log.Fatal()
			os.Exit(1)
		}
		go handleTCPServer(conn)
	}
}

func handleTCPServer(myConn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := myConn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	myTime := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	myConn.Write([]byte("Hello come back in"))
	myConn.Write([]byte(myTime))

	myConn.Close()
}
