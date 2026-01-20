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
	SetSnapshot(v *FeatureViewConfigValueSnapshot) *FeatureViewConfigValue
	GetSnapshot() *FeatureViewConfigValueSnapshot
}

type FeatureViewConfigValue struct {
	Partitions map[string]*FeatureViewConfigValuePartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	// example:
	//
	// 1721186536
	EventTime *string                         `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Equal     *bool                           `json:"Equal,omitempty" xml:"Equal,omitempty"`
	UseMock   *bool                           `json:"UseMock,omitempty" xml:"UseMock,omitempty"`
	Snapshot  *FeatureViewConfigValueSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Struct"`
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

func (s *FeatureViewConfigValue) GetSnapshot() *FeatureViewConfigValueSnapshot {
	return s.Snapshot
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

func (s *FeatureViewConfigValue) SetSnapshot(v *FeatureViewConfigValueSnapshot) *FeatureViewConfigValue {
	s.Snapshot = v
	return s
}

func (s *FeatureViewConfigValue) Validate() error {
	if s.Snapshot != nil {
		if err := s.Snapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FeatureViewConfigValueSnapshot struct {
	Partitions map[string]*FeatureViewConfigValueSnapshotPartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	// example:
	//
	// table_name
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s FeatureViewConfigValueSnapshot) String() string {
	return dara.Prettify(s)
}

func (s FeatureViewConfigValueSnapshot) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValueSnapshot) GetPartitions() map[string]*FeatureViewConfigValueSnapshotPartitionsValue {
	return s.Partitions
}

func (s *FeatureViewConfigValueSnapshot) GetTable() *string {
	return s.Table
}

func (s *FeatureViewConfigValueSnapshot) SetPartitions(v map[string]*FeatureViewConfigValueSnapshotPartitionsValue) *FeatureViewConfigValueSnapshot {
	s.Partitions = v
	return s
}

func (s *FeatureViewConfigValueSnapshot) SetTable(v string) *FeatureViewConfigValueSnapshot {
	s.Table = &v
	return s
}

func (s *FeatureViewConfigValueSnapshot) Validate() error {
	return dara.Validate(s)
}
