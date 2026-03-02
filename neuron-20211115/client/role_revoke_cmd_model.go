// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoleRevokeCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *RoleRevokeCmd
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *RoleRevokeCmd
	GetAuthorizerType() *string
	SetRoleId(v int64) *RoleRevokeCmd
	GetRoleId() *int64
}

type RoleRevokeCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizerId *string `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// developer
	AuthorizerType *string `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s RoleRevokeCmd) String() string {
	return dara.Prettify(s)
}

func (s RoleRevokeCmd) GoString() string {
	return s.String()
}

func (s *RoleRevokeCmd) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *RoleRevokeCmd) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *RoleRevokeCmd) GetRoleId() *int64 {
	return s.RoleId
}

func (s *RoleRevokeCmd) SetAuthorizerId(v string) *RoleRevokeCmd {
	s.AuthorizerId = &v
	return s
}

func (s *RoleRevokeCmd) SetAuthorizerType(v string) *RoleRevokeCmd {
	s.AuthorizerType = &v
	return s
}

func (s *RoleRevokeCmd) SetRoleId(v int64) *RoleRevokeCmd {
	s.RoleId = &v
	return s
}

func (s *RoleRevokeCmd) Validate() error {
	return dara.Validate(s)
}
