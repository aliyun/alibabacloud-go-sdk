// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeClusterTasksResponseBodyPageInfo) *DescribeClusterTasksResponseBody
	GetPageInfo() *DescribeClusterTasksResponseBodyPageInfo
	SetRequestId(v string) *DescribeClusterTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeClusterTasksResponseBodyTasks) *DescribeClusterTasksResponseBody
	GetTasks() []*DescribeClusterTasksResponseBodyTasks
}

type DescribeClusterTasksResponseBody struct {
	// The pagination information.
	PageInfo *DescribeClusterTasksResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0527ac9a-c899-4341-a21a-xxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the tasks.
	Tasks []*DescribeClusterTasksResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s DescribeClusterTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksResponseBody) GetPageInfo() *DescribeClusterTasksResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeClusterTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterTasksResponseBody) GetTasks() []*DescribeClusterTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeClusterTasksResponseBody) SetPageInfo(v *DescribeClusterTasksResponseBodyPageInfo) *DescribeClusterTasksResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeClusterTasksResponseBody) SetRequestId(v string) *DescribeClusterTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterTasksResponseBody) SetTasks(v []*DescribeClusterTasksResponseBodyTasks) *DescribeClusterTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeClusterTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterTasksResponseBodyPageInfo struct {
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClusterTasksResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksResponseBodyPageInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeClusterTasksResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeClusterTasksResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeClusterTasksResponseBodyPageInfo) SetPageNumber(v int64) *DescribeClusterTasksResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyPageInfo) SetPageSize(v int64) *DescribeClusterTasksResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyPageInfo) SetTotalCount(v int64) *DescribeClusterTasksResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterTasksResponseBodyTasks struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2022-08-03T10:11:33+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The error returned for the task.
	Error *DescribeClusterTasksResponseBodyTasksError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The status of the task.
	//
	// example:
	//
	// success
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The task ID.
	//
	// example:
	//
	// install-addons-c3xxxxxx
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// The type of task.
	//
	// example:
	//
	// cluster_addon_install
	TaskType *string `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 2022-08-03T10:12:03.482+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeClusterTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksResponseBodyTasks) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterTasksResponseBodyTasks) GetError() *DescribeClusterTasksResponseBodyTasksError {
	return s.Error
}

func (s *DescribeClusterTasksResponseBodyTasks) GetState() *string {
	return s.State
}

func (s *DescribeClusterTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeClusterTasksResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeClusterTasksResponseBodyTasks) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClusterTasksResponseBodyTasks) SetCreated(v string) *DescribeClusterTasksResponseBodyTasks {
	s.Created = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) SetError(v *DescribeClusterTasksResponseBodyTasksError) *DescribeClusterTasksResponseBodyTasks {
	s.Error = v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) SetState(v string) *DescribeClusterTasksResponseBodyTasks {
	s.State = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) SetTaskId(v string) *DescribeClusterTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) SetTaskType(v string) *DescribeClusterTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) SetUpdated(v string) *DescribeClusterTasksResponseBodyTasks {
	s.Updated = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterTasksResponseBodyTasksError struct {
	// The error code returned.
	//
	// example:
	//
	// BadRequest
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Addon status not match
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DescribeClusterTasksResponseBodyTasksError) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksResponseBodyTasksError) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksResponseBodyTasksError) GetCode() *string {
	return s.Code
}

func (s *DescribeClusterTasksResponseBodyTasksError) GetMessage() *string {
	return s.Message
}

func (s *DescribeClusterTasksResponseBodyTasksError) SetCode(v string) *DescribeClusterTasksResponseBodyTasksError {
	s.Code = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasksError) SetMessage(v string) *DescribeClusterTasksResponseBodyTasksError {
	s.Message = &v
	return s
}

func (s *DescribeClusterTasksResponseBodyTasksError) Validate() error {
	return dara.Validate(s)
}
