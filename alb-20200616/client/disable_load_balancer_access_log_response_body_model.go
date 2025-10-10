// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLoadBalancerAccessLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableLoadBalancerAccessLogResponseBody
	GetRequestId() *string
}

type DisableLoadBalancerAccessLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerAccessLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *DisableLoadBalancerAccessLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableLoadBalancerAccessLogResponseBody) Validate() error {
	return dara.Validate(s)
}
