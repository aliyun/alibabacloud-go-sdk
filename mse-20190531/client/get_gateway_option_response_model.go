// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayOptionResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayOptionResponseBody) *GetGatewayOptionResponse
	GetBody() *GetGatewayOptionResponseBody
}

type GetGatewayOptionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayOptionResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayOptionResponse) GetBody() *GetGatewayOptionResponseBody {
	return s.Body
}

func (s *GetGatewayOptionResponse) SetHeaders(v map[string]*string) *GetGatewayOptionResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayOptionResponse) SetStatusCode(v int32) *GetGatewayOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayOptionResponse) SetBody(v *GetGatewayOptionResponseBody) *GetGatewayOptionResponse {
	s.Body = v
	return s
}

func (s *GetGatewayOptionResponse) Validate() error {
	return dara.Validate(s)
}
