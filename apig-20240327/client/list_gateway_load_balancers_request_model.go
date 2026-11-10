// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListGatewayLoadBalancersRequest
	GetAll() *bool
	SetLoadBalancerId(v string) *ListGatewayLoadBalancersRequest
	GetLoadBalancerId() *string
	SetNetwork(v string) *ListGatewayLoadBalancersRequest
	GetNetwork() *string
	SetRelated(v bool) *ListGatewayLoadBalancersRequest
	GetRelated() *bool
	SetType(v string) *ListGatewayLoadBalancersRequest
	GetType() *string
	SetVpcId(v string) *ListGatewayLoadBalancersRequest
	GetVpcId() *string
}

type ListGatewayLoadBalancersRequest struct {
	// Specifies whether to return all instances in the optional list mode. This parameter affects the pagination and filtering strategies.
	//
	// example:
	//
	// false
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The exact load balancer instance ID used to query details by ID in the optional list mode.
	//
	// example:
	//
	// lb-xxxx
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The network type filter.
	//
	// example:
	//
	// Internet
	Network *string `json:"network,omitempty" xml:"network,omitempty"`
	// Specifies whether to return the list of load balancers associated with the gateway. If set to true, the associated load balancers are returned based on database relationships. We recommend setting this parameter to true to retrieve the associated load balancers. If set to false or left empty, the optional list mode is used to call the SLB or NLB API and filter bindable instances. In the default mode where related is false, calling the SLB or NLB API may return a 500 error due to backend data issues in common scenarios.
	//
	// example:
	//
	// false
	Related *bool `json:"related,omitempty" xml:"related,omitempty"`
	// The load balancer type filter.
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The VPC filter used to filter instances within the specified VPC in the optional list mode. This parameter is required when type is set to NLB and related is not set to true. Otherwise, the API returns a 500 error.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListGatewayLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersRequest) GetAll() *bool {
	return s.All
}

func (s *ListGatewayLoadBalancersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListGatewayLoadBalancersRequest) GetNetwork() *string {
	return s.Network
}

func (s *ListGatewayLoadBalancersRequest) GetRelated() *bool {
	return s.Related
}

func (s *ListGatewayLoadBalancersRequest) GetType() *string {
	return s.Type
}

func (s *ListGatewayLoadBalancersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayLoadBalancersRequest) SetAll(v bool) *ListGatewayLoadBalancersRequest {
	s.All = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) SetLoadBalancerId(v string) *ListGatewayLoadBalancersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) SetNetwork(v string) *ListGatewayLoadBalancersRequest {
	s.Network = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) SetRelated(v bool) *ListGatewayLoadBalancersRequest {
	s.Related = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) SetType(v string) *ListGatewayLoadBalancersRequest {
	s.Type = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) SetVpcId(v string) *ListGatewayLoadBalancersRequest {
	s.VpcId = &v
	return s
}

func (s *ListGatewayLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}
