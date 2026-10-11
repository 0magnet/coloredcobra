//go:build tinygo

package coloredcobra

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestManualHelpSections(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	root := &cobra.Command{Use: "app", Run: func(*cobra.Command, []string) {}}
	sub := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Example: "  app serve -a :8080",
		Short:   "serve it",
		Run:     func(*cobra.Command, []string) {},
	}
	root.AddCommand(sub)
	Init(&Config{RootCmd: root, Headings: HiBlue})

	var b bytes.Buffer
	sub.SetOut(&b)
	sub.SetErr(&b)
	if err := sub.Usage(); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Usage:", "Aliases:\n  serve, s", "Examples:\n  app serve -a :8080"} {
		if !strings.Contains(b.String(), want) {
			t.Errorf("usage lacks %q:\n%s", want, b.String())
		}
	}
}
