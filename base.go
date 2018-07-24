package goworker

import "fmt"

func Run() {
	client := NewClient()
	fmt.Println(client.Ping().String())
}
