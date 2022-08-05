package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := balance
	balance = b + amount
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg)
	}
	wg.Wait()
	fmt.Println(Balance())
}
