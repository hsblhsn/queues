# Queues

[![CI](https://github.com/hsblhsn/queues/actions/workflows/ci.yml/badge.svg)](https://github.com/hsblhsn/queues/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/hsblhsn/queues)](https://goreportcard.com/report/github.com/hsblhsn/queues) [![Go Reference](https://pkg.go.dev/badge/github.com/hsblhsn/queues.svg)](https://pkg.go.dev/github.com/hsblhsn/queues) 

A simple `sync.WaitGroup` like queue and goroutine execution controller.

## Features

- Limit maximum goroutine execution/count.
- Wait for all the goroutines to finish.
- Run goroutines like a queue batch.

## Import

```go
import "github.com/hsblhsn/queues"
```

## Example

### [🔗 Go Playground](https://go.dev/play/p/WGS9b6I7KFd): Limit maximum goroutine counts to 3.

```go
package main

import (
	"fmt"
	"time"

	"github.com/hsblhsn/queues"
)

func main() {
	q := queues.New(3)
	for i := 0; i < 30; i++ {
		q.Add(1)
		go func(n int) {
			defer q.Done()
			time.Sleep(time.Second)
			fmt.Println(n)
		}(i)
	}
	q.Wait()
}
```

### [🔗 Go Playground](https://go.dev/play/p/ZZg5zCvVqaB): Batched queue for async jobs.

```go
package main

import (
	"fmt"
	"time"

	"github.com/hsblhsn/queues"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://twitter.com",
		"https://facebook.com",
		"https://youtube.com",
	}
	q := queues.New(2)
	for _, v := range urls {
		q.Add(1)
		go crawl(q, v)
	}
	q.Wait()
}

func crawl(q *queues.Q, url string) {
	defer q.Done()
	fmt.Println("Crawling: ", url)
	time.Sleep(time.Second)
}
```
