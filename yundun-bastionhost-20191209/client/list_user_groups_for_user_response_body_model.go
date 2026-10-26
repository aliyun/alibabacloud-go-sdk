// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListUserGroupsForUserResponseBody
	GetCount() *int32
	SetRequestId(v string) *ListUserGroupsForUserResponseBody
	GetRequestId() *string
	SetUserGroups(v []*ListUserGroupsForUserResponseBodyUserGroups) *ListUserGroupsForUserResponseBody
	GetUserGroups() []*ListUserGroupsForUserResponseBodyUserGroups
}

type ListUserGroupsForUserResponseBody struct {
	// The total number of user groups returned.
	//
	// example:
	//
	// 12
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of user groups.
	UserGroups []*ListUserGroupsForUserResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForUserResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListUserGroupsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsForUserResponseBody) GetUserGroups() []*ListUserGroupsForUserResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListUserGroupsForUserResponseBody) SetCount(v int32) *ListUserGroupsForUserResponseBody {
	s.Count = &v
	return s
}

func (s *ListUserGroupsForUserResponseBody) SetRequestId(v string) *ListUserGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsForUserResponseBody) SetUserGroups(v []*ListUserGroupsForUserResponseBodyUserGroups) *ListUserGroupsForUserResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListUserGroupsForUserResponseBody) Validate() error {
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

type ListUserGroupsForUserResponseBodyUserGroups struct {
	// The remarks of the user group.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the project to which the user group belongs.
	//
	// example:
	//
	// 2
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project to which the user group belongs.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 2
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsForUserResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForUserResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) GetComment() *string {
	return s.Comment
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) SetComment(v string) *ListUserGroupsForUserResponseBodyUserGroups {
	s.Comment = &v
	return s
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) SetProjectId(v string) *ListUserGroupsForUserResponseBodyUserGroups {
	s.ProjectId = &v
	return s
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) SetProjectName(v string) *ListUserGroupsForUserResponseBodyUserGroups {
	s.ProjectName = &v
	return s
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) SetUserGroupId(v string) *ListUserGroupsForUserResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) SetUserGroupName(v string) *ListUserGroupsForUserResponseBodyUserGroups {
	s.UserGroupName = &v
	return s
}

func (s *ListUserGroupsForUserResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}
