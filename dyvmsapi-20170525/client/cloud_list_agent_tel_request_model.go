// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentTelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudListAgentTelRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudListAgentTelRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudListAgentTelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListAgentTelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListAgentTelRequest
	GetResourceOwnerId() *int64
	SetTel(v string) *CloudListAgentTelRequest
	GetTel() *string
}

type CloudListAgentTelRequest struct {
	// 座席工号；3-10位数字
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
	// 7122600
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 座席电话
	//
	// example:
	//
	// 41008502
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s CloudListAgentTelRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentTelRequest) GoString() string {
	return s.String()
}

func (s *CloudListAgentTelRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudListAgentTelRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListAgentTelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListAgentTelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListAgentTelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListAgentTelRequest) GetTel() *string {
	return s.Tel
}

func (s *CloudListAgentTelRequest) SetCno(v string) *CloudListAgentTelRequest {
	s.Cno = &v
	return s
}

func (s *CloudListAgentTelRequest) SetEnterpriseId(v int64) *CloudListAgentTelRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAgentTelRequest) SetOwnerId(v int64) *CloudListAgentTelRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListAgentTelRequest) SetResourceOwnerAccount(v string) *CloudListAgentTelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListAgentTelRequest) SetResourceOwnerId(v int64) *CloudListAgentTelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListAgentTelRequest) SetTel(v string) *CloudListAgentTelRequest {
	s.Tel = &v
	return s
}

func (s *CloudListAgentTelRequest) Validate() error {
	return dara.Validate(s)
}
