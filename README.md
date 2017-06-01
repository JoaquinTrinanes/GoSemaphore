# Go Semaphore

A simple semaphore implementation for Go.

## Installation
The installation process is pretty straightforward. Open your terminal and type:

```
go get github.com/JoaquinTrinanes/GoSemaphore
```

To use the package in your Go code, simply add `"github.com/JoaquinTrinanes/GoSemaphore"` to the import at the beginning of the code.

Once imported, you'll be set.

## Usage

To create a semaphore, simply create a `*Semaphore` type variable:

```
var s Semaphore
```

This will create a semaphore with 0 value. If you want to initialize it at any other value, use the `SemInit` function, which will return a pointer to a semaphore:


```
var s *Semaphore = SemInit(value)
```

or

```
s := SemInit(value)
```

Note that `value` is an uint32.

Now, you can perform the following methods on the semaphore:

 - `Down()` will block the caller goroutine if the semaphore value is 0, or decrement it otherwise. A blocked goroutine will attempt to perform the Down operation once it's unlocked.

 - `Up()` increments the semaphore value, unlocking every possible goroutine previously blocked by it.

 - `Value()` returns the semaphore current value.

 - `TryDown()` will attempt to decrement the semaphore value. If the operation was a success it will return true. Otherwise, it won't do anything and return false. Note that the goroutine won't block under any circunstances when calling this method.
