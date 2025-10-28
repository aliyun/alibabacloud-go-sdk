// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateConfigTemplateRequest
	GetContent() *string
	SetDescription(v string) *UpdateConfigTemplateRequest
	GetDescription() *string
	SetFormat(v string) *UpdateConfigTemplateRequest
	GetFormat() *string
	SetId(v int64) *UpdateConfigTemplateRequest
	GetId() *int64
	SetName(v string) *UpdateConfigTemplateRequest
	GetName() *string
}

type UpdateConfigTemplateRequest struct {
	// The content of the configuration template. The value must be in the format that is specified by the Format parameter.
	//
	// example:
	//
	// {"name":"william","age":18}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the configuration template. The description can be up to 255 characters in length.
	//
	// example:
	//
	// Test configuration template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data format of the configuration template. Valid values:
	//
	// 	- JSON: JSON format
	//
	// 	- XML: XML format
	//
	// 	- YAML: YAML format
	//
	// 	- Properties: .properties format
	//
	// 	- KeyValue: key-value pairs
	//
	// 	- Custom: custom format
	//
	// example:
	//
	// JSON
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The ID of the configuration template.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the configuration template. The name can be up to 64 characters in length.
	//
	// example:
	//
	// configtmpl1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateConfigTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateConfigTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConfigTemplateRequest) GetFormat() *string {
	return s.Format
}

func (s *UpdateConfigTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateConfigTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateConfigTemplateRequest) SetContent(v string) *UpdateConfigTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetDescription(v string) *UpdateConfigTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetFormat(v string) *UpdateConfigTemplateRequest {
	s.Format = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetId(v int64) *UpdateConfigTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetName(v string) *UpdateConfigTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateConfigTemplateRequest) Validate() error {
	return dara.Validate(s)
}
