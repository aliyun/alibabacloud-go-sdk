// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialPublicConfig(v *CredentialPublicConfig) *UpdateCredentialInput
	GetCredentialPublicConfig() *CredentialPublicConfig
	SetCredentialSecret(v string) *UpdateCredentialInput
	GetCredentialSecret() *string
	SetDescription(v string) *UpdateCredentialInput
	GetDescription() *string
	SetEnabled(v bool) *UpdateCredentialInput
	GetEnabled() *bool
}

type UpdateCredentialInput struct {
	CredentialPublicConfig *CredentialPublicConfig `json:"credentialPublicConfig,omitempty" xml:"credentialPublicConfig,omitempty"`
	CredentialSecret       *string                 `json:"credentialSecret,omitempty" xml:"credentialSecret,omitempty"`
	Description            *string                 `json:"description,omitempty" xml:"description,omitempty"`
	Enabled                *bool                   `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCredentialInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialInput) GoString() string {
	return s.String()
}

func (s *UpdateCredentialInput) GetCredentialPublicConfig() *CredentialPublicConfig {
	return s.CredentialPublicConfig
}

func (s *UpdateCredentialInput) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *UpdateCredentialInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCredentialInput) SetCredentialPublicConfig(v *CredentialPublicConfig) *UpdateCredentialInput {
	s.CredentialPublicConfig = v
	return s
}

func (s *UpdateCredentialInput) SetCredentialSecret(v string) *UpdateCredentialInput {
	s.CredentialSecret = &v
	return s
}

func (s *UpdateCredentialInput) SetDescription(v string) *UpdateCredentialInput {
	s.Description = &v
	return s
}

func (s *UpdateCredentialInput) SetEnabled(v bool) *UpdateCredentialInput {
	s.Enabled = &v
	return s
}

func (s *UpdateCredentialInput) Validate() error {
	if s.CredentialPublicConfig != nil {
		if err := s.CredentialPublicConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
