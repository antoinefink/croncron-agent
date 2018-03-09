package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Waiting a bit makes sure we don't fail if there's no input at the beginning
	time.Sleep(time.Millisecond * 300)

	for i := 0; i < 5; i++ {
		fmt.Println(time.Now())
		time.Sleep(time.Millisecond * 100)
	}

	log.Println("Finished")
}
