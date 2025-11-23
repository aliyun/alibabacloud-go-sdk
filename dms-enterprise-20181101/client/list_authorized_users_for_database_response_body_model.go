// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAuthorizedUsersForDatabaseResponseBody
	GetRequestId() *string
	SetUsers(v []*ListAuthorizedUsersForDatabaseResponseBodyUsers) *ListAuthorizedUsersForDatabaseResponseBody
	GetUsers() []*ListAuthorizedUsersForDatabaseResponseBodyUsers
}

type ListAuthorizedUsersForDatabaseResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of users that have permissions on the specified instance.
	Users []*ListAuthorizedUsersForDatabaseResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListAuthorizedUsersForDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedUsersForDatabaseResponseBody) GetUsers() []*ListAuthorizedUsersForDatabaseResponseBodyUsers {
	return s.Users
}

func (s *ListAuthorizedUsersForDatabaseResponseBody) SetRequestId(v string) *ListAuthorizedUsersForDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponseBody) SetUsers(v []*ListAuthorizedUsersForDatabaseResponseBodyUsers) *ListAuthorizedUsersForDatabaseResponseBody {
	s.Users = v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedUsersForDatabaseResponseBodyUsers struct {
	// The UID of the user\\"s Alibaba Cloud account.
	//
	// example:
	//
	// 164882191****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The user IDs.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// user_test
	UserNickName *string `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
}

func (s ListAuthorizedUsersForDatabaseResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForDatabaseResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) GetUid() *string {
	return s.Uid
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) GetUserNickName() *string {
	return s.UserNickName
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) SetUid(v string) *ListAuthorizedUsersForDatabaseResponseBodyUsers {
	s.Uid = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) SetUserId(v string) *ListAuthorizedUsersForDatabaseResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) SetUserNickName(v string) *ListAuthorizedUsersForDatabaseResponseBodyUsers {
	s.UserNickName = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
