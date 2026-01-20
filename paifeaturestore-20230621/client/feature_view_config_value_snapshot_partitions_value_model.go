// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeatureViewConfigValueSnapshotPartitionsValue interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue
	GetValue() *string
	SetValues(v []*string) *FeatureViewConfigValueSnapshotPartitionsValue
	GetValues() []*string
	SetStartValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue
	GetStartValue() *string
	SetEndValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue
	GetEndValue() *string
}

type FeatureViewConfigValueSnapshotPartitionsValue struct {
	// example:
	//
	// 20260101
	Value      *string   `json:"Value,omitempty" xml:"Value,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	StartValue *string   `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	EndValue   *string   `json:"EndValue,omitempty" xml:"EndValue,omitempty"`
}

func (s FeatureViewConfigValueSnapshotPartitionsValue) String() string {
	return dara.Prettify(s)
}

func (s FeatureViewConfigValueSnapshotPartitionsValue) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) GetValue() *string {
	return s.Value
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) GetValues() []*string {
	return s.Values
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) GetStartValue() *string {
	return s.StartValue
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) GetEndValue() *string {
	return s.EndValue
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) SetValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue {
	s.Value = &v
	return s
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) SetValues(v []*string) *FeatureViewConfigValueSnapshotPartitionsValue {
	s.Values = v
	return s
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) SetStartValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue {
	s.StartValue = &v
	return s
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) SetEndValue(v string) *FeatureViewConfigValueSnapshotPartitionsValue {
	s.EndValue = &v
	return s
}

func (s *FeatureViewConfigValueSnapshotPartitionsValue) Validate() error {
	return dara.Validate(s)
}
