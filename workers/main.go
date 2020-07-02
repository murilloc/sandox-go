package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := make(chan int)
	go func() {
		for i := 1; i <= 1000; i++ {
			go worker(channel, i)

		}
	}()

	for i := 0; i < 4000; i++ {
		channel <- i
	}
}

func worker(channel chan int, worker int) {
	for i := range channel {
		fmt.Println("O número ", i, " foi processado worker:", worker)
		time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	}
}
