// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package io provides gates for manipulating Go's I/O types.
package io

import (
	"io"
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
	kio "github.com/gocircuit/escher/kit/io"
)

func init() {
	faculty.Register(be.NewMaterializer(Chunk{}), "io", "ChunkUp")
}

// Chunk…
type Chunk struct{}

func (Chunk) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Chunk) CognizeReader(eye *be.Eye, v interface{}) {
	r := kio.NewChunkReader(v.(io.Reader))
	for {
		chunk, err := r.Read()
		if err != nil {
			return
		}
		eye.Show("Chunk", chunk)
	}
}

func (Chunk) CognizeChunk(_ *be.Eye, v interface{}) {}
