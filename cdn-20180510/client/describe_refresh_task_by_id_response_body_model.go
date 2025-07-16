// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTaskByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRefreshTaskByIdResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeRefreshTaskByIdResponseBodyTasks) *DescribeRefreshTaskByIdResponseBody
	GetTasks() []*DescribeRefreshTaskByIdResponseBodyTasks
	SetTotalCount(v int64) *DescribeRefreshTaskByIdResponseBody
	GetTotalCount() *int64
}

type DescribeRefreshTaskByIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E0C2EF95-B1EC-4C93-855E-2059A7DA2B7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about tasks.
	Tasks []*DescribeRefreshTaskByIdResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRefreshTaskByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefreshTaskByIdResponseBody) GetTasks() []*DescribeRefreshTaskByIdResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeRefreshTaskByIdResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRefreshTaskByIdResponseBody) SetRequestId(v string) *DescribeRefreshTaskByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBody) SetTasks(v []*DescribeRefreshTaskByIdResponseBodyTasks) *DescribeRefreshTaskByIdResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBody) SetTotalCount(v int64) *DescribeRefreshTaskByIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRefreshTaskByIdResponseBodyTasks struct {
	// The time when the task was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-08-03T08:54:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error returned when the refresh or prefetch task failed. Valid values:
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
	// The path of the object refreshed by the refresh task.
	//
	// example:
	//
	// http://example.com/abc.jpg
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **file**: refreshes an individual file.
	//
	// 	- **directory**: refreshes files in the specified directory.
	//
	// 	- **preload**: prefetches an individual file.
	//
	// 	- **regex**: refreshes content based on a regular expression.
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The progress of the task, in percentage.
	//
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **Complete**
	//
	// 	- **Pending**
	//
	// 	- **Refreshing**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 24840
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTaskByIdResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetProcess() *string {
	return s.Process
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetCreationTime(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetDescription(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetObjectPath(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetObjectType(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetProcess(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetStatus(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetTaskId(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
