package concstore

import (
	"sync"
)

// UMap is a concurrent key-value in-memory storage. Keys are uin64 and returned when an object is added.
type UMap struct {
	c     uint64
	data  sync.Map
	mutex sync.Mutex
}

// NewUMap() returns a new *UMap.
func NewUMap() *UMap {
	return &UMap{}
}

// Add obj to the storage, return the object's key as uint64.
func (s *UMap) Add(datum interface{}) uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.c++
	s.data.Store(s.c, datum)
	return s.c
}

// Get the object by key.
func (s *UMap) Get(key uint64) (interface{}, bool) {
	return s.data.Load(key)
}

func (s *UMap) Set(key uint64, value interface{}) {
	s.data.Store(key, value)
}

// Remove the object with given key.
func (s *UMap) Delete(key uint64) {
	s.data.Delete(key)
}

// Free the memory for the storage.
func (s *UMap) Free() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data.Range(func(k, v interface{}) bool {
		s.data.Delete(k)
		return true
	})
}

// Range over all values until the function returns false.
func (s *UMap) Range(f func(k uint64, v interface{}) bool) {
	s.data.Range(func(k, v interface{}) bool {
		return f(k.(uint64), v)
	})
}

// IMap is a concurrent int - interface{} key - value store.
type IMap struct {
	c     int
	data  sync.Map
	mutex sync.Mutex
}

// NewIMap returns a new *IMap.
func NewIMap() *IMap {
	return &IMap{}
}

// Add obj to the storage, return the object's key as int.
func (s *IMap) Add(datum interface{}) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.c++
	s.data.Store(s.c, datum)
	return s.c
}

// Get the object by key.
func (s *IMap) Get(key int) (interface{}, bool) {
	return s.data.Load(key)
}

func (s *IMap) Set(key int, value interface{}) {
	s.data.Store(key, value)
}

// Remove the object with given key.
func (s *IMap) Delete(key int) {
	s.data.Delete(key)
}

// Free the memory for the storage.
func (s *IMap) Free() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data.Range(func(k, v interface{}) bool {
		s.data.Delete(k)
		return true
	})
}

// Range over all values until the function returns false.
func (s *IMap) Range(f func(k int, v interface{}) bool) {
	s.data.Range(func(k, v interface{}) bool {
		return f(k.(int), v)
	})
}
