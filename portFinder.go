package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

const maxPort int = 65535
const minPort int = 1

func main() {
	var ipToScan string

	flag.StringVar(&ipToScan, "value of ip", "", "IP INPUT")

	if ipToScan == "" {
		ipToScan = "127.0.0.1"
	}

	var activeThreads int = 0
	doneChannel := make(chan bool)
	defer close(doneChannel)

	for port := minPort; port <= maxPort; port++ {
		go testTCPConn(ipToScan, port, doneChannel)
		activeThreads++
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func testTCPConn(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second*10)
	if err == nil {
		fmt.Printf("Host %s has open port: %d\n", ip, port)
	}

	doneChannel <- true
}
