// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialAuthType(v string) *CreateCredentialInput
	GetCredentialAuthType() *string
	SetCredentialName(v string) *CreateCredentialInput
	GetCredentialName() *string
	SetCredentialPublicConfig(v *CredentialPublicConfig) *CreateCredentialInput
	GetCredentialPublicConfig() *CredentialPublicConfig
	SetCredentialSecret(v string) *CreateCredentialInput
	GetCredentialSecret() *string
	SetCredentialSourceType(v string) *CreateCredentialInput
	GetCredentialSourceType() *string
	SetDescription(v string) *CreateCredentialInput
	GetDescription() *string
	SetEnabled(v bool) *CreateCredentialInput
	GetEnabled() *bool
}

type CreateCredentialInput struct {
	// This parameter is required.
	CredentialAuthType *string `json:"credentialAuthType,omitempty" xml:"credentialAuthType,omitempty"`
	// This parameter is required.
	CredentialName         *string                 `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	CredentialPublicConfig *CredentialPublicConfig `json:"credentialPublicConfig,omitempty" xml:"credentialPublicConfig,omitempty"`
	CredentialSecret       *string                 `json:"credentialSecret,omitempty" xml:"credentialSecret,omitempty"`
	// This parameter is required.
	CredentialSourceType *string `json:"credentialSourceType,omitempty" xml:"credentialSourceType,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateCredentialInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialInput) GoString() string {
	return s.String()
}

func (s *CreateCredentialInput) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *CreateCredentialInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateCredentialInput) GetCredentialPublicConfig() *CredentialPublicConfig {
	return s.CredentialPublicConfig
}

func (s *CreateCredentialInput) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *CreateCredentialInput) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *CreateCredentialInput) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCredentialInput) SetCredentialAuthType(v string) *CreateCredentialInput {
	s.CredentialAuthType = &v
	return s
}

func (s *CreateCredentialInput) SetCredentialName(v string) *CreateCredentialInput {
	s.CredentialName = &v
	return s
}

func (s *CreateCredentialInput) SetCredentialPublicConfig(v *CredentialPublicConfig) *CreateCredentialInput {
	s.CredentialPublicConfig = v
	return s
}

func (s *CreateCredentialInput) SetCredentialSecret(v string) *CreateCredentialInput {
	s.CredentialSecret = &v
	return s
}

func (s *CreateCredentialInput) SetCredentialSourceType(v string) *CreateCredentialInput {
	s.CredentialSourceType = &v
	return s
}

func (s *CreateCredentialInput) SetDescription(v string) *CreateCredentialInput {
	s.Description = &v
	return s
}

func (s *CreateCredentialInput) SetEnabled(v bool) *CreateCredentialInput {
	s.Enabled = &v
	return s
}

func (s *CreateCredentialInput) Validate() error {
	if s.CredentialPublicConfig != nil {
		if err := s.CredentialPublicConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
