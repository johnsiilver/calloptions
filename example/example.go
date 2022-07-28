package main

import (
	"log"

	"github.com/johnsiilver/calloptions"
)

type Client struct{}

type aCallOptions struct {
	accountName string
	id int
}

type methodAOption interface{
	a()
}

// WithAccount sets the AccountName to name. It can be used on A() or B().
func WithAccount(name string) interface{
	calloptions.CallOption
	methodAOption
	methodBOption
}{
	return struct{
		methodAOption
		methodBOption
		calloptions.CallOption
	}{
		CallOption: calloptions.New(
			func(a any) error{
				switch t := a.(type) {
				case *aCallOptions:
					t.accountName = name
				case *bCallOptions:
					t.accountName = name
				default:
					panic("bug")
				}
				return nil
			},
		),
	}
}

// WithID sets the ID. It can be used on A(). 
func WithID(id int) interface{
	methodAOption
	calloptions.CallOption
}{
	return struct{
		methodAOption
		calloptions.CallOption
	}{
		CallOption: calloptions.New(
			func(a any) error{
				a.(*aCallOptions).id = id
				return nil
			},
		),
	}
}
	

func (c Client) A(options ...methodAOption) error {
	opts := aCallOptions{}
	if err := calloptions.ApplyOptions(&opts, options); err != nil {
		return err
	}
	log.Printf("method A received: %#+v", opts)
	return nil
}

type bCallOptions struct {
	accountName string
	turnOn bool
}

type methodBOption interface{
	b()
}

// WithTurnOn turns on something. It only works with B().
func WithTurnOn() interface{
	methodBOption
	calloptions.CallOption
}{
	return struct{
		methodBOption
		calloptions.CallOption
	}{
		CallOption: calloptions.New(
			func(a any) error{
				a.(*bCallOptions).turnOn = true
				return nil
			},
		),
	}
}


func (c Client) B(options ...methodBOption) error {
	opts := bCallOptions{}
	if err := calloptions.ApplyOptions(&opts, options); err != nil {
		return err
	}
	log.Printf("method B received: %#+v", opts)
	return nil
}

func main() {
	c := Client{}
	c.A(WithAccount("John Doak"), WithID(3))
	c.B(WithAccount("David Luyer"), WithTurnOn())

	// Uncommenting this will not cause a compile error.
	// c.B(WithID(2))
}
