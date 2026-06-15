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
	SetSecondJoinKey(v string) *FeatureViewConfigValue
	GetSecondJoinKey() *string
	SetJoinWithLabel(v bool) *FeatureViewConfigValue
	GetJoinWithLabel() *bool
}

type FeatureViewConfigValue struct {
	// The list of partitions.
	Partitions map[string]*FeatureViewConfigValuePartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	// The event time.
	//
	// example:
	//
	// 1721186536
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// Specifies whether the feature\\"s timestamp must match the event time.
	//
	// example:
	//
	// False
	Equal *bool `json:"Equal,omitempty" xml:"Equal,omitempty"`
	// Specifies whether to use a mock data table.
	//
	// example:
	//
	// True
	UseMock *bool `json:"UseMock,omitempty" xml:"UseMock,omitempty"`
	// The snapshot configuration.
	Snapshot *FeatureViewConfigValueSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Struct"`
	// The second join key.
	//
	// example:
	//
	// request_id
	SecondJoinKey *string `json:"SecondJoinKey,omitempty" xml:"SecondJoinKey,omitempty"`
	JoinWithLabel *bool   `json:"JoinWithLabel,omitempty" xml:"JoinWithLabel,omitempty"`
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

func (s *FeatureViewConfigValue) GetSecondJoinKey() *string {
	return s.SecondJoinKey
}

func (s *FeatureViewConfigValue) GetJoinWithLabel() *bool {
	return s.JoinWithLabel
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

func (s *FeatureViewConfigValue) SetSecondJoinKey(v string) *FeatureViewConfigValue {
	s.SecondJoinKey = &v
	return s
}

func (s *FeatureViewConfigValue) SetJoinWithLabel(v bool) *FeatureViewConfigValue {
	s.JoinWithLabel = &v
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
	// The partition configuration for the snapshot.
	Partitions map[string]*FeatureViewConfigValueSnapshotPartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	// The name of the snapshot table.
	//
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
