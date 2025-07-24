// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOSUser interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v string) *OSUser
	GetGroup() *string
	SetPassword(v string) *OSUser
	GetPassword() *string
	SetUser(v string) *OSUser
	GetUser() *string
}

type OSUser struct {
	// 用户组。
	//
	// example:
	//
	// hadoop
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// 用户密码。
	//
	// example:
	//
	// 12345****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 用户名称。
	//
	// example:
	//
	// 王五
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s OSUser) String() string {
	return dara.Prettify(s)
}

func (s OSUser) GoString() string {
	return s.String()
}

func (s *OSUser) GetGroup() *string {
	return s.Group
}

func (s *OSUser) GetPassword() *string {
	return s.Password
}

func (s *OSUser) GetUser() *string {
	return s.User
}

func (s *OSUser) SetGroup(v string) *OSUser {
	s.Group = &v
	return s
}

func (s *OSUser) SetPassword(v string) *OSUser {
	s.Password = &v
	return s
}

func (s *OSUser) SetUser(v string) *OSUser {
	s.User = &v
	return s
}

func (s *OSUser) Validate() error {
	return dara.Validate(s)
}
