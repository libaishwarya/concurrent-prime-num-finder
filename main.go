package main

import (
	"fmt"
	"math"
	"sync"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrt; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	const (
		jobs = 2
		max  = 100
	)

	job := make(chan int, max)
	result := make(chan string, max)

	var wg sync.WaitGroup
	wg.Add(jobs)
	for i := 1; i <= jobs; i++ {
		go workerPrime(i, job, result, &wg)
	}

	for i := 1; i <= max; i++ {
		job <- i
	}
	close(job)

	wg.Wait()
	close(result)
	for res := range result {
		fmt.Println(res)
	}

}

func workerPrime(workerID int, job chan int, result chan string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for n := range job {
		val := isPrime(n)
		if val {
			result <- fmt.Sprintf("goroutine %d found prime: %d", workerID, n)
		}
	}

}
