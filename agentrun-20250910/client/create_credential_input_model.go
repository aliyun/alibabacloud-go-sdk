// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialInput interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*string) *CreateCredentialInput
	GetConfig() map[string]*string
	SetDescription(v string) *CreateCredentialInput
	GetDescription() *string
	SetName(v string) *CreateCredentialInput
	GetName() *string
	SetSecret(v string) *CreateCredentialInput
	GetSecret() *string
	SetType(v string) *CreateCredentialInput
	GetType() *string
}

type CreateCredentialInput struct {
	Config      map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty"`
	Secret      *string            `json:"secret,omitempty" xml:"secret,omitempty"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCredentialInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialInput) GoString() string {
	return s.String()
}

func (s *CreateCredentialInput) GetConfig() map[string]*string {
	return s.Config
}

func (s *CreateCredentialInput) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialInput) GetName() *string {
	return s.Name
}

func (s *CreateCredentialInput) GetSecret() *string {
	return s.Secret
}

func (s *CreateCredentialInput) GetType() *string {
	return s.Type
}

func (s *CreateCredentialInput) SetConfig(v map[string]*string) *CreateCredentialInput {
	s.Config = v
	return s
}

func (s *CreateCredentialInput) SetDescription(v string) *CreateCredentialInput {
	s.Description = &v
	return s
}

func (s *CreateCredentialInput) SetName(v string) *CreateCredentialInput {
	s.Name = &v
	return s
}

func (s *CreateCredentialInput) SetSecret(v string) *CreateCredentialInput {
	s.Secret = &v
	return s
}

func (s *CreateCredentialInput) SetType(v string) *CreateCredentialInput {
	s.Type = &v
	return s
}

func (s *CreateCredentialInput) Validate() error {
	return dara.Validate(s)
}
