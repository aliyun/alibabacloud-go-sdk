// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *TaskCallListShrinkRequest
	GetBatchId() *string
	SetCallDate(v string) *TaskCallListShrinkRequest
	GetCallDate() *string
	SetEndCallDate(v string) *TaskCallListShrinkRequest
	GetEndCallDate() *string
	SetIntentTagsShrink(v string) *TaskCallListShrinkRequest
	GetIntentTagsShrink() *string
	SetNumbersShrink(v string) *TaskCallListShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *TaskCallListShrinkRequest
	GetOwnerId() *int64
	SetPage(v int64) *TaskCallListShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *TaskCallListShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *TaskCallListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCallListShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *TaskCallListShrinkRequest
	GetTaskId() *int64
}

type TaskCallListShrinkRequest struct {
	// 导入号码时返回的批次号
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// 结束外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	EndCallDate *string `json:"EndCallDate,omitempty" xml:"EndCallDate,omitempty"`
	// 意向标签
	IntentTagsShrink *string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty"`
	// 号码列表
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// This parameter is required.
	//
	// example:
	//
	// 39
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 每页外呼记录数,正整数，默认10000
	//
	// example:
	//
	// 97
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskCallListShrinkRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *TaskCallListShrinkRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *TaskCallListShrinkRequest) GetEndCallDate() *string {
	return s.EndCallDate
}

func (s *TaskCallListShrinkRequest) GetIntentTagsShrink() *string {
	return s.IntentTagsShrink
}

func (s *TaskCallListShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *TaskCallListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCallListShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *TaskCallListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *TaskCallListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCallListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCallListShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCallListShrinkRequest) SetBatchId(v string) *TaskCallListShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetCallDate(v string) *TaskCallListShrinkRequest {
	s.CallDate = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetEndCallDate(v string) *TaskCallListShrinkRequest {
	s.EndCallDate = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetIntentTagsShrink(v string) *TaskCallListShrinkRequest {
	s.IntentTagsShrink = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetNumbersShrink(v string) *TaskCallListShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetOwnerId(v int64) *TaskCallListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetPage(v int64) *TaskCallListShrinkRequest {
	s.Page = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetPageSize(v int64) *TaskCallListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetResourceOwnerAccount(v string) *TaskCallListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetResourceOwnerId(v int64) *TaskCallListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetTaskId(v int64) *TaskCallListShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCallListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
