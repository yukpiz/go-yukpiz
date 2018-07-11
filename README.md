# go-yukpiz

# Install

```
$ go get github.com/yukpiz/go-yukpiz
```

*main.go*
```go
package main

import yuk "github.com/yukpiz/go-yukpiz"

...
```

# UniqInts

```go
func main() {
    i := []int{1, 2, 3, 1, 2, 3}
    ii := yuk.UniqInts(i)
    fmt.Println(ii) // => []int {1, 2, 3}
}
```