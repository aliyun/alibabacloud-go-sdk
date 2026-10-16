// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTokenVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTokenVaultRequest
	GetDescription() *string
	SetRoleArn(v string) *UpdateTokenVaultRequest
	GetRoleArn() *string
	SetTokenVaultName(v string) *UpdateTokenVaultRequest
	GetTokenVaultName() *string
}

type UpdateTokenVaultRequest struct {
	// The description.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ARN of the RAM role used to grant the AgentIdentity service access to other cloud resources such as KMS.
	//
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The name of the token vault. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s UpdateTokenVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTokenVaultRequest) GoString() string {
	return s.String()
}

func (s *UpdateTokenVaultRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTokenVaultRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateTokenVaultRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *UpdateTokenVaultRequest) SetDescription(v string) *UpdateTokenVaultRequest {
	s.Description = &v
	return s
}

func (s *UpdateTokenVaultRequest) SetRoleArn(v string) *UpdateTokenVaultRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateTokenVaultRequest) SetTokenVaultName(v string) *UpdateTokenVaultRequest {
	s.TokenVaultName = &v
	return s
}

func (s *UpdateTokenVaultRequest) Validate() error {
	return dara.Validate(s)
}
