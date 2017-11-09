package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

//
// Task states
//
type TaskState byte

const (
	IDEA TaskState = iota
	TODO
	WORK
	BLOCK
	DONE
)

// Task structure
// TODO - Be more elegant with the times here?
type Task struct {
	Text    string
	State   TaskState
	AddedAt time.Time
}

func AddTask(text string, state TaskState) {
	err := DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		id, _ := b.NextSequence()

		b.Put(id, Task{text, state, time.Now()})

		return nil
	})
}
