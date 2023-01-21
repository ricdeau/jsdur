# jsdur - package with type Duration

Contains type Duration, that embeds standard *time.Duration*.

This type implements json and text Marshaler/Unmarshaler interfaces.

### Marshaling rules
* Zero duration will marshal into empty string.
* Unmarshal from "0" or empty string leads to zero duration.
* In other cases follow *time.Duration* *Parse* and *String* function behavior.


### Example
```go
package main

import (
    "github.com/ricdeau/jsdur"
    "encoding/json"
    "fmt"
)

type SomeStruct struct {
    Duration jsdur.Duration `json:"duration"`
}

func main() {
    var data SomeStruct
    _ = json.Unmarshal([]byte(`{"duration":"10h15m"}`), &data)
    fmt.Println(data.Duration)
}
```