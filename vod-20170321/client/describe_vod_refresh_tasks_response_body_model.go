// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeVodRefreshTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodRefreshTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeVodRefreshTasksResponseBody
	GetRequestId() *string
	SetTasks(v *DescribeVodRefreshTasksResponseBodyTasks) *DescribeVodRefreshTasksResponseBody
	GetTasks() *DescribeVodRefreshTasksResponseBodyTasks
	SetTotalCount(v int64) *DescribeVodRefreshTasksResponseBody
	GetTotalCount() *int64
}

type DescribeVodRefreshTasksResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 174F6032-AA26-470D-****-36F0EB205BEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the returned tasks.
	Tasks *DescribeVodRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodRefreshTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodRefreshTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodRefreshTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodRefreshTasksResponseBody) GetTasks() *DescribeVodRefreshTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeVodRefreshTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeVodRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeVodRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetPageSize(v int64) *DescribeVodRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetRequestId(v string) *DescribeVodRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetTasks(v *DescribeVodRefreshTasksResponseBodyTasks) *DescribeVodRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeVodRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodRefreshTasksResponseBodyTasks struct {
	Task []*DescribeVodRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeVodRefreshTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBodyTasks) GetTask() []*DescribeVodRefreshTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeVodRefreshTasksResponseBodyTasks) SetTask(v []*DescribeVodRefreshTasksResponseBodyTasksTask) *DescribeVodRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasks) Validate() error {
	if s.Task != nil {
		for _, item := range s.Task {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodRefreshTasksResponseBodyTasksTask struct {
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-11-27T08:23:22Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The type of error returned when the refresh or prefetch task failed. Valid values: Valid values:
	//
	// 	- **Internal Error**: An internal error occurred.
	//
	// 	- **Origin Timeout**: The response from the origin server timed out.
	//
	// 	- **Origin Return StatusCode 5XX**: The origin server returned an HTTP status code 5xx.
	//
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the object refreshed.
	//
	// example:
	//
	// http://example.com/****.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task. Default value: file. Valid values:
	//
	// 	- **file**: refreshes one or more files.
	//
	// 	- **directory**: refreshes files in the specified directory.
	//
	// 	- **preload**: prefetches one or more files.
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The progress of the task in percentage.
	//
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Refreshing**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Pending**: The task is pending.
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

func (s DescribeVodRefreshTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetProcess() *string {
	return s.Process
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
