package main

import (
	"context"
	"fmt"
	"sync"
)

type Service struct {
}

func (s *Service) DoFizzBuzz(ctx context.Context, from, to int) []string {
	pool := newPool(ctx, 10)
	result := make([]string, to-from+1)

	for i := from; i <= to; i++ {
		num := i
		pool.addJob(func() {
			fizzBuzzResult := s.singleFizzBuzz(num)
			result[num-from] = fizzBuzzResult
		})

	}

	pool.close()
	pool.wait()

	return result
}

func (s *Service) singleFizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func newPool(ctx context.Context, numWorker int) *Pool {
	pool := Pool{
		workQueue: make(chan Job),
	}

	pool.wg.Add(numWorker)
	for j := 0; j < numWorker; j++ {
		go func() {
			defer pool.wg.Done()
			select {
			case <-ctx.Done():
				pool.close()
				break
			default:
				for job := range pool.workQueue {
					job()
				}
			}

		}()
	}

	return &pool
}

func (p *Pool) addJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) close() {
	close(p.workQueue)
}

func (p *Pool) wait() {
	p.wg.Wait()
}
