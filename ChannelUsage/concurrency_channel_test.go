package channelusage

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	ctx := context.Background()
	ch := make(chan []int)

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		ChannelProducer(ctx, ch)
	}()

	select {
	case list := <-ch:
		if len(list) != 10 {
			t.Errorf("Expected 10 items, got %d", len(list))
		}
	}

	wg.Wait()
	// 输出执行时间
	t.Errorf("Execution time with channel:%+v", time.Since(start))
}

func TestConcurrencySharingList(t *testing.T) {
	start := time.Now()
	list := []int{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			list = append(list, i)
		}
	}()

	wg.Wait()
	// 输出执行时间
	t.Errorf("Execution time with sharing variable:%+v", time.Since(start))
}

func TestMultiConcurrency(t *testing.T) {
	ctx := context.Background()
	ch := make(chan []int, 3)

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ChannelProducer(ctx, ch)
		}()
	}

	t.Logf("Execution time with channel for wating: %+v", time.Since(start))
	wg.Wait()
	t.Logf("Execution time with channel for merging: %+v", time.Since(start))
	for i := 0; i < 3; i++ {
		select {
		case list := <-ch:
			t.Logf("Received list from channel: %+v", list)
		}
	}
	t.Logf("Execution time with channel for merging: %+v", time.Since(start))

	// 输出执行时间
	t.Errorf("Execution time with channel:%+v", time.Since(start))
}

func TestMultiConcurrencySharingList(t *testing.T) {
	start := time.Now()
	list := []int{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				list = append(list, i*10+j)
			}
		}
	}()

	wg.Wait()
	// 输出执行时间
	t.Errorf("Execution time with sharing variable:%+v", time.Since(start))
}

func TestMultiConcurrencySharingListMultiSlice(t *testing.T) {
	start := time.Now()
	listA := []int{}
	listB := []int{}
	listC := []int{}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			listA = append(listA, j)
		}
	}()
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			listB = append(listB, j)
		}
	}()
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			listC = append(listC, j)
		}
	}()

	wg.Wait()
	for i := 0; i < 3; i++ {
		listA = append(listA, i)
		listB = append(listB, i)
		listC = append(listC, i)
	}
	// 输出执行时间
	t.Errorf("Execution time with sharing variable:%+v", time.Since(start))
}
