// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcGatewayEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcGatewayEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcGatewayEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcGatewayEndpointResponseBody) *DeleteVpcGatewayEndpointResponse
	GetBody() *DeleteVpcGatewayEndpointResponseBody
}

type DeleteVpcGatewayEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcGatewayEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcGatewayEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcGatewayEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcGatewayEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcGatewayEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcGatewayEndpointResponse) GetBody() *DeleteVpcGatewayEndpointResponseBody {
	return s.Body
}

func (s *DeleteVpcGatewayEndpointResponse) SetHeaders(v map[string]*string) *DeleteVpcGatewayEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcGatewayEndpointResponse) SetStatusCode(v int32) *DeleteVpcGatewayEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcGatewayEndpointResponse) SetBody(v *DeleteVpcGatewayEndpointResponseBody) *DeleteVpcGatewayEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcGatewayEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
