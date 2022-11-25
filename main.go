package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	workerNumber := 5
	wg.Add(workerNumber)
	for i := 0; i < workerNumber; i++ {
		go work(i, wg)
	}
	wg.Wait()
	fmt.Println("all workers finished working")
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d is working...\n", i)
}
