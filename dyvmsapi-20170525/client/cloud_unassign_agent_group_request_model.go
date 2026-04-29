// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUnassignAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudUnassignAgentGroupRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudUnassignAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudUnassignAgentGroupRequest
	GetGno() *string
	SetOwnerId(v int64) *CloudUnassignAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudUnassignAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudUnassignAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudUnassignAgentGroupRequest struct {
	// 座席编号；说明：解除此cno对应座席 与 gno对应外呼组 的绑定关系
	//
	// This parameter is required.
	//
	// example:
	//
	// 12221
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号；说明：解除此gno对应外呼组 与 cno对应座席 的绑定关系
	//
	// This parameter is required.
	//
	// example:
	//
	// WH123
	Gno                  *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudUnassignAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudUnassignAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudUnassignAgentGroupRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudUnassignAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudUnassignAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudUnassignAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudUnassignAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudUnassignAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudUnassignAgentGroupRequest) SetCno(v string) *CloudUnassignAgentGroupRequest {
	s.Cno = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) SetEnterpriseId(v int64) *CloudUnassignAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) SetGno(v string) *CloudUnassignAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) SetOwnerId(v int64) *CloudUnassignAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudUnassignAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) SetResourceOwnerId(v int64) *CloudUnassignAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudUnassignAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
