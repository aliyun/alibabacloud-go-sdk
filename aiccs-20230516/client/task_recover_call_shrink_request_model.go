// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskRecoverCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginImportTime(v string) *TaskRecoverCallShrinkRequest
	GetBeginImportTime() *string
	SetEndImportTime(v string) *TaskRecoverCallShrinkRequest
	GetEndImportTime() *string
	SetNumbersShrink(v string) *TaskRecoverCallShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *TaskRecoverCallShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskRecoverCallShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskRecoverCallShrinkRequest
	GetResourceOwnerId() *int64
	SetTagsShrink(v string) *TaskRecoverCallShrinkRequest
	GetTagsShrink() *string
	SetTaskId(v int64) *TaskRecoverCallShrinkRequest
	GetTaskId() *int64
}

type TaskRecoverCallShrinkRequest struct {
	// 查询开始导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskRecoverCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskRecoverCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallShrinkRequest) GetBeginImportTime() *string {
	return s.BeginImportTime
}

func (s *TaskRecoverCallShrinkRequest) GetEndImportTime() *string {
	return s.EndImportTime
}

func (s *TaskRecoverCallShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *TaskRecoverCallShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskRecoverCallShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskRecoverCallShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskRecoverCallShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *TaskRecoverCallShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskRecoverCallShrinkRequest) SetBeginImportTime(v string) *TaskRecoverCallShrinkRequest {
	s.BeginImportTime = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetEndImportTime(v string) *TaskRecoverCallShrinkRequest {
	s.EndImportTime = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetNumbersShrink(v string) *TaskRecoverCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetOwnerId(v int64) *TaskRecoverCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetResourceOwnerAccount(v string) *TaskRecoverCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetResourceOwnerId(v int64) *TaskRecoverCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetTagsShrink(v string) *TaskRecoverCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetTaskId(v int64) *TaskRecoverCallShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
