package main

import (
	"log"

	"github.com/aergoio/aergo-lib/db"
)

func main() {
	aergoDB, err := db.NewBadgerDB("data")
	if err != nil {
		log.Fatal(err)
	}
	defer aergoDB.Close()
}
