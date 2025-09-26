// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolInput interface {
	dara.Model
	String() string
	GoString() string
	SetCAPConfig(v *CAPConfig) *CreateToolInput
	GetCAPConfig() *CAPConfig
	SetDescription(v string) *CreateToolInput
	GetDescription() *string
	SetName(v string) *CreateToolInput
	GetName() *string
	SetSchema(v string) *CreateToolInput
	GetSchema() *string
	SetSourceType(v string) *CreateToolInput
	GetSourceType() *string
	SetToolType(v string) *CreateToolInput
	GetToolType() *string
}

type CreateToolInput struct {
	CAPConfig   *CAPConfig `json:"CAPConfig,omitempty" xml:"CAPConfig,omitempty"`
	Description *string    `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// This parameter is required.
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// This parameter is required.
	ToolType *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
}

func (s CreateToolInput) String() string {
	return dara.Prettify(s)
}

func (s CreateToolInput) GoString() string {
	return s.String()
}

func (s *CreateToolInput) GetCAPConfig() *CAPConfig {
	return s.CAPConfig
}

func (s *CreateToolInput) GetDescription() *string {
	return s.Description
}

func (s *CreateToolInput) GetName() *string {
	return s.Name
}

func (s *CreateToolInput) GetSchema() *string {
	return s.Schema
}

func (s *CreateToolInput) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateToolInput) GetToolType() *string {
	return s.ToolType
}

func (s *CreateToolInput) SetCAPConfig(v *CAPConfig) *CreateToolInput {
	s.CAPConfig = v
	return s
}

func (s *CreateToolInput) SetDescription(v string) *CreateToolInput {
	s.Description = &v
	return s
}

func (s *CreateToolInput) SetName(v string) *CreateToolInput {
	s.Name = &v
	return s
}

func (s *CreateToolInput) SetSchema(v string) *CreateToolInput {
	s.Schema = &v
	return s
}

func (s *CreateToolInput) SetSourceType(v string) *CreateToolInput {
	s.SourceType = &v
	return s
}

func (s *CreateToolInput) SetToolType(v string) *CreateToolInput {
	s.ToolType = &v
	return s
}

func (s *CreateToolInput) Validate() error {
	return dara.Validate(s)
}
