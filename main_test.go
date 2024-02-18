package main

import (
	"testing"
)

func BenchmarkPipeline(b *testing.B) {
	rawMessage := []string{"apple", "orange", "pear", "banana"}
	for i := 0; i < b.N; i++ {
		firstProcess := ProcessMessageOne(rawMessage)
		_ = ProcessMessageTwo(firstProcess)
	}
}

func BenchmarkProcessOne(b *testing.B) {
	rawMessage := []string{"apple", "orange", "pear", "banana"}
	for i := 0; i < b.N; i++ {
		_ = ProcessMessageOne(rawMessage)
	}
}

func BenchmarkProcessTwo(b *testing.B) {
	rawMessage := []string{"apple", "orange", "pear", "banana"}
	firstProcess := ProcessMessageOne(rawMessage)
	for i := 0; i < b.N; i++ {
		_ = ProcessMessageTwo(firstProcess)
	}
}

func TestProcessOne(t *testing.T) {
	// Define test input
	input := []string{"apple", "orange", "pear", "banana"}

	// Call the function
	outputChannel := ProcessMessageOne(input)

	// Collect the output
	var output []string
	for msg := range outputChannel {
		output = append(output, msg.Output)
	}

	// Verify the output
	expected := []string{"process message one apple", "process message one orange", "process message one pear", "process message one banana"}
	if len(output) != len(expected) {
		t.Errorf("Expected output length %d, but got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf("Expected %s at index %d, but got %s", expected[i], i, output[i])
		}
	}
}

func TestProcessTwo(t *testing.T) {
	// Define test input
	input := make(chan Message)
	go func() {
		input <- Message{Output: "test message one"}
		input <- Message{Output: "test message two"}
		close(input)
	}()

	// Call the function
	outputChannel := ProcessMessageTwo(input)

	// Collect the output
	var output []string
	for msg := range outputChannel {
		output = append(output, msg.Output)
	}

	// Verify the output
	expected := []string{"process message two {test message one}", "process message two {test message two}"}
	if len(output) != len(expected) {
		t.Errorf("Expected output length %d, but got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf("Expected %s at index %d, but got %s", expected[i], i, output[i])
		}
	}
}
