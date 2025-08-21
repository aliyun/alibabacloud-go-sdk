// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *StopLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *StopLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *StopLoadBalancerListenerRequest
	GetLoadBalancerId() *string
}

type StopLoadBalancerListenerRequest struct {
	// The listener port that you want to disable. Valid values: **1*	- to **65535**.
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
	// lb-5sc1s9zrui8lpb8u7cl4f****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s StopLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *StopLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *StopLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *StopLoadBalancerListenerRequest) SetListenerPort(v int32) *StopLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetListenerProtocol(v string) *StopLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StopLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
