// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayAuthResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayAuthResponseBody) *AddGatewayAuthResponse
	GetBody() *AddGatewayAuthResponseBody
}

type AddGatewayAuthResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayAuthResponse) GetBody() *AddGatewayAuthResponseBody {
	return s.Body
}

func (s *AddGatewayAuthResponse) SetHeaders(v map[string]*string) *AddGatewayAuthResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayAuthResponse) SetStatusCode(v int32) *AddGatewayAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayAuthResponse) SetBody(v *AddGatewayAuthResponseBody) *AddGatewayAuthResponse {
	s.Body = v
	return s
}

func (s *AddGatewayAuthResponse) Validate() error {
	return dara.Validate(s)
}
