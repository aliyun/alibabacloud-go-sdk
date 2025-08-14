// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMediaLiveInputSecurityGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaLiveInputSecurityGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMediaLiveInputSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroups(v []*ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) *ListMediaLiveInputSecurityGroupsResponseBody
	GetSecurityGroups() []*ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups
	SetTotalCount(v int32) *ListMediaLiveInputSecurityGroupsResponseBody
	GetTotalCount() *int32
}

type ListMediaLiveInputSecurityGroupsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security groups.
	SecurityGroups []*ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMediaLiveInputSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) GetSecurityGroups() []*ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	return s.SecurityGroups
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) SetMaxResults(v int32) *ListMediaLiveInputSecurityGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) SetNextToken(v string) *ListMediaLiveInputSecurityGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) SetRequestId(v string) *ListMediaLiveInputSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) SetSecurityGroups(v []*ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) *ListMediaLiveInputSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) SetTotalCount(v int32) *ListMediaLiveInputSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups struct {
	// The time when the security group was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-03T06:56:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The IDs of the inputs associated with the security group.
	InputIds []*string `json:"InputIds,omitempty" xml:"InputIds,omitempty" type:"Repeated"`
	// The security group name.
	//
	// example:
	//
	// mysg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group rules.
	WhitelistRules []*string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty" type:"Repeated"`
}

func (s ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GetInputIds() []*string {
	return s.InputIds
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) GetWhitelistRules() []*string {
	return s.WhitelistRules
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) SetCreateTime(v string) *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	s.CreateTime = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) SetInputIds(v []*string) *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	s.InputIds = v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) SetName(v string) *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	s.Name = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupId(v string) *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) SetWhitelistRules(v []*string) *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups {
	s.WhitelistRules = v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponseBodySecurityGroups) Validate() error {
	return dara.Validate(s)
}
