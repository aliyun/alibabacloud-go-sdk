// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *StartLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *StartLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *StartLoadBalancerListenerRequest
	GetLoadBalancerId() *string
}

type StartLoadBalancerListenerRequest struct {
	// The listener port to be enabled. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the ELB instance. Valid values:
	//
	// 	- tcp
	//
	// 	- udp
	//
	// 	- http
	//
	// 	- https
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
	// lb-5saivuir6b1mupxjfbhmk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s StartLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *StartLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *StartLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *StartLoadBalancerListenerRequest) SetListenerPort(v int32) *StartLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetListenerProtocol(v string) *StartLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StartLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
