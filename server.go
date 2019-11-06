package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := "Hello World"

		_, err = c.Write([]byte(string(result)))
		if err != nil {
			CheckTcpError(err)
		}
	}
	_ = c.Close()
}