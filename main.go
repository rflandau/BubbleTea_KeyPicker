package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	p := tea.NewProgram(tui{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type tui struct {
}

func (t tui) Init() tea.Cmd {
	return nil
}

func (t tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyEscape {
			return t, tea.Quit
		}

		return t, tea.Printf("%v | alt? %v | type: %v | paste? %v | runes: %v",
			keyMsg.String(), keyMsg.Alt, keyMsg.Type, keyMsg.Paste, keyMsg.Runes)
	}

	return t, nil
}

func (t tui) View() string {
	return "Please press a key.\nPress escape to quit."
}
