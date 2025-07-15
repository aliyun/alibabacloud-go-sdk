// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeAclEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAclEntriesRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeAclEntriesRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeAclEntriesRequest
	GetRegionId() *string
	SetSourceId(v string) *DescribeAclEntriesRequest
	GetSourceId() *string
	SetSourceType(v string) *DescribeAclEntriesRequest
	GetSourceType() *string
}

type DescribeAclEntriesRequest struct {
	// The number of entries per page.
	//
	// 	- Maximum value: 1600.
	//
	// 	- Default value: 1600.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kRxd1mKkNnHlUy14zdjl/I
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-shanghai+dir-631324****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance to which the ACL applies. You can specify an office network ID or a cloud computer ID.
	//
	// example:
	//
	// cn-hangzhou+dir-****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The granularity of the ACL.
	//
	// Valid values:
	//
	// 	- desktop: cloud computer
	//
	// 	- vpc: office network
	//
	// example:
	//
	// desktop
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeAclEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAclEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAclEntriesRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeAclEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAclEntriesRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeAclEntriesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAclEntriesRequest) SetMaxResults(v int32) *DescribeAclEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAclEntriesRequest) SetNextToken(v string) *DescribeAclEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAclEntriesRequest) SetOfficeSiteId(v string) *DescribeAclEntriesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeAclEntriesRequest) SetRegionId(v string) *DescribeAclEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAclEntriesRequest) SetSourceId(v string) *DescribeAclEntriesRequest {
	s.SourceId = &v
	return s
}

func (s *DescribeAclEntriesRequest) SetSourceType(v string) *DescribeAclEntriesRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeAclEntriesRequest) Validate() error {
	return dara.Validate(s)
}
