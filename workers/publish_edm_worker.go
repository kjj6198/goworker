package workers

import "fmt"

type publishEdmWorker struct {
	Run func(edmID string)
}

func newWorker() *publishEdmWorker {
	return &publishEdmWorker{
		Run: func(edmID string) {
			fmt.Println("work", edmID)
		},
	}
}

var PublishEdmWorker = newWorker()
