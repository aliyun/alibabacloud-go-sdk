// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudGetAgentStatusRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudGetAgentStatusRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudGetAgentStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudGetAgentStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudGetAgentStatusRequest
	GetResourceOwnerId() *int64
}

type CloudGetAgentStatusRequest struct {
	// 单个座席号
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

func (s CloudGetAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *CloudGetAgentStatusRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudGetAgentStatusRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetAgentStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudGetAgentStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudGetAgentStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudGetAgentStatusRequest) SetCno(v string) *CloudGetAgentStatusRequest {
	s.Cno = &v
	return s
}

func (s *CloudGetAgentStatusRequest) SetEnterpriseId(v int64) *CloudGetAgentStatusRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAgentStatusRequest) SetOwnerId(v int64) *CloudGetAgentStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudGetAgentStatusRequest) SetResourceOwnerAccount(v string) *CloudGetAgentStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudGetAgentStatusRequest) SetResourceOwnerId(v int64) *CloudGetAgentStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudGetAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
