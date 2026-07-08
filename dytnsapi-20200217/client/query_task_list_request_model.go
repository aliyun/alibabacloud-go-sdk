// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *QueryTaskListRequest
	GetCurrentPage() *int64
	SetOwnerId(v int64) *QueryTaskListRequest
	GetOwnerId() *int64
	SetPageSize(v int64) *QueryTaskListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryTaskListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTaskListRequest
	GetResourceOwnerId() *int64
	SetResult(v []*int64) *QueryTaskListRequest
	GetResult() []*int64
	SetTagId(v int64) *QueryTaskListRequest
	GetTagId() *int64
	SetTaskId(v int64) *QueryTaskListRequest
	GetTaskId() *int64
	SetTaskName(v string) *QueryTaskListRequest
	GetTaskName() *string
	SetTaskType(v []*int64) *QueryTaskListRequest
	GetTaskType() []*int64
}

type QueryTaskListRequest struct {
	// The current page number.
	//
	// example:
	//
	// 80
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page size. Maximum value: 1000.
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task statuses.
	Result []*int64 `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The tag ID. You can call the [QueryTagListPage](~~QueryTagListPage~~) operation to query tag IDs.
	//
	// example:
	//
	// 15
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The task ID. You can call the DescribeTasks operation to query the list of task IDs.
	//
	// example:
	//
	// 91
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// Example
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// An array of task types.
	TaskType []*int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty" type:"Repeated"`
}

func (s QueryTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryTaskListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTaskListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTaskListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTaskListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTaskListRequest) GetResult() []*int64 {
	return s.Result
}

func (s *QueryTaskListRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryTaskListRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryTaskListRequest) GetTaskType() []*int64 {
	return s.TaskType
}

func (s *QueryTaskListRequest) SetCurrentPage(v int64) *QueryTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTaskListRequest) SetOwnerId(v int64) *QueryTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTaskListRequest) SetPageSize(v int64) *QueryTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListRequest) SetResourceOwnerAccount(v string) *QueryTaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTaskListRequest) SetResourceOwnerId(v int64) *QueryTaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTaskListRequest) SetResult(v []*int64) *QueryTaskListRequest {
	s.Result = v
	return s
}

func (s *QueryTaskListRequest) SetTagId(v int64) *QueryTaskListRequest {
	s.TagId = &v
	return s
}

func (s *QueryTaskListRequest) SetTaskId(v int64) *QueryTaskListRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskListRequest) SetTaskName(v string) *QueryTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *QueryTaskListRequest) SetTaskType(v []*int64) *QueryTaskListRequest {
	s.TaskType = v
	return s
}

func (s *QueryTaskListRequest) Validate() error {
	return dara.Validate(s)
}
