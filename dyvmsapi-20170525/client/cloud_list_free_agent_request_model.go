// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListFreeAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudListFreeAgentRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudListFreeAgentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListFreeAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListFreeAgentRequest
	GetResourceOwnerId() *int64
}

type CloudListFreeAgentRequest struct {
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

func (s CloudListFreeAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListFreeAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudListFreeAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListFreeAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListFreeAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListFreeAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListFreeAgentRequest) SetEnterpriseId(v int64) *CloudListFreeAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListFreeAgentRequest) SetOwnerId(v int64) *CloudListFreeAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListFreeAgentRequest) SetResourceOwnerAccount(v string) *CloudListFreeAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListFreeAgentRequest) SetResourceOwnerId(v int64) *CloudListFreeAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListFreeAgentRequest) Validate() error {
	return dara.Validate(s)
}
