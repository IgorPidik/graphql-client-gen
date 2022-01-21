package gcg

var enumsTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/gcg on {{.Created.Format "Mon, 02 Jan 2006 15:04:05 MST" }} 
// DO NOT EDIT

{{- range .Enums}}

// {{.Name}} is an Enum
type {{.Name}} string {{$enumName := .Name}}
const (
{{- range .Values}}
{{$enumName}}_{{.}} = {{$enumName}}("{{.}}")
{{- end}}
)
{{- end}}
`
