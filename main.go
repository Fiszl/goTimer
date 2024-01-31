package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	Timer int
)

func init() {
	flag.IntVar(&Timer, "time", 0, "How many seconds to count.")
	flag.Parse()

	if Timer < 0 {
		fmt.Println("Please use a positive number.")
		os.Exit(1)
	}
}

func main() {
	if Timer == 0 {
		fmt.Println("No number specified, using 60 seconds.")
		Timer = 60
	}

	for i := Timer; i != -1; i-- {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	fmt.Println("Times Up!")
}
