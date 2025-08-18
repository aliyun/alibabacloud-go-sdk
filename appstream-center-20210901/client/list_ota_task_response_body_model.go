// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOtaTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListOtaTaskResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOtaTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOtaTaskResponseBody
	GetRequestId() *string
	SetTaskList(v []*ListOtaTaskResponseBodyTaskList) *ListOtaTaskResponseBody
	GetTaskList() []*ListOtaTaskResponseBodyTaskList
	SetTotalCount(v int32) *ListOtaTaskResponseBody
	GetTotalCount() *int32
}

type ListOtaTaskResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The OTA update tasks.
	TaskList []*ListOtaTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of OTA update tasks.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOtaTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOtaTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOtaTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOtaTaskResponseBody) GetTaskList() []*ListOtaTaskResponseBodyTaskList {
	return s.TaskList
}

func (s *ListOtaTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOtaTaskResponseBody) SetPageNumber(v int32) *ListOtaTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetPageSize(v int32) *ListOtaTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetRequestId(v string) *ListOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetTaskList(v []*ListOtaTaskResponseBodyTaskList) *ListOtaTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *ListOtaTaskResponseBody) SetTotalCount(v int32) *ListOtaTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOtaTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOtaTaskResponseBodyTaskList struct {
	// The OTA version.
	//
	// example:
	//
	// 0.0.1-R-20220708.110604
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The task status.
	//
	// Valid values:
	//
	// 	- FAILED
	//
	// 	- RUNNING
	//
	// 	- TERMINATED
	//
	// 	- PART_FINISHED
	//
	// 	- STANDBY
	//
	// 	- FINISHED
	//
	// example:
	//
	// RUNNING
	TaskDisplayStatus *string `json:"TaskDisplayStatus,omitempty" xml:"TaskDisplayStatus,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The start time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s ListOtaTaskResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListOtaTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBodyTaskList) GetOtaVersion() *string {
	return s.OtaVersion
}

func (s *ListOtaTaskResponseBodyTaskList) GetTaskDisplayStatus() *string {
	return s.TaskDisplayStatus
}

func (s *ListOtaTaskResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOtaTaskResponseBodyTaskList) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *ListOtaTaskResponseBodyTaskList) SetOtaVersion(v string) *ListOtaTaskResponseBodyTaskList {
	s.OtaVersion = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskDisplayStatus(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskDisplayStatus = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskId(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskStartTime(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskStartTime = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
