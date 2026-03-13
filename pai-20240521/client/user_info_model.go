// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserInfo interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *UserInfo
	GetUserId() *string
	SetUserName(v string) *UserInfo
	GetUserName() *string
}

type UserInfo struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UserInfo) String() string {
	return dara.Prettify(s)
}

func (s UserInfo) GoString() string {
	return s.String()
}

func (s *UserInfo) GetUserId() *string {
	return s.UserId
}

func (s *UserInfo) GetUserName() *string {
	return s.UserName
}

func (s *UserInfo) SetUserId(v string) *UserInfo {
	s.UserId = &v
	return s
}

func (s *UserInfo) SetUserName(v string) *UserInfo {
	s.UserName = &v
	return s
}

func (s *UserInfo) Validate() error {
	return dara.Validate(s)
}
