// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialInput interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*string) *UpdateCredentialInput
	GetConfig() map[string]*string
	SetDescription(v string) *UpdateCredentialInput
	GetDescription() *string
	SetName(v string) *UpdateCredentialInput
	GetName() *string
	SetSecret(v string) *UpdateCredentialInput
	GetSecret() *string
	SetType(v string) *UpdateCredentialInput
	GetType() *string
}

type UpdateCredentialInput struct {
	Config      map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty"`
	Secret      *string            `json:"secret,omitempty" xml:"secret,omitempty"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCredentialInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialInput) GoString() string {
	return s.String()
}

func (s *UpdateCredentialInput) GetConfig() map[string]*string {
	return s.Config
}

func (s *UpdateCredentialInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialInput) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialInput) GetSecret() *string {
	return s.Secret
}

func (s *UpdateCredentialInput) GetType() *string {
	return s.Type
}

func (s *UpdateCredentialInput) SetConfig(v map[string]*string) *UpdateCredentialInput {
	s.Config = v
	return s
}

func (s *UpdateCredentialInput) SetDescription(v string) *UpdateCredentialInput {
	s.Description = &v
	return s
}

func (s *UpdateCredentialInput) SetName(v string) *UpdateCredentialInput {
	s.Name = &v
	return s
}

func (s *UpdateCredentialInput) SetSecret(v string) *UpdateCredentialInput {
	s.Secret = &v
	return s
}

func (s *UpdateCredentialInput) SetType(v string) *UpdateCredentialInput {
	s.Type = &v
	return s
}

func (s *UpdateCredentialInput) Validate() error {
	return dara.Validate(s)
}
