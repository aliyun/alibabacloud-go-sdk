// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *AddAgentRequest
	GetAccount() *string
	SetAgentTag(v string) *AddAgentRequest
	GetAgentTag() *string
	SetExtensionPwd(v string) *AddAgentRequest
	GetExtensionPwd() *string
	SetName(v string) *AddAgentRequest
	GetName() *string
	SetOwnerId(v int64) *AddAgentRequest
	GetOwnerId() *int64
	SetPassword(v string) *AddAgentRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *AddAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAgentRequest
	GetResourceOwnerId() *int64
}

type AddAgentRequest struct {
	// 坐席账号名(必须唯一)
	//
	// This parameter is required.
	//
	// example:
	//
	// seat_001
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// 合作方唯一标识
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 分机密码
	//
	// This parameter is required.
	//
	// example:
	//
	// gwegwe23t32
	ExtensionPwd *string `json:"ExtensionPwd,omitempty" xml:"ExtensionPwd,omitempty"`
	// 坐席名称
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 坐席账号密码 (可以跟分机密码一致)
	//
	// This parameter is required.
	//
	// example:
	//
	// gwegwe23t32
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAgentRequest) GoString() string {
	return s.String()
}

func (s *AddAgentRequest) GetAccount() *string {
	return s.Account
}

func (s *AddAgentRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AddAgentRequest) GetExtensionPwd() *string {
	return s.ExtensionPwd
}

func (s *AddAgentRequest) GetName() *string {
	return s.Name
}

func (s *AddAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAgentRequest) GetPassword() *string {
	return s.Password
}

func (s *AddAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAgentRequest) SetAccount(v string) *AddAgentRequest {
	s.Account = &v
	return s
}

func (s *AddAgentRequest) SetAgentTag(v string) *AddAgentRequest {
	s.AgentTag = &v
	return s
}

func (s *AddAgentRequest) SetExtensionPwd(v string) *AddAgentRequest {
	s.ExtensionPwd = &v
	return s
}

func (s *AddAgentRequest) SetName(v string) *AddAgentRequest {
	s.Name = &v
	return s
}

func (s *AddAgentRequest) SetOwnerId(v int64) *AddAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAgentRequest) SetPassword(v string) *AddAgentRequest {
	s.Password = &v
	return s
}

func (s *AddAgentRequest) SetResourceOwnerAccount(v string) *AddAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAgentRequest) SetResourceOwnerId(v int64) *AddAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAgentRequest) Validate() error {
	return dara.Validate(s)
}
