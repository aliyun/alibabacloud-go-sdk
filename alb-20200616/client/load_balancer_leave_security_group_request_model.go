// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadBalancerLeaveSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *LoadBalancerLeaveSecurityGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *LoadBalancerLeaveSecurityGroupRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *LoadBalancerLeaveSecurityGroupRequest
	GetLoadBalancerId() *string
	SetSecurityGroupIds(v []*string) *LoadBalancerLeaveSecurityGroupRequest
	GetSecurityGroupIds() []*string
}

type LoadBalancerLeaveSecurityGroupRequest struct {
	// The task result.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-iv9gj3lpak6fbj****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The security group IDs.
	//
	// This parameter is required.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s LoadBalancerLeaveSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadBalancerLeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LoadBalancerLeaveSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *LoadBalancerLeaveSecurityGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *LoadBalancerLeaveSecurityGroupRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *LoadBalancerLeaveSecurityGroupRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetClientToken(v string) *LoadBalancerLeaveSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetDryRun(v bool) *LoadBalancerLeaveSecurityGroupRequest {
	s.DryRun = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetLoadBalancerId(v string) *LoadBalancerLeaveSecurityGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetSecurityGroupIds(v []*string) *LoadBalancerLeaveSecurityGroupRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
