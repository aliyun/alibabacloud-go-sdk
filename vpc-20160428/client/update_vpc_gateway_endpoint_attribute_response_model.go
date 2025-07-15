// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcGatewayEndpointAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcGatewayEndpointAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcGatewayEndpointAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcGatewayEndpointAttributeResponseBody) *UpdateVpcGatewayEndpointAttributeResponse
	GetBody() *UpdateVpcGatewayEndpointAttributeResponseBody
}

type UpdateVpcGatewayEndpointAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcGatewayEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcGatewayEndpointAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcGatewayEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) GetBody() *UpdateVpcGatewayEndpointAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcGatewayEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) SetStatusCode(v int32) *UpdateVpcGatewayEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) SetBody(v *UpdateVpcGatewayEndpointAttributeResponseBody) *UpdateVpcGatewayEndpointAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeResponse) Validate() error {
	return dara.Validate(s)
}
