// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsForApplicationResponseBodyGroups) *ListGroupsForApplicationResponseBody
	GetGroups() []*ListGroupsForApplicationResponseBodyGroups
	SetRequestId(v string) *ListGroupsForApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsForApplicationResponseBody
	GetTotalCount() *int64
}

type ListGroupsForApplicationResponseBody struct {
	// The group IDs.
	Groups []*ListGroupsForApplicationResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBody) GetGroups() []*ListGroupsForApplicationResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForApplicationResponseBody) SetGroups(v []*ListGroupsForApplicationResponseBodyGroups) *ListGroupsForApplicationResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetRequestId(v string) *ListGroupsForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetTotalCount(v int64) *ListGroupsForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForApplicationResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForApplicationResponseBodyGroups struct {
	// 应用角色列表。
	ApplicationRoles []*ListGroupsForApplicationResponseBodyGroupsApplicationRoles `json:"ApplicationRoles,omitempty" xml:"ApplicationRoles,omitempty" type:"Repeated"`
	// The group ID.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListGroupsForApplicationResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBodyGroups) GetApplicationRoles() []*ListGroupsForApplicationResponseBodyGroupsApplicationRoles {
	return s.ApplicationRoles
}

func (s *ListGroupsForApplicationResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForApplicationResponseBodyGroups) SetApplicationRoles(v []*ListGroupsForApplicationResponseBodyGroupsApplicationRoles) *ListGroupsForApplicationResponseBodyGroups {
	s.ApplicationRoles = v
	return s
}

func (s *ListGroupsForApplicationResponseBodyGroups) SetGroupId(v string) *ListGroupsForApplicationResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBodyGroups) Validate() error {
	if s.ApplicationRoles != nil {
		for _, item := range s.ApplicationRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForApplicationResponseBodyGroupsApplicationRoles struct {
	// 应用角色标识。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
}

func (s ListGroupsForApplicationResponseBodyGroupsApplicationRoles) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponseBodyGroupsApplicationRoles) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBodyGroupsApplicationRoles) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListGroupsForApplicationResponseBodyGroupsApplicationRoles) SetApplicationRoleId(v string) *ListGroupsForApplicationResponseBodyGroupsApplicationRoles {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBodyGroupsApplicationRoles) Validate() error {
	return dara.Validate(s)
}
