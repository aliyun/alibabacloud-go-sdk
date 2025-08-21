// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DeleteLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest
	GetLoadBalancerId() *string
}

type DeleteLoadBalancerListenerRequest struct {
	// The frontend port that is used by the Edge Load Balance (ELB) instance. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the ELB instance.
	//
	// >  This parameter is required if the same port is used by listeners that use different protocols.
	//
	// example:
	//
	// tcp
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DeleteLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DeleteLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DeleteLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerProtocol(v string) *DeleteLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
