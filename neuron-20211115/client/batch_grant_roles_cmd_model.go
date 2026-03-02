// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantRolesCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *BatchGrantRolesCmd
	GetAccountId() *string
	SetRoleIds(v []*int64) *BatchGrantRolesCmd
	GetRoleIds() []*int64
}

type BatchGrantRolesCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AccountId *string  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	RoleIds   []*int64 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
}

func (s BatchGrantRolesCmd) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantRolesCmd) GoString() string {
	return s.String()
}

func (s *BatchGrantRolesCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchGrantRolesCmd) GetRoleIds() []*int64 {
	return s.RoleIds
}

func (s *BatchGrantRolesCmd) SetAccountId(v string) *BatchGrantRolesCmd {
	s.AccountId = &v
	return s
}

func (s *BatchGrantRolesCmd) SetRoleIds(v []*int64) *BatchGrantRolesCmd {
	s.RoleIds = v
	return s
}

func (s *BatchGrantRolesCmd) Validate() error {
	return dara.Validate(s)
}
