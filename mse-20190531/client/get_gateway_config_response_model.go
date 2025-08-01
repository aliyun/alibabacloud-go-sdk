// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayConfigResponseBody) *GetGatewayConfigResponse
	GetBody() *GetGatewayConfigResponseBody
}

type GetGatewayConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayConfigResponse) GetBody() *GetGatewayConfigResponseBody {
	return s.Body
}

func (s *GetGatewayConfigResponse) SetHeaders(v map[string]*string) *GetGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayConfigResponse) SetStatusCode(v int32) *GetGatewayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayConfigResponse) SetBody(v *GetGatewayConfigResponseBody) *GetGatewayConfigResponse {
	s.Body = v
	return s
}

func (s *GetGatewayConfigResponse) Validate() error {
	return dara.Validate(s)
}
