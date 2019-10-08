## Installation

As a library

```shell
go get github.com/GishaJamesB/pnp
```

## Usage

```go
package main

import (
    "github.com/GishaJamesB/pnp"
)

func main() {
  result := pnp.GetToken("endPoint", "clientId", "clientSecret")
	fmt.Println(result)
	fmt.Println(result["access_token"])
}
```