# filecache

A filecache used directly with default setting. Based on —— https://github.com/huntsman-li/go-cache

## How To Use

    go get github.com/vcgo/filecache

```go
package main

import (
	"fmt"
	"time"

	. "github.com/vcgo/filecache"
)

func main() {
	Cache.Put("github.com/wilon", 2333, 5)
	for i := 0; i < 10; i++ {
		fmt.Println("...", Cache.Get("github.com/wilon"))
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
```

And more func——

```go
	// Put puts value into cache with key and expire time.
	Put(key string, val interface{}, timeout int64) error
	// Get gets cached value by given key.
	Get(key string) interface{}
	// Delete deletes cached value by given key.
	Delete(key string) error
	// Incr increases cached int-type value by given key as a counter.
	Incr(key string) error
	// Decr decreases cached int-type value by given key as a counter.
	Decr(key string) error
	// IsExist returns true if cached value exists.
	IsExist(key string) bool
	// Flush deletes all cached data.
	Flush() error
	// StartAndGC starts GC routine based on config string settings.
	StartAndGC(opt Options) error
```

## Test

    go run example/cache.go

## More requirements?

Just use https://github.com/huntsman-li/go-cache.