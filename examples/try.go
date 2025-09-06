package main

import (
	"errors"
	try "go-operators"
	"log"
)

var (
	MyCustomError = errors.New("CustomError")
	AnotherError  = errors.New("AnotherError")
	ThirdError    = errors.New("ThirdError")
)

func main() {
	try.Try(
		try.Do(func() error {
			log.Printf("Executing try block")
			//return MyCustomError
			//return AnotherError
			//return ThirdError
			return errors.New("Some other error")
			//return nil
		}),
		try.Catch(MyCustomError, func(err error) {
			log.Printf("Caught MyCustomError: %v", err)
		}),
		try.Catch(AnotherError, func(err error) {
			log.Printf("Caught AnotherError: %v", err)
		}),
		try.Catch(ThirdError, func(err error) {
			log.Printf("Caught ThirdError: %v", err)
		}),
		try.Default(func(err error) {
			log.Printf("Default handler (any error): %v", err)
		}),
		try.Finally(func() {
			log.Printf("Finally block executed")
		}),
	)
}
