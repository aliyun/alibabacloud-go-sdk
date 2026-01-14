// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCancelCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumbersShrink(v string) *TaskCancelCallShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *TaskCancelCallShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskCancelCallShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCancelCallShrinkRequest
	GetResourceOwnerId() *int64
	SetTagsShrink(v string) *TaskCancelCallShrinkRequest
	GetTagsShrink() *string
	SetTaskId(v int64) *TaskCancelCallShrinkRequest
	GetTaskId() *int64
}

type TaskCancelCallShrinkRequest struct {
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagsShrink           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCancelCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCancelCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskCancelCallShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *TaskCancelCallShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCancelCallShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCancelCallShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCancelCallShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *TaskCancelCallShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCancelCallShrinkRequest) SetNumbersShrink(v string) *TaskCancelCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetOwnerId(v int64) *TaskCancelCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetResourceOwnerAccount(v string) *TaskCancelCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetResourceOwnerId(v int64) *TaskCancelCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetTagsShrink(v string) *TaskCancelCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetTaskId(v int64) *TaskCancelCallShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
