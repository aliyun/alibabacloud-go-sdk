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
	// The total number of entries that match the query.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of account objects.
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
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForGroupResponseBodyUsers struct {
	// The source ID of the group member relationship.
	//
	// If the group is created in EIAM, the value of this parameter is the instance ID. For other types of groups, the value is the enterprise ID from the source. For example, if the group is imported from DingTalk, the value is the corpId of the DingTalk enterprise.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupMemberRelationSourceId *string `json:"GroupMemberRelationSourceId,omitempty" xml:"GroupMemberRelationSourceId,omitempty"`
	// The source type of the group member relationship. Valid values:
	//
	// build_in: The group is created in EIAM.
	//
	// ding_talk: The group is imported from DingTalk.
	//
	// ad: The group is imported from Active Directory (AD).
	//
	// ldap: The group is imported from LDAP.
	//
	// we_com: The group is imported from WeCom.
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
