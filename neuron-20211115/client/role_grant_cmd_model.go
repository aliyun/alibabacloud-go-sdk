// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoleGrantCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *RoleGrantCmd
	GetAccountId() *string
	SetAuthorizationResourceList(v []*AuthorizationResource) *RoleGrantCmd
	GetAuthorizationResourceList() []*AuthorizationResource
	SetRoleId(v int64) *RoleGrantCmd
	GetRoleId() *int64
}

type RoleGrantCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AccountId                 *string                  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AuthorizationResourceList []*AuthorizationResource `json:"authorizationResourceList,omitempty" xml:"authorizationResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s RoleGrantCmd) String() string {
	return dara.Prettify(s)
}

func (s RoleGrantCmd) GoString() string {
	return s.String()
}

func (s *RoleGrantCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *RoleGrantCmd) GetAuthorizationResourceList() []*AuthorizationResource {
	return s.AuthorizationResourceList
}

func (s *RoleGrantCmd) GetRoleId() *int64 {
	return s.RoleId
}

func (s *RoleGrantCmd) SetAccountId(v string) *RoleGrantCmd {
	s.AccountId = &v
	return s
}

func (s *RoleGrantCmd) SetAuthorizationResourceList(v []*AuthorizationResource) *RoleGrantCmd {
	s.AuthorizationResourceList = v
	return s
}

func (s *RoleGrantCmd) SetRoleId(v int64) *RoleGrantCmd {
	s.RoleId = &v
	return s
}

func (s *RoleGrantCmd) Validate() error {
	if s.AuthorizationResourceList != nil {
		for _, item := range s.AuthorizationResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
