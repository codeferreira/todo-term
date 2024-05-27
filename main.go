package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type task struct {
	title       string
	description string
}

type keymap struct {
	up, down, left, right, quit, add key.Binding
}

type model struct {
	keymap keymap
	help   help.Model
	focus  int
}

func initialModel() model {
	return model{
		keymap: keymap{
			up: key.NewBinding(
				key.WithKeys("k", "up"),
				key.WithHelp("↑/k", "move up"),
			),
			down: key.NewBinding(
				key.WithKeys("j", "down"),
				key.WithHelp("↓/j", "move down"),
			),
			left: key.NewBinding(
				key.WithKeys("h", "left"),
				key.WithHelp("←/h", "move left"),
			),
			right: key.NewBinding(
				key.WithKeys("l", "right"),
				key.WithHelp("→/l", "move right"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			add: key.NewBinding(
				key.WithKeys("a"),
				key.WithHelp("a", "add new"),
			),
		},
		help: help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.left,
		m.keymap.right,
		m.keymap.quit,
	})
	return "Hello world\n\n" + help
}

func main() {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
