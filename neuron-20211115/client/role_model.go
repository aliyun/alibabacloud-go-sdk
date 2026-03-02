// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRole interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Role
	GetAccountId() *string
	SetEnterpriseId(v int64) *Role
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *Role
	GetGmtCreate() *string
	SetId(v int64) *Role
	GetId() *int64
	SetName(v string) *Role
	GetName() *string
	SetPlatform(v string) *Role
	GetPlatform() *string
	SetRequestId(v string) *Role
	GetRequestId() *string
	SetType(v string) *Role
	GetType() *string
}

type Role struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 2022-04-01
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 开发者
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PDP
	Platform  *string `json:"platform,omitempty" xml:"platform,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Role) String() string {
	return dara.Prettify(s)
}

func (s Role) GoString() string {
	return s.String()
}

func (s *Role) GetAccountId() *string {
	return s.AccountId
}

func (s *Role) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *Role) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Role) GetId() *int64 {
	return s.Id
}

func (s *Role) GetName() *string {
	return s.Name
}

func (s *Role) GetPlatform() *string {
	return s.Platform
}

func (s *Role) GetRequestId() *string {
	return s.RequestId
}

func (s *Role) GetType() *string {
	return s.Type
}

func (s *Role) SetAccountId(v string) *Role {
	s.AccountId = &v
	return s
}

func (s *Role) SetEnterpriseId(v int64) *Role {
	s.EnterpriseId = &v
	return s
}

func (s *Role) SetGmtCreate(v string) *Role {
	s.GmtCreate = &v
	return s
}

func (s *Role) SetId(v int64) *Role {
	s.Id = &v
	return s
}

func (s *Role) SetName(v string) *Role {
	s.Name = &v
	return s
}

func (s *Role) SetPlatform(v string) *Role {
	s.Platform = &v
	return s
}

func (s *Role) SetRequestId(v string) *Role {
	s.RequestId = &v
	return s
}

func (s *Role) SetType(v string) *Role {
	s.Type = &v
	return s
}

func (s *Role) Validate() error {
	return dara.Validate(s)
}
