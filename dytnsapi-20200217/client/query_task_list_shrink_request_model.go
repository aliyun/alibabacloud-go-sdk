// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *QueryTaskListShrinkRequest
	GetCurrentPage() *int64
	SetOwnerId(v int64) *QueryTaskListShrinkRequest
	GetOwnerId() *int64
	SetPageSize(v int64) *QueryTaskListShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryTaskListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTaskListShrinkRequest
	GetResourceOwnerId() *int64
	SetResultShrink(v string) *QueryTaskListShrinkRequest
	GetResultShrink() *string
	SetTagId(v int64) *QueryTaskListShrinkRequest
	GetTagId() *int64
	SetTaskId(v int64) *QueryTaskListShrinkRequest
	GetTaskId() *int64
	SetTaskName(v string) *QueryTaskListShrinkRequest
	GetTaskName() *string
	SetTaskTypeShrink(v string) *QueryTaskListShrinkRequest
	GetTaskTypeShrink() *string
}

type QueryTaskListShrinkRequest struct {
	// example:
	//
	// 80
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResultShrink         *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 15
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// 91
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskName       *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskTypeShrink *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryTaskListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskListShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryTaskListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTaskListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTaskListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTaskListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTaskListShrinkRequest) GetResultShrink() *string {
	return s.ResultShrink
}

func (s *QueryTaskListShrinkRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryTaskListShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskListShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryTaskListShrinkRequest) GetTaskTypeShrink() *string {
	return s.TaskTypeShrink
}

func (s *QueryTaskListShrinkRequest) SetCurrentPage(v int64) *QueryTaskListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetOwnerId(v int64) *QueryTaskListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetPageSize(v int64) *QueryTaskListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetResourceOwnerAccount(v string) *QueryTaskListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetResourceOwnerId(v int64) *QueryTaskListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetResultShrink(v string) *QueryTaskListShrinkRequest {
	s.ResultShrink = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetTagId(v int64) *QueryTaskListShrinkRequest {
	s.TagId = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetTaskId(v int64) *QueryTaskListShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetTaskName(v string) *QueryTaskListShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *QueryTaskListShrinkRequest) SetTaskTypeShrink(v string) *QueryTaskListShrinkRequest {
	s.TaskTypeShrink = &v
	return s
}

func (s *QueryTaskListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
