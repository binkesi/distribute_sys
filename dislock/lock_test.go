package dislock

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestSyncLock(t *testing.T) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var count = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count += 1
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestTryLock(t *testing.T) {
	var count int
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				log.Fatal("lock failed")
				return
			}
			count++
			fmt.Println("current count:", count)
			l.Unlock()
		}()
		wg.Wait()
	}
}
