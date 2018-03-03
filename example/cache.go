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
