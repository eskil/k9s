// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/tview"
)

// Clocks displays one or more timezone clocks in the header.
type Clocks struct {
	*tview.TextView

	configs []config.ClockConfig
	app     *tview.Application
}

// NewClocks returns a new Clocks widget for the given timezone configs.
func NewClocks(app *tview.Application, configs []config.ClockConfig, styles *config.Styles) *Clocks {
	c := &Clocks{
		TextView: tview.NewTextView(),
		configs:  configs,
		app:      app,
	}
	c.SetDynamicColors(true)
	c.SetTextAlign(tview.AlignRight)
	c.SetWrap(false)
	c.SetBackgroundColor(styles.BgColor())
	c.refresh()

	return c
}

// Width returns the preferred display width for the widget.
func (c *Clocks) Width() int {
	max := 0
	for _, cfg := range c.configs {
		label := cfg.Label
		if label == "" {
			label = cfg.Timezone
		}
		// "LABEL HH:MM:SS+" (extra char for day indicator)
		if w := len(label) + 1 + 8 + 1; w > max {
			max = w
		}
	}
	if max < 10 {
		max = 10
	}
	return max + 2 // side padding
}

// Watch starts the per-second refresh goroutine, stopping when ctx is cancelled.
func (c *Clocks) Watch(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				c.app.QueueUpdateDraw(func() {
					c.refresh()
				})
			}
		}
	}()
}

func (c *Clocks) refresh() {
	c.Clear()
	now := time.Now()
	localYear, localMonth, localDay := now.Date()
	for i, cfg := range c.configs {
		if i > 0 {
			fmt.Fprint(c, "\n")
		}
		loc, err := time.LoadLocation(cfg.Timezone)
		if err != nil {
			loc = time.UTC
		}
		t := now.In(loc)
		label := cfg.Label
		if label == "" {
			label = cfg.Timezone
		}
		clockYear, clockMonth, clockDay := t.Date()
		localDate := time.Date(localYear, localMonth, localDay, 0, 0, 0, 0, time.Local)
		clockDate := time.Date(clockYear, clockMonth, clockDay, 0, 0, 0, 0, time.Local)
		dayIndicator := " "
		switch {
		case clockDate.After(localDate):
			dayIndicator = "+"
		case clockDate.Before(localDate):
			dayIndicator = "-"
		}
		fmt.Fprintf(c, "[::b]%s[::-] %s%s", label, t.Format("15:04:05"), dayIndicator)
	}
}
