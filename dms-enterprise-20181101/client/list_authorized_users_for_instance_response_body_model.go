// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAuthorizedUsersForInstanceResponseBody
	GetRequestId() *string
	SetUsers(v []*ListAuthorizedUsersForInstanceResponseBodyUsers) *ListAuthorizedUsersForInstanceResponseBody
	GetUsers() []*ListAuthorizedUsersForInstanceResponseBodyUsers
}

type ListAuthorizedUsersForInstanceResponseBody struct {
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*ListAuthorizedUsersForInstanceResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListAuthorizedUsersForInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedUsersForInstanceResponseBody) GetUsers() []*ListAuthorizedUsersForInstanceResponseBodyUsers {
	return s.Users
}

func (s *ListAuthorizedUsersForInstanceResponseBody) SetRequestId(v string) *ListAuthorizedUsersForInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBody) SetUsers(v []*ListAuthorizedUsersForInstanceResponseBodyUsers) *ListAuthorizedUsersForInstanceResponseBody {
	s.Users = v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedUsersForInstanceResponseBodyUsers struct {
	// example:
	//
	// 164882191****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// 51***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// user_test
	UserNickName *string `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
	// example:
	//
	// user01
	UserRealName *string `json:"UserRealName,omitempty" xml:"UserRealName,omitempty"`
}

func (s ListAuthorizedUsersForInstanceResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForInstanceResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) GetUid() *string {
	return s.Uid
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) GetUserNickName() *string {
	return s.UserNickName
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) GetUserRealName() *string {
	return s.UserRealName
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) SetUid(v string) *ListAuthorizedUsersForInstanceResponseBodyUsers {
	s.Uid = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) SetUserId(v string) *ListAuthorizedUsersForInstanceResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) SetUserNickName(v string) *ListAuthorizedUsersForInstanceResponseBodyUsers {
	s.UserNickName = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) SetUserRealName(v string) *ListAuthorizedUsersForInstanceResponseBodyUsers {
	s.UserRealName = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
