// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUserGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserGroupsResponseBody
	GetTotalCount() *int32
	SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody
	GetUserGroups() []*ListUserGroupsResponseBodyUserGroups
}

type ListUserGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of user groups returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The user groups returned.
	UserGroups []*ListUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserGroupsResponseBody) GetUserGroups() []*ListUserGroupsResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetTotalCount(v int32) *ListUserGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListUserGroupsResponseBody) Validate() error {
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsResponseBodyUserGroups struct {
	// The description of the user group.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The number of users in the user group.
	//
	// example:
	//
	// 5
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// TestGroup01
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroups) GetComment() *string {
	return s.Comment
}

func (s *ListUserGroupsResponseBodyUserGroups) GetMemberCount() *int32 {
	return s.MemberCount
}

func (s *ListUserGroupsResponseBodyUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupsResponseBodyUserGroups) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListUserGroupsResponseBodyUserGroups) SetComment(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Comment = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetMemberCount(v int32) *ListUserGroupsResponseBodyUserGroups {
	s.MemberCount = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupId(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupName(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupName = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}
