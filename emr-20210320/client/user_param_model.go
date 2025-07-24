// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserParam interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *UserParam
	GetPassword() *string
	SetUserId(v string) *UserParam
	GetUserId() *string
	SetUserName(v string) *UserParam
	GetUserName() *string
}

type UserParam struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UserParam) String() string {
	return dara.Prettify(s)
}

func (s UserParam) GoString() string {
	return s.String()
}

func (s *UserParam) GetPassword() *string {
	return s.Password
}

func (s *UserParam) GetUserId() *string {
	return s.UserId
}

func (s *UserParam) GetUserName() *string {
	return s.UserName
}

func (s *UserParam) SetPassword(v string) *UserParam {
	s.Password = &v
	return s
}

func (s *UserParam) SetUserId(v string) *UserParam {
	s.UserId = &v
	return s
}

func (s *UserParam) SetUserName(v string) *UserParam {
	s.UserName = &v
	return s
}

func (s *UserParam) Validate() error {
	return dara.Validate(s)
}
