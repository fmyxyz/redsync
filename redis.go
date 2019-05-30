package redsync

import "github.com/go-redis/redis"

// A Pool maintains a pool of Redis connections.
type Pool interface {
	Get() *redis.Client
}
