// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewaysResponseBody) *ListGatewaysResponse
	GetBody() *ListGatewaysResponseBody
}

type ListGatewaysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewaysResponse) GetBody() *ListGatewaysResponseBody {
	return s.Body
}

func (s *ListGatewaysResponse) SetHeaders(v map[string]*string) *ListGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListGatewaysResponse) SetStatusCode(v int32) *ListGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewaysResponse) SetBody(v *ListGatewaysResponseBody) *ListGatewaysResponse {
	s.Body = v
	return s
}

func (s *ListGatewaysResponse) Validate() error {
	return dara.Validate(s)
}
