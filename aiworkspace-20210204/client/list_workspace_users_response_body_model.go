// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspaceUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListWorkspaceUsersResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListWorkspaceUsersResponseBodyUsers) *ListWorkspaceUsersResponseBody
	GetUsers() []*ListWorkspaceUsersResponseBodyUsers
}

type ListWorkspaceUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of users who meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The users.
	Users []*ListWorkspaceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListWorkspaceUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkspaceUsersResponseBody) GetUsers() []*ListWorkspaceUsersResponseBodyUsers {
	return s.Users
}

func (s *ListWorkspaceUsersResponseBody) SetRequestId(v string) *ListWorkspaceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetTotalCount(v int64) *ListWorkspaceUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetUsers(v []*ListWorkspaceUsersResponseBodyUsers) *ListWorkspaceUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListWorkspaceUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWorkspaceUsersResponseBodyUsers struct {
	// The user ID.
	//
	// example:
	//
	// 1611******3000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// she******mo
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListWorkspaceUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspaceUsersResponseBodyUsers) GetUserName() *string {
	return s.UserName
}

func (s *ListWorkspaceUsersResponseBodyUsers) SetUserId(v string) *ListWorkspaceUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyUsers) SetUserName(v string) *ListWorkspaceUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
