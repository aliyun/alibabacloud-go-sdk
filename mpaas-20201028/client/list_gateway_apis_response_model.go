// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayApisResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayApisResponseBody) *ListGatewayApisResponse
	GetBody() *ListGatewayApisResponseBody
}

type ListGatewayApisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayApisResponse) GetBody() *ListGatewayApisResponseBody {
	return s.Body
}

func (s *ListGatewayApisResponse) SetHeaders(v map[string]*string) *ListGatewayApisResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayApisResponse) SetStatusCode(v int32) *ListGatewayApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayApisResponse) SetBody(v *ListGatewayApisResponseBody) *ListGatewayApisResponse {
	s.Body = v
	return s
}

func (s *ListGatewayApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
