package tester

import (
	"io"

	"github.com/clipperhouse/typewriter"
	"github.com/davecgh/go-spew/spew"
)

type Tester struct{}

func (tw *Tester) Name() string {
	return "tester"
}

func (tw *Tester) Imports(t typewriter.Type) (res []typewriter.ImportSpec) {
	return
}

func (tw *Tester) Write(w io.Writer, t typewriter.Type) error {
	spew.Dump(t)
	return nil
}
