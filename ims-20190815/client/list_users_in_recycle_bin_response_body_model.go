// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUsersInRecycleBinResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListUsersInRecycleBinResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListUsersInRecycleBinResponseBody
	GetRequestId() *string
	SetUsers(v *ListUsersInRecycleBinResponseBodyUsers) *ListUsersInRecycleBinResponseBody
	GetUsers() *ListUsersInRecycleBinResponseBodyUsers
}

type ListUsersInRecycleBinResponseBody struct {
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
	// The parameter that is used to obtain the truncated part. It takes effect only when `IsTruncated` is set to `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3687BD52-49FD-585B-AB14-CD05B7C76963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersInRecycleBinResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUsersInRecycleBinResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersInRecycleBinResponseBody) GetUsers() *ListUsersInRecycleBinResponseBodyUsers {
	return s.Users
}

func (s *ListUsersInRecycleBinResponseBody) SetIsTruncated(v bool) *ListUsersInRecycleBinResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetMarker(v string) *ListUsersInRecycleBinResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetRequestId(v string) *ListUsersInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetUsers(v *ListUsersInRecycleBinResponseBodyUsers) *ListUsersInRecycleBinResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersInRecycleBinResponseBodyUsers struct {
	User []*ListUsersInRecycleBinResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersInRecycleBinResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBodyUsers) GetUser() []*ListUsersInRecycleBinResponseBodyUsersUser {
	return s.User
}

func (s *ListUsersInRecycleBinResponseBodyUsers) SetUser(v []*ListUsersInRecycleBinResponseBodyUsersUser) *ListUsersInRecycleBinResponseBodyUsers {
	s.User = v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}

type ListUsersInRecycleBinResponseBodyUsersUser struct {
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the RAM user will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-15T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-15T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
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

func (s ListUsersInRecycleBinResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetRecycleDate() *string {
	return s.RecycleDate
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetCreateDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetDeleteDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.DeleteDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetDisplayName(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetRecycleDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.RecycleDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetUserId(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
