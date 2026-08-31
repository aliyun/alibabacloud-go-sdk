// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchInstancesRequest
	GetDBInstanceName() *string
	SetMaxResults(v int32) *DescribeOpenSearchInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeOpenSearchInstancesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeOpenSearchInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeOpenSearchInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeOpenSearchInstancesRequest
	GetRegionId() *string
}

type DescribeOpenSearchInstancesRequest struct {
	// The instance name.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The maximum number of entries per page for a paging query. Maximum value: 100. Default value: If you do not specify a value or the value is less than 10, the default value is 10. If the value is greater than 100, the default value is 100.
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeOpenSearchInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOpenSearchInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeOpenSearchInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOpenSearchInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchInstancesRequest) SetDBInstanceName(v string) *DescribeOpenSearchInstancesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) SetMaxResults(v int32) *DescribeOpenSearchInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) SetNextToken(v string) *DescribeOpenSearchInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) SetPageNumber(v int32) *DescribeOpenSearchInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) SetPageSize(v int32) *DescribeOpenSearchInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) SetRegionId(v string) *DescribeOpenSearchInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchInstancesRequest) Validate() error {
	return dara.Validate(s)
}
