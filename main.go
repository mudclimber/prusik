package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/mudclimber/relay/pkg/handler"
	"github.com/mudclimber/relay/pkg/server"
)

type PrusikHandler struct{}

func (s PrusikHandler) HandleInit(a *handler.HandlerActions, login string) error {
	createCommand := fmt.Sprintf("create %s %s\n", login, THROWAWAY_PW)
  a.SendBytes([]byte(createCommand))

	time.Sleep(50 * time.Millisecond)
	result := a.ReadUntilSize(135, 1000)
	if bytes.Contains(result.Bytes(), []byte("Is this what you intended")) {
    a.SendBytes([]byte("y\n"))
		result.Reset()
		result := a.ReadUntilSize(30, 1000)
		if bytes.Contains(result.Bytes(), []byte("that username is already taken")) {
		} else {
		}
	} else {
    a.SendBytes([]byte(result.String()))
		return nil
	}

  connectCommand := fmt.Sprintf("connect %s %s\n", login, THROWAWAY_PW)
  a.SendBytes([]byte(connectCommand))
  introMsg := a.ReadUntilSize(20, 1000)
  a.Intro = introMsg.Bytes()
	return nil
}

func (s PrusikHandler) ParseOutput(buf *[]byte) { /* noop */ }

func main() {
  h := PrusikHandler{}
  opts := handler.HandlerOptions {
    DisplayName: "(Demo) Prusik - Basically a stock MUD",
    Port: 12383,
  }
  server.Run(h, opts)
}
