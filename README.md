# CallOptions - One Option to Rule them All

## Introduction

Most of the time, optional arguments to methods are done with the help of option methods.  You see functions like `WithAccount()` that you pass in as a variadic argument.

But what happens when you end up needing to share an option? I've seen a few methods, like where you make a WithSharable(WithAccount()) kind of deal, but its ugly. You want the same semantics for the user as when they aren't shared and you don't want to mess up the godoc with a bunch of useless types. Other methods have a generic call option and give runtime errors when they are the wrong type of option passed.  But we want it to be a compile error.

Some people fall back to the passing structs like `AMethodOptions{}`. But that is passee for hardcore Go developer and we can do better.

This package and its methodology provide the same user semantics as non-shareable call options, but with support to have some call options work with multiple methods and compiler type safety.

The godoc has everything you need to understand how to you can provide:

```go
	client.UseToken(WithName("jdoak"), WithToken(token))
	client.UsePass(WithName("jdoak"), WithPass(pass))
	// Uncommenting this is a compile error, as WithPas isn't valid with UseToken()
	client.UseToken(WithName("jdoak"), WithPass(pass))
```
