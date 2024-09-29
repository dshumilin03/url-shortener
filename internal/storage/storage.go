package storage

// TODO write for sqlite

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL already exists")
)
