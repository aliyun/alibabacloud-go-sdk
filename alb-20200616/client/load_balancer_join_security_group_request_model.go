// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadBalancerJoinSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *LoadBalancerJoinSecurityGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *LoadBalancerJoinSecurityGroupRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *LoadBalancerJoinSecurityGroupRequest
	GetLoadBalancerId() *string
	SetSecurityGroupIds(v []*string) *LoadBalancerJoinSecurityGroupRequest
	GetSecurityGroupIds() []*string
}

type LoadBalancerJoinSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-h7kcw4g4nnvtqp****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The security group IDs.
	//
	// This parameter is required.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s LoadBalancerJoinSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadBalancerJoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LoadBalancerJoinSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *LoadBalancerJoinSecurityGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *LoadBalancerJoinSecurityGroupRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *LoadBalancerJoinSecurityGroupRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetClientToken(v string) *LoadBalancerJoinSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetDryRun(v bool) *LoadBalancerJoinSecurityGroupRequest {
	s.DryRun = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetLoadBalancerId(v string) *LoadBalancerJoinSecurityGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetSecurityGroupIds(v []*string) *LoadBalancerJoinSecurityGroupRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
