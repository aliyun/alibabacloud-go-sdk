// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerTCPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerTCPListenerAttributeResponseBody
	GetRequestId() *string
}

type SetLoadBalancerTCPListenerAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
