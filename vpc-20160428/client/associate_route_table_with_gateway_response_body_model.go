// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableWithGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateRouteTableWithGatewayResponseBody
	GetRequestId() *string
}

type AssociateRouteTableWithGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F89C2176-8F10-55EF-90CF-CF99D1E3F816
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateRouteTableWithGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableWithGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableWithGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateRouteTableWithGatewayResponseBody) SetRequestId(v string) *AssociateRouteTableWithGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
