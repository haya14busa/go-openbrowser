package openbrowser_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haya14busa/go-openbrowser"
)

func ExampleWaitAndStart() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi!")
	})

	port := "8080"

	go openbrowser.WaitAndStart(fmt.Sprintf("http://localhost:%s", port))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
