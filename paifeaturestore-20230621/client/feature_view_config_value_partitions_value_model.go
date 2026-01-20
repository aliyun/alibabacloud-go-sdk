// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeatureViewConfigValuePartitionsValue interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *FeatureViewConfigValuePartitionsValue
	GetValue() *string
	SetValues(v []*string) *FeatureViewConfigValuePartitionsValue
	GetValues() []*string
	SetStartValue(v string) *FeatureViewConfigValuePartitionsValue
	GetStartValue() *string
	SetEndValue(v string) *FeatureViewConfigValuePartitionsValue
	GetEndValue() *string
}

type FeatureViewConfigValuePartitionsValue struct {
	// example:
	//
	// 20250101
	Value  *string   `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	// example:
	//
	// 20250101
	StartValue *string `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	// example:
	//
	// 20250201
	EndValue *string `json:"EndValue,omitempty" xml:"EndValue,omitempty"`
}

func (s FeatureViewConfigValuePartitionsValue) String() string {
	return dara.Prettify(s)
}

func (s FeatureViewConfigValuePartitionsValue) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValuePartitionsValue) GetValue() *string {
	return s.Value
}

func (s *FeatureViewConfigValuePartitionsValue) GetValues() []*string {
	return s.Values
}

func (s *FeatureViewConfigValuePartitionsValue) GetStartValue() *string {
	return s.StartValue
}

func (s *FeatureViewConfigValuePartitionsValue) GetEndValue() *string {
	return s.EndValue
}

func (s *FeatureViewConfigValuePartitionsValue) SetValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.Value = &v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetValues(v []*string) *FeatureViewConfigValuePartitionsValue {
	s.Values = v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetStartValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.StartValue = &v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetEndValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.EndValue = &v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) Validate() error {
	return dara.Validate(s)
}
