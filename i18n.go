package fyne

var GetI18NString func(string, string) string = func(k string, def string) string {
	return def
}
