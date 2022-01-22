package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		// time.Sleep(2000 * time.Millisecond)
		_ = <-ch
	}()

	time.Sleep(2000 * time.Millisecond)

	//ch <- 100
	//fmt.Println("sent")

	select {
	case ch <- 100:
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
