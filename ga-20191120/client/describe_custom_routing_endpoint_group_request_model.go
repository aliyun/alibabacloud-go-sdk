// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *DescribeCustomRoutingEndpointGroupRequest
	GetRegionId() *string
}

type DescribeCustomRoutingEndpointGroupRequest struct {
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaua****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomRoutingEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
