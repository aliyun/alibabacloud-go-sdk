// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerZonesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED8905C2-A4F6-5E43-87B7-6A5DC8757146
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerZonesResponseBody) SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
