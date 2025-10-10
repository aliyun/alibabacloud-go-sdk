// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerEditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateLoadBalancerEditionRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateLoadBalancerEditionRequest
	GetDryRun() *bool
	SetLoadBalancerEdition(v string) *UpdateLoadBalancerEditionRequest
	GetLoadBalancerEdition() *string
	SetLoadBalancerId(v string) *UpdateLoadBalancerEditionRequest
	GetLoadBalancerId() *string
}

type UpdateLoadBalancerEditionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The edition of the ALB instance. Different editions have different limits and support different billing methods.
	//
	// 	- **Basic**: basic
	//
	// 	- **Standard**: standard
	//
	// 	- **StandardWithWaf**: WAF-enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// Standard
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s UpdateLoadBalancerEditionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerEditionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerEditionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateLoadBalancerEditionRequest) GetLoadBalancerEdition() *string {
	return s.LoadBalancerEdition
}

func (s *UpdateLoadBalancerEditionRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerEditionRequest) SetClientToken(v string) *UpdateLoadBalancerEditionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetDryRun(v bool) *UpdateLoadBalancerEditionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerEdition(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) Validate() error {
	return dara.Validate(s)
}
