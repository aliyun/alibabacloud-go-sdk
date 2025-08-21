// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPSListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestId() *string
}

type SetLoadBalancerHTTPSListenerAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
