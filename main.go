package main

import "fmt"

type Message struct {
	Output string
}

func ProcessMessageOne(message []string) <-chan Message {
	out := make(chan Message)
	go func() {
		for _, msg := range message {
			o := Message{Output: fmt.Sprintf("process message one %s", msg)}
			out <- o
		}
		close(out)
	}()
	return out
}

func ProcessMessageTwo(data <-chan Message) <-chan Message {
	out := make(chan Message)
	go func() {
		for data := range data {
			o := Message{Output: fmt.Sprintf("process message two %s", data)}
			out <- o
		}
		close(out)
	}()
	return out
}

func main() {
	rawMessage := []string{"apple", "orange", "pear", "banana"}
	firstProcess := ProcessMessageOne(rawMessage)
	secondProcess := ProcessMessageTwo(firstProcess)

	for result := range secondProcess {
		fmt.Println(result)
	}
}
