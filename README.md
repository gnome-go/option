# Gnome-Go Option

![logo](./assets/icon.png)

A functional option package for handling optional values.

```go
import(
    "fmt"
    "github.com/gnome-go/option"
)


func Example() {
    s := option.Some(10)

    fmt.Printf("IsSome() %v", s.IsSome());
    fmt.Printf("IsNone() %v", s.IsNone());

    if s.IsSomeAnd(func(c int) { c > 5 }) {
        s.Inspect(func(c int) {
            fmt.Printf("%v is greater than 5\n", c)
        })
    } else {
        s.Inspect(func(c int) {
            fmt.Printf("%v is 5 or less\n", c)
        })
    }

    v := option.Unwrap()
    fmt.Printf("%v\n", v)
}

```
