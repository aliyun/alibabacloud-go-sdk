// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *ClinkListClientsRequest
	GetActive() *int64
	SetBindTel(v string) *ClinkListClientsRequest
	GetBindTel() *string
	SetClid(v string) *ClinkListClientsRequest
	GetClid() *string
	SetEndTime(v int64) *ClinkListClientsRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *ClinkListClientsRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *ClinkListClientsRequest
	GetLimit() *int64
	SetOffset(v int64) *ClinkListClientsRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListClientsRequest
	GetOwnerId() *int64
	SetQno(v string) *ClinkListClientsRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *ClinkListClientsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListClientsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ClinkListClientsRequest
	GetStartTime() *int64
	SetUpdateEndTime(v int64) *ClinkListClientsRequest
	GetUpdateEndTime() *int64
	SetUpdateStartTime(v int64) *ClinkListClientsRequest
	GetUpdateStartTime() *int64
}

type ClinkListClientsRequest struct {
	// 是否激活，1: 激活；0: 未激活
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 绑定电话
	//
	// example:
	//
	// xxx
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 外显号码
	//
	// example:
	//
	// xxx
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 创建时间查询范围-结束时间,秒级时间戳
	//
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
	// 查询记录条数，默认值为 10，最大范围 100
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 偏移量，默认值为 0，最大范围 10000
	//
	// example:
	//
	// 0
	Offset  *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号
	//
	// example:
	//
	// 12112
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 创建时间查询范围-开始时间,秒级时间戳
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 修改时间查询范围-结束时间,秒级时间戳
	//
	// example:
	//
	// 1775024775
	UpdateEndTime *int64 `json:"UpdateEndTime,omitempty" xml:"UpdateEndTime,omitempty"`
	// 修改时间查询范围-开始时间,秒级时间戳
	//
	// example:
	//
	// 1775024775
	UpdateStartTime *int64 `json:"UpdateStartTime,omitempty" xml:"UpdateStartTime,omitempty"`
}

func (s ClinkListClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListClientsRequest) GoString() string {
	return s.String()
}

func (s *ClinkListClientsRequest) GetActive() *int64 {
	return s.Active
}

func (s *ClinkListClientsRequest) GetBindTel() *string {
	return s.BindTel
}

func (s *ClinkListClientsRequest) GetClid() *string {
	return s.Clid
}

func (s *ClinkListClientsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListClientsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListClientsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListClientsRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListClientsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListClientsRequest) GetQno() *string {
	return s.Qno
}

func (s *ClinkListClientsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListClientsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListClientsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListClientsRequest) GetUpdateEndTime() *int64 {
	return s.UpdateEndTime
}

func (s *ClinkListClientsRequest) GetUpdateStartTime() *int64 {
	return s.UpdateStartTime
}

func (s *ClinkListClientsRequest) SetActive(v int64) *ClinkListClientsRequest {
	s.Active = &v
	return s
}

func (s *ClinkListClientsRequest) SetBindTel(v string) *ClinkListClientsRequest {
	s.BindTel = &v
	return s
}

func (s *ClinkListClientsRequest) SetClid(v string) *ClinkListClientsRequest {
	s.Clid = &v
	return s
}

func (s *ClinkListClientsRequest) SetEndTime(v int64) *ClinkListClientsRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListClientsRequest) SetEnterpriseId(v int64) *ClinkListClientsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListClientsRequest) SetLimit(v int64) *ClinkListClientsRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListClientsRequest) SetOffset(v int64) *ClinkListClientsRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListClientsRequest) SetOwnerId(v int64) *ClinkListClientsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListClientsRequest) SetQno(v string) *ClinkListClientsRequest {
	s.Qno = &v
	return s
}

func (s *ClinkListClientsRequest) SetResourceOwnerAccount(v string) *ClinkListClientsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListClientsRequest) SetResourceOwnerId(v int64) *ClinkListClientsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListClientsRequest) SetStartTime(v int64) *ClinkListClientsRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListClientsRequest) SetUpdateEndTime(v int64) *ClinkListClientsRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *ClinkListClientsRequest) SetUpdateStartTime(v int64) *ClinkListClientsRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *ClinkListClientsRequest) Validate() error {
	return dara.Validate(s)
}
