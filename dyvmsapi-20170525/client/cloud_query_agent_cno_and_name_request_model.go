// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentCnoAndNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudQueryAgentCnoAndNameRequest
	GetCnos() *string
	SetEnterpriseId(v int64) *CloudQueryAgentCnoAndNameRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudQueryAgentCnoAndNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudQueryAgentCnoAndNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudQueryAgentCnoAndNameRequest
	GetResourceOwnerId() *int64
}

type CloudQueryAgentCnoAndNameRequest struct {
	// 座席号之间用,分隔 例cnos=2000,20001 最多支持1000个座席
	//
	// example:
	//
	// 2000,20001
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
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

func (s CloudQueryAgentCnoAndNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentCnoAndNameRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentCnoAndNameRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudQueryAgentCnoAndNameRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentCnoAndNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudQueryAgentCnoAndNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudQueryAgentCnoAndNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudQueryAgentCnoAndNameRequest) SetCnos(v string) *CloudQueryAgentCnoAndNameRequest {
	s.Cnos = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameRequest) SetEnterpriseId(v int64) *CloudQueryAgentCnoAndNameRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameRequest) SetOwnerId(v int64) *CloudQueryAgentCnoAndNameRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameRequest) SetResourceOwnerAccount(v string) *CloudQueryAgentCnoAndNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameRequest) SetResourceOwnerId(v int64) *CloudQueryAgentCnoAndNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameRequest) Validate() error {
	return dara.Validate(s)
}
