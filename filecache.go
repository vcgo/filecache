package cache

import (
	"os"
	"path"

	"github.com/huntsman-li/go-cache"
)

var Cache cache.Cache

func init() {
	pwd, _ := os.Getwd()
	dir := path.Join(pwd, "caches")
	cache, err := cache.Cacher(cache.Options{
		Adapter:       "file",
		AdapterConfig: dir,
		Interval:      2,
	})
	if err != nil {
		panic(err)
	}
	Cache = cache
}
