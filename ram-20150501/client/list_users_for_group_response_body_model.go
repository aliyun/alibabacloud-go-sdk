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
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.````
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     *ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JoinDate    *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *ListUsersForGroupResponseBodyUsersUser) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetDisplayName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetJoinDate(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.JoinDate = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
