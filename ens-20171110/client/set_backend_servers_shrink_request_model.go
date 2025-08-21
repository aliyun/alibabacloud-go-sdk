// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackendServersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServersShrink(v string) *SetBackendServersShrinkRequest
	GetBackendServersShrink() *string
	SetLoadBalancerId(v string) *SetBackendServersShrinkRequest
	GetLoadBalancerId() *string
}

type SetBackendServersShrinkRequest struct {
	// The list of backend servers that you added. You can modify the weights of up to 20 backend servers in each request.
	//
	// This parameter is required.
	BackendServersShrink *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5s7crik3yo3bp03gqrbp5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s SetBackendServersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetBackendServersShrinkRequest) GetBackendServersShrink() *string {
	return s.BackendServersShrink
}

func (s *SetBackendServersShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetBackendServersShrinkRequest) SetBackendServersShrink(v string) *SetBackendServersShrinkRequest {
	s.BackendServersShrink = &v
	return s
}

func (s *SetBackendServersShrinkRequest) SetLoadBalancerId(v string) *SetBackendServersShrinkRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
