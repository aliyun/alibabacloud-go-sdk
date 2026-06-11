// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossRegion(v bool) *DescribeZonesRequest
	GetCrossRegion() *bool
	SetCrossRegionSide(v string) *DescribeZonesRequest
	GetCrossRegionSide() *string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetServiceResourceType(v string) *DescribeZonesRequest
	GetServiceResourceType() *string
}

type DescribeZonesRequest struct {
	// Specifies whether this is a cross-region scenario. Default value: false.
	//
	// example:
	//
	// false
	CrossRegion *bool `json:"CrossRegion,omitempty" xml:"CrossRegion,omitempty"`
	// Specifies whether to query the active zones for the initiator side or the service side in a cross-region connection. Valid values:
	//
	// - **Endpoint*	- (default): endpoint.
	//
	// - **EndpointService**: endpoint service.
	//
	// > This parameter takes effect only when CrossRegion is set to true.
	//
	// example:
	//
	// EndpointService
	CrossRegionSide *string `json:"CrossRegionSide,omitempty" xml:"CrossRegionSide,omitempty"`
	// The ID of the region where the zones reside. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zones supported by PrivateLink in a region depend on the backend service resource type. You can specify the service resource type when querying the zones supported by PrivateLink. Valid values:
	//
	// - **slb*	- (default): Classic Load Balancer (CLB), the service resource type is classic load balancing.
	//
	// - **alb**: Application Load Balancer (ALB), the service resource type is application load balancing.
	//
	// - **nlb**: Network Load Balancer (NLB), the service resource type is network load balancing.
	//
	// - **gwlb**: Gateway Load Balancer (GWLB), the service resource type is gateway load balancing.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetCrossRegion() *bool {
	return s.CrossRegion
}

func (s *DescribeZonesRequest) GetCrossRegionSide() *string {
	return s.CrossRegionSide
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *DescribeZonesRequest) SetCrossRegion(v bool) *DescribeZonesRequest {
	s.CrossRegion = &v
	return s
}

func (s *DescribeZonesRequest) SetCrossRegionSide(v string) *DescribeZonesRequest {
	s.CrossRegionSide = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetServiceResourceType(v string) *DescribeZonesRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
