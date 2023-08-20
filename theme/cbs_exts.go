package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
)

func LookupColor(s string) (color.Color, bool) {
	col := current().Color(fyne.ThemeColorName(s), currentVariant())
	if col == nil {
		return color.Black, false
	}
	return col, true
}

func LookupResource(n string) (fyne.Resource, bool) {
	icon := current().Icon(fyne.ThemeIconName(n))
	if icon == nil {
		return nil, false
	}
	return icon, true
}
