// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerListenerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerListenerStatusResponseBody
	GetRequestId() *string
}

type SetLoadBalancerListenerStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerListenerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerListenerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerListenerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerListenerStatusResponseBody) SetRequestId(v string) *SetLoadBalancerListenerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerListenerStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
