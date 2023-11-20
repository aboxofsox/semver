# Semver

A Go package for parsing and comparing semantic versions

## Installation
```shell
go get github.com/aboxofsox/semver
```

### Usage
```go
import (
    "fmt"
    "github.com/aboxofsox/semver"
)

func main() {
    v1, v2 := "1.0.0", "1.0.1"
    result, err := semver.Compare(v1, v2)
    if err != nil {
        panic(err)
    }

    switch result {
        case -1:
            // upgrade?
        case 0:
            // do nothing?
        case 1:
            // downgrade?
    }
}
```

### Functions
- `Compare(v1, v2 string) (int, error)`: Compares two semantic versions. Returns -1 if v1 < v2, 1 if v1 > v2, and 0 if v1 == v2.
- `ParseVersion(v string) (Semver, error)`: Parses a semantic version string into a `Semver` struct.

### Testing
```shell
go test
```

