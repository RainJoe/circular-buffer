package main

import (
	"sync"
	"testing"
)

func TestNewAndCapacity(t *testing.T) {
	buf := New(10)
	if buf.Capacity() != 10 {
		t.Error("expected: 10, got: ", buf.Capacity())
	}
}

func TestCircularBuffer_PutAndGet(t *testing.T) {
	var tests = []struct {
		Puts     []int
		Expected []int
	} {
		{[]int{1,2,3,4,5}, []int{1,2,3,4,5}},
	}
	buf := New(10)
	for _, tt := range tests {
		for _, p := range tt.Puts {
			buf.Put(p)
		}
		for _, e := range tt.Expected {
			val := buf.Get().(int)
			if  val != e {
				t.Errorf("expected: %d, got: %d", e, val)
			}
		}
	}
}

func TestCircularBuffer_Reset(t *testing.T) {
	buf := New(10)
	buf.Put(1)
	buf.Put(2)
	buf.Reset()
	if !buf.Empty() {
		t.Errorf("expected: true, got: %v", buf.Empty())
	}
}

func TestCircularBuffer_Size(t *testing.T) {
	buf := New(10)
	buf.Put(1)
	buf.Put(2)
	size := buf.Size()
	if size != 2 {
		t.Errorf("expected: 2, got: %d", size)
	}
}

func TestCircularBuffer_Full(t *testing.T) {
	buf := New(3)
	buf.Put(1)
	buf.Put(2)
	buf.Put(3)
	if !buf.Full() {
		t.Errorf("expected: true, got: %v", buf.Empty())
	}
}

func TestCircularBuffer_parallel(t *testing.T) {
	buf := New(10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf.Put(1)
			buf.Size()
			buf.Capacity()
			buf.Full()
			buf.Empty()
		}()
	}
	wg.Wait()
}