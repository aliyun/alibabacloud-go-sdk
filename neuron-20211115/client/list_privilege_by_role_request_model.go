// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivilegeByRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListPrivilegeByRoleRequest
	GetAccountId() *string
}

type ListPrivilegeByRoleRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListPrivilegeByRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegeByRoleRequest) GoString() string {
	return s.String()
}

func (s *ListPrivilegeByRoleRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListPrivilegeByRoleRequest) SetAccountId(v string) *ListPrivilegeByRoleRequest {
	s.AccountId = &v
	return s
}

func (s *ListPrivilegeByRoleRequest) Validate() error {
	return dara.Validate(s)
}
