// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTasks(v []*ListAsyncTaskResponseBodyAsyncTasks) *ListAsyncTaskResponseBody
	GetAsyncTasks() []*ListAsyncTaskResponseBodyAsyncTasks
	SetRequestId(v string) *ListAsyncTaskResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAsyncTaskResponseBody
	GetTotal() *int32
}

type ListAsyncTaskResponseBody struct {
	AsyncTasks []*ListAsyncTaskResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseBody) GetAsyncTasks() []*ListAsyncTaskResponseBodyAsyncTasks {
	return s.AsyncTasks
}

func (s *ListAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsyncTaskResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAsyncTaskResponseBody) SetAsyncTasks(v []*ListAsyncTaskResponseBodyAsyncTasks) *ListAsyncTaskResponseBody {
	s.AsyncTasks = v
	return s
}

func (s *ListAsyncTaskResponseBody) SetRequestId(v string) *ListAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTaskResponseBody) SetTotal(v int32) *ListAsyncTaskResponseBody {
	s.Total = &v
	return s
}

func (s *ListAsyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAsyncTaskResponseBodyAsyncTasks struct {
	// example:
	//
	// 1533866201000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1533866201000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {"instanceId": "ddoscoo-1234-qrq2134"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// {"instanceId": "ddoscoo-1234-qrq2134", "url": "https://oss.xxx.xxx"}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListAsyncTaskResponseBodyAsyncTasks) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTaskResponseBodyAsyncTasks) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetTaskParams() *string {
	return s.TaskParams
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetTaskResult() *string {
	return s.TaskResult
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) GetTaskType() *int32 {
	return s.TaskType
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetEndTime(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.EndTime = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetStartTime(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskId(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskParams(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskResult(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskStatus(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskType(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) Validate() error {
	return dara.Validate(s)
}
