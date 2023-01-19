package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Database is a struct that contains the database connection
type Product struct {
	Id        int
	Name      string
	Inventory int
	Price     int
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	fmt.Println("Id\tName\tInventory\tPrice")
}
