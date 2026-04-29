// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAssignedAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *CloudListAssignedAgentGroupRequest
	GetCname() *string
	SetCno(v string) *CloudListAssignedAgentGroupRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudListAssignedAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudListAssignedAgentGroupRequest
	GetGno() *string
	SetOwnerId(v int64) *CloudListAssignedAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListAssignedAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListAssignedAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudListAssignedAgentGroupRequest struct {
	// 查询条件座席名称；说明：以座席名称为筛选条件查询
	//
	// example:
	//
	// 111111111
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// 查询条件座席编号；说明：以座席编号为筛选条件查询
	//
	// example:
	//
	// 111111111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号；说明：获取此gno下绑定的座席信息列表
	//
	// This parameter is required.
	//
	// example:
	//
	// WH113
	Gno                  *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudListAssignedAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListAssignedAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudListAssignedAgentGroupRequest) GetCname() *string {
	return s.Cname
}

func (s *CloudListAssignedAgentGroupRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudListAssignedAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListAssignedAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudListAssignedAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListAssignedAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListAssignedAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListAssignedAgentGroupRequest) SetCname(v string) *CloudListAssignedAgentGroupRequest {
	s.Cname = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetCno(v string) *CloudListAssignedAgentGroupRequest {
	s.Cno = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetEnterpriseId(v int64) *CloudListAssignedAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetGno(v string) *CloudListAssignedAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetOwnerId(v int64) *CloudListAssignedAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudListAssignedAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) SetResourceOwnerId(v int64) *CloudListAssignedAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListAssignedAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
