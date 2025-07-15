// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6GatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpv6GatewayResponseBody
	GetRequestId() *string
}

type DeleteIpv6GatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E9A8AABE-A84B-4AF2-A68A-8E2EA190E7AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpv6GatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6GatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpv6GatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpv6GatewayResponseBody) SetRequestId(v string) *DeleteIpv6GatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpv6GatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
