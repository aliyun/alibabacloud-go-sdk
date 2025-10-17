// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfigItem interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *CredentialConfigItem
	GetKey() *string
	SetRoles(v []*CredentialRole) *CredentialConfigItem
	GetRoles() []*CredentialRole
	SetType(v string) *CredentialConfigItem
	GetType() *string
}

type CredentialConfigItem struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// if can be null:
	// true
	Roles []*CredentialRole `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	Type  *string           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CredentialConfigItem) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigItem) GoString() string {
	return s.String()
}

func (s *CredentialConfigItem) GetKey() *string {
	return s.Key
}

func (s *CredentialConfigItem) GetRoles() []*CredentialRole {
	return s.Roles
}

func (s *CredentialConfigItem) GetType() *string {
	return s.Type
}

func (s *CredentialConfigItem) SetKey(v string) *CredentialConfigItem {
	s.Key = &v
	return s
}

func (s *CredentialConfigItem) SetRoles(v []*CredentialRole) *CredentialConfigItem {
	s.Roles = v
	return s
}

func (s *CredentialConfigItem) SetType(v string) *CredentialConfigItem {
	s.Type = &v
	return s
}

func (s *CredentialConfigItem) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
