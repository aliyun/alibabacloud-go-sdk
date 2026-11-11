// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoadBalancerProtectionResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerProtectionResponseBody) SetRequestId(v string) *UpdateLoadBalancerProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
