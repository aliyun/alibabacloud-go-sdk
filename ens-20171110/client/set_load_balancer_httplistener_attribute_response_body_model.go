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
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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
