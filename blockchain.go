package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const blocksBucket =  "blocksBucket"
const dbFile = "dbFile.db"


type Blockchain struct {
	tip []byte
	db *bolt.DB
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}




func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				panic(err)
			}

			if err != nil {
				panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())

			if err != nil {
				panic(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				panic(err)
			}
			tip = genesis.Hash


		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err !=nil {
		panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}



func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}