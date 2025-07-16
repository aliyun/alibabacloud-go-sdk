// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse
	GetBody() *CreateGatewayResponseBody
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayResponse) GetBody() *CreateGatewayResponseBody {
	return s.Body
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayResponse) Validate() error {
	return dara.Validate(s)
}
