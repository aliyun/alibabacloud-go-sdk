// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenVaultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTokenVaultShrinkRequest
	GetDescription() *string
	SetEncryptionConfigShrink(v string) *CreateTokenVaultShrinkRequest
	GetEncryptionConfigShrink() *string
	SetRoleArn(v string) *CreateTokenVaultShrinkRequest
	GetRoleArn() *string
	SetTokenVaultName(v string) *CreateTokenVaultShrinkRequest
	GetTokenVaultName() *string
}

type CreateTokenVaultShrinkRequest struct {
	// example:
	//
	// example description
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptionConfigShrink *string `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty"`
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateTokenVaultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTokenVaultShrinkRequest) GetEncryptionConfigShrink() *string {
	return s.EncryptionConfigShrink
}

func (s *CreateTokenVaultShrinkRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTokenVaultShrinkRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateTokenVaultShrinkRequest) SetDescription(v string) *CreateTokenVaultShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTokenVaultShrinkRequest) SetEncryptionConfigShrink(v string) *CreateTokenVaultShrinkRequest {
	s.EncryptionConfigShrink = &v
	return s
}

func (s *CreateTokenVaultShrinkRequest) SetRoleArn(v string) *CreateTokenVaultShrinkRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateTokenVaultShrinkRequest) SetTokenVaultName(v string) *CreateTokenVaultShrinkRequest {
	s.TokenVaultName = &v
	return s
}

func (s *CreateTokenVaultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
