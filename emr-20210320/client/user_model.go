// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUser interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v string) *User
	GetGroup() *string
	SetPassword(v string) *User
	GetPassword() *string
	SetUserId(v string) *User
	GetUserId() *string
	SetUserName(v string) *User
	GetUserName() *string
	SetUserType(v string) *User
	GetUserType() *string
}

type User struct {
	// 用户组。
	//
	// example:
	//
	// hadoop
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// 用户密码。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 用户ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1238539****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// 王五
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 用户类型。
	//
	// example:
	//
	// LDAP
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s User) String() string {
	return dara.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) GetGroup() *string {
	return s.Group
}

func (s *User) GetPassword() *string {
	return s.Password
}

func (s *User) GetUserId() *string {
	return s.UserId
}

func (s *User) GetUserName() *string {
	return s.UserName
}

func (s *User) GetUserType() *string {
	return s.UserType
}

func (s *User) SetGroup(v string) *User {
	s.Group = &v
	return s
}

func (s *User) SetPassword(v string) *User {
	s.Password = &v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

func (s *User) SetUserType(v string) *User {
	s.UserType = &v
	return s
}

func (s *User) Validate() error {
	return dara.Validate(s)
}
