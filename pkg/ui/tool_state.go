package ui

import (
	"github.com/miu200521358/mlib_go/pkg/interface/app"
	"github.com/miu200521358/mlib_go/pkg/interface/controller"
	"github.com/miu200521358/mlib_go/pkg/interface/controller/widget"
	"github.com/miu200521358/pmx_translator/pkg/domain"
	"github.com/miu200521358/walk/pkg/walk"
)

type ToolState struct {
	App                  *app.MApp
	ControlWindow        *controller.ControlWindow
	TranslateModel       *domain.TranslateModel
	TranslateTab         *widget.MTabPage
	OriginalPmxPicker    *widget.FilePicker
	LangCsvPicker        *widget.FilePicker
	OutputPmxPicker      *widget.FilePicker
	OriginalCsvPmxPicker *widget.FilePicker
	OutputCsvPicker      *widget.FilePicker
	SaveButton           *walk.PushButton
	CsvTab               *widget.MTabPage
	CsvTableView         *CsvTableView
	TranslateTableView   *TranslateTableView
}

func NewToolState(app *app.MApp, controlWindow *controller.ControlWindow) *ToolState {
	toolState := &ToolState{
		App:            app,
		ControlWindow:  controlWindow,
		TranslateModel: domain.NewTranslateModel(),
	}

	newTranslateTab(controlWindow, toolState)
	newCsvTab(controlWindow, toolState)

	return toolState
}
