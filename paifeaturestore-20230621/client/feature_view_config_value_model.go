// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeatureViewConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetPartitions(v map[string]*FeatureViewConfigValuePartitionsValue) *FeatureViewConfigValue
	GetPartitions() map[string]*FeatureViewConfigValuePartitionsValue
	SetEventTime(v string) *FeatureViewConfigValue
	GetEventTime() *string
	SetEqual(v bool) *FeatureViewConfigValue
	GetEqual() *bool
	SetUseMock(v bool) *FeatureViewConfigValue
	GetUseMock() *bool
}

type FeatureViewConfigValue struct {
	Partitions map[string]*FeatureViewConfigValuePartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	EventTime  *string                                           `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Equal      *bool                                             `json:"Equal,omitempty" xml:"Equal,omitempty"`
	UseMock    *bool                                             `json:"UseMock,omitempty" xml:"UseMock,omitempty"`
}

func (s FeatureViewConfigValue) String() string {
	return dara.Prettify(s)
}

func (s FeatureViewConfigValue) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValue) GetPartitions() map[string]*FeatureViewConfigValuePartitionsValue {
	return s.Partitions
}

func (s *FeatureViewConfigValue) GetEventTime() *string {
	return s.EventTime
}

func (s *FeatureViewConfigValue) GetEqual() *bool {
	return s.Equal
}

func (s *FeatureViewConfigValue) GetUseMock() *bool {
	return s.UseMock
}

func (s *FeatureViewConfigValue) SetPartitions(v map[string]*FeatureViewConfigValuePartitionsValue) *FeatureViewConfigValue {
	s.Partitions = v
	return s
}

func (s *FeatureViewConfigValue) SetEventTime(v string) *FeatureViewConfigValue {
	s.EventTime = &v
	return s
}

func (s *FeatureViewConfigValue) SetEqual(v bool) *FeatureViewConfigValue {
	s.Equal = &v
	return s
}

func (s *FeatureViewConfigValue) SetUseMock(v bool) *FeatureViewConfigValue {
	s.UseMock = &v
	return s
}

func (s *FeatureViewConfigValue) Validate() error {
	return dara.Validate(s)
}
