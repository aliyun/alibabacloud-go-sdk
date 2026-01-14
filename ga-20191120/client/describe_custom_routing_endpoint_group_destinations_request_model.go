// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest
	GetDestinationId() *string
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest
	GetRegionId() *string
}

type DescribeCustomRoutingEndpointGroupDestinationsRequest struct {
	// The ID of the endpoint group mapping configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// dst-123abc****
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1j184jhb9i9ubwf****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) GetDestinationId() *string {
	return s.DestinationId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) Validate() error {
	return dara.Validate(s)
}
