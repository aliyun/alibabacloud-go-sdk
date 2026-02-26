// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProperty interface {
	dara.Model
	String() string
	GoString() string
	SetItemsType(v string) *Property
	GetItemsType() *string
	SetName(v string) *Property
	GetName() *string
	SetValue(v string) *Property
	GetValue() *string
	SetValueType(v string) *Property
	GetValueType() *string
}

type Property struct {
	// If you set the ValueType field to array, you must specify the type of the elements within the array. The enumerated values include float, integer, and string.
	//
	// example:
	//
	// float
	ItemsType *string `json:"ItemsType,omitempty" xml:"ItemsType,omitempty"`
	// The property name.
	//
	// example:
	//
	// channels
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value.
	//
	// example:
	//
	// [40, 80, 160, 320]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of the property. Supported enumerated values: float, integer, string, and array.
	//
	// example:
	//
	// array
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s Property) String() string {
	return dara.Prettify(s)
}

func (s Property) GoString() string {
	return s.String()
}

func (s *Property) GetItemsType() *string {
	return s.ItemsType
}

func (s *Property) GetName() *string {
	return s.Name
}

func (s *Property) GetValue() *string {
	return s.Value
}

func (s *Property) GetValueType() *string {
	return s.ValueType
}

func (s *Property) SetItemsType(v string) *Property {
	s.ItemsType = &v
	return s
}

func (s *Property) SetName(v string) *Property {
	s.Name = &v
	return s
}

func (s *Property) SetValue(v string) *Property {
	s.Value = &v
	return s
}

func (s *Property) SetValueType(v string) *Property {
	s.ValueType = &v
	return s
}

func (s *Property) Validate() error {
	return dara.Validate(s)
}
