package tea

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewLoop configure loop with default settings.
func NewLoop() *Loop {
	return &Loop{
		ShutdownTimeout: 10 * time.Second,
		QuitSignals:     []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		onQuit: func(s os.Signal) bool {
			return true
		},
	}
}

// Loop
type Loop struct {
	ShutdownTimeout time.Duration
	QuitSignals     []os.Signal

	onQuit     func(os.Signal) bool
	onShutdown func(context.Context)
}

// OnQuit callback. Return true to break loop.
func (l *Loop) OnQuit(fn func(os.Signal) bool) {
	l.onQuit = fn
}

// OnShutdown callback is run after OnQuit returns true.
func (l *Loop) OnShutdown(fn func(context.Context)) {
	l.onShutdown = fn
}

// Run beings main loop.
func (l *Loop) Run() {
	quit := make(chan os.Signal)
	signal.Notify(quit, l.QuitSignals...)

LOOP:
	for {
		select {
		case sig := <-quit:
			if l.onQuit(sig) {
				break LOOP
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), l.ShutdownTimeout)
	defer cancel()
	l.onShutdown(ctx)
}
