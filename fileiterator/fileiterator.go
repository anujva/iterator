package fileiterator

import (
	"bufio"
	"context"
	"os"

	"github.com/anujva/iterator"
	"github.com/anujva/utils"
	"github.com/thumbtack/go/lib/logging"
)

var log = logging.NewLogger()

// FileIterator implements the iterator interface for a file
type FileIterator struct {
	// A file iterator will need a file scanner
	f       *os.File
	scanner *bufio.Scanner
	line    *string
}

// New returns an implementation of file Iterator
// for the given file
func New(f *os.File) iterator.Iterator {
	scanner := bufio.NewScanner(f)
	// read in  a line to initialize the state of the
	// file iterator.
	var line *string
	if scanner.Scan() {
		line = utils.StrPtr(scanner.Text())
	}

	return &FileIterator{
		f:       f,
		scanner: scanner,
		line:    line,
	}
}

var _ iterator.Iterator = &FileIterator{}

// HasNext returns true/false if the iterator has more values
// to return
func (fi *FileIterator) HasNext() bool {
	if fi.line == nil {
		fi.f.Close()
		return false
	}
	return true
}

// Next actually returns the value, the value will be string wrapped
// in interface{}
func (fi *FileIterator) Next(ctx context.Context) interface{} {
	tempLine := fi.line
	if fi.scanner.Scan() {
		fi.line = utils.StrPtr(fi.scanner.Text())
	} else {
		fi.line = nil
	}
	return tempLine
}

// Close closes the iterator and the resources that it might be consuming
func (fi *FileIterator) Close() bool {
	err := fi.f.Close()
	if err != nil {
		log.Errorf("error occurred while closing the file: %+v", err)
		return false
	}
	return true
}

// Reset resets the iterator to the start
func (fi *FileIterator) Reset() bool {
	f, _ := os.Open(fi.f.Name())
	fi.f.Close()
	scanner := bufio.NewScanner(f)
	// read in  a line to initialize the state of the
	// file iterator.
	var line *string

	if scanner.Scan() {
		line = utils.StrPtr(scanner.Text())
	}

	fi.f = f
	fi.line = line
	fi.scanner = scanner

	return true
}
