// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTokenVaultRequest
	GetDescription() *string
	SetEncryptionConfig(v *EncryptionConfig) *CreateTokenVaultRequest
	GetEncryptionConfig() *EncryptionConfig
	SetRoleArn(v string) *CreateTokenVaultRequest
	GetRoleArn() *string
	SetTokenVaultName(v string) *CreateTokenVaultRequest
	GetTokenVaultName() *string
}

type CreateTokenVaultRequest struct {
	// example:
	//
	// example description
	Description      *string           `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptionConfig *EncryptionConfig `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty"`
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateTokenVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTokenVaultRequest) GetEncryptionConfig() *EncryptionConfig {
	return s.EncryptionConfig
}

func (s *CreateTokenVaultRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTokenVaultRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateTokenVaultRequest) SetDescription(v string) *CreateTokenVaultRequest {
	s.Description = &v
	return s
}

func (s *CreateTokenVaultRequest) SetEncryptionConfig(v *EncryptionConfig) *CreateTokenVaultRequest {
	s.EncryptionConfig = v
	return s
}

func (s *CreateTokenVaultRequest) SetRoleArn(v string) *CreateTokenVaultRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateTokenVaultRequest) SetTokenVaultName(v string) *CreateTokenVaultRequest {
	s.TokenVaultName = &v
	return s
}

func (s *CreateTokenVaultRequest) Validate() error {
	if s.EncryptionConfig != nil {
		if err := s.EncryptionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
