// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataPathConfigsFieldPathsValue interface {
	dara.Model
	String() string
	GoString() string
	SetFieldKey(v string) *DataPathConfigsFieldPathsValue
	GetFieldKey() *string
	SetJsonPath(v string) *DataPathConfigsFieldPathsValue
	GetJsonPath() *string
	SetRule(v string) *DataPathConfigsFieldPathsValue
	GetRule() *string
	SetName(v string) *DataPathConfigsFieldPathsValue
	GetName() *string
	SetDescription(v string) *DataPathConfigsFieldPathsValue
	GetDescription() *string
	SetSource(v string) *DataPathConfigsFieldPathsValue
	GetSource() *string
	SetCategory(v string) *DataPathConfigsFieldPathsValue
	GetCategory() *string
	SetIo(v string) *DataPathConfigsFieldPathsValue
	GetIo() *string
	SetSensitive(v bool) *DataPathConfigsFieldPathsValue
	GetSensitive() *bool
	SetRecordEnabled(v bool) *DataPathConfigsFieldPathsValue
	GetRecordEnabled() *bool
	SetToggleAllowed(v bool) *DataPathConfigsFieldPathsValue
	GetToggleAllowed() *bool
	SetJsonPathRewriteAllowed(v bool) *DataPathConfigsFieldPathsValue
	GetJsonPathRewriteAllowed() *bool
	SetFromTemplate(v bool) *DataPathConfigsFieldPathsValue
	GetFromTemplate() *bool
}

type DataPathConfigsFieldPathsValue struct {
	// The identity key of the field, which is unique within the same request path, such as input_token.
	//
	// example:
	//
	// input_token
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// The JSONPath expression of the field in the request or response message, such as usage.prompt_tokens.
	//
	// example:
	//
	// usage.prompt_tokens
	JsonPath *string `json:"jsonPath,omitempty" xml:"jsonPath,omitempty"`
	// The value extraction rule for the field.
	//
	// example:
	//
	// first
	Rule *string `json:"rule,omitempty" xml:"rule,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// Input token
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The description of the field.
	//
	// example:
	//
	// Input token statistics
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The source of the field value, indicating which part of the message the field is read from, such as response_body.
	//
	// example:
	//
	// response_body
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The statistical category to which the field belongs.
	//
	// example:
	//
	// usage
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The message direction to which the field belongs, indicating whether it is extracted from the request or the response.
	//
	// example:
	//
	// response
	Io *string `json:"io,omitempty" xml:"io,omitempty"`
	// Indicates whether the field contains sensitive information.
	//
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// Indicates whether collection is enabled for the field.
	//
	// example:
	//
	// true
	RecordEnabled *bool `json:"recordEnabled,omitempty" xml:"recordEnabled,omitempty"`
	// Indicates whether the collection status of the field can be toggled.
	//
	// example:
	//
	// true
	ToggleAllowed *bool `json:"toggleAllowed,omitempty" xml:"toggleAllowed,omitempty"`
	// Indicates whether the JSONPath of the field can be rewritten.
	//
	// example:
	//
	// true
	JsonPathRewriteAllowed *bool `json:"jsonPathRewriteAllowed,omitempty" xml:"jsonPathRewriteAllowed,omitempty"`
	// Indicates whether the field configuration is derived from a field template.
	//
	// example:
	//
	// true
	FromTemplate *bool `json:"fromTemplate,omitempty" xml:"fromTemplate,omitempty"`
}

func (s DataPathConfigsFieldPathsValue) String() string {
	return dara.Prettify(s)
}

func (s DataPathConfigsFieldPathsValue) GoString() string {
	return s.String()
}

func (s *DataPathConfigsFieldPathsValue) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DataPathConfigsFieldPathsValue) GetJsonPath() *string {
	return s.JsonPath
}

func (s *DataPathConfigsFieldPathsValue) GetRule() *string {
	return s.Rule
}

func (s *DataPathConfigsFieldPathsValue) GetName() *string {
	return s.Name
}

func (s *DataPathConfigsFieldPathsValue) GetDescription() *string {
	return s.Description
}

func (s *DataPathConfigsFieldPathsValue) GetSource() *string {
	return s.Source
}

func (s *DataPathConfigsFieldPathsValue) GetCategory() *string {
	return s.Category
}

func (s *DataPathConfigsFieldPathsValue) GetIo() *string {
	return s.Io
}

func (s *DataPathConfigsFieldPathsValue) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DataPathConfigsFieldPathsValue) GetRecordEnabled() *bool {
	return s.RecordEnabled
}

func (s *DataPathConfigsFieldPathsValue) GetToggleAllowed() *bool {
	return s.ToggleAllowed
}

func (s *DataPathConfigsFieldPathsValue) GetJsonPathRewriteAllowed() *bool {
	return s.JsonPathRewriteAllowed
}

func (s *DataPathConfigsFieldPathsValue) GetFromTemplate() *bool {
	return s.FromTemplate
}

func (s *DataPathConfigsFieldPathsValue) SetFieldKey(v string) *DataPathConfigsFieldPathsValue {
	s.FieldKey = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetJsonPath(v string) *DataPathConfigsFieldPathsValue {
	s.JsonPath = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetRule(v string) *DataPathConfigsFieldPathsValue {
	s.Rule = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetName(v string) *DataPathConfigsFieldPathsValue {
	s.Name = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetDescription(v string) *DataPathConfigsFieldPathsValue {
	s.Description = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetSource(v string) *DataPathConfigsFieldPathsValue {
	s.Source = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetCategory(v string) *DataPathConfigsFieldPathsValue {
	s.Category = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetIo(v string) *DataPathConfigsFieldPathsValue {
	s.Io = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetSensitive(v bool) *DataPathConfigsFieldPathsValue {
	s.Sensitive = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetRecordEnabled(v bool) *DataPathConfigsFieldPathsValue {
	s.RecordEnabled = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetToggleAllowed(v bool) *DataPathConfigsFieldPathsValue {
	s.ToggleAllowed = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetJsonPathRewriteAllowed(v bool) *DataPathConfigsFieldPathsValue {
	s.JsonPathRewriteAllowed = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) SetFromTemplate(v bool) *DataPathConfigsFieldPathsValue {
	s.FromTemplate = &v
	return s
}

func (s *DataPathConfigsFieldPathsValue) Validate() error {
	return dara.Validate(s)
}
