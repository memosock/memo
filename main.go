package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"os"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Started TCP Server...")
	defer l.Close()
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go HandleConnection(c)
	}
}