// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateLoadBalancerResponseBody
	GetAddress() *string
	SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerName() *string
	SetRequestId(v string) *CreateLoadBalancerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerResponseBody struct {
	Address          *string `json:"Address,omitempty" xml:"Address,omitempty"`
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) GetAddress() *string {
	return s.Address
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerResponseBody) SetAddress(v string) *CreateLoadBalancerResponseBody {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
