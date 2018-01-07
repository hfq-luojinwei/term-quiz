package quiz

import (
	"log"
	"github.com/jroimartin/gocui"
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/crazcalm/term-quiz/answers"
)

//FBInit -- Initializes the ABCD gui interface
func FBInit(g *gocui.Gui, q *questions.Question, as answers.Answers, count string) (err error){
	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame", count)
	question := gui.NewQuestion("question", "question", q.Question)	
	answerBlank := gui.NewAnswer(gui.BoxBlank, gui.BoxBlank, as.Answers[0].Answer)

	g.SetManager(questionFrame, question, answerBlank)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	return nil
}