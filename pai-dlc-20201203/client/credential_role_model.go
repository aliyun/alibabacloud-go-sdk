// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRole interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeRoleFor(v string) *CredentialRole
	GetAssumeRoleFor() *string
	SetAssumeUserInfo(v *AssumeUserInfo) *CredentialRole
	GetAssumeUserInfo() *AssumeUserInfo
	SetPolicy(v string) *CredentialRole
	GetPolicy() *string
	SetRoleArn(v string) *CredentialRole
	GetRoleArn() *string
	SetRoleType(v string) *CredentialRole
	GetRoleType() *string
}

type CredentialRole struct {
	AssumeRoleFor  *string         `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	AssumeUserInfo *AssumeUserInfo `json:"AssumeUserInfo,omitempty" xml:"AssumeUserInfo,omitempty"`
	Policy         *string         `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RoleArn        *string         `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType       *string         `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CredentialRole) String() string {
	return dara.Prettify(s)
}

func (s CredentialRole) GoString() string {
	return s.String()
}

func (s *CredentialRole) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CredentialRole) GetAssumeUserInfo() *AssumeUserInfo {
	return s.AssumeUserInfo
}

func (s *CredentialRole) GetPolicy() *string {
	return s.Policy
}

func (s *CredentialRole) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CredentialRole) GetRoleType() *string {
	return s.RoleType
}

func (s *CredentialRole) SetAssumeRoleFor(v string) *CredentialRole {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialRole) SetAssumeUserInfo(v *AssumeUserInfo) *CredentialRole {
	s.AssumeUserInfo = v
	return s
}

func (s *CredentialRole) SetPolicy(v string) *CredentialRole {
	s.Policy = &v
	return s
}

func (s *CredentialRole) SetRoleArn(v string) *CredentialRole {
	s.RoleArn = &v
	return s
}

func (s *CredentialRole) SetRoleType(v string) *CredentialRole {
	s.RoleType = &v
	return s
}

func (s *CredentialRole) Validate() error {
	return dara.Validate(s)
}
