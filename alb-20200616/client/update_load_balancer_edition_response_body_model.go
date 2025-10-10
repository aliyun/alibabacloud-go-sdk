// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerEditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoadBalancerEditionResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerEditionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerEditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerEditionResponseBody) SetRequestId(v string) *UpdateLoadBalancerEditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerEditionResponseBody) Validate() error {
	return dara.Validate(s)
}
