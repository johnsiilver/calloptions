# CallOptions - One Option to Rule them All

## Introduction

Most of the time, optional arguments to methods are done with the help of option methods.  You see functions like `WithAccount()` that you pass in as a variadic argument.

But what happens when you end up needing to share an option? I've seen a few methods to achieve this functionality, but they generally are lacking in some feature.

You can use a method like WithSharable(WithAccount()), but its ugly as your wrapping a With() inside a WithSharable(). 

We want the same semantics for the user as when they aren't shared.

Other methods have a generic call option and give runtime errors when they are the wrong type of option passed.  But we want it to be a compile error.

Some people fall back to the passing structs like `AMethodOptions{}`. But that has other problems with setting defaults correctly and having "optional" values that aren't actually optional. We can do better.

This package and its methodology provide the same user semantics as non-shareable call options, but with support to have some call options work with multiple methods and compiler type safety.

The godoc has everything you need to understand how ou can provide a user experience like this (with compiler checks):

```go
	client.UseToken(WithName("jdoak"), WithToken(token))
	client.UsePass(WithName("jdoak"), WithPass(pass))
	// Uncommenting this is a compile error, as WithPass isn't valid with UseToken()
	// client.UseToken(WithName("jdoak"), WithPass(pass))
```
