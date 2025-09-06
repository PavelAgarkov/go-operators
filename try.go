package try

import (
	"errors"
	"log"
)

type catch struct {
	target error
	fn     func(err error)
}

type tryOperator struct {
	try       func() error
	catches   []catch
	defaultFn func(err error)
	finally   func()
}

type option func(*tryOperator)

func Do(fn func() error) option {
	return func(t *tryOperator) {
		t.try = fn
	}
}

func Catch(target error, fn func(err error)) option {
	return func(t *tryOperator) {
		t.catches = append(t.catches, catch{target: target, fn: fn})
	}
}

func Default(fn func(err error)) option {
	return func(t *tryOperator) {
		t.defaultFn = fn
	}
}

func Finally(fn func()) option {
	return func(t *tryOperator) { t.finally = fn }
}

func Try(opts ...option) {
	t := &tryOperator{}
	for _, o := range opts {
		o(t)
	}
	if t.try == nil {
		panic("Try block is required")
	}
	t.do()
}

func (t *tryOperator) do() {
	defer func() {
		if t.finally != nil {
			t.finally()
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	err := t.try()
	if err != nil {
		handled := false
		for _, c := range t.catches {
			if errors.Is(err, c.target) {
				c.fn(err)
				handled = true
				break
			}
		}
		if !handled && t.defaultFn != nil {
			t.defaultFn(err)
		} else if !handled {
			log.Printf("Unhandled error: %v", err)
		}
	}
}
