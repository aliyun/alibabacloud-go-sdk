// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndPointTrafficPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest
	GetEndpointId() *string
	SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest
	GetRegionId() *string
}

type DescribeCustomRoutingEndPointTrafficPolicyRequest struct {
	// The ID of the traffic policy to be queried.
	//
	// example:
	//
	// ep-bp1d2utp8qqe2a44t****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ply-bptest2****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The operation that you want to perform. Set the value to **DescribeCustomRoutingEndPointTrafficPolicy**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) SetEndpointId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) SetRegionId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) Validate() error {
	return dara.Validate(s)
}
