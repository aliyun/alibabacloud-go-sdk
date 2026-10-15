// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttributeDetailDTO interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttributeDetailDTO
	GetCode() *string
	SetDefaultValue(v string) *AttributeDetailDTO
	GetDefaultValue() *string
	SetStatus(v string) *AttributeDetailDTO
	GetStatus() *string
	SetValues(v []*AttributeDetailDTOValues) *AttributeDetailDTO
	GetValues() []*AttributeDetailDTOValues
}

type AttributeDetailDTO struct {
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default value.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The property status.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of valid values for the property.
	Values []*AttributeDetailDTOValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s AttributeDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s AttributeDetailDTO) GoString() string {
	return s.String()
}

func (s *AttributeDetailDTO) GetCode() *string {
	return s.Code
}

func (s *AttributeDetailDTO) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *AttributeDetailDTO) GetStatus() *string {
	return s.Status
}

func (s *AttributeDetailDTO) GetValues() []*AttributeDetailDTOValues {
	return s.Values
}

func (s *AttributeDetailDTO) SetCode(v string) *AttributeDetailDTO {
	s.Code = &v
	return s
}

func (s *AttributeDetailDTO) SetDefaultValue(v string) *AttributeDetailDTO {
	s.DefaultValue = &v
	return s
}

func (s *AttributeDetailDTO) SetStatus(v string) *AttributeDetailDTO {
	s.Status = &v
	return s
}

func (s *AttributeDetailDTO) SetValues(v []*AttributeDetailDTOValues) *AttributeDetailDTO {
	s.Values = v
	return s
}

func (s *AttributeDetailDTO) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttributeDetailDTOValues struct {
	// The list of cascading property values (recursive structure).
	CascadingAttributeValues []*AttributeDetailDTO `json:"CascadingAttributeValues,omitempty" xml:"CascadingAttributeValues,omitempty" type:"Repeated"`
	// The property value, which corresponds to the identifyKey in the property value definition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AttributeDetailDTOValues) String() string {
	return dara.Prettify(s)
}

func (s AttributeDetailDTOValues) GoString() string {
	return s.String()
}

func (s *AttributeDetailDTOValues) GetCascadingAttributeValues() []*AttributeDetailDTO {
	return s.CascadingAttributeValues
}

func (s *AttributeDetailDTOValues) GetValue() *string {
	return s.Value
}

func (s *AttributeDetailDTOValues) SetCascadingAttributeValues(v []*AttributeDetailDTO) *AttributeDetailDTOValues {
	s.CascadingAttributeValues = v
	return s
}

func (s *AttributeDetailDTOValues) SetValue(v string) *AttributeDetailDTOValues {
	s.Value = &v
	return s
}

func (s *AttributeDetailDTOValues) Validate() error {
	if s.CascadingAttributeValues != nil {
		for _, item := range s.CascadingAttributeValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
