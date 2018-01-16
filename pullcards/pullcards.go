// Pull card data from JSON into native DB
// Single shot only - no need to run unless a new set is in

package main

import (
	"fmt"
	"log"
	//"github.com/MagicTheGathering/mtg-sdk-go"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Cleanly connect to a mySQL database
func mySQLConnect(userDb string, pwDb string, connDb string, schemaDb string) (*sql.DB, error) {
	db, err := sql.Open("mysql", userDb+":"+pwDb+"@tcp"+"("+connDb+")/"+schemaDb)
	fmt.Println("Connecting to db...")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	return db, err
}

func main() {

	userDb := "admin"
	pwDb := "True-faith1"
	connDb := "localhost:3306"
	schemaDb := "mtgcards_test"

	db, err := mySQLConnect(userDb, pwDb, connDb, schemaDb)
	r, err := db.Query("SELECT name FROM mtgcards_test.cards_all;")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for r.Next() {
		var name string
		if err := r.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}
}
