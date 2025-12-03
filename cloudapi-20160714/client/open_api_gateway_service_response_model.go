// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenApiGatewayServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenApiGatewayServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenApiGatewayServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenApiGatewayServiceResponseBody) *OpenApiGatewayServiceResponse
	GetBody() *OpenApiGatewayServiceResponseBody
}

type OpenApiGatewayServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenApiGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenApiGatewayServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenApiGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenApiGatewayServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenApiGatewayServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenApiGatewayServiceResponse) GetBody() *OpenApiGatewayServiceResponseBody {
	return s.Body
}

func (s *OpenApiGatewayServiceResponse) SetHeaders(v map[string]*string) *OpenApiGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenApiGatewayServiceResponse) SetStatusCode(v int32) *OpenApiGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenApiGatewayServiceResponse) SetBody(v *OpenApiGatewayServiceResponseBody) *OpenApiGatewayServiceResponse {
	s.Body = v
	return s
}

func (s *OpenApiGatewayServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
