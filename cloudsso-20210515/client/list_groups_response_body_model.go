// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody
	GetGroups() []*ListGroupsResponseBodyGroups
	SetIsTruncated(v bool) *ListGroupsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListGroupsResponseBody
	GetTotalCounts() *int32
}

type ListGroupsResponseBody struct {
	// The groups.
	Groups []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
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
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetGroups() []*ListGroupsResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListGroupsResponseBody) SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetIsTruncated(v bool) *ListGroupsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponseBody) SetMaxResults(v int32) *ListGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsResponseBody) SetNextToken(v string) *ListGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCounts(v int32) *ListGroupsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
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

type ListGroupsResponseBodyGroups struct {
	// The time when the group was created.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// The type of the group. Valid values:
	//
	// 	- Manual: The group is manually created.
	//
	// 	- Synchronized: The group is synchronized from an external IdP.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *ListGroupsResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyGroups) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListGroupsResponseBodyGroups) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListGroupsResponseBodyGroups) SetCreateTime(v string) *ListGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetDescription(v string) *ListGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupId(v string) *ListGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupName(v string) *ListGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetProvisionType(v string) *ListGroupsResponseBodyGroups {
	s.ProvisionType = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetUpdateTime(v string) *ListGroupsResponseBodyGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
