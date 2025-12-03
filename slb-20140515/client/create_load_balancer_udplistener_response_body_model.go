// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerUDPListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLoadBalancerUDPListenerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerUDPListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 06F00FBB-3D9E-4CCE-9D43-1A6946A75556
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerUDPListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerUDPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerUDPListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
