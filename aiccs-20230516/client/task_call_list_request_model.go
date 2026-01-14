// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *TaskCallListRequest
	GetBatchId() *string
	SetCallDate(v string) *TaskCallListRequest
	GetCallDate() *string
	SetEndCallDate(v string) *TaskCallListRequest
	GetEndCallDate() *string
	SetIntentTags(v []*string) *TaskCallListRequest
	GetIntentTags() []*string
	SetNumbers(v []*string) *TaskCallListRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *TaskCallListRequest
	GetOwnerId() *int64
	SetPage(v int64) *TaskCallListRequest
	GetPage() *int64
	SetPageSize(v int64) *TaskCallListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *TaskCallListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCallListRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *TaskCallListRequest
	GetTaskId() *int64
}

type TaskCallListRequest struct {
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
	IntentTags []*string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 号码列表
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s TaskCallListRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListRequest) GoString() string {
	return s.String()
}

func (s *TaskCallListRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *TaskCallListRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *TaskCallListRequest) GetEndCallDate() *string {
	return s.EndCallDate
}

func (s *TaskCallListRequest) GetIntentTags() []*string {
	return s.IntentTags
}

func (s *TaskCallListRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *TaskCallListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCallListRequest) GetPage() *int64 {
	return s.Page
}

func (s *TaskCallListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *TaskCallListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCallListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCallListRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCallListRequest) SetBatchId(v string) *TaskCallListRequest {
	s.BatchId = &v
	return s
}

func (s *TaskCallListRequest) SetCallDate(v string) *TaskCallListRequest {
	s.CallDate = &v
	return s
}

func (s *TaskCallListRequest) SetEndCallDate(v string) *TaskCallListRequest {
	s.EndCallDate = &v
	return s
}

func (s *TaskCallListRequest) SetIntentTags(v []*string) *TaskCallListRequest {
	s.IntentTags = v
	return s
}

func (s *TaskCallListRequest) SetNumbers(v []*string) *TaskCallListRequest {
	s.Numbers = v
	return s
}

func (s *TaskCallListRequest) SetOwnerId(v int64) *TaskCallListRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallListRequest) SetPage(v int64) *TaskCallListRequest {
	s.Page = &v
	return s
}

func (s *TaskCallListRequest) SetPageSize(v int64) *TaskCallListRequest {
	s.PageSize = &v
	return s
}

func (s *TaskCallListRequest) SetResourceOwnerAccount(v string) *TaskCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallListRequest) SetResourceOwnerId(v int64) *TaskCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallListRequest) SetTaskId(v int64) *TaskCallListRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCallListRequest) Validate() error {
	return dara.Validate(s)
}
