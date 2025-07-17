// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *RegisterUserRequest
	GetMobile() *string
	SetRoleNames(v string) *RegisterUserRequest
	GetRoleNames() *string
	SetTid(v int64) *RegisterUserRequest
	GetTid() *int64
	SetUid(v string) *RegisterUserRequest
	GetUid() *string
	SetUserNick(v string) *RegisterUserRequest
	GetUserNick() *string
}

type RegisterUserRequest struct {
	// The mobile number of the user.
	//
	// example:
	//
	// 130000000xx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The role that you want to assign to the user. Valid values:
	//
	// 	- **USER**: a regular user role
	//
	// 	- **DBA**: a database administrator (DBA) role
	//
	// 	- **ADMIN**: a DMS administrator role
	//
	// 	- **SECURITY_ADMIN**: a security administrator role
	//
	// >  If you do not specify this parameter, the regular user role is assigned to the user by default. You can assign one or more roles to the user. Separate multiple roles with commas (,).
	//
	// example:
	//
	// USER,DBA
	RoleNames *string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	// The ID of the tenant.
	//
	// >  To query ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The UID of the Alibaba Cloud account of the user that you want to register.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// dmstest
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s RegisterUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterUserRequest) GoString() string {
	return s.String()
}

func (s *RegisterUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *RegisterUserRequest) GetRoleNames() *string {
	return s.RoleNames
}

func (s *RegisterUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RegisterUserRequest) GetUid() *string {
	return s.Uid
}

func (s *RegisterUserRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *RegisterUserRequest) SetMobile(v string) *RegisterUserRequest {
	s.Mobile = &v
	return s
}

func (s *RegisterUserRequest) SetRoleNames(v string) *RegisterUserRequest {
	s.RoleNames = &v
	return s
}

func (s *RegisterUserRequest) SetTid(v int64) *RegisterUserRequest {
	s.Tid = &v
	return s
}

func (s *RegisterUserRequest) SetUid(v string) *RegisterUserRequest {
	s.Uid = &v
	return s
}

func (s *RegisterUserRequest) SetUserNick(v string) *RegisterUserRequest {
	s.UserNick = &v
	return s
}

func (s *RegisterUserRequest) Validate() error {
	return dara.Validate(s)
}
