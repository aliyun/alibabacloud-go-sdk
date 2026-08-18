// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRegionsRequest
	GetRegionId() *string
	SetServiceResourceType(v string) *DescribeRegionsRequest
	GetServiceResourceType() *string
}

type DescribeRegionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The EPS resource type supported by PrivateLink. The available regions vary based on the service resource type. You can specify a service resource type when querying the regions where PrivateLink is available. Valid values:
	//
	// - **slb*	- (default): the service resource type is Classic Load Balancer (CLB).
	//
	// - **alb**: the service resource type is Application Load Balancer (ALB).
	//
	// - **nlb**: the service resource type is Network Load Balancer (NLB).
	//
	// - **gwlb**: the service resource type is Gateway Load Balancer (GWLB).
	//
	// - **ALL**: all of the preceding types.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsRequest) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetServiceResourceType(v string) *DescribeRegionsRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
