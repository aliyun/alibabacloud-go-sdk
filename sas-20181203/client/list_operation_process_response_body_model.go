// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListOperationProcessResponseBodyPageInfo) *ListOperationProcessResponseBody
	GetPageInfo() *ListOperationProcessResponseBodyPageInfo
	SetProcesses(v []*ListOperationProcessResponseBodyProcesses) *ListOperationProcessResponseBody
	GetProcesses() []*ListOperationProcessResponseBodyProcesses
	SetRequestId(v string) *ListOperationProcessResponseBody
	GetRequestId() *string
}

type ListOperationProcessResponseBody struct {
	// The pagination information.
	PageInfo *ListOperationProcessResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The information about the operation tasks.
	Processes []*ListOperationProcessResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOperationProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationProcessResponseBody) GetPageInfo() *ListOperationProcessResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListOperationProcessResponseBody) GetProcesses() []*ListOperationProcessResponseBodyProcesses {
	return s.Processes
}

func (s *ListOperationProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationProcessResponseBody) SetPageInfo(v *ListOperationProcessResponseBodyPageInfo) *ListOperationProcessResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListOperationProcessResponseBody) SetProcesses(v []*ListOperationProcessResponseBodyProcesses) *ListOperationProcessResponseBody {
	s.Processes = v
	return s
}

func (s *ListOperationProcessResponseBody) SetRequestId(v string) *ListOperationProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationProcessResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Processes != nil {
		for _, item := range s.Processes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationProcessResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 263
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationProcessResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListOperationProcessResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListOperationProcessResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationProcessResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationProcessResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOperationProcessResponseBodyPageInfo) SetCount(v int32) *ListOperationProcessResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListOperationProcessResponseBodyPageInfo) SetCurrentPage(v int32) *ListOperationProcessResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationProcessResponseBodyPageInfo) SetPageSize(v int32) *ListOperationProcessResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListOperationProcessResponseBodyPageInfo) SetTotalCount(v int32) *ListOperationProcessResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListOperationProcessResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListOperationProcessResponseBodyProcesses struct {
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1674388824000
	CreateTime           *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DetailTaskReadyCount *int32 `json:"DetailTaskReadyCount,omitempty" xml:"DetailTaskReadyCount,omitempty"`
	DetailTaskTotalCount *int32 `json:"DetailTaskTotalCount,omitempty" xml:"DetailTaskTotalCount,omitempty"`
	// The end time of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1705467559000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of tasks that are complete.
	//
	// example:
	//
	// 197
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The start time of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1705457102000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status code. Valid values:
	//
	// 	- 0: not started.
	//
	// 	- 1: running.
	//
	// 	- 2: complete.
	//
	// 	- 3: times out.
	//
	// example:
	//
	// 1
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The ID of the operation task.
	//
	// example:
	//
	// 3d7a1b68-599f-4e16-9b45-e920a183b***
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
	// The task type. Valid values:
	//
	// 	- CHECK_ALL: full check.
	//
	// 	- CHECK_POLICY: policy-based check for which check items are configured.
	//
	// 	- CHECK_SCHEDULE: scheduled check.
	//
	// 	- CHECK_ITEM: specific check item-based check.
	//
	// 	- CHECK_INSTANCE: specific check item-based check on specific instances.
	//
	// example:
	//
	// CHECK_POLICY
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 337
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationProcessResponseBodyProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *ListOperationProcessResponseBodyProcesses) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListOperationProcessResponseBodyProcesses) GetDetailTaskReadyCount() *int32 {
	return s.DetailTaskReadyCount
}

func (s *ListOperationProcessResponseBodyProcesses) GetDetailTaskTotalCount() *int32 {
	return s.DetailTaskTotalCount
}

func (s *ListOperationProcessResponseBodyProcesses) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationProcessResponseBodyProcesses) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *ListOperationProcessResponseBodyProcesses) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationProcessResponseBodyProcesses) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationProcessResponseBodyProcesses) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOperationProcessResponseBodyProcesses) GetTaskSource() *string {
	return s.TaskSource
}

func (s *ListOperationProcessResponseBodyProcesses) GetTaskType() *string {
	return s.TaskType
}

func (s *ListOperationProcessResponseBodyProcesses) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOperationProcessResponseBodyProcesses) SetCreateTime(v int64) *ListOperationProcessResponseBodyProcesses {
	s.CreateTime = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetDetailTaskReadyCount(v int32) *ListOperationProcessResponseBodyProcesses {
	s.DetailTaskReadyCount = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetDetailTaskTotalCount(v int32) *ListOperationProcessResponseBodyProcesses {
	s.DetailTaskTotalCount = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetEndTime(v int64) *ListOperationProcessResponseBodyProcesses {
	s.EndTime = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetFinishCount(v int32) *ListOperationProcessResponseBodyProcesses {
	s.FinishCount = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetStartTime(v int64) *ListOperationProcessResponseBodyProcesses {
	s.StartTime = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetStatusCode(v int32) *ListOperationProcessResponseBodyProcesses {
	s.StatusCode = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetTaskId(v string) *ListOperationProcessResponseBodyProcesses {
	s.TaskId = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetTaskSource(v string) *ListOperationProcessResponseBodyProcesses {
	s.TaskSource = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetTaskType(v string) *ListOperationProcessResponseBodyProcesses {
	s.TaskType = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) SetTotalCount(v int32) *ListOperationProcessResponseBodyProcesses {
	s.TotalCount = &v
	return s
}

func (s *ListOperationProcessResponseBodyProcesses) Validate() error {
	return dara.Validate(s)
}
