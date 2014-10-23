// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
)

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize(*Matter) (Reflex, Value) {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize)
	}()
	return Reflex{DefaultValve: t}, nil
}

func MaterializeNoun(matter *Matter, v interface{}) (Reflex, Value) {
	return MaterializeNative(matter, Noun{}, v)
}

func NewNoun(v interface{}) Materializer {
	return NewNativeMaterializer(Noun{}, v)
}

// Noun
type Noun struct{}

func (n Noun) Spark(eye *Eye, matter *Matter, aux ...interface{}) Value {
	go func() {
		for vlv, _ := range matter.View.Gate {
			eye.Show(vlv, aux[0])
		}
	}()
	if matter.View.Len() == 0 {
		return aux[0]
	}
	return nil
}

func (n Noun) OverCognize(*Eye, Name, interface{}) {}

// Future
type Future struct {
	eye *Eye
	view Circuit
}

func (f *Future) Spark(eye *Eye, matter *Matter, _ ...interface{}) Value {
	f.eye = eye
	f.view = matter.View
	return nil
}

func (f *Future) Charge(v Value) {
	go func() {
		for vlv, _ := range f.view.Gate {
			f.eye.Show(vlv, DeepCopy(v))
		}
	}()
}

func (f *Future) OverCognize(*Eye, Name, interface{}) {}
