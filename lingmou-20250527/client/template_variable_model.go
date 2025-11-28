// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateVariable interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *TemplateVariable
	GetName() *string
	SetProperties(v interface{}) *TemplateVariable
	GetProperties() interface{}
	SetType(v string) *TemplateVariable
	GetType() *string
}

type TemplateVariable struct {
	Name       *string     `json:"name,omitempty" xml:"name,omitempty"`
	Properties interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	Type       *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TemplateVariable) String() string {
	return dara.Prettify(s)
}

func (s TemplateVariable) GoString() string {
	return s.String()
}

func (s *TemplateVariable) GetName() *string {
	return s.Name
}

func (s *TemplateVariable) GetProperties() interface{} {
	return s.Properties
}

func (s *TemplateVariable) GetType() *string {
	return s.Type
}

func (s *TemplateVariable) SetName(v string) *TemplateVariable {
	s.Name = &v
	return s
}

func (s *TemplateVariable) SetProperties(v interface{}) *TemplateVariable {
	s.Properties = v
	return s
}

func (s *TemplateVariable) SetType(v string) *TemplateVariable {
	s.Type = &v
	return s
}

func (s *TemplateVariable) Validate() error {
	return dara.Validate(s)
}
