// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayServiceResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayServiceResponseBody) *ListGatewayServiceResponse
	GetBody() *ListGatewayServiceResponseBody
}

type ListGatewayServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayServiceResponse) GetBody() *ListGatewayServiceResponseBody {
	return s.Body
}

func (s *ListGatewayServiceResponse) SetHeaders(v map[string]*string) *ListGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayServiceResponse) SetStatusCode(v int32) *ListGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayServiceResponse) SetBody(v *ListGatewayServiceResponseBody) *ListGatewayServiceResponse {
	s.Body = v
	return s
}

func (s *ListGatewayServiceResponse) Validate() error {
	return dara.Validate(s)
}
