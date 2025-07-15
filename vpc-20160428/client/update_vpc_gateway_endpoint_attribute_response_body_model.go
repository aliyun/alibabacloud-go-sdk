// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcGatewayEndpointAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVpcGatewayEndpointAttributeResponseBody
	GetRequestId() *string
}

type UpdateVpcGatewayEndpointAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E9654534-5A38-5545-813F-0403D49042FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcGatewayEndpointAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcGatewayEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcGatewayEndpointAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVpcGatewayEndpointAttributeResponseBody) SetRequestId(v string) *UpdateVpcGatewayEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
