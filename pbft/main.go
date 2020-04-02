package main

import (
	"fmt"
	"os"

	"github.com/bigsakh001/consensusPBFT/pbft/network"
)

func main() {
	fmt.Println("!************preparing to start********************")
	nodeID := os.Args[0]
	server := network.NewServer(nodeID)

	server.Start()
}
