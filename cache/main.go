package main

import (
	"fmt"
	"time"
)

// Function to calculate fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exist := m.cache[key]
	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

// Function to be used in the cache
func GetFibonacci(key int) (interface{}, error) {
	return Fibonacci(key), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38, 50, 13, 50}

	for _, v := range fibo {
		start := time.Now()

		value, err := cache.Get(v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %d t: %s \n", v, value, time.Since(start))
	}
}
