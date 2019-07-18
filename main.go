package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6e8d7fef09c9485daac48699bea64f66")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(4753195))
	if err != nil {
		log.Fatal(err)
	}
	pubkey, err := crypto.Ecrecover(clique.SealHash(block.Header()).Bytes(), block.Extra()[len(block.Extra())-65:])
	if err != nil {
		log.Println(err)
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
	log.Println(signer.Hex())
	// fmt.Printf("blk num: %v", latest.Number())
	// fmt.Println()
	// bingo := big.NewInt(7158774)
	// for blockNumber = latest.Number().Int64(); blockNumber >= 7159043; blockNumber-- {
	// 	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if new(big.Int).Mod(block.Hash().Big(), big.NewInt(1000)).Cmp(big.NewInt(6)) == 0 {
	// 		fmt.Printf("BINGO..... %v,%v %v %v", block.Time(), block.NumberU64(), block.Hash().Hex(), block.Hash().Big())
	// 		fmt.Println()
	// 		fmt.Printf("%v blocks ago", latest.NumberU64()-block.NumberU64())
	// 		os.Exit(0)
	// 	}
	// }
	// fmt.Printf("BINGO..... %v", bingo)
	// fmt.Println()
	// fmt.Printf("%v blocks ago", latest.NumberU64()-bingo.Uint64())

	// for _, tx := range latest.Transactions() {
	// 	fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	// 	fmt.Println(tx.Value().String())    // 10000000000000000
	// 	fmt.Println(tx.Gas())               // 105000
	// 	fmt.Println(tx.GasPrice().Uint64()) // 102000000000
	// 	fmt.Println(tx.Nonce())             // 110644
	// 	fmt.Println(tx.Data())              // []
	// 	fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

	// 	signer := types.NewEIP155Signer(tx.ChainId())
	// 	sender, err := signer.Sender(tx)
	// 	if err == nil {
	// 		fmt.Printf("sender: %v", sender.Hex())
	// 		fmt.Println()
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// 	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(receipt.Status) // 1
	// }

	// blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	// count, err := client.TransactionCount(context.Background(), blockHash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for idx := uint(0); idx < count; idx++ {
	// 	tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	// }

	// txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	// tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	// fmt.Println(isPending)       // false
}
