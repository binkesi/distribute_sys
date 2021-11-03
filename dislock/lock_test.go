package dislock

import (
	"fmt"
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
