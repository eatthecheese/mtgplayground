package main

import (
	"fmt"
	"log"

	mtg "github.com/MagicTheGathering/mtg-sdk-go"
)

func main() {
	pgNum, pgSize := 2, 25
	log.Println("Fetch Page", pgNum, "with a page size of", pgSize)

	cards, totalCards, err := mtg.NewQuery().Where(mtg.CardSetName, "Rivals of Ixalan").PageS(pgNum, pgSize)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "RIX cards")
	for _, card := range cards {
		fmt.Println(card.Name, card.ManaCost, card.Colors)
	}
}
