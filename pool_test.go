package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Bobi")
	pool.Put("Doon")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			// setelah get jangan lupa di put untuk mengembalikan data ke pool
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Complete")
}
