package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// DB connection, global handle
var DB *bolt.DB

func main() {
	fmt.Println("Welcome to taskman!")

	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	AddTask("Finish this program", TODO)
}
