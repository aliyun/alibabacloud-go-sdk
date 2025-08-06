// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRefreshTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRefreshTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRefreshTasksResponseBody
	GetRequestId() *string
	SetTasks(v *DescribeRefreshTasksResponseBodyTasks) *DescribeRefreshTasksResponseBody
	GetTasks() *DescribeRefreshTasksResponseBodyTasks
	SetTotalCount(v int32) *DescribeRefreshTasksResponseBody
	GetTotalCount() *int32
}

type DescribeRefreshTasksResponseBody struct {
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      *DescribeRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRefreshTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRefreshTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRefreshTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefreshTasksResponseBody) GetTasks() *DescribeRefreshTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeRefreshTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRefreshTasksResponseBody) SetPageNumber(v int32) *DescribeRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetPageSize(v int32) *DescribeRefreshTasksResponseBody {
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

func (s *DescribeRefreshTasksResponseBody) SetTotalCount(v int32) *DescribeRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRefreshTasksResponseBodyTasks struct {
	Task []*DescribeRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeRefreshTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasks) GetTask() []*DescribeRefreshTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeRefreshTasksResponseBodyTasks) SetTask(v []*DescribeRefreshTasksResponseBodyTasksTask) *DescribeRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeRefreshTasksResponseBodyTasksTask struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetProcess() *string {
	return s.Process
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
