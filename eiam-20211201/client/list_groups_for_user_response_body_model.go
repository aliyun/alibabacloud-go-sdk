// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody
	GetGroups() []*ListGroupsForUserResponseBodyGroups
	SetRequestId(v string) *ListGroupsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsForUserResponseBody
	GetTotalCount() *int64
}

type ListGroupsForUserResponseBody struct {
	// The queried account groups.
	Groups []*ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries returned at a time depends on the value of PageSize.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) GetGroups() []*ListGroupsForUserResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForUserResponseBody) SetGroups(v []*ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetRequestId(v string) *ListGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetTotalCount(v int64) *ListGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForUserResponseBody) Validate() error {
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

type ListGroupsForUserResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Account membership source ID
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
}

func (s ListGroupsForUserResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForUserResponseBodyGroups) GetGroupMemberRelationSourceId() *string {
	return s.GroupMemberRelationSourceId
}

func (s *ListGroupsForUserResponseBodyGroups) GetGroupMemberRelationSourceType() *string {
	return s.GroupMemberRelationSourceType
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupId(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupMemberRelationSourceId(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupMemberRelationSourceType(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupMemberRelationSourceType = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
