package main

import (
	"changeme/handler"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx           context.Context
	serialHandler *handler.SerialHandler
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSerial() []string {
	list := a.serialHandler.GetSerial()
	return list
}

func (a *App) OpenPort(portName string) error {
	return a.serialHandler.OpenPort(a.ctx, portName)
}

func (a *App) ClosePort() {
	a.serialHandler.ClosePort()
}
