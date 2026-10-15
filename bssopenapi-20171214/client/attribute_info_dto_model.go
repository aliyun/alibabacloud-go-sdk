// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttributeInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeCategories(v []*string) *AttributeInfoDTO
	GetAttributeCategories() []*string
	SetAttributeSource(v string) *AttributeInfoDTO
	GetAttributeSource() *string
	SetAttributeType(v string) *AttributeInfoDTO
	GetAttributeType() *string
	SetCode(v string) *AttributeInfoDTO
	GetCode() *string
	SetCompositeAttributes(v []*BaseAttributeInfoDTO) *AttributeInfoDTO
	GetCompositeAttributes() []*BaseAttributeInfoDTO
	SetDescription(v string) *AttributeInfoDTO
	GetDescription() *string
	SetDescriptionAttributes(v []*BaseAttributeInfoDTO) *AttributeInfoDTO
	GetDescriptionAttributes() []*BaseAttributeInfoDTO
	SetName(v string) *AttributeInfoDTO
	GetName() *string
	SetOperationalAttributeConfig(v string) *AttributeInfoDTO
	GetOperationalAttributeConfig() *string
	SetOperationalAttributeType(v string) *AttributeInfoDTO
	GetOperationalAttributeType() *string
	SetSubAttributes(v []*AttributeInfoDTO) *AttributeInfoDTO
	GetSubAttributes() []*AttributeInfoDTO
	SetUnit(v string) *AttributeInfoDTO
	GetUnit() *string
	SetValueType(v string) *AttributeInfoDTO
	GetValueType() *string
	SetValues(v []*AttributeValueInfoDTO) *AttributeInfoDTO
	GetValues() []*AttributeValueInfoDTO
}

type AttributeInfoDTO struct {
	// The property categories.
	AttributeCategories []*string `json:"AttributeCategories,omitempty" xml:"AttributeCategories,omitempty" type:"Repeated"`
	// The property source.
	AttributeSource *string `json:"AttributeSource,omitempty" xml:"AttributeSource,omitempty"`
	// The property type.
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// The property code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The composite properties (basic information).
	CompositeAttributes []*BaseAttributeInfoDTO `json:"CompositeAttributes,omitempty" xml:"CompositeAttributes,omitempty" type:"Repeated"`
	// The property description (translated).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description properties (basic information).
	DescriptionAttributes []*BaseAttributeInfoDTO `json:"DescriptionAttributes,omitempty" xml:"DescriptionAttributes,omitempty" type:"Repeated"`
	// The property name (translated).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The O&M property configuration.
	OperationalAttributeConfig *string `json:"OperationalAttributeConfig,omitempty" xml:"OperationalAttributeConfig,omitempty"`
	// The O&M property type.
	OperationalAttributeType *string `json:"OperationalAttributeType,omitempty" xml:"OperationalAttributeType,omitempty"`
	// The sub-properties (complete information).
	SubAttributes []*AttributeInfoDTO `json:"SubAttributes,omitempty" xml:"SubAttributes,omitempty" type:"Repeated"`
	// The unit.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value type.
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The property value list.
	Values []*AttributeValueInfoDTO `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s AttributeInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s AttributeInfoDTO) GoString() string {
	return s.String()
}

func (s *AttributeInfoDTO) GetAttributeCategories() []*string {
	return s.AttributeCategories
}

func (s *AttributeInfoDTO) GetAttributeSource() *string {
	return s.AttributeSource
}

func (s *AttributeInfoDTO) GetAttributeType() *string {
	return s.AttributeType
}

func (s *AttributeInfoDTO) GetCode() *string {
	return s.Code
}

func (s *AttributeInfoDTO) GetCompositeAttributes() []*BaseAttributeInfoDTO {
	return s.CompositeAttributes
}

func (s *AttributeInfoDTO) GetDescription() *string {
	return s.Description
}

func (s *AttributeInfoDTO) GetDescriptionAttributes() []*BaseAttributeInfoDTO {
	return s.DescriptionAttributes
}

func (s *AttributeInfoDTO) GetName() *string {
	return s.Name
}

func (s *AttributeInfoDTO) GetOperationalAttributeConfig() *string {
	return s.OperationalAttributeConfig
}

func (s *AttributeInfoDTO) GetOperationalAttributeType() *string {
	return s.OperationalAttributeType
}

func (s *AttributeInfoDTO) GetSubAttributes() []*AttributeInfoDTO {
	return s.SubAttributes
}

func (s *AttributeInfoDTO) GetUnit() *string {
	return s.Unit
}

func (s *AttributeInfoDTO) GetValueType() *string {
	return s.ValueType
}

func (s *AttributeInfoDTO) GetValues() []*AttributeValueInfoDTO {
	return s.Values
}

func (s *AttributeInfoDTO) SetAttributeCategories(v []*string) *AttributeInfoDTO {
	s.AttributeCategories = v
	return s
}

func (s *AttributeInfoDTO) SetAttributeSource(v string) *AttributeInfoDTO {
	s.AttributeSource = &v
	return s
}

func (s *AttributeInfoDTO) SetAttributeType(v string) *AttributeInfoDTO {
	s.AttributeType = &v
	return s
}

func (s *AttributeInfoDTO) SetCode(v string) *AttributeInfoDTO {
	s.Code = &v
	return s
}

func (s *AttributeInfoDTO) SetCompositeAttributes(v []*BaseAttributeInfoDTO) *AttributeInfoDTO {
	s.CompositeAttributes = v
	return s
}

func (s *AttributeInfoDTO) SetDescription(v string) *AttributeInfoDTO {
	s.Description = &v
	return s
}

func (s *AttributeInfoDTO) SetDescriptionAttributes(v []*BaseAttributeInfoDTO) *AttributeInfoDTO {
	s.DescriptionAttributes = v
	return s
}

func (s *AttributeInfoDTO) SetName(v string) *AttributeInfoDTO {
	s.Name = &v
	return s
}

func (s *AttributeInfoDTO) SetOperationalAttributeConfig(v string) *AttributeInfoDTO {
	s.OperationalAttributeConfig = &v
	return s
}

func (s *AttributeInfoDTO) SetOperationalAttributeType(v string) *AttributeInfoDTO {
	s.OperationalAttributeType = &v
	return s
}

func (s *AttributeInfoDTO) SetSubAttributes(v []*AttributeInfoDTO) *AttributeInfoDTO {
	s.SubAttributes = v
	return s
}

func (s *AttributeInfoDTO) SetUnit(v string) *AttributeInfoDTO {
	s.Unit = &v
	return s
}

func (s *AttributeInfoDTO) SetValueType(v string) *AttributeInfoDTO {
	s.ValueType = &v
	return s
}

func (s *AttributeInfoDTO) SetValues(v []*AttributeValueInfoDTO) *AttributeInfoDTO {
	s.Values = v
	return s
}

func (s *AttributeInfoDTO) Validate() error {
	if s.CompositeAttributes != nil {
		for _, item := range s.CompositeAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DescriptionAttributes != nil {
		for _, item := range s.DescriptionAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubAttributes != nil {
		for _, item := range s.SubAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
