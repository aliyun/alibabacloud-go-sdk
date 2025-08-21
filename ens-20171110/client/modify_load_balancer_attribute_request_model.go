// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *ModifyLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *ModifyLoadBalancerAttributeRequest
	GetLoadBalancerName() *string
}

type ModifyLoadBalancerAttributeRequest struct {
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4gtm8
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ELB instance. The name must be **2*	- to **128*	- characters in length.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
}

func (s ModifyLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyLoadBalancerAttributeRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *ModifyLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *ModifyLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *ModifyLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
