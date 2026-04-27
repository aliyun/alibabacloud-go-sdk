// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudGetAgentRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudGetAgentRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudGetAgentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudGetAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudGetAgentRequest
	GetResourceOwnerId() *int64
}

type CloudGetAgentRequest struct {
	// 座席工号；查询指定工号对应的座席信息
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

func (s CloudGetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudGetAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudGetAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudGetAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudGetAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudGetAgentRequest) SetCno(v string) *CloudGetAgentRequest {
	s.Cno = &v
	return s
}

func (s *CloudGetAgentRequest) SetEnterpriseId(v int64) *CloudGetAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAgentRequest) SetOwnerId(v int64) *CloudGetAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudGetAgentRequest) SetResourceOwnerAccount(v string) *CloudGetAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudGetAgentRequest) SetResourceOwnerId(v int64) *CloudGetAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudGetAgentRequest) Validate() error {
	return dara.Validate(s)
}
