// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServersShrink(v string) *AddBackendServersShrinkRequest
	GetBackendServersShrink() *string
	SetLoadBalancerId(v string) *AddBackendServersShrinkRequest
	GetLoadBalancerId() *string
}

type AddBackendServersShrinkRequest struct {
	// The list of backend servers to add. You can add up to 20 backend servers per invoke.
	//
	// > Backend servers (ENS instances) must be in the Running state before they can be added to a load balancing instance.
	//
	// This parameter is required.
	BackendServersShrink *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The frontend port used by the load balancing instance. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5qzdmxefgrpxd7oz2mefonvtx
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s AddBackendServersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddBackendServersShrinkRequest) GetBackendServersShrink() *string {
	return s.BackendServersShrink
}

func (s *AddBackendServersShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddBackendServersShrinkRequest) SetBackendServersShrink(v string) *AddBackendServersShrinkRequest {
	s.BackendServersShrink = &v
	return s
}

func (s *AddBackendServersShrinkRequest) SetLoadBalancerId(v string) *AddBackendServersShrinkRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
