// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeRefreshTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRefreshTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeRefreshTasksResponseBody
	GetRequestId() *string
	SetTasks(v *DescribeRefreshTasksResponseBodyTasks) *DescribeRefreshTasksResponseBody
	GetTasks() *DescribeRefreshTasksResponseBodyTasks
	SetTotalCount(v int64) *DescribeRefreshTasksResponseBody
	GetTotalCount() *int64
}

type DescribeRefreshTasksResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 174F6032-AA26-470D-B90E-36F0EB205BEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about tasks.
	Tasks *DescribeRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRefreshTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRefreshTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRefreshTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefreshTasksResponseBody) GetTasks() *DescribeRefreshTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeRefreshTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetPageSize(v int64) *DescribeRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetRequestId(v string) *DescribeRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetTasks(v *DescribeRefreshTasksResponseBodyTasks) *DescribeRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRefreshTasksResponseBodyTasks struct {
	CDNTask []*DescribeRefreshTasksResponseBodyTasksCDNTask `json:"CDNTask,omitempty" xml:"CDNTask,omitempty" type:"Repeated"`
}

func (s DescribeRefreshTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasks) GetCDNTask() []*DescribeRefreshTasksResponseBodyTasksCDNTask {
	return s.CDNTask
}

func (s *DescribeRefreshTasksResponseBodyTasks) SetCDNTask(v []*DescribeRefreshTasksResponseBodyTasksCDNTask) *DescribeRefreshTasksResponseBodyTasks {
	s.CDNTask = v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasks) Validate() error {
	if s.CDNTask != nil {
		for _, item := range s.CDNTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRefreshTasksResponseBodyTasksCDNTask struct {
	// The time when the task was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-11-27T08:23:22Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The type of the error returned when the refresh or prefetch task failed. Valid values:
	//
	// 	- **InternalError**: An internal error occurred.
	//
	// 	- **OriginTimeout**: The response from the origin server timed out.
	//
	// 	- **OriginReturnStatusCode 5XX**: The origin server returned a 5XX error.
	//
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the object refreshed.
	//
	// example:
	//
	// http://example.com/1.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task.
	//
	// 	- **file**: refreshes one or more files.
	//
	// 	- **directory**: refreshes files in the specified directories.
	//
	// 	- **regex**: refreshes content based on a regular expression.
	//
	// 	- **preload**: prefetches one or more files.
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
	// 	- **Complete**: The task has completed.
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
	// 704225667
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTasksResponseBodyTasksCDNTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasksCDNTask) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetProcess() *string {
	return s.Process
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetCreationTime(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetDescription(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Description = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetObjectPath(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetObjectType(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetProcess(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Process = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetStatus(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetTaskId(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) Validate() error {
	return dara.Validate(s)
}
