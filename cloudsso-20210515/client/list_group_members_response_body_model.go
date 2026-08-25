// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupMembers(v []*ListGroupMembersResponseBodyGroupMembers) *ListGroupMembersResponseBody
	GetGroupMembers() []*ListGroupMembersResponseBodyGroupMembers
	SetIsTruncated(v bool) *ListGroupMembersResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListGroupMembersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGroupMembersResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListGroupMembersResponseBody
	GetTotalCounts() *int32
}

type ListGroupMembersResponseBody struct {
	// The users in the group.
	GroupMembers []*ListGroupMembersResponseBodyGroupMembers `json:"GroupMembers,omitempty" xml:"GroupMembers,omitempty" type:"Repeated"`
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
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when the value of the `IsTruncated` parameter is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB759F84-2C64-5C36-B6C6-253172C5C370
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponseBody) GetGroupMembers() []*ListGroupMembersResponseBodyGroupMembers {
	return s.GroupMembers
}

func (s *ListGroupMembersResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListGroupMembersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupMembersResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListGroupMembersResponseBody) SetGroupMembers(v []*ListGroupMembersResponseBodyGroupMembers) *ListGroupMembersResponseBody {
	s.GroupMembers = v
	return s
}

func (s *ListGroupMembersResponseBody) SetIsTruncated(v bool) *ListGroupMembersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetMaxResults(v int32) *ListGroupMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetNextToken(v string) *ListGroupMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetRequestId(v string) *ListGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupMembersResponseBody) SetTotalCounts(v int32) *ListGroupMembersResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListGroupMembersResponseBody) Validate() error {
	if s.GroupMembers != nil {
		for _, item := range s.GroupMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupMembersResponseBodyGroupMembers struct {
	// The description of the user.
	//
	// example:
	//
	// This is a user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// Alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// AliceLee@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// g-00jqzghi2n3o5hkh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the user was added to the group.
	//
	// example:
	//
	// 2021-11-01T06:58:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The type of the user. Valid values:
	//
	// 	- Manual: The user is manually created.
	//
	// 	- Synchronized: The user is synchronized from an external identity provider (IdP).
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// 	- Enabled: The logon of the user is enabled.
	//
	// 	- Disabled: The logon of the user is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListGroupMembersResponseBodyGroupMembers) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMembersResponseBodyGroupMembers) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetDescription() *string {
	return s.Description
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetEmail() *string {
	return s.Email
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetStatus() *string {
	return s.Status
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetUserId() *string {
	return s.UserId
}

func (s *ListGroupMembersResponseBodyGroupMembers) GetUserName() *string {
	return s.UserName
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetDescription(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Description = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetDisplayName(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.DisplayName = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetEmail(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Email = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetGroupId(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.GroupId = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetJoinTime(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.JoinTime = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetProvisionType(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.ProvisionType = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetStatus(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.Status = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetUserId(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.UserId = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) SetUserName(v string) *ListGroupMembersResponseBodyGroupMembers {
	s.UserName = &v
	return s
}

func (s *ListGroupMembersResponseBodyGroupMembers) Validate() error {
	return dara.Validate(s)
}
