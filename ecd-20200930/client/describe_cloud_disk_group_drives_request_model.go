// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupDrivesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCloudDiskGroupDrivesRequest
	GetCdsId() *string
	SetGroupName(v string) *DescribeCloudDiskGroupDrivesRequest
	GetGroupName() *string
	SetMaxResults(v int32) *DescribeCloudDiskGroupDrivesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudDiskGroupDrivesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeCloudDiskGroupDrivesRequest
	GetRegionId() *string
}

type DescribeCloudDiskGroupDrivesRequest struct {
	// The enterprise cloud disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-6805637***
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The team space name.
	//
	// example:
	//
	// Team1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100
	//
	// - Default value: 20
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. If NextToken is empty, no more results are available.
	//
	// example:
	//
	// MTA0MjA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudDiskGroupDrivesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupDrivesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupDrivesRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCloudDiskGroupDrivesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDiskGroupDrivesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudDiskGroupDrivesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDiskGroupDrivesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudDiskGroupDrivesRequest) SetCdsId(v string) *DescribeCloudDiskGroupDrivesRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesRequest) SetGroupName(v string) *DescribeCloudDiskGroupDrivesRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesRequest) SetMaxResults(v int32) *DescribeCloudDiskGroupDrivesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesRequest) SetNextToken(v string) *DescribeCloudDiskGroupDrivesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesRequest) SetRegionId(v string) *DescribeCloudDiskGroupDrivesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesRequest) Validate() error {
	return dara.Validate(s)
}
