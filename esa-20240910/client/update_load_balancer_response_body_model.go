// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoadBalancerResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerResponseBody) SetRequestId(v string) *UpdateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
