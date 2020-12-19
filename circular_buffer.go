package main

import "sync"

// Circular Buffer ensures that we are always consuming the most recent data
type CircularBuffer struct {
	buffer   []interface{}
	head     int
	tail     int
	capacity int
	full     bool
	mu       sync.RWMutex
}

// New circular Buffer with fixed size
func New(capacity int) *CircularBuffer {
	return &CircularBuffer{
		buffer:   make([]interface{}, capacity),
		capacity: capacity,
	}
}

// Reset circular Buffer
func (c *CircularBuffer) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.head = c.tail
	c.full = false
}

// Empty returns true if the buffer is empty
func (c *CircularBuffer) Empty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return !c.full && (c.head == c.tail)
}

// Full returns true if the buffer is full
func (c *CircularBuffer) Full() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.full
}

// Capacity returns the capacity of buffer
func (c *CircularBuffer) Capacity() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.capacity
}

// Size returns the size of buffer
func (c *CircularBuffer) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	size := c.capacity
	if !c.full {
		if c.head >= c.tail {
			size = c.head - c.tail
		} else {
			size = c.capacity + c.head - c.tail
		}
	}
	return size
}


// Put item to buffer
func (c *CircularBuffer) Put(item interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.buffer[c.head] = item
	if c.full {
		c.tail = (c.tail + 1) % c.capacity
	}
	c.head = (c.head + 1) % c.capacity
	c.full = c.head == c.tail
}

// Get item from buffer
func (c *CircularBuffer) Get() interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.full && (c.head == c.tail) {
		return nil
	}
	item := c.buffer[c.tail]
	c.full = false
	c.tail = (c.tail + 1) % c.capacity
	return item
}
