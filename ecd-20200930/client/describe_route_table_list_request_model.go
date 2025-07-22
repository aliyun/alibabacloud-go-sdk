// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRouteTableListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRouteTableListRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeRouteTableListRequest
	GetRegionId() *string
	SetRouteTableId(v string) *DescribeRouteTableListRequest
	GetRouteTableId() *string
	SetRouteTableName(v string) *DescribeRouteTableListRequest
	GetRouteTableName() *string
	SetVpcId(v string) *DescribeRouteTableListRequest
	GetVpcId() *string
}

type DescribeRouteTableListRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteTableId   *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRouteTableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRouteTableListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteTableListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteTableListRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTableListRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeRouteTableListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRouteTableListRequest) SetMaxResults(v int32) *DescribeRouteTableListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetNextToken(v string) *DescribeRouteTableListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRegionId(v string) *DescribeRouteTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouteTableId(v string) *DescribeRouteTableListRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouteTableName(v string) *DescribeRouteTableListRequest {
	s.RouteTableName = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetVpcId(v string) *DescribeRouteTableListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRouteTableListRequest) Validate() error {
	return dara.Validate(s)
}
