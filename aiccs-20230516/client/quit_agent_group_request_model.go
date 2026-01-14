// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuitAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *QuitAgentGroupRequest
	GetAgentGroupId() *int64
	SetAgentIds(v []*int64) *QuitAgentGroupRequest
	GetAgentIds() []*int64
	SetOwnerId(v int64) *QuitAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuitAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuitAgentGroupRequest
	GetResourceOwnerId() *int64
}

type QuitAgentGroupRequest struct {
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
	AgentIds             []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuitAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QuitAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *QuitAgentGroupRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *QuitAgentGroupRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *QuitAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuitAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuitAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuitAgentGroupRequest) SetAgentGroupId(v int64) *QuitAgentGroupRequest {
	s.AgentGroupId = &v
	return s
}

func (s *QuitAgentGroupRequest) SetAgentIds(v []*int64) *QuitAgentGroupRequest {
	s.AgentIds = v
	return s
}

func (s *QuitAgentGroupRequest) SetOwnerId(v int64) *QuitAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *QuitAgentGroupRequest) SetResourceOwnerAccount(v string) *QuitAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuitAgentGroupRequest) SetResourceOwnerId(v int64) *QuitAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuitAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
