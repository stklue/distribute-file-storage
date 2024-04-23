package main

import "testing"

func TestStorage(t *testing.T) {
	opts := StorageOpts{
		PathTransformFunc: DefaultTransformFunc,
	}

	storage = NewStorage(opts)

}
