package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Channel() {
	var a chan int
	fmt.Println(a)
	val := <-a
	fmt.Println(val)
}
func main() {
	// ch := make(chan int)
	// fmt.Println("Sending value to channel")
	// go send(ch)

	// fmt.Println("Receiving from channel")
	// go receiver(ch)

	// time.Sleep(time.Second * 1)
	fmt.Print("I am Awaus Buaz")
	go parallism()
	fmt.Printf("I am End position")

	// All Datatype of the System

	// Pointer in Golang

	a := new(int)
	*a = 10
	fmt.Print(&a)

	// Enum in Golang

	type Size int

	const (
		small Size = iota
		medium
		large
		extraLarge
	)
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)

}

func goRoutineCheck() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sleep(&wg, time.Second*1)
	go sleep(&wg, time.Second*2)
	wg.Wait()
	fmt.Println("All Goroutines finished")
}

func sleep(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	time.Sleep(duration)
	fmt.Println("Finished Execution")
}

func checkTheString() {
	x := "123"
	_, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("Supplied value is not a valid integer")
	} else {
		fmt.Println()
	}
}
func parallism() {
	fmt.Printf("I am Goroutine")
}

func send(ch chan int) {
	ch <- 1
}
func receiver(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
