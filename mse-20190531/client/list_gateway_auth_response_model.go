// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayAuthResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayAuthResponseBody) *ListGatewayAuthResponse
	GetBody() *ListGatewayAuthResponseBody
}

type ListGatewayAuthResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayAuthResponse) GetBody() *ListGatewayAuthResponseBody {
	return s.Body
}

func (s *ListGatewayAuthResponse) SetHeaders(v map[string]*string) *ListGatewayAuthResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayAuthResponse) SetStatusCode(v int32) *ListGatewayAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayAuthResponse) SetBody(v *ListGatewayAuthResponseBody) *ListGatewayAuthResponse {
	s.Body = v
	return s
}

func (s *ListGatewayAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
