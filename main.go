package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const timeout = 10 * time.Second // 5 * time.Minute

type model struct {
	timer timer.Model
	vp    viewport.Model
	quit  bool
}

type responseMsg struct{}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case tea.MouseMsg:
		var cmd tea.Cmd
		m.timer.Timeout = timeout
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		m.quit = true
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	s := m.timer.View()

	if m.timer.Timedout() {
		s = "More than 5 mins"
	}

	if m.quit {
		s += "\n"
	}

	return s
}

func initModel() model {
	return model{
		timer: timer.NewWithInterval(timeout, time.Second),
		vp:    viewport.New(0, 0),
	}
}

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
