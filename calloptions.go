/* Package calloptions provides a way to construct method call options that can be used between multiple methods
while providing type safety guarantees. This is a little verbose, but the godoc is short and the client is clean.

Example:

	type Client struct{}

	// aCallOptions holds the options used in the A() method call.
	type aCallOptions struct {
		accountName string
		id int
	}

	// methodAOption represents an options argument to A().
	type methodAOption interface{
		a() // This ensures only this package can implement this.
	}

	// WithAccount sets the AccountName to name. It can be used on A() or B().
	func WithAccount(name string) interface{ // The returned interface should implement CallOption and any method option that can use it.
		calloptions.CallOption
		methodAOption
		methodBOption
	}{
		return struct{ // This implements the returned interface.
			methodAOption
			methodBOption
			calloptions.CallOption
		}{
			CallOption: calloptions.New(
				func(a any) error{ // This applies the option
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

Now above we have an option, WithAccount() that can be used with .A() or .B() . We also have 
WithID() that can only be used with .A() and WithTurnOn() that can only be used with .B(). These
are all compile time checked. This example is runable in the example directory.
*/
package calloptions

// CallOption is a type that is implementing an optional argument to a method call.
type CallOption interface {
	Do(a any) error
	callOption()
}

// New returns a new CallOption where a call to the Do() method calls function "f".
func New(f func(a any) error) CallOption {
	if f == nil {
		panic("cannot pass a nil function")
	}
	return callOption(f)
}

// callOption is an adapter for a function to a CallOption.
type callOption func(a any) error

func (c callOption) Do(a any) error {
	return c(a)
}

func (c callOption) callOption() {}

// ApplyOptions applies all the callOptions to options. options must be a pointer to a struct and
// callOptions must be a list of objects that implement CallOption or it will panic.
func ApplyOptions[O any, C any](options O, callOptions []C) error {
	for _, o := range callOptions {
		t := any(o).(CallOption)
		if err := t.Do(options); err != nil {
			return err
		}
	}
	return nil
}
