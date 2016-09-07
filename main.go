// Package openbrowser provides a way to wait http server and open
// browser. It's mainly for http server development or command line tool.
package openbrowser

// Package openbrowser borrows gotour code.
// ref: https://github.com/golang/tour/blob/c9941e54e5b8e9618a8c951bc89798f85f6a7a71/gotour/local.go
//
// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var (
	// The number of count to retry to wait server.
	Tries = 20
	// Duration to wait before retry.
	Sleep = 100 * time.Millisecond
)

var sleepFunc = time.Sleep

// WaitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
func WaitServer(url string) bool {
	for Tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		sleepFunc(Sleep)
		Tries--
	}
	return false
}

// Start tries to open the URL in a browser.
func Start(url string) error {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start()
}

// WaitAndStart waits server up and tries to open the URL in a browser.
func WaitAndStart(url string) error {
	if WaitServer(url) {
		return Start(url)
	}
	return fmt.Errorf("%v didn't start", url)
}
