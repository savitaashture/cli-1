// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package triggerbinding

import (
	"fmt"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/jonboulle/clockwork"
	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/actions"
	"github.com/tektoncd/cli/pkg/cli"
	"github.com/tektoncd/cli/pkg/formatted"
	"github.com/tektoncd/cli/pkg/validate"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	cliopts "k8s.io/cli-runtime/pkg/genericclioptions"
)

const describeTemplate = `{{decorate "bold" "Name"}}:	{{ .Triggerbinding.Name }}
{{decorate "bold" "Namespace"}}:	{{ .Triggerbinding.Namespace }}

{{decorate "params" ""}}{{decorate "underline bold" "Params\n"}}
{{- $l := len .Triggerbinding.Spec.Params }}{{ if eq $l 0 }}
No params
{{- else }}
 NAME	VALUE
{{- range $i, $p := .Triggerbinding.Spec.Params }}
{{- if eq $p.Value.Type "string" }}
 {{decorate "bullet" $p.Name }}	{{ $p.Value.StringVal }}
{{- else }}
 {{decorate "bullet" $p.Name }}	{{ $p.Value.ArrayVal }}
{{- end }}
{{- end }}
{{- end }}
`
//const describeTemplate = `{{decorate "bold" "Name"}}:	{{ .Triggerbinding.Name }}
//{{decorate "bold" "Namespace"}}:	{{ .Triggerbinding.Namespace }}
//
//{{decorate "params" ""}}{{decorate "underline bold" "Params\n"}}
//{{- $l := len .Triggerbinding.Spec.Params }}{{ if eq $l 0 }}
//No params
//{{- else }}
// NAME     VALUE
//{{- range $i, $p := .Triggerbinding.Spec.Params }}
//{{- if eq $p.Value.Type "string" }}
// {{decorate "bullet" $p.Name }}     {{ $p.Value.StringVal }}
//{{- else }}
// {{decorate "bullet" $p.Name }}     {{ $p.Value.ArrayVal }}
//{{- end }}
//{{- end }}
//{{- end }}
//`

/*

{{- $l := len .Triggerbinding.Spec.Params }}{{ if eq $l 0 }}
No params
{{- else }}
 NAME	VALUE
{{- range $i, $p := .Triggerbinding.Spec.Params }}
{{- if eq $p.Value.Type "string" }}
 {{decorate "bullet" $p.Name }} {{ $p.Value.StringVal }}
{{- else }}
 {{decorate "bullet" $p.Name }} {{ $p.Value.ArrayVal }}
{{- end }}
{{- end }}
{{- end }}
*/

func describeCommand(p cli.Params) *cobra.Command {
	f := cliopts.NewPrintFlags("describe")
	eg := `Describe a Triggerbinding of name 'foo' in namespace 'bar':

    tkn triggerbinding describe foo -n bar

or

   tkn tb desc foo -n bar
`

	c := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"desc"},
		Short:   "Describes a triggerbinding in a namespace",
		Example: eg,
		Annotations: map[string]string{
			"commandType": "main",
		},
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s := &cli.Stream{
				Out: cmd.OutOrStdout(),
				Err: cmd.OutOrStderr(),
			}

			if err := validate.NamespaceExists(p); err != nil {
				return err
			}

			output, err := cmd.LocalFlags().GetString("output")
			if err != nil {
				fmt.Fprint(os.Stderr, "Error: output option not set properly \n")
				return err
			}

			if output != "" {
				tbGroupResource := schema.GroupVersionResource{Group: "trigger.tekton.dev", Resource: "triggerbindings"}
				return actions.PrintObject(tbGroupResource, args[0], cmd.OutOrStdout(), p, f, p.Namespace())
			}

			return printTriigerBindingDescription(s, p, args[0])
		},
	}

	_ = c.MarkZshCompPositionalArgumentCustom(1, "__tkn_get_triggerbindings")
	f.AddFlags(c)
	return c
}

func printTriigerBindingDescription(s *cli.Stream, p cli.Params, tbName string) error {
	cs, err := p.Clients()
	if err != nil {
		return fmt.Errorf("failed to create tekton client")
	}

	tb, err := cs.Triggers.TriggersV1alpha1().TriggerBindings(p.Namespace()).Get(tbName, metav1.GetOptions{})
	if err != nil {
		fmt.Fprintf(s.Err, "failed to get triggerbinding %s\n", tbName)
		return err
	}

	var data = struct {
		Triggerbinding     *v1alpha1.TriggerBinding
		Time     clockwork.Clock
	}{
		Triggerbinding:     tb,
		Time:     p.Time(),
	}

	fmt.Println("TRIGGERE BINDING VALUES", tb.Spec.Params)
	funcMap := template.FuncMap{
		"decorate":        formatted.DecorateAttr,
	}

	w := tabwriter.NewWriter(s.Out, 0, 5, 3, ' ', tabwriter.TabIndent)
	tparsed := template.Must(template.New("Describe Triggerbinding").Funcs(funcMap).Parse(describeTemplate))
	err = tparsed.Execute(w, data)
	//fmt.Println("DDDDDDDDDDDDDDDDDDDDD", data.Triggerbinding.Spec.Params)
	if err != nil {
		fmt.Fprintf(s.Err, "Failed to execute template \n")
		return err
	}
	return nil
}