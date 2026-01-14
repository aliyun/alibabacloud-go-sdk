// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *DescribeEndpointGroupRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *DescribeEndpointGroupRequest
	GetRegionId() *string
}

type DescribeEndpointGroupRequest struct {
	// The ID of the endpoint group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where your Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeEndpointGroupRequest) SetRegionId(v string) *DescribeEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
