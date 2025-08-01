// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGatewayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGatewayTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryGatewayTypeResponseBody) *QueryGatewayTypeResponse
	GetBody() *QueryGatewayTypeResponseBody
}

type QueryGatewayTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGatewayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGatewayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryGatewayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGatewayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGatewayTypeResponse) GetBody() *QueryGatewayTypeResponseBody {
	return s.Body
}

func (s *QueryGatewayTypeResponse) SetHeaders(v map[string]*string) *QueryGatewayTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryGatewayTypeResponse) SetStatusCode(v int32) *QueryGatewayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGatewayTypeResponse) SetBody(v *QueryGatewayTypeResponseBody) *QueryGatewayTypeResponse {
	s.Body = v
	return s
}

func (s *QueryGatewayTypeResponse) Validate() error {
	return dara.Validate(s)
}
