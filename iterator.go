package iterator

import "context"

// Iterator defines the interface for any iterative behavior of a code
type Iterator interface {
	HasNext() bool
	Next(context.Context) interface{}
	Reset() bool
	Close() bool
}
