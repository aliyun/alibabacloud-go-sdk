// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcGatewayEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcGatewayEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcGatewayEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcGatewayEndpointResponseBody) *CreateVpcGatewayEndpointResponse
	GetBody() *CreateVpcGatewayEndpointResponseBody
}

type CreateVpcGatewayEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcGatewayEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcGatewayEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcGatewayEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcGatewayEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcGatewayEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcGatewayEndpointResponse) GetBody() *CreateVpcGatewayEndpointResponseBody {
	return s.Body
}

func (s *CreateVpcGatewayEndpointResponse) SetHeaders(v map[string]*string) *CreateVpcGatewayEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcGatewayEndpointResponse) SetStatusCode(v int32) *CreateVpcGatewayEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponse) SetBody(v *CreateVpcGatewayEndpointResponseBody) *CreateVpcGatewayEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateVpcGatewayEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
