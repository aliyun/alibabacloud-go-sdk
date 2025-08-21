// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerUDPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerUDPListenerAttributeResponseBody
	GetRequestId() *string
}

type SetLoadBalancerUDPListenerAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 06F00FBB-3D9E-4CCE-9D43-1A6946A75456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerUDPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerUDPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
