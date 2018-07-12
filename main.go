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
	"errors"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const socketAddr = "/tmp/calculator.sock"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Run as client or server")
	}
	if strings.HasPrefix(os.Args[1], "c") {
		runClient(os.Args[2:])
	} else if strings.HasPrefix(os.Args[1], "s") {
		runServer(os.Args[2:])
	} else {
		log.Fatal("Run as client or server")
	}
}

func runClient(args []string) {
	log.Printf("Client: %v\n", args)
	client, err := rpc.Dial("unix", socketAddr)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	in := &Input{1, 2}
	var out int
	client.Call("Calculator.Add", in, &out)
	log.Printf("Result: %v", out)
}

func runServer(args []string) {
	log.Printf("Server: %v\n", args)
	server := rpc.NewServer()
	server.Register(new(Calculator))
	listener, err := net.Listen("unix", socketAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	cleanupOnSigTerm(listener)
	server.Accept(listener)
	log.Println("server done")
}

func cleanupOnSigTerm(listener net.Listener) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := listener.Close()
		if err != nil {
			log.Fatal("listener close error:", err)
			os.Exit(1)
		}
		log.Println("listener closed")
		os.Exit(0)
	}()
}

// Calculator can add
type Calculator int

// Input to add
type Input struct {
	A, B int
}

// Add integers
func (t *Calculator) Add(in *Input, out *int) error {
	log.Printf("Adding %v and %v", in.A, in.B)
	if in == nil {
		return errors.New("must supply input")
	}
	*out = in.A + in.B
	log.Printf("out: %v", *out)
	return nil
}
