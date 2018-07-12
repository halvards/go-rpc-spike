// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Run as client or server")
		os.Exit(1)
	}
	if strings.HasPrefix(os.Args[1], "c") {
		client(os.Args[2:])
	} else if strings.HasPrefix(os.Args[1], "s") {
		server(os.Args[2:])
	} else {
		fmt.Println("Run as client or server")
		os.Exit(1)
	}
}

func client(args []string) {
	fmt.Printf("Client: %v\n", args)
}

func server(args []string) {
	fmt.Printf("Server: %v\n", args)
}
