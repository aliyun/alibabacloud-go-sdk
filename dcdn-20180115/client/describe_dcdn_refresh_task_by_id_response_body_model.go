// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTaskByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnRefreshTaskByIdResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeDcdnRefreshTaskByIdResponseBodyTasks) *DescribeDcdnRefreshTaskByIdResponseBody
	GetTasks() []*DescribeDcdnRefreshTaskByIdResponseBodyTasks
	SetTotalCount(v int64) *DescribeDcdnRefreshTaskByIdResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnRefreshTaskByIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E0C2EF95-B1EC-4C93-855E-2059A7DA2B7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of prefetch or refresh tasks.
	Tasks []*DescribeDcdnRefreshTaskByIdResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) GetTasks() []*DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetRequestId(v string) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetTasks(v []*DescribeDcdnRefreshTaskByIdResponseBodyTasks) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetTotalCount(v int64) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRefreshTaskByIdResponseBodyTasks struct {
	// The time when the task was created. The time follows the ISO8601 standard in the YYYY-MM-DDThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-10T08:54:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error returned when the refresh or prefetch task failed. Valid values:
	//
	// 	- **Internal Error**: An internal error occurred.
	//
	// 	- **Origin Timeout**: The response from the origin server timed out.
	//
	// 	- **Origin Return StatusCode 5XX**: The origin server returned a 5XX error.
	//
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The path of the refresh or prefetch object.
	//
	// example:
	//
	// http://example.com/image_01.png
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the refresh or prefetch task. Valid values:
	//
	// 	- **file**: refreshes an individual file.
	//
	// 	- **directory**: refreshes files under the specified directory.
	//
	// 	- **preload**: prefetches an individual file.
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
	// The task status. Valid values:
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Pending**: The task is pending.
	//
	// 	- **Refreshing**: The task is running.
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
	// 113681**
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetProcess() *string {
	return s.Process
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetCreationTime(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetDescription(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetObjectPath(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetObjectType(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetProcess(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetStatus(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetTaskId(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
