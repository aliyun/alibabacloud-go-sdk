// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHyperParameterDefinition interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *HyperParameterDefinition
	GetDefaultValue() *string
	SetDescription(v string) *HyperParameterDefinition
	GetDescription() *string
	SetDisplayName(v string) *HyperParameterDefinition
	GetDisplayName() *string
	SetName(v string) *HyperParameterDefinition
	GetName() *string
	SetRange(v *HyperParameterRange) *HyperParameterDefinition
	GetRange() *HyperParameterRange
	SetRequired(v bool) *HyperParameterDefinition
	GetRequired() *bool
	SetType(v string) *HyperParameterDefinition
	GetType() *string
}

type HyperParameterDefinition struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	Name     *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	Range    *HyperParameterRange `json:"Range,omitempty" xml:"Range,omitempty"`
	Required *bool                `json:"Required,omitempty" xml:"Required,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s HyperParameterDefinition) String() string {
	return dara.Prettify(s)
}

func (s HyperParameterDefinition) GoString() string {
	return s.String()
}

func (s *HyperParameterDefinition) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *HyperParameterDefinition) GetDescription() *string {
	return s.Description
}

func (s *HyperParameterDefinition) GetDisplayName() *string {
	return s.DisplayName
}

func (s *HyperParameterDefinition) GetName() *string {
	return s.Name
}

func (s *HyperParameterDefinition) GetRange() *HyperParameterRange {
	return s.Range
}

func (s *HyperParameterDefinition) GetRequired() *bool {
	return s.Required
}

func (s *HyperParameterDefinition) GetType() *string {
	return s.Type
}

func (s *HyperParameterDefinition) SetDefaultValue(v string) *HyperParameterDefinition {
	s.DefaultValue = &v
	return s
}

func (s *HyperParameterDefinition) SetDescription(v string) *HyperParameterDefinition {
	s.Description = &v
	return s
}

func (s *HyperParameterDefinition) SetDisplayName(v string) *HyperParameterDefinition {
	s.DisplayName = &v
	return s
}

func (s *HyperParameterDefinition) SetName(v string) *HyperParameterDefinition {
	s.Name = &v
	return s
}

func (s *HyperParameterDefinition) SetRange(v *HyperParameterRange) *HyperParameterDefinition {
	s.Range = v
	return s
}

func (s *HyperParameterDefinition) SetRequired(v bool) *HyperParameterDefinition {
	s.Required = &v
	return s
}

func (s *HyperParameterDefinition) SetType(v string) *HyperParameterDefinition {
	s.Type = &v
	return s
}

func (s *HyperParameterDefinition) Validate() error {
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}
