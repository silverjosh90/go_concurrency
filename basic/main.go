package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// One Go routine can run at at time
	sleepDur, _ := time.ParseDuration("10ms")

	// How many simultaneous cores can run
	runtime.GOMAXPROCS(1)

	go /* Go keyword is used to declare a go routine */ func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(sleepDur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(sleepDur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
