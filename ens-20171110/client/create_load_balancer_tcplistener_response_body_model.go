// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerTCPListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLoadBalancerTCPListenerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerTCPListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 06F00FBB-3D9E-4CCE-9D43-1A6946A75456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerTCPListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerTCPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerTCPListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
