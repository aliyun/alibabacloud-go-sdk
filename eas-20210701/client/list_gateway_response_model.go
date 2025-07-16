// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayResponseBody) *ListGatewayResponse
	GetBody() *ListGatewayResponseBody
}

type ListGatewayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayResponse) GetBody() *ListGatewayResponseBody {
	return s.Body
}

func (s *ListGatewayResponse) SetHeaders(v map[string]*string) *ListGatewayResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayResponse) SetStatusCode(v int32) *ListGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayResponse) SetBody(v *ListGatewayResponseBody) *ListGatewayResponse {
	s.Body = v
	return s
}

func (s *ListGatewayResponse) Validate() error {
	return dara.Validate(s)
}
