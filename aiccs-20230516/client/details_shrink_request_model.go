// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *DetailsShrinkRequest
	GetBatchId() *int64
	SetEndTime(v string) *DetailsShrinkRequest
	GetEndTime() *string
	SetNumberStatus(v int64) *DetailsShrinkRequest
	GetNumberStatus() *int64
	SetNumbersShrink(v string) *DetailsShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *DetailsShrinkRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *DetailsShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DetailsShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *DetailsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetailsShrinkRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DetailsShrinkRequest
	GetStartTime() *string
	SetTaskId(v int64) *DetailsShrinkRequest
	GetTaskId() *int64
}

type DetailsShrinkRequest struct {
	// 批次id
	//
	// example:
	//
	// 75
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 结束导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 号码状态
	//
	// example:
	//
	// 1
	NumberStatus *int64 `json:"NumberStatus,omitempty" xml:"NumberStatus,omitempty"`
	// 号码列表
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// example:
	//
	// 77
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 开始导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetailsShrinkRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *DetailsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DetailsShrinkRequest) GetNumberStatus() *int64 {
	return s.NumberStatus
}

func (s *DetailsShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *DetailsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetailsShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DetailsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DetailsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetailsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetailsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DetailsShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DetailsShrinkRequest) SetBatchId(v int64) *DetailsShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *DetailsShrinkRequest) SetEndTime(v string) *DetailsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DetailsShrinkRequest) SetNumberStatus(v int64) *DetailsShrinkRequest {
	s.NumberStatus = &v
	return s
}

func (s *DetailsShrinkRequest) SetNumbersShrink(v string) *DetailsShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *DetailsShrinkRequest) SetOwnerId(v int64) *DetailsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DetailsShrinkRequest) SetPageNo(v int64) *DetailsShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DetailsShrinkRequest) SetPageSize(v int64) *DetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DetailsShrinkRequest) SetResourceOwnerAccount(v string) *DetailsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetailsShrinkRequest) SetResourceOwnerId(v int64) *DetailsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetailsShrinkRequest) SetStartTime(v string) *DetailsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DetailsShrinkRequest) SetTaskId(v int64) *DetailsShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *DetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
