package main

import (
	"fmt"
	"log"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

func main() {
	pgNum, pgSize := 2, 25
	log.Println("Fetch Page", pgNum, "with a page size of", pgSize)

	cards, totalCards, err := mtg.NewQuery().Where(mtg.CardSetName, "Rivals of Ixalan|Ixalan").Where(mtg.CardColors, "white").PageS(pgNum, pgSize)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "white RIX cards")
	for _, card := range cards {
		fmt.Println(card)
	}
}
