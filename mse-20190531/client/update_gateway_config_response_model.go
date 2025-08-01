// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayConfigResponseBody) *UpdateGatewayConfigResponse
	GetBody() *UpdateGatewayConfigResponseBody
}

type UpdateGatewayConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayConfigResponse) GetBody() *UpdateGatewayConfigResponseBody {
	return s.Body
}

func (s *UpdateGatewayConfigResponse) SetHeaders(v map[string]*string) *UpdateGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayConfigResponse) SetStatusCode(v int32) *UpdateGatewayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayConfigResponse) SetBody(v *UpdateGatewayConfigResponseBody) *UpdateGatewayConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayConfigResponse) Validate() error {
	return dara.Validate(s)
}
