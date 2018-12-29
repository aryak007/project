package main

import (
	cache "github.com/aryak93/project/internal/cache"
)

func main() {
	cache.SaveToCache(".cache", "cache")
}
