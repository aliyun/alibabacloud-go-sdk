// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUsersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUsersResponseBody
	GetTotalCount() *int32
	SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() *ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the users.
	Users *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) GetUsers() *ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersResponseBodyUsers struct {
	UserInfo []*ListUsersResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetUserInfo() []*ListUsersResponseBodyUsersUserInfo {
	return s.UserInfo
}

func (s *ListUsersResponseBodyUsers) SetUserInfo(v []*ListUsersResponseBodyUsersUserInfo) *ListUsersResponseBodyUsers {
	s.UserInfo = v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
	if s.UserInfo != nil {
		for _, item := range s.UserInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyUsersUserInfo struct {
	// The time when the user was first added.
	//
	// example:
	//
	// 2014-08-22T17:46:47
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The name of the permission group. Valid values:
	//
	// users: ordinary permissions, which are suitable for regular users that need only to submit and debug jobs.
	//
	// wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The permission group ID.
	//
	// example:
	//
	// 100
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsersUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserInfo) GetAddTime() *string {
	return s.AddTime
}

func (s *ListUsersResponseBodyUsersUserInfo) GetGroup() *string {
	return s.Group
}

func (s *ListUsersResponseBodyUsersUserInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *ListUsersResponseBodyUsersUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsersUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersResponseBodyUsersUserInfo) SetAddTime(v string) *ListUsersResponseBodyUsersUserInfo {
	s.AddTime = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetGroup(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Group = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetGroupId(v string) *ListUsersResponseBodyUsersUserInfo {
	s.GroupId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetUserId(v string) *ListUsersResponseBodyUsersUserInfo {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetUserName(v string) *ListUsersResponseBodyUsersUserInfo {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) Validate() error {
	return dara.Validate(s)
}
