// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeInstanceEndpointsRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceEndpointsRequest
	GetRegionId() *string
}

type DescribeInstanceEndpointsRequest struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceEndpointsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceEndpointsRequest) SetInstanceName(v string) *DescribeInstanceEndpointsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceEndpointsRequest) SetRegionId(v string) *DescribeInstanceEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
