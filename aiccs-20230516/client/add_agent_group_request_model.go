// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupName(v string) *AddAgentGroupRequest
	GetAgentGroupName() *string
	SetOwnerId(v int64) *AddAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAgentGroupRequest
	GetResourceOwnerId() *int64
}

type AddAgentGroupRequest struct {
	// 坐席组名称，不能为空且最大长度为20个字符
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	AgentGroupName       *string `json:"AgentGroupName,omitempty" xml:"AgentGroupName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *AddAgentGroupRequest) GetAgentGroupName() *string {
	return s.AgentGroupName
}

func (s *AddAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAgentGroupRequest) SetAgentGroupName(v string) *AddAgentGroupRequest {
	s.AgentGroupName = &v
	return s
}

func (s *AddAgentGroupRequest) SetOwnerId(v int64) *AddAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAgentGroupRequest) SetResourceOwnerAccount(v string) *AddAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAgentGroupRequest) SetResourceOwnerId(v int64) *AddAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
