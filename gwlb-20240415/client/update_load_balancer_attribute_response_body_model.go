// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B956C629-0E8C-5EFF-BAC1-B0E3A8C5CBDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
