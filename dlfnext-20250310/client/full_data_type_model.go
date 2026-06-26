// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFullDataType interface {
	dara.Model
	String() string
	GoString() string
	SetElement(v *FullDataType) *FullDataType
	GetElement() *FullDataType
	SetFields(v []*DataField) *FullDataType
	GetFields() []*DataField
	SetKey(v *FullDataType) *FullDataType
	GetKey() *FullDataType
	SetType(v string) *FullDataType
	GetType() *string
	SetValue(v *FullDataType) *FullDataType
	GetValue() *FullDataType
}

type FullDataType struct {
	// The data type.
	Element *FullDataType `json:"element,omitempty" xml:"element,omitempty"`
	// The list of field definitions. This parameter is required when `type` is `ROW`.
	Fields []*DataField `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The data type of the key. This parameter is required when `type` is `MAP`.
	Key *FullDataType `json:"key,omitempty" xml:"key,omitempty"`
	// The data type. This parameter is required when `type` is `ARRAY` or `MULTISET`.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The data type of the value. This parameter is required when `type` is `MAP`.
	Value *FullDataType `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FullDataType) String() string {
	return dara.Prettify(s)
}

func (s FullDataType) GoString() string {
	return s.String()
}

func (s *FullDataType) GetElement() *FullDataType {
	return s.Element
}

func (s *FullDataType) GetFields() []*DataField {
	return s.Fields
}

func (s *FullDataType) GetKey() *FullDataType {
	return s.Key
}

func (s *FullDataType) GetType() *string {
	return s.Type
}

func (s *FullDataType) GetValue() *FullDataType {
	return s.Value
}

func (s *FullDataType) SetElement(v *FullDataType) *FullDataType {
	s.Element = v
	return s
}

func (s *FullDataType) SetFields(v []*DataField) *FullDataType {
	s.Fields = v
	return s
}

func (s *FullDataType) SetKey(v *FullDataType) *FullDataType {
	s.Key = v
	return s
}

func (s *FullDataType) SetType(v string) *FullDataType {
	s.Type = &v
	return s
}

func (s *FullDataType) SetValue(v *FullDataType) *FullDataType {
	s.Value = v
	return s
}

func (s *FullDataType) Validate() error {
	if s.Element != nil {
		if err := s.Element.Validate(); err != nil {
			return err
		}
	}
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}
