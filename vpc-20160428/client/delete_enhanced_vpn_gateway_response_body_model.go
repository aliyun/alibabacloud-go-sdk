// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnhancedVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnhancedVpnGatewayResponseBody
	GetRequestId() *string
}

type DeleteEnhancedVpnGatewayResponseBody struct {
	// example:
	//
	// E9A8AABE-A84B-4AF2-A68A-8E2EA190E7AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnhancedVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnhancedVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnhancedVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnhancedVpnGatewayResponseBody) SetRequestId(v string) *DeleteEnhancedVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
