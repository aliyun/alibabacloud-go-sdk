// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeRCClusterNodesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRCClusterNodesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeRCClusterNodesRequest
	GetRegionId() *string
	SetVpcId(v string) *DescribeRCClusterNodesRequest
	GetVpcId() *string
}

type DescribeRCClusterNodesRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// >  This is a reserved parameter.
	//
	// example:
	//
	// None
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterNodesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCClusterNodesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCClusterNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCClusterNodesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCClusterNodesRequest) SetPageNumber(v int64) *DescribeRCClusterNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCClusterNodesRequest) SetPageSize(v int64) *DescribeRCClusterNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCClusterNodesRequest) SetRegionId(v string) *DescribeRCClusterNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCClusterNodesRequest) SetVpcId(v string) *DescribeRCClusterNodesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRCClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
