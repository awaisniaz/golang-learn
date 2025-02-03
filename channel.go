package main

import (
	"fmt"
	"time"
)

func Channel() {
	var a chan int
	fmt.Println(a)
	val := <-a
	fmt.Println(val)
}
func main() {
	ch := make(chan int)
	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving from channel")
	go receiver(ch)

	time.Sleep(time.Second * 1)

}

func send(ch chan int) {
	ch <- 1
}

func receiver(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
