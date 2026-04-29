// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudListAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudListAgentGroupRequest
	GetGno() *string
	SetGroupName(v string) *CloudListAgentGroupRequest
	GetGroupName() *string
	SetLimit(v string) *CloudListAgentGroupRequest
	GetLimit() *string
	SetOwnerId(v int64) *CloudListAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListAgentGroupRequest
	GetResourceOwnerId() *int64
	SetStart(v string) *CloudListAgentGroupRequest
	GetStart() *string
}

type CloudListAgentGroupRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼总组编号；说明：以外呼组编号为筛选条件查询
	//
	// example:
	//
	// WH123
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 外呼总组名称；说明：以外呼组名称为筛选条件查询
	//
	// example:
	//
	// gpName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 需要取出的数据条数；说明：大于0，最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit                *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 从第几条开始取；说明：大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s CloudListAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudListAgentGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudListAgentGroupRequest) GetLimit() *string {
	return s.Limit
}

func (s *CloudListAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListAgentGroupRequest) GetStart() *string {
	return s.Start
}

func (s *CloudListAgentGroupRequest) SetEnterpriseId(v int64) *CloudListAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetGno(v string) *CloudListAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetGroupName(v string) *CloudListAgentGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetLimit(v string) *CloudListAgentGroupRequest {
	s.Limit = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetOwnerId(v int64) *CloudListAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudListAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetResourceOwnerId(v int64) *CloudListAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListAgentGroupRequest) SetStart(v string) *CloudListAgentGroupRequest {
	s.Start = &v
	return s
}

func (s *CloudListAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
