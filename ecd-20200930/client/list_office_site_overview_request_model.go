// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceRefresh(v bool) *ListOfficeSiteOverviewRequest
	GetForceRefresh() *bool
	SetMaxResults(v int32) *ListOfficeSiteOverviewRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOfficeSiteOverviewRequest
	GetNextToken() *string
	SetOfficeSiteId(v []*string) *ListOfficeSiteOverviewRequest
	GetOfficeSiteId() []*string
	SetQueryRange(v int32) *ListOfficeSiteOverviewRequest
	GetQueryRange() *int32
	SetRegionId(v string) *ListOfficeSiteOverviewRequest
	GetRegionId() *string
}

type ListOfficeSiteOverviewRequest struct {
	// Specifies whether to refresh the cache.
	//
	// example:
	//
	// false
	ForceRefresh *bool `json:"ForceRefresh,omitempty" xml:"ForceRefresh,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Valid values: 1 to 100
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. If this is your first query or no next query is to be sent, skip this parameter. If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network IDs. You can specify IDs of 1 to 100 office networks.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	// The query scope. Cloud computers in a cloud computer pool are pooled cloud computers.
	//
	// Default values:
	//
	// 	- 1 (default): queries non-pooled cloud computers in the office network.
	//
	// 	- 2: queries pooled cloud computers in the office network.
	//
	// 	- 3: queries all cloud computers in the office network.
	//
	// example:
	//
	// 1
	QueryRange *int32 `json:"QueryRange,omitempty" xml:"QueryRange,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOfficeSiteOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteOverviewRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewRequest) GetForceRefresh() *bool {
	return s.ForceRefresh
}

func (s *ListOfficeSiteOverviewRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOfficeSiteOverviewRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOfficeSiteOverviewRequest) GetOfficeSiteId() []*string {
	return s.OfficeSiteId
}

func (s *ListOfficeSiteOverviewRequest) GetQueryRange() *int32 {
	return s.QueryRange
}

func (s *ListOfficeSiteOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfficeSiteOverviewRequest) SetForceRefresh(v bool) *ListOfficeSiteOverviewRequest {
	s.ForceRefresh = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetMaxResults(v int32) *ListOfficeSiteOverviewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetNextToken(v string) *ListOfficeSiteOverviewRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetOfficeSiteId(v []*string) *ListOfficeSiteOverviewRequest {
	s.OfficeSiteId = v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetQueryRange(v int32) *ListOfficeSiteOverviewRequest {
	s.QueryRange = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetRegionId(v string) *ListOfficeSiteOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) Validate() error {
	return dara.Validate(s)
}
