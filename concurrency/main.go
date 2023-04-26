package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {

	fmt.Println("Concurrency")

	//we can see that after adding go keyword the func executes much faster
	go compute(5)
	go compute(5)

	fmt.Scanln()
}
