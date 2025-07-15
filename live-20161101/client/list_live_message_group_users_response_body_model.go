// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListLiveMessageGroupUsersResponseBody
	GetGroupId() *string
	SetHasmore(v bool) *ListLiveMessageGroupUsersResponseBody
	GetHasmore() *bool
	SetNextPageToken(v int64) *ListLiveMessageGroupUsersResponseBody
	GetNextPageToken() *int64
	SetRequestId(v string) *ListLiveMessageGroupUsersResponseBody
	GetRequestId() *string
	SetUserList(v []*ListLiveMessageGroupUsersResponseBodyUserList) *ListLiveMessageGroupUsersResponseBody
	GetUserList() []*ListLiveMessageGroupUsersResponseBodyUserList
}

type ListLiveMessageGroupUsersResponseBody struct {
	// The ID of the group queried.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the current page is followed by another page.
	//
	// example:
	//
	// false
	Hasmore *bool `json:"Hasmore,omitempty" xml:"Hasmore,omitempty"`
	// The starting page number for the next query. A value of 0 indicates that no further pages can be queried.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1D75BEA-1329-116F-B29C-76F3F200****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the users.
	UserList []*ListLiveMessageGroupUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListLiveMessageGroupUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupUsersResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupUsersResponseBody) GetHasmore() *bool {
	return s.Hasmore
}

func (s *ListLiveMessageGroupUsersResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageGroupUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveMessageGroupUsersResponseBody) GetUserList() []*ListLiveMessageGroupUsersResponseBodyUserList {
	return s.UserList
}

func (s *ListLiveMessageGroupUsersResponseBody) SetGroupId(v string) *ListLiveMessageGroupUsersResponseBody {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBody) SetHasmore(v bool) *ListLiveMessageGroupUsersResponseBody {
	s.Hasmore = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBody) SetNextPageToken(v int64) *ListLiveMessageGroupUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBody) SetRequestId(v string) *ListLiveMessageGroupUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBody) SetUserList(v []*ListLiveMessageGroupUsersResponseBodyUserList) *ListLiveMessageGroupUsersResponseBody {
	s.UserList = v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveMessageGroupUsersResponseBodyUserList struct {
	// The ID of the user.
	//
	// example:
	//
	// uid1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The additional information about the user.
	//
	// example:
	//
	// info1
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListLiveMessageGroupUsersResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupUsersResponseBodyUserList) GetUserId() *string {
	return s.UserId
}

func (s *ListLiveMessageGroupUsersResponseBodyUserList) GetUserInfo() *string {
	return s.UserInfo
}

func (s *ListLiveMessageGroupUsersResponseBodyUserList) SetUserId(v string) *ListLiveMessageGroupUsersResponseBodyUserList {
	s.UserId = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBodyUserList) SetUserInfo(v string) *ListLiveMessageGroupUsersResponseBodyUserList {
	s.UserInfo = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
