// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Lessons() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"Page\" hx-trigger=\"load\" hx-get=\"/lessons/prev\" hx-target=\"#lesson\"><div class=\"pgButtons\"><a id=\"prev\" hx-get=\"/lessons/prev\" hx-target=\"#lesson\" hx-swap=\"innerHTML scroll:#lesson:top\">&lt;-- Previous lesson</a> <a id=\"next\" hx-get=\"/lessons/next\" hx-target=\"#lesson\" hx-swap=\"innerHTML scroll:#lesson:top\">Next lesson --&gt;</a></div><br><div class=\"LessonWrapper\"><div id=\"lesson\"></div></div><div class=\"pgButtons\"><a id=\"prev\" hx-get=\"/lessons/prev\" hx-target=\"#lesson\" hx-swap=\"innerHTML scroll:#lesson:top\">&lt;-- Previous lesson</a> <a id=\"next\" hx-get=\"/lessons/next\" hx-target=\"#lesson\" hx-swap=\"innerHTML scroll:#lesson:top\">Next lesson --&gt;</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
