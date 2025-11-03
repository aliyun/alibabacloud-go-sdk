// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcGatewayEndpointAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcGatewayEndpointAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcGatewayEndpointAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcGatewayEndpointAttributeResponseBody) *GetVpcGatewayEndpointAttributeResponse
	GetBody() *GetVpcGatewayEndpointAttributeResponseBody
}

type GetVpcGatewayEndpointAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcGatewayEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcGatewayEndpointAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcGatewayEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcGatewayEndpointAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcGatewayEndpointAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcGatewayEndpointAttributeResponse) GetBody() *GetVpcGatewayEndpointAttributeResponseBody {
	return s.Body
}

func (s *GetVpcGatewayEndpointAttributeResponse) SetHeaders(v map[string]*string) *GetVpcGatewayEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponse) SetStatusCode(v int32) *GetVpcGatewayEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponse) SetBody(v *GetVpcGatewayEndpointAttributeResponseBody) *GetVpcGatewayEndpointAttributeResponse {
	s.Body = v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
