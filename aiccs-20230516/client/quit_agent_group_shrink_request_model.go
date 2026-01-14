// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuitAgentGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *QuitAgentGroupShrinkRequest
	GetAgentGroupId() *int64
	SetAgentIdsShrink(v string) *QuitAgentGroupShrinkRequest
	GetAgentIdsShrink() *string
	SetOwnerId(v int64) *QuitAgentGroupShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuitAgentGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuitAgentGroupShrinkRequest
	GetResourceOwnerId() *int64
}

type QuitAgentGroupShrinkRequest struct {
	// 坐席组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 坐席ID列表
	//
	// This parameter is required.
	AgentIdsShrink       *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuitAgentGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuitAgentGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuitAgentGroupShrinkRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *QuitAgentGroupShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *QuitAgentGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuitAgentGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuitAgentGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuitAgentGroupShrinkRequest) SetAgentGroupId(v int64) *QuitAgentGroupShrinkRequest {
	s.AgentGroupId = &v
	return s
}

func (s *QuitAgentGroupShrinkRequest) SetAgentIdsShrink(v string) *QuitAgentGroupShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *QuitAgentGroupShrinkRequest) SetOwnerId(v int64) *QuitAgentGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuitAgentGroupShrinkRequest) SetResourceOwnerAccount(v string) *QuitAgentGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuitAgentGroupShrinkRequest) SetResourceOwnerId(v int64) *QuitAgentGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuitAgentGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
