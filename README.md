# concstore
 Concurrent uint64, interface{} and int, interface{} maps for Go. `UMap` is a concurrent uint64 key - value in-memory map for Go. `IMap` is a concurrent int key - value in-memory map for Go.

These are currently just wrappers around `sync.Map`, but the point is that if you use the abstraction layer you can easily change the implementation later. That's the point of this package.
