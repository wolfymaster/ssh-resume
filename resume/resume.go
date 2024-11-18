package resume

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Program struct{}

func (p Program) Init() tea.Cmd {
	return nil
}

func (p Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p Program) View() string {
	return "hello world"
}
