package main

import (
	"fmt"
	"math/big"
	"sync"
)

// calculateFactorial menghitung faktorial dari sebuah angka
func calculateFactorial(n int64, results chan<- *big.Int) {
	result := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	results <- result
}

func main() {

	// Pengaturan untuk menghitung faktorial
	numbers := []int64{20, 25, 30, 35}
	results := make(chan *big.Int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			calculateFactorial(n, results)
		}(num)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()
	close(results)

	// Cetak hasil
	for result := range results {
		fmt.Println("Faktorial: ", result)
	}

	fmt.Println("Main application finished")
}
