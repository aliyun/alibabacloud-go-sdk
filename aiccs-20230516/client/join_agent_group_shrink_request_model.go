// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinAgentGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *JoinAgentGroupShrinkRequest
	GetAgentGroupId() *int64
	SetAgentIdsShrink(v string) *JoinAgentGroupShrinkRequest
	GetAgentIdsShrink() *string
	SetOwnerId(v int64) *JoinAgentGroupShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *JoinAgentGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *JoinAgentGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type JoinAgentGroupShrinkRequest struct {
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
	AgentIdsShrink       *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s JoinAgentGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinAgentGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *JoinAgentGroupShrinkRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *JoinAgentGroupShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *JoinAgentGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *JoinAgentGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *JoinAgentGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *JoinAgentGroupShrinkRequest) SetAgentGroupId(v int64) *JoinAgentGroupShrinkRequest {
	s.AgentGroupId = &v
	return s
}

func (s *JoinAgentGroupShrinkRequest) SetAgentIdsShrink(v string) *JoinAgentGroupShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *JoinAgentGroupShrinkRequest) SetOwnerId(v int64) *JoinAgentGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *JoinAgentGroupShrinkRequest) SetResourceOwnerAccount(v string) *JoinAgentGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JoinAgentGroupShrinkRequest) SetResourceOwnerId(v int64) *JoinAgentGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JoinAgentGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
