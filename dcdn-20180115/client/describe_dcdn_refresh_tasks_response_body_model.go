// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeDcdnRefreshTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnRefreshTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnRefreshTasksResponseBody
	GetRequestId() *string
	SetTasks(v *DescribeDcdnRefreshTasksResponseBodyTasks) *DescribeDcdnRefreshTasksResponseBody
	GetTasks() *DescribeDcdnRefreshTasksResponseBodyTasks
	SetTotalCount(v int64) *DescribeDcdnRefreshTasksResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnRefreshTasksResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 174F6032-AA26-470D-B90E-36F0EB205BEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about tasks.
	Tasks *DescribeDcdnRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	// The number of tasks.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnRefreshTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnRefreshTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnRefreshTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnRefreshTasksResponseBody) GetTasks() *DescribeDcdnRefreshTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeDcdnRefreshTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetPageSize(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetRequestId(v string) *DescribeDcdnRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetTasks(v *DescribeDcdnRefreshTasksResponseBodyTasks) *DescribeDcdnRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRefreshTasksResponseBodyTasks struct {
	Task []*DescribeDcdnRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRefreshTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasks) GetTask() []*DescribeDcdnRefreshTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasks) SetTask(v []*DescribeDcdnRefreshTasksResponseBodyTasksTask) *DescribeDcdnRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRefreshTasksResponseBodyTasksTask struct {
	// The time when the task was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-11-27T08:23:22Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The type of error returned when the refresh or prefetch task has failed.
	//
	// 	- **InternalError**: An internal error occurred.
	//
	// 	- **OriginTimeout**: The response from the origin server timed out.
	//
	// 	- **OriginReturn StatusCode 5XX**: The origin server returned a 5XX error.
	//
	// example:
	//
	// InternalError
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the object refreshed.
	//
	// example:
	//
	// http://example.com/examplefile.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task.
	//
	// 	- **file**: URL-based refresh
	//
	// 	- **path**: directory-based refresh
	//
	// 	- **preload**: URL-based prefetch
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The progress of the task in percentage.
	//
	// example:
	//
	// 10
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The status of the task.
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
	// The ID of the task.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetProcess() *string {
	return s.Process
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
