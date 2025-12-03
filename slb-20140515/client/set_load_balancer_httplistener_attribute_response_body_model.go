// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestId() *string
}

type SetLoadBalancerHTTPListenerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
