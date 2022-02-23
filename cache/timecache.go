package cache

import (
	"sync"
	"time"
)

type TimeCache struct {
	*timecache
}

type timecache struct {
	defaultExpiration time.Duration
	items             sync.Map
	mu                sync.Mutex
	onEvicted         func(string, interface{}, bool)
	// janitor           *janitor
}
