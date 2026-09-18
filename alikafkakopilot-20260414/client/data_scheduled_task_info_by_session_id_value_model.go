// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataScheduledTaskInfoBySessionIdValue interface {
	dara.Model
	String() string
	GoString() string
	SetAsOf(v string) *DataScheduledTaskInfoBySessionIdValue
	GetAsOf() *string
	SetEnabledCount(v int64) *DataScheduledTaskInfoBySessionIdValue
	GetEnabledCount() *int64
	SetHasScheduledTask(v bool) *DataScheduledTaskInfoBySessionIdValue
	GetHasScheduledTask() *bool
	SetTaskCount(v int64) *DataScheduledTaskInfoBySessionIdValue
	GetTaskCount() *int64
}

type DataScheduledTaskInfoBySessionIdValue struct {
	// The time when the overview was generated, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	AsOf *string `json:"AsOf,omitempty" xml:"AsOf,omitempty"`
	// The number of associated tasks in the ENABLED status.
	//
	// example:
	//
	// 1
	EnabledCount *int64 `json:"EnabledCount,omitempty" xml:"EnabledCount,omitempty"`
	// Indicates whether the current session has associated scheduled tasks that are in the ENABLED, PAUSED, or NEEDS_AUTH status.
	//
	// example:
	//
	// true
	HasScheduledTask *bool `json:"HasScheduledTask,omitempty" xml:"HasScheduledTask,omitempty"`
	// The total number of associated tasks. Only tasks in the ENABLED, PAUSED, or NEEDS_AUTH status are counted.
	//
	// example:
	//
	// 1
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s DataScheduledTaskInfoBySessionIdValue) String() string {
	return dara.Prettify(s)
}

func (s DataScheduledTaskInfoBySessionIdValue) GoString() string {
	return s.String()
}

func (s *DataScheduledTaskInfoBySessionIdValue) GetAsOf() *string {
	return s.AsOf
}

func (s *DataScheduledTaskInfoBySessionIdValue) GetEnabledCount() *int64 {
	return s.EnabledCount
}

func (s *DataScheduledTaskInfoBySessionIdValue) GetHasScheduledTask() *bool {
	return s.HasScheduledTask
}

func (s *DataScheduledTaskInfoBySessionIdValue) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *DataScheduledTaskInfoBySessionIdValue) SetAsOf(v string) *DataScheduledTaskInfoBySessionIdValue {
	s.AsOf = &v
	return s
}

func (s *DataScheduledTaskInfoBySessionIdValue) SetEnabledCount(v int64) *DataScheduledTaskInfoBySessionIdValue {
	s.EnabledCount = &v
	return s
}

func (s *DataScheduledTaskInfoBySessionIdValue) SetHasScheduledTask(v bool) *DataScheduledTaskInfoBySessionIdValue {
	s.HasScheduledTask = &v
	return s
}

func (s *DataScheduledTaskInfoBySessionIdValue) SetTaskCount(v int64) *DataScheduledTaskInfoBySessionIdValue {
	s.TaskCount = &v
	return s
}

func (s *DataScheduledTaskInfoBySessionIdValue) Validate() error {
	return dara.Validate(s)
}
