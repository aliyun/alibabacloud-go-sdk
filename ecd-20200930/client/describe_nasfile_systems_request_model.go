// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNASFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v []*string) *DescribeNASFileSystemsRequest
	GetFileSystemId() []*string
	SetMatchCompatibleProfile(v bool) *DescribeNASFileSystemsRequest
	GetMatchCompatibleProfile() *bool
	SetMaxResults(v int32) *DescribeNASFileSystemsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNASFileSystemsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeNASFileSystemsRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeNASFileSystemsRequest
	GetRegionId() *string
}

type DescribeNASFileSystemsRequest struct {
	// The IDs of the NAS file systems.
	//
	// example:
	//
	// 04f314****
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
	// Specifies whether to include only NAS file systems that support the User Profile Management (UPM) feature in the query results.
	//
	// example:
	//
	// false
	MatchCompatibleProfile *bool `json:"MatchCompatibleProfile,omitempty" xml:"MatchCompatibleProfile,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNASFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsRequest) GetFileSystemId() []*string {
	return s.FileSystemId
}

func (s *DescribeNASFileSystemsRequest) GetMatchCompatibleProfile() *bool {
	return s.MatchCompatibleProfile
}

func (s *DescribeNASFileSystemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNASFileSystemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNASFileSystemsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeNASFileSystemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNASFileSystemsRequest) SetFileSystemId(v []*string) *DescribeNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetMatchCompatibleProfile(v bool) *DescribeNASFileSystemsRequest {
	s.MatchCompatibleProfile = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetMaxResults(v int32) *DescribeNASFileSystemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetNextToken(v string) *DescribeNASFileSystemsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetOfficeSiteId(v string) *DescribeNASFileSystemsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetRegionId(v string) *DescribeNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
