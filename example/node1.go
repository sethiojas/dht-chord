package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	chord "github.com/sethiojas/dht-chord"
)

func main() {
	node, err := chord.CreateNewNode("localhost:12334", "localhost:12336")
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan struct{})
	go func() {
		count := 20
		for {
			ticker := time.NewTicker(5 * time.Second)
			select {
			case <-ticker.C:
				count++
				key := strconv.Itoa(count)
				value := "hello-" + strconv.Itoa(count)
				node.Save(key, value)
				if count >= 30 {
					close(ch)
				}
			case <-ch:
				ticker.Stop()
				return
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	_, ok := <-ch
	if ok {
		close(ch)
	}
	node.Stop()
}