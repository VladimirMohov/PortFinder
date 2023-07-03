package PortFinder

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const maxPort int = 65535
const minPort int = 1

func PortFinder(ipToScan string) {

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
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		fmt.Printf("Host %s has open port: %d\n", ip, port)
	}

	doneChannel <- true
}
