// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttributeValueInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBoundedValue(v *AttributeValueInfoDTOBoundedValue) *AttributeValueInfoDTO
	GetBoundedValue() *AttributeValueInfoDTOBoundedValue
	SetCode(v string) *AttributeValueInfoDTO
	GetCode() *string
	SetDefaultValue(v string) *AttributeValueInfoDTO
	GetDefaultValue() *string
	SetDescription(v string) *AttributeValueInfoDTO
	GetDescription() *string
	SetDescriptionAttributeValues(v []*AttributeValueInfoDTO) *AttributeValueInfoDTO
	GetDescriptionAttributeValues() []*AttributeValueInfoDTO
	SetIdentifyKey(v string) *AttributeValueInfoDTO
	GetIdentifyKey() *string
	SetMax(v string) *AttributeValueInfoDTO
	GetMax() *string
	SetMin(v string) *AttributeValueInfoDTO
	GetMin() *string
	SetOperationTags(v []*string) *AttributeValueInfoDTO
	GetOperationTags() []*string
	SetStep(v string) *AttributeValueInfoDTO
	GetStep() *string
	SetSubAttributeValues(v map[string][]*AttributeValueInfoDTO) *AttributeValueInfoDTO
	GetSubAttributeValues() map[string][]*AttributeValueInfoDTO
	SetValue(v string) *AttributeValueInfoDTO
	GetValue() *string
	SetValueName(v string) *AttributeValueInfoDTO
	GetValueName() *string
}

type AttributeValueInfoDTO struct {
	// The upper and lower bounds.
	BoundedValue *AttributeValueInfoDTOBoundedValue `json:"BoundedValue,omitempty" xml:"BoundedValue,omitempty" type:"Struct"`
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default value.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description (translated).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description property values.
	DescriptionAttributeValues []*AttributeValueInfoDTO `json:"DescriptionAttributeValues,omitempty" xml:"DescriptionAttributeValues,omitempty" type:"Repeated"`
	// The unique identifier.
	IdentifyKey *string `json:"IdentifyKey,omitempty" xml:"IdentifyKey,omitempty"`
	// The maximum value.
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum value.
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
	// The operation tags.
	OperationTags []*string `json:"OperationTags,omitempty" xml:"OperationTags,omitempty" type:"Repeated"`
	// The step size.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// The sub-property values.
	SubAttributeValues map[string][]*AttributeValueInfoDTO `json:"SubAttributeValues,omitempty" xml:"SubAttributeValues,omitempty"`
	// The enum value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The value name (translated).
	ValueName *string `json:"ValueName,omitempty" xml:"ValueName,omitempty"`
}

func (s AttributeValueInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s AttributeValueInfoDTO) GoString() string {
	return s.String()
}

func (s *AttributeValueInfoDTO) GetBoundedValue() *AttributeValueInfoDTOBoundedValue {
	return s.BoundedValue
}

func (s *AttributeValueInfoDTO) GetCode() *string {
	return s.Code
}

func (s *AttributeValueInfoDTO) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *AttributeValueInfoDTO) GetDescription() *string {
	return s.Description
}

func (s *AttributeValueInfoDTO) GetDescriptionAttributeValues() []*AttributeValueInfoDTO {
	return s.DescriptionAttributeValues
}

func (s *AttributeValueInfoDTO) GetIdentifyKey() *string {
	return s.IdentifyKey
}

func (s *AttributeValueInfoDTO) GetMax() *string {
	return s.Max
}

func (s *AttributeValueInfoDTO) GetMin() *string {
	return s.Min
}

func (s *AttributeValueInfoDTO) GetOperationTags() []*string {
	return s.OperationTags
}

func (s *AttributeValueInfoDTO) GetStep() *string {
	return s.Step
}

func (s *AttributeValueInfoDTO) GetSubAttributeValues() map[string][]*AttributeValueInfoDTO {
	return s.SubAttributeValues
}

func (s *AttributeValueInfoDTO) GetValue() *string {
	return s.Value
}

