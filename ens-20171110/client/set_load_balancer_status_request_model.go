// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerId() *string
	SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerStatus() *string
}

type SetLoadBalancerStatusRequest struct {
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5t18quoohsrc3xkf86spmnu77
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The new instance status. Valid values:
	//
	// 	- **Active**
	//
	// 	- **InActive**
	//
	// This parameter is required.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
}

func (s SetLoadBalancerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) Validate() error {
	return dara.Validate(s)
}
