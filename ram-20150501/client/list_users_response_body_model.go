// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUsersResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListUsersResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() *ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set `Marker` to obtain the truncated part.``
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUsersResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetUsers() *ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetIsTruncated(v bool) *ListUsersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponseBody) SetMarker(v string) *ListUsersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
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
	User []*ListUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetUser() []*ListUsersResponseBodyUsersUser {
	return s.User
}

func (s *ListUsersResponseBodyUsers) SetUser(v []*ListUsersResponseBodyUsersUser) *ListUsersResponseBodyUsers {
	s.User = v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
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

type ListUsersResponseBodyUsersUser struct {
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUser) GetComments() *string {
	return s.Comments
}

func (s *ListUsersResponseBodyUsersUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListUsersResponseBodyUsersUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyUsersUser) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsersUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ListUsersResponseBodyUsersUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListUsersResponseBodyUsersUser) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsersUser) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersResponseBodyUsersUser) SetComments(v string) *ListUsersResponseBodyUsersUser {
	s.Comments = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetCreateDate(v string) *ListUsersResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetDisplayName(v string) *ListUsersResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetEmail(v string) *ListUsersResponseBodyUsersUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetMobilePhone(v string) *ListUsersResponseBodyUsersUser {
	s.MobilePhone = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUpdateDate(v string) *ListUsersResponseBodyUsersUser {
	s.UpdateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserId(v string) *ListUsersResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserName(v string) *ListUsersResponseBodyUsersUser {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
