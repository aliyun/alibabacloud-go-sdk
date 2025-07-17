// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceLinkedRole interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *ServiceLinkedRole
	GetArn() *string
	SetAssumeRolePolicyDocument(v string) *ServiceLinkedRole
	GetAssumeRolePolicyDocument() *string
	SetCreateDate(v string) *ServiceLinkedRole
	GetCreateDate() *string
	SetDescription(v string) *ServiceLinkedRole
	GetDescription() *string
	SetIsServiceLinkedRole(v bool) *ServiceLinkedRole
	GetIsServiceLinkedRole() *bool
	SetRoleId(v string) *ServiceLinkedRole
	GetRoleId() *string
	SetRoleName(v string) *ServiceLinkedRole
	GetRoleName() *string
	SetRolePrincipalName(v string) *ServiceLinkedRole
	GetRolePrincipalName() *string
}

type ServiceLinkedRole struct {
	Arn                      *string `json:"arn,omitempty" xml:"arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty" xml:"assumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	Description              *string `json:"description,omitempty" xml:"description,omitempty"`
	IsServiceLinkedRole      *bool   `json:"isServiceLinkedRole,omitempty" xml:"isServiceLinkedRole,omitempty"`
	RoleId                   *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName                 *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	RolePrincipalName        *string `json:"rolePrincipalName,omitempty" xml:"rolePrincipalName,omitempty"`
}

func (s ServiceLinkedRole) String() string {
	return dara.Prettify(s)
}

func (s ServiceLinkedRole) GoString() string {
	return s.String()
}

func (s *ServiceLinkedRole) GetArn() *string {
	return s.Arn
}

func (s *ServiceLinkedRole) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *ServiceLinkedRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ServiceLinkedRole) GetDescription() *string {
	return s.Description
}

func (s *ServiceLinkedRole) GetIsServiceLinkedRole() *bool {
	return s.IsServiceLinkedRole
}

func (s *ServiceLinkedRole) GetRoleId() *string {
	return s.RoleId
}

func (s *ServiceLinkedRole) GetRoleName() *string {
	return s.RoleName
}

func (s *ServiceLinkedRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *ServiceLinkedRole) SetArn(v string) *ServiceLinkedRole {
	s.Arn = &v
	return s
}

func (s *ServiceLinkedRole) SetAssumeRolePolicyDocument(v string) *ServiceLinkedRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *ServiceLinkedRole) SetCreateDate(v string) *ServiceLinkedRole {
	s.CreateDate = &v
	return s
}

func (s *ServiceLinkedRole) SetDescription(v string) *ServiceLinkedRole {
	s.Description = &v
	return s
}

func (s *ServiceLinkedRole) SetIsServiceLinkedRole(v bool) *ServiceLinkedRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ServiceLinkedRole) SetRoleId(v string) *ServiceLinkedRole {
	s.RoleId = &v
	return s
}

func (s *ServiceLinkedRole) SetRoleName(v string) *ServiceLinkedRole {
	s.RoleName = &v
	return s
}

func (s *ServiceLinkedRole) SetRolePrincipalName(v string) *ServiceLinkedRole {
	s.RolePrincipalName = &v
	return s
}

func (s *ServiceLinkedRole) Validate() error {
	return dara.Validate(s)
}
