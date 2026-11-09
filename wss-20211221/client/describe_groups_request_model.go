// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeGroupsRequest
	GetEndTime() *int64
	SetGroupResourceTypes(v []*string) *DescribeGroupsRequest
	GetGroupResourceTypes() []*string
	SetGroupType(v string) *DescribeGroupsRequest
	GetGroupType() *string
	SetInstanceIds(v []*string) *DescribeGroupsRequest
	GetInstanceIds() []*string
	SetMaxResults(v string) *DescribeGroupsRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeGroupsRequest
	GetNextToken() *string
	SetResourceIds(v []*string) *DescribeGroupsRequest
	GetResourceIds() []*string
	SetStartTime(v int64) *DescribeGroupsRequest
	GetStartTime() *int64
	SetStates(v []*string) *DescribeGroupsRequest
	GetStates() []*string
	SetTimeUsingPeriod(v bool) *DescribeGroupsRequest
	GetTimeUsingPeriod() *bool
}

type DescribeGroupsRequest struct {
	// The end time of the period.
	//
	// example:
	//
	// 2026-03-28T02:06:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of group resource types.
	GroupResourceTypes []*string `json:"GroupResourceTypes,omitempty" xml:"GroupResourceTypes,omitempty" type:"Repeated"`
	// The group type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The Lingxiao instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first request. For subsequent requests, use the `nextToken` value returned in the previous response.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The start time of the period.
	//
	// example:
	//
	// 2026-03-25 10:26:44
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of states.
	States []*string `json:"States,omitempty" xml:"States,omitempty" type:"Repeated"`
	// The effective period used for the query.
	//
	// example:
	//
	// true/false
	TimeUsingPeriod *bool `json:"TimeUsingPeriod,omitempty" xml:"TimeUsingPeriod,omitempty"`
}

func (s DescribeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeGroupsRequest) GetGroupResourceTypes() []*string {
	return s.GroupResourceTypes
}

func (s *DescribeGroupsRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeGroupsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeGroupsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGroupsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeGroupsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeGroupsRequest) GetStates() []*string {
	return s.States
}

func (s *DescribeGroupsRequest) GetTimeUsingPeriod() *bool {
	return s.TimeUsingPeriod
}

func (s *DescribeGroupsRequest) SetEndTime(v int64) *DescribeGroupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeGroupsRequest) SetGroupResourceTypes(v []*string) *DescribeGroupsRequest {
	s.GroupResourceTypes = v
	return s
}

func (s *DescribeGroupsRequest) SetGroupType(v string) *DescribeGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeGroupsRequest) SetInstanceIds(v []*string) *DescribeGroupsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeGroupsRequest) SetMaxResults(v string) *DescribeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGroupsRequest) SetNextToken(v string) *DescribeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGroupsRequest) SetResourceIds(v []*string) *DescribeGroupsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeGroupsRequest) SetStartTime(v int64) *DescribeGroupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupsRequest) SetStates(v []*string) *DescribeGroupsRequest {
	s.States = v
	return s
}

func (s *DescribeGroupsRequest) SetTimeUsingPeriod(v bool) *DescribeGroupsRequest {
	s.TimeUsingPeriod = &v
	return s
}

func (s *DescribeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
