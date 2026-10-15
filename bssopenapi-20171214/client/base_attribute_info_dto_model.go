// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseAttributeInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeCategories(v []*string) *BaseAttributeInfoDTO
	GetAttributeCategories() []*string
	SetAttributeSource(v string) *BaseAttributeInfoDTO
	GetAttributeSource() *string
	SetAttributeType(v string) *BaseAttributeInfoDTO
	GetAttributeType() *string
	SetCode(v string) *BaseAttributeInfoDTO
	GetCode() *string
	SetDescription(v string) *BaseAttributeInfoDTO
	GetDescription() *string
	SetName(v string) *BaseAttributeInfoDTO
	GetName() *string
	SetOperationalAttributeType(v string) *BaseAttributeInfoDTO
	GetOperationalAttributeType() *string
	SetUnit(v string) *BaseAttributeInfoDTO
	GetUnit() *string
	SetValueType(v string) *BaseAttributeInfoDTO
	GetValueType() *string
	SetValues(v []*AttributeValueInfoDTO) *BaseAttributeInfoDTO
	GetValues() []*AttributeValueInfoDTO
}

type BaseAttributeInfoDTO struct {
	// The property categories.
	AttributeCategories []*string `json:"AttributeCategories,omitempty" xml:"AttributeCategories,omitempty" type:"Repeated"`
	// The property source.
	AttributeSource *string `json:"AttributeSource,omitempty" xml:"AttributeSource,omitempty"`
	// The property type.
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property description (translated).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The property name (translated).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The O&M property type.
	OperationalAttributeType *string `json:"OperationalAttributeType,omitempty" xml:"OperationalAttributeType,omitempty"`
	// The unit.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value type.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The property value list.
	Values []*AttributeValueInfoDTO `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s BaseAttributeInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s BaseAttributeInfoDTO) GoString() string {
	return s.String()
}

func (s *BaseAttributeInfoDTO) GetAttributeCategories() []*string {
	return s.AttributeCategories
}

func (s *BaseAttributeInfoDTO) GetAttributeSource() *string {
	return s.AttributeSource
}

func (s *BaseAttributeInfoDTO) GetAttributeType() *string {
	return s.AttributeType
}

func (s *BaseAttributeInfoDTO) GetCode() *string {
	return s.Code
}

func (s *BaseAttributeInfoDTO) GetDescription() *string {
	return s.Description
}

func (s *BaseAttributeInfoDTO) GetName() *string {
	return s.Name
}

func (s *BaseAttributeInfoDTO) GetOperationalAttributeType() *string {
	return s.OperationalAttributeType
}

func (s *BaseAttributeInfoDTO) GetUnit() *string {
	return s.Unit
}

func (s *BaseAttributeInfoDTO) GetValueType() *string {
	return s.ValueType
}

func (s *BaseAttributeInfoDTO) GetValues() []*AttributeValueInfoDTO {
	return s.Values
}

func (s *BaseAttributeInfoDTO) SetAttributeCategories(v []*string) *BaseAttributeInfoDTO {
	s.AttributeCategories = v
	return s
}

func (s *BaseAttributeInfoDTO) SetAttributeSource(v string) *BaseAttributeInfoDTO {
	s.AttributeSource = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetAttributeType(v string) *BaseAttributeInfoDTO {
	s.AttributeType = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetCode(v string) *BaseAttributeInfoDTO {
	s.Code = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetDescription(v string) *BaseAttributeInfoDTO {
	s.Description = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetName(v string) *BaseAttributeInfoDTO {
	s.Name = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetOperationalAttributeType(v string) *BaseAttributeInfoDTO {
	s.OperationalAttributeType = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetUnit(v string) *BaseAttributeInfoDTO {
	s.Unit = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetValueType(v string) *BaseAttributeInfoDTO {
	s.ValueType = &v
	return s
}

func (s *BaseAttributeInfoDTO) SetValues(v []*AttributeValueInfoDTO) *BaseAttributeInfoDTO {
	s.Values = v
	return s
}

func (s *BaseAttributeInfoDTO) Validate() error {
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
