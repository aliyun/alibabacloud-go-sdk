// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
}

type GetLoadBalancerAttributeRequest struct {
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
