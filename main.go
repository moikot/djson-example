package main

import (
	"os"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/moikot/djson"
)

type testCmd struct {
	values         []string
	strings		   []string
}

// NewGenerateCmd creates a new instance of generate command.
func NewTestCmd() *cobra.Command {
	inst := &testCmd{
	}

	cmd := &cobra.Command{
		Use:   "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			return inst.run(cmd)
		},
	}

	f := cmd.Flags()

	f.StringArrayVar(&inst.values, "set-value", []string{}, "set values on the command line")
	f.StringArrayVar(&inst.strings, "set-string", []string{}, "set strings on the command line")

	return cmd
}

func (c *testCmd) run(cmd *cobra.Command) error {
	m := map[string]interface{}{}
	for _, val := range c.values {
		err := djson.MergeValue(m, val)
		if err != nil {
			return err
		}
	}
	for _, str := range c.strings {
		err := djson.MergeString(m, str)
		if err != nil {
			return err
		}
	}
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Print(string(b))
	return nil
}

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "djson-example",
	}

	flags := cmd.PersistentFlags()
	cmd.AddCommand(
		NewTestCmd(),
	)

	_ = flags.Parse(args) // To make the linter happy.
	return cmd
}

func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
