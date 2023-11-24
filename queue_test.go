package queues_test

import (
	"sync/atomic"
	"testing"

	"github.com/hsblhsn/queues"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	var (
		queue   = queues.New(10)
		counter int32
	)

	for range make([]struct{}, 10) {
		queue.Add(1)

		go func() {
			defer queue.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	queue.Wait()
	require.EqualValues(t, 10, atomic.LoadInt32(&counter))
}
