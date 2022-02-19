package queues

import (
	"runtime"
	"sync"
)

// Q is a queue.
// Uses Go Channel and sync.WaitGroup under the hood.
type Q struct {
	ch chan struct{}
	wg sync.WaitGroup
}

// New returns a new queue with the given max size.
// No more than max size items can be added to the queue.
// It uses Golang's channel under the hood.
// Think this function as make(chan int, maxSize).
// So, if you set the max size to 0, it will not work as a queue or buffer.
func New(max uint) *Q {
	return &Q{
		ch: make(chan struct{}, max),
		wg: sync.WaitGroup{},
	}
}

// Add adds an item to the queue.
// It uses wg.Add(n) to keep track of the number of items in the queue.
func (q *Q) Add(delta int) {
	for range make([]struct{}, delta) {
		q.wg.Add(1)
		q.ch <- struct{}{}
	}
}

// Done removes an item from the queue.
// It uses wg.Done() to keep track of the number of items in the queue.
func (q *Q) Done() {
	defer q.wg.Done()
	<-q.ch
}

// Wait for the queue to be empty.
// It uses wg.Wait() to wait for the queue to be empty.
func (q *Q) Wait() {
	q.wg.Wait()
}

// Exit calls q.Done() first then calls runtime.Goexit().
// If called inside a goroutine, the goroutine will exit immediately.
// See https://golang.org/pkg/runtime/#Goexit for the documentation.
func (q *Q) Exit() {
	q.Done()
	runtime.Goexit()
}
