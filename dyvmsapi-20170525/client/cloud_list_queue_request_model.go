// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CloudListQueueRequest
	GetDescription() *string
	SetEndTime(v string) *CloudListQueueRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudListQueueRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudListQueueRequest
	GetLimit() *int64
	SetOffset(v int64) *CloudListQueueRequest
	GetOffset() *int64
	SetOrder(v int64) *CloudListQueueRequest
	GetOrder() *int64
	SetOwnerId(v int64) *CloudListQueueRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudListQueueRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudListQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListQueueRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CloudListQueueRequest
	GetStartTime() *string
}

type CloudListQueueRequest struct {
	// 描述
	//
	// example:
	//
	// 示例值示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建时间终止时间点；取值说明："2019-10-11 23:59:59"
	//
	// example:
	//
	// 2026-04-20 23:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 条数；正整数，大于0，最大不能超过500， 默认为500
	//
	// example:
	//
	// 500
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 第几条开始；正整数 大于等于0 默认为0
	//
	// example:
	//
	// 0
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// 排序方式,按照创建时间排序；0正序 1倒序
	//
	// example:
	//
	// 0
	Order   *int64 `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号
	//
	// example:
	//
	// 233
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 创建时间起始时间点；取值说明："2019-10-11 00:00:00"
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CloudListQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueRequest) GoString() string {
	return s.String()
}

func (s *CloudListQueueRequest) GetDescription() *string {
	return s.Description
}

func (s *CloudListQueueRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudListQueueRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListQueueRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudListQueueRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *CloudListQueueRequest) GetOrder() *int64 {
	return s.Order
}

func (s *CloudListQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListQueueRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudListQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListQueueRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudListQueueRequest) SetDescription(v string) *CloudListQueueRequest {
	s.Description = &v
	return s
}

func (s *CloudListQueueRequest) SetEndTime(v string) *CloudListQueueRequest {
	s.EndTime = &v
	return s
}

func (s *CloudListQueueRequest) SetEnterpriseId(v int64) *CloudListQueueRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListQueueRequest) SetLimit(v int64) *CloudListQueueRequest {
	s.Limit = &v
	return s
}

func (s *CloudListQueueRequest) SetOffset(v int64) *CloudListQueueRequest {
	s.Offset = &v
	return s
}

func (s *CloudListQueueRequest) SetOrder(v int64) *CloudListQueueRequest {
	s.Order = &v
	return s
}

func (s *CloudListQueueRequest) SetOwnerId(v int64) *CloudListQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListQueueRequest) SetQno(v string) *CloudListQueueRequest {
	s.Qno = &v
	return s
}

func (s *CloudListQueueRequest) SetResourceOwnerAccount(v string) *CloudListQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListQueueRequest) SetResourceOwnerId(v int64) *CloudListQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListQueueRequest) SetStartTime(v string) *CloudListQueueRequest {
	s.StartTime = &v
	return s
}

func (s *CloudListQueueRequest) Validate() error {
	return dara.Validate(s)
}
