package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// pake atomic kalo tipe data yang diubah tipe data primitif
// kalau struct, dll baru pake mutexx
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter :", x)
}
