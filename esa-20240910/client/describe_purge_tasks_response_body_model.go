// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurgeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribePurgeTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePurgeTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePurgeTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribePurgeTasksResponseBodyTasks) *DescribePurgeTasksResponseBody
	GetTasks() []*DescribePurgeTasksResponseBodyTasks
	SetTotalCount(v int64) *DescribePurgeTasksResponseBody
	GetTotalCount() *int64
}

type DescribePurgeTasksResponseBody struct {
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
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*DescribePurgeTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurgeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePurgeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePurgeTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePurgeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePurgeTasksResponseBody) GetTasks() []*DescribePurgeTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribePurgeTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePurgeTasksResponseBody) SetPageNumber(v int64) *DescribePurgeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetPageSize(v int64) *DescribePurgeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetRequestId(v string) *DescribePurgeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetTasks(v []*DescribePurgeTasksResponseBodyTasks) *DescribePurgeTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetTotalCount(v int64) *DescribePurgeTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePurgeTasksResponseBodyTasks struct {
	// The purged content.
	//
	// example:
	//
	// http://a.com/1.jpg?b=1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2023-07-26T01:56:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error description returned when the purge task failed.
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
	// The task ID.
	//
	// example:
	//
	// 16346513304
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the purge task. Valid values:
	//
	// 	- **file*	- (default): purges the cache by file.
	//
	// 	- **cachetag**: purges the cache by cache tag.
	//
	// 	- **directory**: purges the cache by directory.
	//
	// 	- **ignoreParams**: purges the cache by URL with specified parameters ignored.
	//
	// 	- **hostname**: purges the cache by hostname.
	//
	// 	- **purgeall**: purges all cache.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePurgeTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribePurgeTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponseBodyTasks) GetContent() *string {
	return s.Content
}

func (s *DescribePurgeTasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePurgeTasksResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribePurgeTasksResponseBodyTasks) GetProcess() *string {
	return s.Process
}

func (s *DescribePurgeTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribePurgeTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePurgeTasksResponseBodyTasks) GetType() *string {
	return s.Type
}

func (s *DescribePurgeTasksResponseBodyTasks) SetContent(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Content = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetCreateTime(v string) *DescribePurgeTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetDescription(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetProcess(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetStatus(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetTaskId(v string) *DescribePurgeTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetType(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
