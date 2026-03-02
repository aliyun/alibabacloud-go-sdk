// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegePopCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreatePrivilegePopCmd
	GetAccountId() *string
	SetActions(v string) *CreatePrivilegePopCmd
	GetActions() *string
	SetResource(v string) *CreatePrivilegePopCmd
	GetResource() *string
	SetRoleId(v int64) *CreatePrivilegePopCmd
	GetRoleId() *int64
}

type CreatePrivilegePopCmd struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Actions   *string `json:"actions,omitempty" xml:"actions,omitempty"`
	Resource  *string `json:"resource,omitempty" xml:"resource,omitempty"`
	RoleId    *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s CreatePrivilegePopCmd) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegePopCmd) GoString() string {
	return s.String()
}

func (s *CreatePrivilegePopCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePrivilegePopCmd) GetActions() *string {
	return s.Actions
}

func (s *CreatePrivilegePopCmd) GetResource() *string {
	return s.Resource
}

func (s *CreatePrivilegePopCmd) GetRoleId() *int64 {
	return s.RoleId
}

func (s *CreatePrivilegePopCmd) SetAccountId(v string) *CreatePrivilegePopCmd {
	s.AccountId = &v
	return s
}

func (s *CreatePrivilegePopCmd) SetActions(v string) *CreatePrivilegePopCmd {
	s.Actions = &v
	return s
}

func (s *CreatePrivilegePopCmd) SetResource(v string) *CreatePrivilegePopCmd {
	s.Resource = &v
	return s
}

func (s *CreatePrivilegePopCmd) SetRoleId(v int64) *CreatePrivilegePopCmd {
	s.RoleId = &v
	return s
}

func (s *CreatePrivilegePopCmd) Validate() error {
	return dara.Validate(s)
}
