package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/1makarov/gen-prv-keys/file"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var count int

	fmt.Printf("How many keys to create? ")
	fmt.Scanf("%v", &count)

	name := fmt.Sprintf("%v", time.Now().Unix())

	file, err := file.New(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i != count; i++ {
		prvKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		prvKeyBytes := crypto.FromECDSA(prvKey)
		prvKeyString := hexutil.Encode(prvKeyBytes)
		addr := crypto.PubkeyToAddress(prvKey.PublicKey)

		text := fmt.Sprintf("%v,%v", prvKeyString, addr)

		if err := file.Write(text); err != nil {
			log.Fatal(err)
		}
	}
}
