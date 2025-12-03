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
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
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
