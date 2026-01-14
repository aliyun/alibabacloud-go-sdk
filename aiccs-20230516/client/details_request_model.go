// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *DetailsRequest
	GetBatchId() *int64
	SetEndTime(v string) *DetailsRequest
	GetEndTime() *string
	SetNumberStatus(v int64) *DetailsRequest
	GetNumberStatus() *int64
	SetNumbers(v []*string) *DetailsRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *DetailsRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *DetailsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DetailsRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *DetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetailsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DetailsRequest
	GetStartTime() *string
	SetTaskId(v int64) *DetailsRequest
	GetTaskId() *int64
}

type DetailsRequest struct {
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
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailsRequest) GoString() string {
	return s.String()
}

func (s *DetailsRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *DetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DetailsRequest) GetNumberStatus() *int64 {
	return s.NumberStatus
}

func (s *DetailsRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *DetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetailsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DetailsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DetailsRequest) SetBatchId(v int64) *DetailsRequest {
	s.BatchId = &v
	return s
}

func (s *DetailsRequest) SetEndTime(v string) *DetailsRequest {
	s.EndTime = &v
	return s
}

func (s *DetailsRequest) SetNumberStatus(v int64) *DetailsRequest {
	s.NumberStatus = &v
	return s
}

func (s *DetailsRequest) SetNumbers(v []*string) *DetailsRequest {
	s.Numbers = v
	return s
}

func (s *DetailsRequest) SetOwnerId(v int64) *DetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetailsRequest) SetPageNo(v int64) *DetailsRequest {
	s.PageNo = &v
	return s
}

func (s *DetailsRequest) SetPageSize(v int64) *DetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DetailsRequest) SetResourceOwnerAccount(v string) *DetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetailsRequest) SetResourceOwnerId(v int64) *DetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetailsRequest) SetStartTime(v string) *DetailsRequest {
	s.StartTime = &v
	return s
}

func (s *DetailsRequest) SetTaskId(v int64) *DetailsRequest {
	s.TaskId = &v
	return s
}

func (s *DetailsRequest) Validate() error {
	return dara.Validate(s)
}
