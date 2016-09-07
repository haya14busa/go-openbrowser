## go-openbrower

[![GoDoc](https://godoc.org/github.com/haya14busa/go-openbrowser?status.svg)](https://godoc.org/github.com/haya14busa/go-openbrowser)

Package openbrowser provides a way to wait http server and open
browser. It's mainly for http server development or command line tool.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haya14busa/go-openbrowser"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi!")
	})

	port := "8080"

	go openbrowser.WaitAndStart(fmt.Sprintf("http://localhost:%s", port))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

### :bird: Author
haya14busa (https://github.com/haya14busa)

