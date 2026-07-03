// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSupabaseInstancesRequest
	GetDBInstanceName() *string
	SetDescription(v string) *DescribeSupabaseInstancesRequest
	GetDescription() *string
	SetMaxResults(v int32) *DescribeSupabaseInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSupabaseInstancesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeSupabaseInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSupabaseInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSupabaseInstancesRequest
	GetRegionId() *string
}

type DescribeSupabaseInstancesRequest struct {
	// The instance name. This parameter is used for exact match filtering. Separate multiple instance names with commas (,).
	//
	// example:
	//
	// pxsp-hz********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance description. This parameter supports fuzzy match.
	//
	// example:
	//
	// project name
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 100. Default value: If you do not specify this parameter or specify a value less than 10, the default value is 10. If you specify a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// xxdds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSupabaseInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeSupabaseInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSupabaseInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSupabaseInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSupabaseInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSupabaseInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseInstancesRequest) SetDBInstanceName(v string) *DescribeSupabaseInstancesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetDescription(v string) *DescribeSupabaseInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetMaxResults(v int32) *DescribeSupabaseInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetNextToken(v string) *DescribeSupabaseInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetPageNumber(v int32) *DescribeSupabaseInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetPageSize(v int32) *DescribeSupabaseInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) SetRegionId(v string) *DescribeSupabaseInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseInstancesRequest) Validate() error {
	return dara.Validate(s)
}
