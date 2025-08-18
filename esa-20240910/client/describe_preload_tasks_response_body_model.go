// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribePreloadTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePreloadTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePreloadTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribePreloadTasksResponseBodyTasks) *DescribePreloadTasksResponseBody
	GetTasks() []*DescribePreloadTasksResponseBodyTasks
	SetTotalCount(v int64) *DescribePreloadTasksResponseBody
	GetTotalCount() *int64
}

type DescribePreloadTasksResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*DescribePreloadTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 83
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePreloadTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePreloadTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePreloadTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreloadTasksResponseBody) GetTasks() []*DescribePreloadTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribePreloadTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePreloadTasksResponseBody) SetPageNumber(v int64) *DescribePreloadTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetPageSize(v int64) *DescribePreloadTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetRequestId(v string) *DescribePreloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetTasks(v []*DescribePreloadTasksResponseBodyTasks) *DescribePreloadTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetTotalCount(v int64) *DescribePreloadTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePreloadTasksResponseBodyTasks struct {
	// The prefetched content.
	//
	// example:
	//
	// http://a.com/1.jpg?b=2
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2023-03-28 14:28:57
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned upon a prefetch task failure. Valid values:
	//
	// 	- **Internal Error**
	//
	// 	- **Origin Timeout**
	//
	// 	- **Origin Return StatusCode 5XX**
	//
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The progress of the task, in percentage.
	//
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The task status.
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Refreshing**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the queried task.
	//
	// example:
	//
	// 1597854579687428
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreloadTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponseBodyTasks) GetContent() *string {
	return s.Content
}

func (s *DescribePreloadTasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePreloadTasksResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribePreloadTasksResponseBodyTasks) GetProcess() *string {
	return s.Process
}

func (s *DescribePreloadTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribePreloadTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreloadTasksResponseBodyTasks) SetContent(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Content = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetCreateTime(v string) *DescribePreloadTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetDescription(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetProcess(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetStatus(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetTaskId(v string) *DescribePreloadTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
