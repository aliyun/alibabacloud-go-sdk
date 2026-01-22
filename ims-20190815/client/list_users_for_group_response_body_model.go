// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUsersForGroupResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListUsersForGroupResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListUsersForGroupResponseBody
	GetRequestId() *string
	SetUsers(v *ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody
	GetUsers() *ListUsersForGroupResponseBodyUsers
}

type ListUsersForGroupResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 789FF581-B3C8-43A8-9115-54304B46D05C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUsersForGroupResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersForGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersForGroupResponseBody) GetUsers() *ListUsersForGroupResponseBodyUsers {
	return s.Users
}

func (s *ListUsersForGroupResponseBody) SetIsTruncated(v bool) *ListUsersForGroupResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetMarker(v string) *ListUsersForGroupResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetRequestId(v string) *ListUsersForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetUsers(v *ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersForGroupResponseBody) Validate() error {
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersForGroupResponseBodyUsers struct {
	User []*ListUsersForGroupResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsers) GetUser() []*ListUsersForGroupResponseBodyUsersUser {
	return s.User
}

func (s *ListUsersForGroupResponseBodyUsers) SetUser(v []*ListUsersForGroupResponseBodyUsersUser) *ListUsersForGroupResponseBodyUsers {
	s.User = v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForGroupResponseBodyUsersUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was added to the RAM user group. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-20T06:57:00Z
	JoinDate *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUsersForGroupResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsersUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersForGroupResponseBodyUsersUser) GetJoinDate() *string {
	return s.JoinDate
}

func (s *ListUsersForGroupResponseBodyUsersUser) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForGroupResponseBodyUsersUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetDisplayName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetJoinDate(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.JoinDate = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserId(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
