package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			go worker(channel)
		}
	}()

	for i := 0; i <= 1000; i++ {
		channel <- i
	}
}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i, " done")
		time.Sleep(time.Second * 5)
	}
}
