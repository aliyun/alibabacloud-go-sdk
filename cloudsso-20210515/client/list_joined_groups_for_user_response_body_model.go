// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJoinedGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListJoinedGroupsForUserResponseBody
	GetIsTruncated() *bool
	SetJoinedGroups(v []*ListJoinedGroupsForUserResponseBodyJoinedGroups) *ListJoinedGroupsForUserResponseBody
	GetJoinedGroups() []*ListJoinedGroupsForUserResponseBodyJoinedGroups
	SetMaxResults(v int32) *ListJoinedGroupsForUserResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListJoinedGroupsForUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListJoinedGroupsForUserResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListJoinedGroupsForUserResponseBody
	GetTotalCounts() *int32
}

type ListJoinedGroupsForUserResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The groups to which the user is added.
	JoinedGroups []*ListJoinedGroupsForUserResponseBodyJoinedGroups `json:"JoinedGroups,omitempty" xml:"JoinedGroups,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when the `IsTruncated` parameter is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9BBB45F-7877-5DE9-96A5-20E6CFA48929
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListJoinedGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListJoinedGroupsForUserResponseBody) GetJoinedGroups() []*ListJoinedGroupsForUserResponseBodyJoinedGroups {
	return s.JoinedGroups
}

func (s *ListJoinedGroupsForUserResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJoinedGroupsForUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJoinedGroupsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJoinedGroupsForUserResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListJoinedGroupsForUserResponseBody) SetIsTruncated(v bool) *ListJoinedGroupsForUserResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetJoinedGroups(v []*ListJoinedGroupsForUserResponseBodyJoinedGroups) *ListJoinedGroupsForUserResponseBody {
	s.JoinedGroups = v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetMaxResults(v int32) *ListJoinedGroupsForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetNextToken(v string) *ListJoinedGroupsForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetRequestId(v string) *ListJoinedGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) SetTotalCounts(v int32) *ListJoinedGroupsForUserResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBody) Validate() error {
	if s.JoinedGroups != nil {
		for _, item := range s.JoinedGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJoinedGroupsForUserResponseBodyJoinedGroups struct {
	// The description of the group.
	//
	// example:
	//
	// This is a group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// g-00jqzghi2n3o5hkh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// TestGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the user was added to the group.
	//
	// example:
	//
	// 2021-11-01T06:58:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The type of the group. Valid values:
	//
	// 	- Manual: The group is manually created.
	//
	// 	- Synchronized: The group is synchronized from an external identity provider (IdP).
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListJoinedGroupsForUserResponseBodyJoinedGroups) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedGroupsForUserResponseBodyJoinedGroups) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetDescription() *string {
	return s.Description
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) GetUserId() *string {
	return s.UserId
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetDescription(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.Description = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetGroupId(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.GroupId = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetGroupName(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.GroupName = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetJoinTime(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.JoinTime = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetProvisionType(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.ProvisionType = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) SetUserId(v string) *ListJoinedGroupsForUserResponseBodyJoinedGroups {
	s.UserId = &v
	return s
}

func (s *ListJoinedGroupsForUserResponseBodyJoinedGroups) Validate() error {
	return dara.Validate(s)
}
