// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServersShrink(v string) *RemoveBackendServersShrinkRequest
	GetBackendServersShrink() *string
	SetLoadBalancerId(v string) *RemoveBackendServersShrinkRequest
	GetLoadBalancerId() *string
}

type RemoveBackendServersShrinkRequest struct {
	// The list of backend servers that you want to remove. You can remove up to 20 backend servers at a time.
	//
	// This parameter is required.
	BackendServersShrink *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5ovkn1piwqmoqrfjdyhq4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s RemoveBackendServersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersShrinkRequest) GetBackendServersShrink() *string {
	return s.BackendServersShrink
}

func (s *RemoveBackendServersShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveBackendServersShrinkRequest) SetBackendServersShrink(v string) *RemoveBackendServersShrinkRequest {
	s.BackendServersShrink = &v
	return s
}

func (s *RemoveBackendServersShrinkRequest) SetLoadBalancerId(v string) *RemoveBackendServersShrinkRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
