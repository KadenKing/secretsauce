# The Secret Sauce Pattern 

## What is it? 

The secret sauce pattern is a way of using generics to make it impossible to default ininitialize a type outside of its package.

## Should I actually do this?

idk 

## The Problem

Suppose we have a struct called `ComplicatedType` that requires setup in order to be behave correctly.

```go
package complicated

// ComplicatedType must be instantiated using NewComplicatedType.
type ComplicatedType struct {
	veryImportantMustBeValid string
}

// DoComplicatedThings requires ComplicatedType to have been constructed with NewComplicatedType otherwise it panics.
func (c *ComplicatedType) DoComplicatedThings() {
	if c.veryImportantMustBeValid != "hey" {
		panic("oh no everything is wrong")
	}
}
```

Normally, we provide a factory function that initializes the type with the necessary setup. 

```go 
package complicated

// NewComplicatedType returns a new ComplicatedType
func NewComplicatedType() ComplicatedType {
	return ComplicatedType{
		veryImportantMustBeValid: "hey",
	}
}
```

This wouldn't stop a user from accidentally using the default initializer and using the type. This can create runtime errors.

```go
package main

bad := complicated.ComplicatedType{}
bad.DoComplicatedThings() // panic
```

## The Solution

If we make `ComplicatedType` generic over `T` and `T` is constrained to be an unexported type, it would be impossible to use the default initializer outside of the `complicated` package. 

```go
package complicated

// secretSauce is an unexported type that must be passed as a type param in order to make protected types.
type secretSauce struct{}

// ComplicatedType must be instantiated using NewComplicatedType.
type ComplicatedType[T secretSauce] struct {
	veryImportantMustBeValid string
}
```

But we can use it in our factory function. 

```go
package complicated

// NewComplicatedType returns a new ComplicatedType
func NewComplicatedType() ComplicatedType[secretSauce] {
	return ComplicatedType[secretSauce]{
		veryImportantMustBeValid: "hey",
	}
}
```

The incorrectly innitialized type no longer compiles, and the user is completely unaware of the `secretSauce` type in the public API. What used to be a runtime error is now a compile time error!

```go
package main

// does not compile
// bad := complicated.ComplicatedType{}

good := complicated.NewComplicatedType()
good.DoComplicatedThings()
fmt.Println("everything went okay")
```