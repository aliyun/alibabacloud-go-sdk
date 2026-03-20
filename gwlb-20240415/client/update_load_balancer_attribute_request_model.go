// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateLoadBalancerAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest
	GetLoadBalancerName() *string
	SetTrafficMode(v string) *UpdateLoadBalancerAttributeRequest
	GetTrafficMode() *string
}

type UpdateLoadBalancerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// testGwlbName
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// Specifies the traffic processing mode. Valid values:
	//
	// 	- **LoadBalance**: load balancing mode. In this mode, GWLB forwards traffic to backend servers.
	//
	// 	- **ByPass**: bypass mode. GWLB returns traffic directly to the GWLB endpoint instead of forwarding the traffic to backend servers.
	//
	// example:
	//
	// LoadBalance
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerAttributeRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *UpdateLoadBalancerAttributeRequest) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetTrafficMode(v string) *UpdateLoadBalancerAttributeRequest {
	s.TrafficMode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
