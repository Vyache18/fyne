package theme

import (
	"testing"

	"github.com/fyne-io/fyne"
	"github.com/stretchr/testify/assert"
)

// TODO figure how to test the default theme if we are messing with it...
//func TestThemeDefault(t *testing.T) {
//	loaded := fyne.GlobalSettings().Theme()
//	assert.NotEqual(t, lightBackground, loaded.BackgroundColor())
//}

func TestThemeChange(t *testing.T) {
	fyne.GlobalSettings().SetTheme(DarkTheme())
	bg := BackgroundColor()

	fyne.GlobalSettings().SetTheme(LightTheme())
	assert.NotEqual(t, bg, BackgroundColor())
}
