package main

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/jroimartin/gocui"
	"log"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	questionFrame := gui.NewQuestionFrame("questionFrame", "count")
	question := gui.NewQuestion("question", "title - question", "question string")
	answerBlank := gui.NewAnswer("answerBlank", gui.BoxBlank, "")

	g.SetManager(questionFrame, question, answerBlank)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}