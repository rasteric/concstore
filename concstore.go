package concstore

import (
	"sync"
)

// Storage is a concurrent key-value in-memory storage. Keys are uin64 and returned when an object is added.
// You create a new storage by using &Storage{}.
type Storage struct {
	c     uint64
	data  sync.Map
	mutex sync.Mutex
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

func (s *Storage) Set(key uint64, value interface{}) {
	s.data.Store(key, value)
}

// Remove the object with given key.
func (s *Storage) Delete(key uint64) {
	s.data.Delete(key)
}

// Free the memory for the storage.
func (s *Storage) Free() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data.Range(func(k, v interface{}) bool {
		s.data.Delete(k)
		return true
	})
}

// Range over all values until the function returns false.
func (s *Storage) Range(f func(k uint64, v interface{}) bool) {
	s.data.Range(func(k, v interface{}) bool {
		return f(k.(uint64), v)
	})
}

// IntStorage is an int - interface{} key - value store.
type IStorage struct {
	c     int
	data  sync.Map
	mutex sync.Mutex
}

// Add obj to the storage, return the object's key as int.
func (s *IStorage) Add(datum interface{}) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.c++
	s.data.Store(s.c, datum)
	return s.c
}

// Get the object by key.
func (s *IStorage) Get(key int) (interface{}, bool) {
	return s.data.Load(key)
}

func (s *IStorage) Set(key int, value interface{}) {
	s.data.Store(key, value)
}

// Remove the object with given key.
func (s *IStorage) Delete(key int) {
	s.data.Delete(key)
}

// Free the memory for the storage.
func (s *IStorage) Free() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data.Range(func(k, v interface{}) bool {
		s.data.Delete(k)
		return true
	})
}

// Range over all values until the function returns false.
func (s *IStorage) Range(f func(k int, v interface{}) bool) {
	s.data.Range(func(k, v interface{}) bool {
		return f(k.(int), v)
	})
}
