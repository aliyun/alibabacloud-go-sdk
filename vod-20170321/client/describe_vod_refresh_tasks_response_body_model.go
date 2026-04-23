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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *DescribeVodRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
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
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
