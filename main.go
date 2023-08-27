package main

import (
	"fmt"
	"net-cat/pkg"
	"os"
)

func main() {
	input := os.Args
	switch len(input) {
	case 1:
		pkg.StartServerDefPort()
	case 2:
		pkg.StartServerMyPort()
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
}
