// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListBenchmarkTaskResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBenchmarkTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBenchmarkTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*ListBenchmarkTaskResponseBodyTasks) *ListBenchmarkTaskResponseBody
	GetTasks() []*ListBenchmarkTaskResponseBodyTasks
	SetTotalCount(v int32) *ListBenchmarkTaskResponseBody
	GetTotalCount() *int32
}

type ListBenchmarkTaskResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stress testing tasks.
	Tasks []*ListBenchmarkTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBenchmarkTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBenchmarkTaskResponseBody) GetTasks() []*ListBenchmarkTaskResponseBodyTasks {
	return s.Tasks
}

func (s *ListBenchmarkTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBenchmarkTaskResponseBody) SetPageNumber(v int32) *ListBenchmarkTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetPageSize(v int32) *ListBenchmarkTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetRequestId(v string) *ListBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetTasks(v []*ListBenchmarkTaskResponseBodyTasks) *ListBenchmarkTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetTotalCount(v int32) *ListBenchmarkTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBenchmarkTaskResponseBodyTasks struct {
	// The number of instances that are available for stress testing.
	//
	// example:
	//
	// 2
	AvailableAgent *int64 `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	// The time when the stress testing task was created.
	//
	// example:
	//
	// 2020-12-04T02:43:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID of the stress testing task.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the service on which you want to perform a stress testing.
	//
	// example:
	//
	// test_quota
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stress testing task.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Error
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the stress testing task.
	//
	// example:
	//
	// eas-b-gv4y86uvgt****i
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-1076
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The time when the stress testing task was updated.
	//
	// example:
	//
	// 2020-06-24T03:11:30Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListBenchmarkTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListBenchmarkTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetAvailableAgent() *int64 {
	return s.AvailableAgent
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetMessage() *string {
	return s.Message
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetRegion() *string {
	return s.Region
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListBenchmarkTaskResponseBodyTasks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetAvailableAgent(v int64) *ListBenchmarkTaskResponseBodyTasks {
	s.AvailableAgent = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetCreateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetMessage(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Message = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetRegion(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Region = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetServiceName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.ServiceName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetStatus(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskId(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetUpdateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.UpdateTime = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
