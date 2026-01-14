// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *JoinAgentGroupRequest
	GetAgentGroupId() *int64
	SetAgentIds(v []*int64) *JoinAgentGroupRequest
	GetAgentIds() []*int64
	SetOwnerId(v int64) *JoinAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *JoinAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *JoinAgentGroupRequest
	GetResourceOwnerId() *int64
}

type JoinAgentGroupRequest struct {
	// 坐席组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 坐席ID列表
	//
	// This parameter is required.
	AgentIds             []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s JoinAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinAgentGroupRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *JoinAgentGroupRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *JoinAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *JoinAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *JoinAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *JoinAgentGroupRequest) SetAgentGroupId(v int64) *JoinAgentGroupRequest {
	s.AgentGroupId = &v
	return s
}

func (s *JoinAgentGroupRequest) SetAgentIds(v []*int64) *JoinAgentGroupRequest {
	s.AgentIds = v
	return s
}

func (s *JoinAgentGroupRequest) SetOwnerId(v int64) *JoinAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *JoinAgentGroupRequest) SetResourceOwnerAccount(v string) *JoinAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JoinAgentGroupRequest) SetResourceOwnerId(v int64) *JoinAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JoinAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
