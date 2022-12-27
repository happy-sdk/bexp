// Copyright 2022 The Happy Authors
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

package cli

// // TmplParser enables to parse templates for cli apps.
// type cliTmplParser struct {
// 	tmpl   string
// 	buffer bytes.Buffer
// 	t      *template.Template
// }

// // SetTemplate sets template to be parsed.
// func (t *cliTmplParser) SetTemplate(tmpl string) {
// 	t.tmpl = tmpl
// }

// // ParseTmpl parses template for cli application
// // arg name is template name, arg info is common passed to template
// // and elapsed is time duration used by specific type of templates and can usually set to "0".
// func (t *cliTmplParser) ParseTmpl(name string, h interface{}, elapsed time.Duration) error {
// 	t.t = template.New(name)
// 	t.t.Funcs(template.FuncMap{
// 		"funcTextBold":    t.textBold,
// 		"funcCmdCategory": t.cmdCategory,
// 		"funcCmdName":     t.cmdName,
// 		"funcFlagName":    t.flagName,
// 		"funcDate":        t.dateOnly,
// 		"funcYear":        t.year,
// 		"funcElapsed":     func() string { return elapsed.String() },
// 	})
// 	tmpl, err := t.t.Parse(t.tmpl)
// 	if err != nil {
// 		return err
// 	}
// 	err = tmpl.Execute(&t.buffer, h)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // String returns parsed template as string.
// func (t *cliTmplParser) String() string {
// 	return t.buffer.String()
// }

// func (t *cliTmplParser) cmdCategory(s string) string {
// 	if s == "" {
// 		return s
// 	}
// 	return strings.ToUpper(s)
// }

// func (t *cliTmplParser) cmdName(s string) string {
// 	if s == "" {
// 		return s
// 	}
// 	return fmt.Sprintf("\033[1m %-20s\033[0m", s)
// }

// func (t *cliTmplParser) flagName(s string, a string) string {
// 	if s == "" {
// 		return s
// 	}
// 	if len(a) > 0 {
// 		s += ", " + a
// 	}
// 	return fmt.Sprintf("%-25s", s)
// }

// func (t *cliTmplParser) textBold(s string) string {
// 	if s == "" {
// 		return s
// 	}
// 	return fmt.Sprintf("\033[1m%s\033[0m", s)
// }

// func (t *cliTmplParser) dateOnly(ts time.Time) string {
// 	y, m, d := ts.Date()
// 	return fmt.Sprintf("%.2d-%.2d-%d", d, m, y)
// }

// func (t *cliTmplParser) year() int {
// 	return time.Now().Year()
// }
