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
	// Set CrossRegion to true and use it together with CrossRegionSide in the following scenarios:
	//
	// - **As a service consumer**: You need to create a cross-region endpoint to connect to an endpoint service in another region. To query the zones in the current region that support cross-region endpoints, set CrossRegion to true and CrossRegionSide to Endpoint.
	//
	// - **As a service provider**: You need to share your endpoint service across regions with service consumers in other regions. To query the zones in the current region that support cross-region sharing, set CrossRegion to true and CrossRegionSide to EndpointService.
	//
	// example:
	//
	// false
	CrossRegion *bool `json:"CrossRegion,omitempty" xml:"CrossRegion,omitempty"`
	// Specifies whether to query the active zone support for the initiator side or the service side in a cross-region connection. Valid values:
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
	// The ID of the region where the zones reside. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone availability of PrivateLink in a region depends on the backend EPS resource type. You can specify the EPS resource type when querying the list of zones supported by PrivateLink. Valid values:
	//
	// - **slb*	- (default): the EPS resource type is Classic Load Balancer (CLB).
	//
	// - **alb**: the EPS resource type is Application Load Balancer (ALB).
	//
	// - **nlb**: the EPS resource type is Network Load Balancer (NLB).
	//
	// - **gwlb**: the EPS resource type is Gateway Load Balancer (GWLB).
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
