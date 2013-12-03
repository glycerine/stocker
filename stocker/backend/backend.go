package backend

import (
	"strings"
)

type Backend interface {
	Get(key string) (string, error)
	Set(key, value string) error
	SetWithTTL(key, value string, ttl int) error
	Remove(key string) error
	Publish(key, message string) error
	Subscribe(pattern string, process func(key, message string) error) error
	Unsubscribe(key string) error
}

func Key(components ...string) string {
	return strings.Join(components, "/")
}
