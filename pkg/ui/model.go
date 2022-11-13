package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-rod/rod"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
)

type Model struct {
	err error

	configFile string
	config     *config.Config

	ttyd *ttyd.Ttyd
	page *rod.Page

	quitting bool

	spinner spinner.Model
}

var _ tea.Model = (*Model)(nil)

func New(configFile string) *Model {
	return &Model{
		configFile: configFile,
		spinner:    spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styleSpinner)),
	}
}

func (m *Model) Err() error {
	return m.err
}

func (m *Model) Close() error {
	if m.ttyd == nil {
		return nil
	}

	if err := m.ttyd.Command.Process.Kill(); err != nil {
		return err
	}

	return nil
}

func (m *Model) running() bool {
	return m.page != nil
}
