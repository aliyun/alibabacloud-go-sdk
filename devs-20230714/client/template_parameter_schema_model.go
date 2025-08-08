// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateParameterSchema interface {
	dara.Model
	String() string
	GoString() string
	SetDefault(v interface{}) *TemplateParameterSchema
	GetDefault() interface{}
	SetDescription(v string) *TemplateParameterSchema
	GetDescription() *string
	SetEnum(v []*string) *TemplateParameterSchema
	GetEnum() []*string
	SetPattern(v string) *TemplateParameterSchema
	GetPattern() *string
	SetRequired(v bool) *TemplateParameterSchema
	GetRequired() *bool
	SetRoleExtension(v *TemplateParameterSchemaRoleExtension) *TemplateParameterSchema
	GetRoleExtension() *TemplateParameterSchemaRoleExtension
	SetSensitive(v bool) *TemplateParameterSchema
	GetSensitive() *bool
	SetTitle(v string) *TemplateParameterSchema
	GetTitle() *string
	SetType(v string) *TemplateParameterSchema
	GetType() *string
}

type TemplateParameterSchema struct {
	// example:
	//
	// defaultValue
	Default interface{} `json:"default,omitempty" xml:"default,omitempty"`
	// example:
	//
	// Parameters for testing
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Enum        []*string `json:"enum,omitempty" xml:"enum,omitempty" type:"Repeated"`
	// example:
	//
	// "^[a-zA-Z._-]+$"
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// example:
	//
	// true
	Required      *bool                                 `json:"required,omitempty" xml:"required,omitempty"`
	RoleExtension *TemplateParameterSchemaRoleExtension `json:"roleExtension,omitempty" xml:"roleExtension,omitempty" type:"Struct"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TemplateParameterSchema) String() string {
	return dara.Prettify(s)
}

func (s TemplateParameterSchema) GoString() string {
	return s.String()
}

func (s *TemplateParameterSchema) GetDefault() interface{} {
	return s.Default
}

func (s *TemplateParameterSchema) GetDescription() *string {
	return s.Description
}

func (s *TemplateParameterSchema) GetEnum() []*string {
	return s.Enum
}

func (s *TemplateParameterSchema) GetPattern() *string {
	return s.Pattern
}

func (s *TemplateParameterSchema) GetRequired() *bool {
	return s.Required
}

func (s *TemplateParameterSchema) GetRoleExtension() *TemplateParameterSchemaRoleExtension {
	return s.RoleExtension
}

func (s *TemplateParameterSchema) GetSensitive() *bool {
	return s.Sensitive
}

func (s *TemplateParameterSchema) GetTitle() *string {
	return s.Title
}

func (s *TemplateParameterSchema) GetType() *string {
	return s.Type
}

func (s *TemplateParameterSchema) SetDefault(v interface{}) *TemplateParameterSchema {
	s.Default = v
	return s
}

func (s *TemplateParameterSchema) SetDescription(v string) *TemplateParameterSchema {
	s.Description = &v
	return s
}

func (s *TemplateParameterSchema) SetEnum(v []*string) *TemplateParameterSchema {
	s.Enum = v
	return s
}

func (s *TemplateParameterSchema) SetPattern(v string) *TemplateParameterSchema {
	s.Pattern = &v
	return s
}

func (s *TemplateParameterSchema) SetRequired(v bool) *TemplateParameterSchema {
	s.Required = &v
	return s
}

func (s *TemplateParameterSchema) SetRoleExtension(v *TemplateParameterSchemaRoleExtension) *TemplateParameterSchema {
	s.RoleExtension = v
	return s
}

func (s *TemplateParameterSchema) SetSensitive(v bool) *TemplateParameterSchema {
	s.Sensitive = &v
	return s
}

func (s *TemplateParameterSchema) SetTitle(v string) *TemplateParameterSchema {
	s.Title = &v
	return s
}

func (s *TemplateParameterSchema) SetType(v string) *TemplateParameterSchema {
	s.Type = &v
	return s
}

func (s *TemplateParameterSchema) Validate() error {
	return dara.Validate(s)
}

type TemplateParameterSchemaRoleExtension struct {
	Authorities []*string `json:"authorities,omitempty" xml:"authorities,omitempty" type:"Repeated"`
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Service     *string   `json:"service,omitempty" xml:"service,omitempty"`
}

func (s TemplateParameterSchemaRoleExtension) String() string {
	return dara.Prettify(s)
}

func (s TemplateParameterSchemaRoleExtension) GoString() string {
	return s.String()
}

func (s *TemplateParameterSchemaRoleExtension) GetAuthorities() []*string {
	return s.Authorities
}

func (s *TemplateParameterSchemaRoleExtension) GetName() *string {
	return s.Name
}

func (s *TemplateParameterSchemaRoleExtension) GetService() *string {
	return s.Service
}

func (s *TemplateParameterSchemaRoleExtension) SetAuthorities(v []*string) *TemplateParameterSchemaRoleExtension {
	s.Authorities = v
	return s
}

func (s *TemplateParameterSchemaRoleExtension) SetName(v string) *TemplateParameterSchemaRoleExtension {
	s.Name = &v
	return s
}

func (s *TemplateParameterSchemaRoleExtension) SetService(v string) *TemplateParameterSchemaRoleExtension {
	s.Service = &v
	return s
}

func (s *TemplateParameterSchemaRoleExtension) Validate() error {
	return dara.Validate(s)
}
