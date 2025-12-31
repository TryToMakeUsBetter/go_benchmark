package channelusage

import "context"

// use channel to send a list of integers from a goroutine
func ChannelProducer(ctx context.Context, ch chan<- []int) {
	list := []int{}
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
	ch <- list
}
