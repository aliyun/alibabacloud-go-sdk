// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTableEntryAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGatewayRouteTableEntryAttributeResponseBody
	GetRequestId() *string
}

type UpdateGatewayRouteTableEntryAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayRouteTableEntryAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTableEntryAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTableEntryAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteTableEntryAttributeResponseBody) SetRequestId(v string) *UpdateGatewayRouteTableEntryAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
