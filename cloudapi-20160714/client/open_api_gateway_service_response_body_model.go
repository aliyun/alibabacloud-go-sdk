// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenApiGatewayServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenApiGatewayServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenApiGatewayServiceResponseBody
	GetRequestId() *string
}

type OpenApiGatewayServiceResponseBody struct {
	// example:
	//
	// 210981***530495
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 2D39D1B3-8548-508A-9CE2-7F4A3F2A7989
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenApiGatewayServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenApiGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenApiGatewayServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenApiGatewayServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenApiGatewayServiceResponseBody) SetOrderId(v string) *OpenApiGatewayServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenApiGatewayServiceResponseBody) SetRequestId(v string) *OpenApiGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenApiGatewayServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
