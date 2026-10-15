// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectedComponentConfigDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeConfigs(v []*SelectedComponentConfigDTOAttributeConfigs) *SelectedComponentConfigDTO
	GetAttributeConfigs() []*SelectedComponentConfigDTOAttributeConfigs
	SetCode(v string) *SelectedComponentConfigDTO
	GetCode() *string
	SetTag(v string) *SelectedComponentConfigDTO
	GetTag() *string
}

type SelectedComponentConfigDTO struct {
	// The property configuration list.
	AttributeConfigs []*SelectedComponentConfigDTOAttributeConfigs `json:"AttributeConfigs,omitempty" xml:"AttributeConfigs,omitempty" type:"Repeated"`
	// The component code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique identifier of the component, used in multi-node and multi-disk scenarios.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s SelectedComponentConfigDTO) String() string {
	return dara.Prettify(s)
}

func (s SelectedComponentConfigDTO) GoString() string {
	return s.String()
}

func (s *SelectedComponentConfigDTO) GetAttributeConfigs() []*SelectedComponentConfigDTOAttributeConfigs {
	return s.AttributeConfigs
}

func (s *SelectedComponentConfigDTO) GetCode() *string {
	return s.Code
}

func (s *SelectedComponentConfigDTO) GetTag() *string {
	return s.Tag
}

func (s *SelectedComponentConfigDTO) SetAttributeConfigs(v []*SelectedComponentConfigDTOAttributeConfigs) *SelectedComponentConfigDTO {
	s.AttributeConfigs = v
	return s
}

func (s *SelectedComponentConfigDTO) SetCode(v string) *SelectedComponentConfigDTO {
	s.Code = &v
	return s
}

func (s *SelectedComponentConfigDTO) SetTag(v string) *SelectedComponentConfigDTO {
	s.Tag = &v
	return s
}

func (s *SelectedComponentConfigDTO) Validate() error {
	if s.AttributeConfigs != nil {
		for _, item := range s.AttributeConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SelectedComponentConfigDTOAttributeConfigs struct {
	// The upper and lower bound values, used for bounded properties.
	BoundedValue *SelectedComponentConfigDTOAttributeConfigsBoundedValue `json:"BoundedValue,omitempty" xml:"BoundedValue,omitempty" type:"Struct"`
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SelectedComponentConfigDTOAttributeConfigs) String() string {
	return dara.Prettify(s)
}

func (s SelectedComponentConfigDTOAttributeConfigs) GoString() string {
	return s.String()
}

func (s *SelectedComponentConfigDTOAttributeConfigs) GetBoundedValue() *SelectedComponentConfigDTOAttributeConfigsBoundedValue {
	return s.BoundedValue
}

func (s *SelectedComponentConfigDTOAttributeConfigs) GetCode() *string {
	return s.Code
}

func (s *SelectedComponentConfigDTOAttributeConfigs) GetValue() *string {
	return s.Value
}

func (s *SelectedComponentConfigDTOAttributeConfigs) SetBoundedValue(v *SelectedComponentConfigDTOAttributeConfigsBoundedValue) *SelectedComponentConfigDTOAttributeConfigs {
	s.BoundedValue = v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigs) SetCode(v string) *SelectedComponentConfigDTOAttributeConfigs {
	s.Code = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigs) SetValue(v string) *SelectedComponentConfigDTOAttributeConfigs {
	s.Value = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigs) Validate() error {
	if s.BoundedValue != nil {
		if err := s.BoundedValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SelectedComponentConfigDTOAttributeConfigsBoundedValue struct {
	// The lower bound.
	LowerBound *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound `json:"LowerBound,omitempty" xml:"LowerBound,omitempty" type:"Struct"`
	// The upper bound.
	UpperBound *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound `json:"UpperBound,omitempty" xml:"UpperBound,omitempty" type:"Struct"`
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValue) String() string {
	return dara.Prettify(s)
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValue) GoString() string {
	return s.String()
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValue) GetLowerBound() *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound {
	return s.LowerBound
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValue) GetUpperBound() *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound {
	return s.UpperBound
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValue) SetLowerBound(v *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) *SelectedComponentConfigDTOAttributeConfigsBoundedValue {
	s.LowerBound = v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValue) SetUpperBound(v *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) *SelectedComponentConfigDTOAttributeConfigsBoundedValue {
	s.UpperBound = v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValue) Validate() error {
	if s.LowerBound != nil {
		if err := s.LowerBound.Validate(); err != nil {
			return err
		}
	}
	if s.UpperBound != nil {
		if err := s.UpperBound.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound struct {
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) String() string {
	return dara.Prettify(s)
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) GoString() string {
	return s.String()
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) GetCode() *string {
	return s.Code
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) GetValue() *string {
	return s.Value
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) SetCode(v string) *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound {
	s.Code = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) SetValue(v string) *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound {
	s.Value = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueLowerBound) Validate() error {
	return dara.Validate(s)
}

type SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound struct {
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) String() string {
	return dara.Prettify(s)
}

func (s SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) GoString() string {
	return s.String()
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) GetCode() *string {
	return s.Code
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) GetValue() *string {
	return s.Value
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) SetCode(v string) *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound {
	s.Code = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) SetValue(v string) *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound {
	s.Value = &v
	return s
}

func (s *SelectedComponentConfigDTOAttributeConfigsBoundedValueUpperBound) Validate() error {
	return dara.Validate(s)
}
