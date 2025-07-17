// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayResponseBody) *GetGatewayResponse
	GetBody() *GetGatewayResponseBody
}

type GetGatewayResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayResponse) GetBody() *GetGatewayResponseBody {
	return s.Body
}

func (s *GetGatewayResponse) SetHeaders(v map[string]*string) *GetGatewayResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayResponse) SetStatusCode(v int32) *GetGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayResponse) SetBody(v *GetGatewayResponseBody) *GetGatewayResponse {
	s.Body = v
	return s
}

func (s *GetGatewayResponse) Validate() error {
	return dara.Validate(s)
}
