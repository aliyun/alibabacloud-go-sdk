// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *UpdateAgentStatusRequest
	GetAgentId() *int64
	SetAgentStatus(v int64) *UpdateAgentStatusRequest
	GetAgentStatus() *int64
	SetAgentTag(v string) *UpdateAgentStatusRequest
	GetAgentTag() *string
	SetOwnerId(v int64) *UpdateAgentStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateAgentStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAgentStatusRequest
	GetResourceOwnerId() *int64
}

type UpdateAgentStatusRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 58
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席状态 1:在线；2:忙碌；3:小休；4:离线
	//
	// example:
	//
	// 1
	AgentStatus *int64 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abac
	AgentTag             *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *UpdateAgentStatusRequest) GetAgentStatus() *int64 {
	return s.AgentStatus
}

func (s *UpdateAgentStatusRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *UpdateAgentStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAgentStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAgentStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAgentStatusRequest) SetAgentId(v int64) *UpdateAgentStatusRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetAgentStatus(v int64) *UpdateAgentStatusRequest {
	s.AgentStatus = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetAgentTag(v string) *UpdateAgentStatusRequest {
	s.AgentTag = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetOwnerId(v int64) *UpdateAgentStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetResourceOwnerAccount(v string) *UpdateAgentStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetResourceOwnerId(v int64) *UpdateAgentStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
