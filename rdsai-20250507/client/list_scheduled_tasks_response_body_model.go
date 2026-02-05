// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListScheduledTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListScheduledTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListScheduledTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListScheduledTasksResponseBody
	GetRequestId() *string
	SetSchedules(v []*ListScheduledTasksResponseBodySchedules) *ListScheduledTasksResponseBody
	GetSchedules() []*ListScheduledTasksResponseBodySchedules
	SetSuccess(v bool) *ListScheduledTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListScheduledTasksResponseBody
	GetTotalCount() *int64
}

type ListScheduledTasksResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schedules []*ListScheduledTasksResponseBodySchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduledTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduledTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListScheduledTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledTasksResponseBody) GetSchedules() []*ListScheduledTasksResponseBodySchedules {
	return s.Schedules
}

func (s *ListScheduledTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduledTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListScheduledTasksResponseBody) SetMessage(v string) *ListScheduledTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetPageNumber(v int64) *ListScheduledTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetPageSize(v int64) *ListScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetRequestId(v string) *ListScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetSchedules(v []*ListScheduledTasksResponseBodySchedules) *ListScheduledTasksResponseBody {
	s.Schedules = v
	return s
}

func (s *ListScheduledTasksResponseBody) SetSuccess(v bool) *ListScheduledTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetTotalCount(v int64) *ListScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduledTasksResponseBody) Validate() error {
	if s.Schedules != nil {
		for _, item := range s.Schedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScheduledTasksResponseBodySchedules struct {
	// example:
	//
	// 2026-02-04T06:51:24Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Monday
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 9d246af2-a0cd-4f69-857d-3785048f****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// example:
	//
	// 18:00:00Z
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	TimeRange     *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ListScheduledTasksResponseBodySchedules) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBodySchedules) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBodySchedules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListScheduledTasksResponseBodySchedules) GetDescription() *string {
	return s.Description
}

func (s *ListScheduledTasksResponseBodySchedules) GetFrequency() *string {
	return s.Frequency
}

func (s *ListScheduledTasksResponseBodySchedules) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListScheduledTasksResponseBodySchedules) GetName() *string {
	return s.Name
}

func (s *ListScheduledTasksResponseBodySchedules) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ListScheduledTasksResponseBodySchedules) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *ListScheduledTasksResponseBodySchedules) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ListScheduledTasksResponseBodySchedules) SetCreateTime(v string) *ListScheduledTasksResponseBodySchedules {
	s.CreateTime = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetDescription(v string) *ListScheduledTasksResponseBodySchedules {
	s.Description = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetFrequency(v string) *ListScheduledTasksResponseBodySchedules {
	s.Frequency = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetInstanceCount(v int64) *ListScheduledTasksResponseBodySchedules {
	s.InstanceCount = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetName(v string) *ListScheduledTasksResponseBodySchedules {
	s.Name = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetScheduledId(v string) *ListScheduledTasksResponseBodySchedules {
	s.ScheduledId = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetTaskStartTime(v string) *ListScheduledTasksResponseBodySchedules {
	s.TaskStartTime = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) SetTimeRange(v string) *ListScheduledTasksResponseBodySchedules {
	s.TimeRange = &v
	return s
}

func (s *ListScheduledTasksResponseBodySchedules) Validate() error {
	return dara.Validate(s)
}
