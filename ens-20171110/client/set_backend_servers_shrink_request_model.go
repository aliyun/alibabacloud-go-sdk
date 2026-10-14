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
	// The list of backend servers to add. You can modify up to 20 backend servers in each call.
	//
	// This parameter is required.
	BackendServersShrink *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the load balancing instance.
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
