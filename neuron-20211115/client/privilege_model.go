// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrivilege interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Privilege
	GetAccountId() *string
	SetActions(v string) *Privilege
	GetActions() *string
	SetGmtCreate(v string) *Privilege
	GetGmtCreate() *string
	SetId(v int64) *Privilege
	GetId() *int64
	SetRequestId(v string) *Privilege
	GetRequestId() *string
	SetResource(v string) *Privilege
	GetResource() *string
	SetRoleId(v int64) *Privilege
	GetRoleId() *int64
}

type Privilege struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "write,edit"
	Actions   *string `json:"actions,omitempty" xml:"actions,omitempty"`
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// neuron:catalog:company/1:pbc/2
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s Privilege) String() string {
	return dara.Prettify(s)
}

func (s Privilege) GoString() string {
	return s.String()
}

func (s *Privilege) GetAccountId() *string {
	return s.AccountId
}

func (s *Privilege) GetActions() *string {
	return s.Actions
}

func (s *Privilege) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Privilege) GetId() *int64 {
	return s.Id
}

func (s *Privilege) GetRequestId() *string {
	return s.RequestId
}

func (s *Privilege) GetResource() *string {
	return s.Resource
}

func (s *Privilege) GetRoleId() *int64 {
	return s.RoleId
}

func (s *Privilege) SetAccountId(v string) *Privilege {
	s.AccountId = &v
	return s
}

func (s *Privilege) SetActions(v string) *Privilege {
	s.Actions = &v
	return s
}

func (s *Privilege) SetGmtCreate(v string) *Privilege {
	s.GmtCreate = &v
	return s
}

func (s *Privilege) SetId(v int64) *Privilege {
	s.Id = &v
	return s
}

func (s *Privilege) SetRequestId(v string) *Privilege {
	s.RequestId = &v
	return s
}

func (s *Privilege) SetResource(v string) *Privilege {
	s.Resource = &v
	return s
}

func (s *Privilege) SetRoleId(v int64) *Privilege {
	s.RoleId = &v
	return s
}

func (s *Privilege) Validate() error {
	return dara.Validate(s)
}
