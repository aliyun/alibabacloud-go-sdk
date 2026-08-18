// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceEndpointRequest
	GetDBInstanceName() *string
	SetMaxResults(v int32) *DescribeDBInstanceEndpointRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDBInstanceEndpointRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDBInstanceEndpointRequest
	GetRegionId() *string
}

type DescribeDBInstanceEndpointRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100. Default value: If the value is not specified or is less than 10, the default value is 10. If the value is greater than 100, the default value is 100.
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
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceEndpointRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDBInstanceEndpointRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDBInstanceEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceEndpointRequest) SetDBInstanceName(v string) *DescribeDBInstanceEndpointRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceEndpointRequest) SetMaxResults(v int32) *DescribeDBInstanceEndpointRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDBInstanceEndpointRequest) SetNextToken(v string) *DescribeDBInstanceEndpointRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDBInstanceEndpointRequest) SetRegionId(v string) *DescribeDBInstanceEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
