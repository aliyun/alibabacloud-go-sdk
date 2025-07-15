// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpnGatewayResponseBody
	GetRequestId() *string
}

type DeleteVpnGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// >0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpnGatewayResponseBody) SetRequestId(v string) *DeleteVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
