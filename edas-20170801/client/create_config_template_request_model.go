// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateConfigTemplateRequest
	GetContent() *string
	SetDescription(v string) *CreateConfigTemplateRequest
	GetDescription() *string
	SetFormat(v string) *CreateConfigTemplateRequest
	GetFormat() *string
	SetName(v string) *CreateConfigTemplateRequest
	GetName() *string
}

type CreateConfigTemplateRequest struct {
	// The content of the configuration template. The value must be in the format that is specified by the Format parameter.
	//
	// example:
	//
	// [{"Key":"name","Value":"william"},{"Key":"age","Value":"12"}]
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the configuration template. The description can be up to 255 characters in length.
	//
	// example:
	//
	// My configuration template
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
	// The name of the configuration template. The name can be up to 64 characters in length.
	//
	// example:
	//
	// my-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateConfigTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreateConfigTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigTemplateRequest) GetFormat() *string {
	return s.Format
}

func (s *CreateConfigTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateConfigTemplateRequest) SetContent(v string) *CreateConfigTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetDescription(v string) *CreateConfigTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetFormat(v string) *CreateConfigTemplateRequest {
	s.Format = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetName(v string) *CreateConfigTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateConfigTemplateRequest) Validate() error {
	return dara.Validate(s)
}
