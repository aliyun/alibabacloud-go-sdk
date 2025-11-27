// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfig interface {
	dara.Model
	String() string
	GoString() string
	SetChain(v []*CredentialConfigChain) *CredentialConfig
	GetChain() []*CredentialConfigChain
	SetPolicy(v string) *CredentialConfig
	GetPolicy() *string
	SetServiceRole(v string) *CredentialConfig
	GetServiceRole() *string
}

type CredentialConfig struct {
	Chain       []*CredentialConfigChain `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	Policy      *string                  `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ServiceRole *string                  `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s CredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) GetChain() []*CredentialConfigChain {
	return s.Chain
}

func (s *CredentialConfig) GetPolicy() *string {
	return s.Policy
}

func (s *CredentialConfig) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CredentialConfig) SetChain(v []*CredentialConfigChain) *CredentialConfig {
	s.Chain = v
	return s
}

func (s *CredentialConfig) SetPolicy(v string) *CredentialConfig {
	s.Policy = &v
	return s
}

func (s *CredentialConfig) SetServiceRole(v string) *CredentialConfig {
	s.ServiceRole = &v
	return s
}

func (s *CredentialConfig) Validate() error {
	if s.Chain != nil {
		for _, item := range s.Chain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CredentialConfigChain struct {
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CredentialConfigChain) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigChain) GoString() string {
	return s.String()
}

func (s *CredentialConfigChain) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CredentialConfigChain) GetRole() *string {
	return s.Role
}

func (s *CredentialConfigChain) GetRoleType() *string {
	return s.RoleType
}

func (s *CredentialConfigChain) SetAssumeRoleFor(v string) *CredentialConfigChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialConfigChain) SetRole(v string) *CredentialConfigChain {
	s.Role = &v
	return s
}

func (s *CredentialConfigChain) SetRoleType(v string) *CredentialConfigChain {
	s.RoleType = &v
	return s
}

func (s *CredentialConfigChain) Validate() error {
	return dara.Validate(s)
}