func (s *AttributeValueInfoDTO) GetValueName() *string {
	return s.ValueName
}

func (s *AttributeValueInfoDTO) SetBoundedValue(v *AttributeValueInfoDTOBoundedValue) *AttributeValueInfoDTO {
	s.BoundedValue = v
	return s
}

func (s *AttributeValueInfoDTO) SetCode(v string) *AttributeValueInfoDTO {
	s.Code = &v
	return s
}

func (s *AttributeValueInfoDTO) SetDefaultValue(v string) *AttributeValueInfoDTO {
	s.DefaultValue = &v
	return s
}

func (s *AttributeValueInfoDTO) SetDescription(v string) *AttributeValueInfoDTO {
	s.Description = &v
	return s
}

func (s *AttributeValueInfoDTO) SetDescriptionAttributeValues(v []*AttributeValueInfoDTO) *AttributeValueInfoDTO {
	s.DescriptionAttributeValues = v
	return s
}

func (s *AttributeValueInfoDTO) SetIdentifyKey(v string) *AttributeValueInfoDTO {
	s.IdentifyKey = &v
	return s
}

func (s *AttributeValueInfoDTO) SetMax(v string) *AttributeValueInfoDTO {
	s.Max = &v
	return s
}

func (s *AttributeValueInfoDTO) SetMin(v string) *AttributeValueInfoDTO {
	s.Min = &v
	return s
}

func (s *AttributeValueInfoDTO) SetOperationTags(v []*string) *AttributeValueInfoDTO {
	s.OperationTags = v
	return s
}

func (s *AttributeValueInfoDTO) SetStep(v string) *AttributeValueInfoDTO {
	s.Step = &v
	return s
}

func (s *AttributeValueInfoDTO) SetSubAttributeValues(v map[string][]*AttributeValueInfoDTO) *AttributeValueInfoDTO {
	s.SubAttributeValues = v
	return s
}

func (s *AttributeValueInfoDTO) SetValue(v string) *AttributeValueInfoDTO {
	s.Value = &v
	return s
}

func (s *AttributeValueInfoDTO) SetValueName(v string) *AttributeValueInfoDTO {
	s.ValueName = &v
	return s
}

func (s *AttributeValueInfoDTO) Validate() error {
	if s.BoundedValue != nil {
		if err := s.BoundedValue.Validate(); err != nil {
			return err
		}
	}
	if s.DescriptionAttributeValues != nil {
		for _, item := range s.DescriptionAttributeValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttributeValueInfoDTOBoundedValue struct {
	// The lower bound list.
	LowerBound []*AttributeValueInfoDTO `json:"LowerBound,omitempty" xml:"LowerBound,omitempty" type:"Repeated"`
	// The upper bound list.
	UpperBound []*AttributeValueInfoDTO `json:"UpperBound,omitempty" xml:"UpperBound,omitempty" type:"Repeated"`
}

func (s AttributeValueInfoDTOBoundedValue) String() string {
	return dara.Prettify(s)
}

func (s AttributeValueInfoDTOBoundedValue) GoString() string {
	return s.String()
}

func (s *AttributeValueInfoDTOBoundedValue) GetLowerBound() []*AttributeValueInfoDTO {
	return s.LowerBound
}

func (s *AttributeValueInfoDTOBoundedValue) GetUpperBound() []*AttributeValueInfoDTO {
	return s.UpperBound
}

func (s *AttributeValueInfoDTOBoundedValue) SetLowerBound(v []*AttributeValueInfoDTO) *AttributeValueInfoDTOBoundedValue {
	s.LowerBound = v
	return s
}

func (s *AttributeValueInfoDTOBoundedValue) SetUpperBound(v []*AttributeValueInfoDTO) *AttributeValueInfoDTOBoundedValue {
	s.UpperBound = v
	return s
}

func (s *AttributeValueInfoDTOBoundedValue) Validate() error {
	if s.LowerBound != nil {
		for _, item := range s.LowerBound {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpperBound != nil {
		for _, item := range s.UpperBound {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
