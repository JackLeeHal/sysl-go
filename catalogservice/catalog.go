package catalogservice

import (
	"net/http"

	"github.com/anz-bank/sysl-go/handlerinitialiser"

	"github.com/go-chi/chi"

	"github.com/anz-bank/sysl-catalog/pkg/catalog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func Enable(s handlerinitialiser.HandlerInitialiser, r chi.Router, specification string) {
	module := []byte(specification)
	m := afero.NewMemMapFs()
	c := catalog.NewProjectFromJson("service", "http://plantuml.com/plantuml", "html", logrus.New(), module, m, "/-/catalog")
	c.Templates = nil
	c.WithTemplateString(NewPackageTemplate)
	c.ProjectTitle = s.Name()
	c.StartTemplateIndex = 0
	c.Run()
	httpfs := afero.NewHttpFs(m)
	fileserver := http.FileServer(httpfs.Dir("/"))
	r.Mount("/-/catalog", fileserver)
}

const NewPackageTemplate = `
{{$RootModule := .RootModule}}
{{with index (.ModuleAsPackages .RootModule) .ProjectTitle }}
{{/* Automatically generated by https://github.com/anz-bank/sysl-catalog it is strongly recommended not to edit this file */}}
{{$packageName := ModulePackageName .}}

# {{$packageName}}

## Integration Diagram
![](/-/catalog//catalog/{{CreateIntegrationDiagram $RootModule $packageName false}})
{{$Apps := .Apps}}

{{$databases := false}}
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if and (eq (hasPattern $app.Attrs "ignore") false) (eq (hasPattern $app.Attrs "db") true)}}
{{$databases = true}}
{{end}}{{end}}

{{if $databases}}
## Database Index
| Database Application Name  | Source Location |
----|----{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if and (eq (hasPattern $app.Attrs "ignore") false) (eq (hasPattern $app.Attrs "db") true)}}
[{{$appName}}](/-/catalog/#Database-{{$appName}}) | [{{SourcePath $app}}](/-/catalog/{{SourcePath $app}})|  {{end}}{{end}}
{{end}}

## Application Index
{{$anyApps := false}}

{{$Apps := .Apps}}{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if eq (hasPattern $app.Attrs "ignore") false}}{{$Endpoints := $app.Endpoints}}{{range $endpointName := SortedKeys $Endpoints}}{{$endpoint := index $Endpoints $endpointName}}{{if eq (hasPattern $endpoint.Attrs "ignore") false}}{{if not $anyApps}}| Application Name | Method | Source Location |
|----|----|----|{{$anyApps = true}}{{end}}
| {{$appName}} | [{{$endpoint.Name}}](/-/catalog/#{{$appName}}-{{SanitiseOutputName $endpoint.Name}}) | [{{SourcePath $app}}](/-/catalog/{{SourcePath $app}})|  {{end}}{{end}}{{end}}{{end}}

{{if not $anyApps}}
<span style="color:grey">No Applications Defined</span>
{{end}}


## Type Index
{{$anyTypes := false}}

{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{$types := $app.Types}}{{if ne (hasPattern $app.Attrs "db") true}}{{range $typeName := SortedKeys $types}}{{$type := index $types $typeName}}{{if not $anyTypes}}| Application Name | Type Name | Source Location |
|----|----|----|{{$anyTypes = true}}{{end}}
| {{$appName}} | [{{$typeName}}](/-/catalog/#{{$appName}}.{{$typeName}}) | [{{SourcePath $type}}](/-/catalog/{{SourcePath $type}})|{{end}}{{end}}{{end}}

{{if not $anyTypes}}
<span style="color:grey">No Types Defined</span>
{{end}}


{{if $databases}}
# Databases
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}
{{if hasPattern $app.GetAttrs "db"}}

<a name=Database-{{$appName}}></a><details>
<summary>Database {{$appName}}</summary>

{{Attribute $app "description"}}
![](/-/catalog/{{GenerateDataModel $app}})
</details>
{{end}}{{end}}
{{end}}


{{if $anyApps}}
# Applications
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}
{{if eq (hasPattern $app.Attrs "ignore") false}}
{{if eq (hasPattern $app.Attrs "db") false}}
{{if ne (len $app.Endpoints) 0}}

## Application {{$appName}}

{{$desc := Attribute $app "description"}}
{{if $desc}}
- {{$desc}}
{{end}}

{{ServiceMetadata $app}}

{{with CreateRedoc $app.SourceContext $appName}}
[View OpenAPI Specs in Redoc](/-/catalog/{{CreateRedoc $app.SourceContext $appName}})
{{end}}

{{range $e := $app.Endpoints}}
{{if eq (hasPattern $e.Attrs "ignore") false}}


### <a name={{$appName}}-{{SanitiseOutputName $e.Name}}></a>{{$appName}} {{$e.Name}}
{{Attribute $e "description"}}

<details>
<summary>Sequence Diagram</summary>

![](/-/catalog/{{CreateSequenceDiagram $appName $e}})
</details>

<details>
<summary>Request types</summary>

{{if and (not $e.Param) (not $e.RestParams) }}
<span style="color:grey">No Request types</span>
{{end}}
{{if not $e.Param}}{{if $e.RestParams }}{{if not $e.RestParams.UrlParam}}
<span style="color:grey">No Request types</span>
{{end}}{{end}}{{end}}

{{range $param := $e.Param}}
{{Attribute $param.Type "description"}}

![](/-/catalog/{{CreateParamDataModel $app $param}})
{{end}}

{{if $e.RestParams}}{{if $e.RestParams.UrlParam}}
{{range $param := $e.RestParams.UrlParam}}
{{$pathDataModel := (CreateParamDataModel $app $param)}}
{{if ne $pathDataModel ""}}
#### Path Parameter

![](/-/catalog/{{$pathDataModel}})
{{end}}{{end}}{{end}}

{{if $e.RestParams.QueryParam}}
{{range $param := $e.RestParams.QueryParam}}
{{$queryDataModel := (CreateParamDataModel $app $param)}}
{{if ne $queryDataModel ""}}
#### Query Parameter

![](/-/catalog/{{$queryDataModel}})
{{end}}{{end}}{{end}}{{end}}
</details>

<details>
<summary>Response types</summary>

{{$responses := false}}
{{range $s := $e.Stmt}}{{$diagram := CreateReturnDataModel  $appName $s $e}}{{if ne $diagram ""}}
{{$responses = true}}
{{$ret := (GetReturnType $e $s)}}{{if $ret }}
{{Attribute $ret "description"}}{{end}}

![](/-/catalog/{{$diagram}})

{{end}}{{end}}

{{if not $responses}}
<span style="color:grey">No Response Types</span>
{{end}}
</details>
{{end}}

---

{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}


{{if $anyTypes}}
# Types


{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{$types := $app.Types}}
{{if ne (hasPattern $app.Attrs "db") true}}


{{range $typeName := SortedKeys $types}}{{$type := index $types $typeName}}
<a name={{$appName}}.{{$typeName}}></a><details>
<summary>{{$appName}}.{{$typeName}}</summary>

### {{$appName}}.{{$typeName}}
{{$typedesc := (Attribute $type "description")}}
{{if ne $typedesc ""}}- {{$typedesc}}{{end}}

![](/-/catalog/{{CreateTypeDiagram $appName $typeName $type false}})

[Full Diagram](/-/catalog/{{CreateTypeDiagram $appName $typeName $type true}})

{{if Fields $type}}
#### Fields
{{$fieldHeader := false}}
{{$fieldMap := Fields $type}}{{range $fieldName := SortedKeys $fieldMap}}{{$field := index $fieldMap $fieldName}}{{if not $fieldHeader}}| Field name | Type | Description |
|----|----|----|{{$fieldHeader = true}}{{end}}
| {{$fieldName}} | {{FieldType $field}} | {{$desc := Attribute $field "description"}}{{if ne $desc $typedesc}}{{$desc}}{{end}}|{{end}}
{{end}}

</details>{{end}}{{end}}{{end}}
{{end}}

<div class="footer">
{{end}}
`
