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
	// Specifies whether to return all instances in the selectable list mode. This parameter affects pagination and filtering behavior.
	//
	// example:
	//
	// false
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The load balancer instance ID. Specify this parameter to query details of a specific load balancer in the selectable list mode.
	//
	// example:
	//
	// lb-xxxx
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The network type used to filter results.
	//
	// example:
	//
	// Internet
	Network *string `json:"network,omitempty" xml:"network,omitempty"`
	// The mode switch. Valid values:
	//
	// - true: Returns the list of load balancers that are already associated with the gateway by reading database relationships.
	//
	// - false (default): Returns the selectable list by calling SLB or NLB OpenAPI to filter bindable instances.
	//
	// Use related=true to retrieve the list of associated load balancers. The default mode (related=false) calls SLB or NLB OpenAPI to filter bindable instances, which may return a 500 error in common scenarios due to backend data issues.
	//
	// example:
	//
	// false
	Related *bool `json:"related,omitempty" xml:"related,omitempty"`
	// The load balancer type used to filter results.
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The VPC ID used to filter instances within the specified VPC in the selectable list mode. This parameter is required when type is set to NLB and related is not set to true. Otherwise, the API returns a 500 error.
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
