// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamMembersResponseBody
	GetCount() *int64
	SetMaxResults(v int32) *ListIpamMembersResponseBody
	GetMaxResults() *int32
	SetMemberInfos(v []*ListIpamMembersResponseBodyMemberInfos) *ListIpamMembersResponseBody
	GetMemberInfos() []*ListIpamMembersResponseBodyMemberInfos
	SetNextToken(v string) *ListIpamMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamMembersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamMembersResponseBody
	GetTotalCount() *int64
}

type ListIpamMembersResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 6
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The maximum number of entries returned on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A list of members managed by the IPAM trusted service.
	MemberInfos []*ListIpamMembersResponseBodyMemberInfos `json:"MemberInfos,omitempty" xml:"MemberInfos,omitempty" type:"Repeated"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no more results are available.
	//
	// - If **NextToken*	- has a value, the value is the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B90776C8-F703-51D5-893A-AD1CA699D535
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamMembersResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamMembersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamMembersResponseBody) GetMemberInfos() []*ListIpamMembersResponseBodyMemberInfos {
	return s.MemberInfos
}

func (s *ListIpamMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamMembersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamMembersResponseBody) SetCount(v int64) *ListIpamMembersResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamMembersResponseBody) SetMaxResults(v int32) *ListIpamMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamMembersResponseBody) SetMemberInfos(v []*ListIpamMembersResponseBodyMemberInfos) *ListIpamMembersResponseBody {
	s.MemberInfos = v
	return s
}

func (s *ListIpamMembersResponseBody) SetNextToken(v string) *ListIpamMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamMembersResponseBody) SetRequestId(v string) *ListIpamMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamMembersResponseBody) SetTotalCount(v int64) *ListIpamMembersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamMembersResponseBody) Validate() error {
	if s.MemberInfos != nil {
		for _, item := range s.MemberInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamMembersResponseBodyMemberInfos struct {
	// The time when the member was added.
	//
	// example:
	//
	// 2025-07-11T07:18:07Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The member ID.
	//
	// - **Folder ID**: The ID of the folder.
	//
	// - **Account UID**: The UID of the member account in the resource directory.
	//
	// example:
	//
	// fd-ccccncASqa
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The type of the member. Valid values:
	//
	// - **Folder**: The member is a folder.
	//
	// - **Account**: The member is a member account in the resource directory.
	//
	// example:
	//
	// Folder
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The status of the member managed by the IPAM trusted service. Valid values:
	//
	// - **Created**: The member is managed.
	//
	// - **Deleted**: The member is removed.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamMembersResponseBodyMemberInfos) String() string {
	return dara.Prettify(s)
}

func (s ListIpamMembersResponseBodyMemberInfos) GoString() string {
	return s.String()
}

func (s *ListIpamMembersResponseBodyMemberInfos) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListIpamMembersResponseBodyMemberInfos) GetMemberId() *string {
	return s.MemberId
}

func (s *ListIpamMembersResponseBodyMemberInfos) GetMemberType() *string {
	return s.MemberType
}

func (s *ListIpamMembersResponseBodyMemberInfos) GetStatus() *string {
	return s.Status
}

func (s *ListIpamMembersResponseBodyMemberInfos) SetCreationTime(v string) *ListIpamMembersResponseBodyMemberInfos {
	s.CreationTime = &v
	return s
}

func (s *ListIpamMembersResponseBodyMemberInfos) SetMemberId(v string) *ListIpamMembersResponseBodyMemberInfos {
	s.MemberId = &v
	return s
}

func (s *ListIpamMembersResponseBodyMemberInfos) SetMemberType(v string) *ListIpamMembersResponseBodyMemberInfos {
	s.MemberType = &v
	return s
}

func (s *ListIpamMembersResponseBodyMemberInfos) SetStatus(v string) *ListIpamMembersResponseBodyMemberInfos {
	s.Status = &v
	return s
}

func (s *ListIpamMembersResponseBodyMemberInfos) Validate() error {
	return dara.Validate(s)
}
