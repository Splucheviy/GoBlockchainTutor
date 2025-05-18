package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 ETH to John")
	bc.AddBlock("Send 2 more EHT to John")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.HashPrevBlock)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
