// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRouteEntryListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRouteEntryListRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeRouteEntryListRequest
	GetRegionId() *string
	SetRouteTableId(v string) *DescribeRouteEntryListRequest
	GetRouteTableId() *string
}

type DescribeRouteEntryListRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeRouteEntryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRouteEntryListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteEntryListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteEntryListRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteEntryListRequest) SetMaxResults(v int32) *DescribeRouteEntryListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetNextToken(v string) *DescribeRouteEntryListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRegionId(v string) *DescribeRouteEntryListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRouteTableId(v string) *DescribeRouteEntryListRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) Validate() error {
	return dara.Validate(s)
}
