package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Variables
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// Error handling
	myValue, err := strconv.ParseInt("a", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println((myValue))
	}

	// Maps
	m := make(map[string]int)
	m["six"] = 6

	fmt.Println(m["six"])

	// Slices
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index, value)
	}

	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}

	// Goroutines & channels
	c := make(chan int)
	go doSomething(c)
	<-c

	// Pointers
	g := 25
	h := &g

	fmt.Println(g, *h)
}

func doSomething(c chan int) {
	time.Sleep((3 * time.Second))
	fmt.Println("Done")
	c <- 1
}
