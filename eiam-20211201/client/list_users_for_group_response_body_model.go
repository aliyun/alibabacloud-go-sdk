// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUsersForGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUsersForGroupResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody
	GetUsers() []*ListUsersForGroupResponseBodyUsers
}

type ListUsersForGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries that can be returned per page is specified by PageSize.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about accounts.
	Users []*ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersForGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersForGroupResponseBody) GetUsers() []*ListUsersForGroupResponseBodyUsers {
	return s.Users
}

func (s *ListUsersForGroupResponseBody) SetRequestId(v string) *ListUsersForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetTotalCount(v int64) *ListUsersForGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetUsers(v []*ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersForGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersForGroupResponseBodyUsers struct {
	// Account membership source id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupMemberRelationSourceId *string `json:"GroupMemberRelationSourceId,omitempty" xml:"GroupMemberRelationSourceId,omitempty"`
	// Account membership source type
	//
	// example:
	//
	// build_in
	GroupMemberRelationSourceType *string `json:"GroupMemberRelationSourceType,omitempty" xml:"GroupMemberRelationSourceType,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersForGroupResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsers) GetGroupMemberRelationSourceId() *string {
	return s.GroupMemberRelationSourceId
}

func (s *ListUsersForGroupResponseBodyUsers) GetGroupMemberRelationSourceType() *string {
	return s.GroupMemberRelationSourceType
}

func (s *ListUsersForGroupResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForGroupResponseBodyUsers) SetGroupMemberRelationSourceId(v string) *ListUsersForGroupResponseBodyUsers {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) SetGroupMemberRelationSourceType(v string) *ListUsersForGroupResponseBodyUsers {
	s.GroupMemberRelationSourceType = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) SetUserId(v string) *ListUsersForGroupResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
