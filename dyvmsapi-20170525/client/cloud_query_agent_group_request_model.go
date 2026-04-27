// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudQueryAgentGroupRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudQueryAgentGroupRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudQueryAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudQueryAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudQueryAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudQueryAgentGroupRequest struct {
	// 座席工号
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudQueryAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentGroupRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudQueryAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudQueryAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudQueryAgentGroupRequest) SetCno(v string) *CloudQueryAgentGroupRequest {
	s.Cno = &v
	return s
}

func (s *CloudQueryAgentGroupRequest) SetEnterpriseId(v int64) *CloudQueryAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentGroupRequest) SetOwnerId(v int64) *CloudQueryAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudQueryAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudQueryAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudQueryAgentGroupRequest) SetResourceOwnerId(v int64) *CloudQueryAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudQueryAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
