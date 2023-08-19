package glfw

func (w *window) doSetMaximize() {
	w.viewport.Maximize()
}

func (w *window) SetMaximize() {
	if w.view() != nil {
		runOnMain(w.doSetMaximize)
		return
	}

	w.afterCreate = append(w.afterCreate, w.doSetMaximize)
}
