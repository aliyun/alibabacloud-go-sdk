// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGatewayRouteResponseBody
	GetRequestId() *string
}

type DeleteGatewayRouteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayRouteResponseBody) SetRequestId(v string) *DeleteGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
