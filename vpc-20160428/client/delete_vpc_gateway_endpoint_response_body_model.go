// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcGatewayEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcGatewayEndpointResponseBody
	GetRequestId() *string
}

type DeleteVpcGatewayEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1122D0F-7B3B-5445-BB19-17F27F97FE1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcGatewayEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcGatewayEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcGatewayEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcGatewayEndpointResponseBody) SetRequestId(v string) *DeleteVpcGatewayEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcGatewayEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
