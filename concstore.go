package concstore

import (
	"sync"
)

// Storage is a concurrent key-value in-memory storage. Keys are uin64 and returned when an object is added.
type Storage struct {
	c     uint64
	data  sync.Map
	mutex sync.Mutex
}

// New returns a new storage with initial size.
func New() *Storage {
	s := Storage{}
	return &s
}

// Add obj to the storage, return the object's key as uint64.
func (s *Storage) Add(datum interface{}) uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.c++
	s.data.Store(s.c, datum)
	return s.c
}

// Get the object by key.
func (s *Storage) Get(key uint64) (interface{}, bool) {
	return s.data.Load(key)
}

// Remove the object with given key.
func (s *Storage) Delete(key uint64) {
	s.data.Delete(key)
}

// Free the memory for the storage.
func (s *Storage) Free() {
	s.data.Range(func(k, v interface{}) bool {
		s.data.Delete(k)
		return true
	})
}
