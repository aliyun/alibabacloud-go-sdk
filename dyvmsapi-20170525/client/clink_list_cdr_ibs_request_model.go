// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientNumber(v string) *ClinkListCdrIbsRequest
	GetClientNumber() *string
	SetCno(v string) *ClinkListCdrIbsRequest
	GetCno() *string
	SetCustomerNumber(v string) *ClinkListCdrIbsRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *ClinkListCdrIbsRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *ClinkListCdrIbsRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkListCdrIbsRequest
	GetHiddenType() *int64
	SetLimit(v int64) *ClinkListCdrIbsRequest
	GetLimit() *int64
	SetMainUniqueId(v string) *ClinkListCdrIbsRequest
	GetMainUniqueId() *string
	SetOffset(v int64) *ClinkListCdrIbsRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListCdrIbsRequest
	GetOwnerId() *int64
	SetQno(v string) *ClinkListCdrIbsRequest
	GetQno() *string
	SetQueueAnswerInTime(v int64) *ClinkListCdrIbsRequest
	GetQueueAnswerInTime() *int64
	SetResourceOwnerAccount(v string) *ClinkListCdrIbsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListCdrIbsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ClinkListCdrIbsRequest
	GetStartTime() *int64
	SetStatus(v int64) *ClinkListCdrIbsRequest
	GetStatus() *int64
}

type ClinkListCdrIbsRequest struct {
	// example:
	//
	// 177xxxx7750
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// example:
	//
	// 0032
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// example:
	//
	// xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// example:
	//
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// example:
	//
	// 14
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// example:
	//
	// 0
	Offset  *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2233
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// example:
	//
	// 0
	QueueAnswerInTime    *int64  `json:"QueueAnswerInTime,omitempty" xml:"QueueAnswerInTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ClinkListCdrIbsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbsRequest) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbsRequest) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrIbsRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrIbsRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrIbsRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkListCdrIbsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListCdrIbsRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrIbsRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListCdrIbsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListCdrIbsRequest) GetQno() *string {
	return s.Qno
}

func (s *ClinkListCdrIbsRequest) GetQueueAnswerInTime() *int64 {
	return s.QueueAnswerInTime
}

func (s *ClinkListCdrIbsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListCdrIbsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListCdrIbsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbsRequest) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkListCdrIbsRequest) SetClientNumber(v string) *ClinkListCdrIbsRequest {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetCno(v string) *ClinkListCdrIbsRequest {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetCustomerNumber(v string) *ClinkListCdrIbsRequest {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetEndTime(v int64) *ClinkListCdrIbsRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetEnterpriseId(v int64) *ClinkListCdrIbsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetHiddenType(v int64) *ClinkListCdrIbsRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetLimit(v int64) *ClinkListCdrIbsRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetMainUniqueId(v string) *ClinkListCdrIbsRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetOffset(v int64) *ClinkListCdrIbsRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetOwnerId(v int64) *ClinkListCdrIbsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetQno(v string) *ClinkListCdrIbsRequest {
	s.Qno = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetQueueAnswerInTime(v int64) *ClinkListCdrIbsRequest {
	s.QueueAnswerInTime = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetResourceOwnerAccount(v string) *ClinkListCdrIbsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetResourceOwnerId(v int64) *ClinkListCdrIbsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetStartTime(v int64) *ClinkListCdrIbsRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbsRequest) SetStatus(v int64) *ClinkListCdrIbsRequest {
	s.Status = &v
	return s
}

func (s *ClinkListCdrIbsRequest) Validate() error {
	return dara.Validate(s)
}
