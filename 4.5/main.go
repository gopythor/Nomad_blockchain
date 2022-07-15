package main

import (
	"fmt"
	"nomadcoin/4.5/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Thrid Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("PrevHash : %s\n", block.PrevHash)
	}

}
