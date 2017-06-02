A tiny Go library for reading passwords typed in a terminal without echoing
output, inspired by Python's [getpass](https://docs.python.org/3.6/library/getpass.html).

It is compatible with modern Windows and *nix systems that use `cmd.exe`,
`powershell`, `bash`, or `sh`.

Example:
```go
import "github.com/princebot/getpass"

func main() {
    pw, err := getpass.Get("prompt to display to user goes here")
    if err != nil {
        // Do something with password.
    }
}
```
