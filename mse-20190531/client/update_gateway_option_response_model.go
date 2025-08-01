// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayOptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayOptionResponseBody) *UpdateGatewayOptionResponse
	GetBody() *UpdateGatewayOptionResponseBody
}

type UpdateGatewayOptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayOptionResponse) GetBody() *UpdateGatewayOptionResponseBody {
	return s.Body
}

func (s *UpdateGatewayOptionResponse) SetHeaders(v map[string]*string) *UpdateGatewayOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayOptionResponse) SetStatusCode(v int32) *UpdateGatewayOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayOptionResponse) SetBody(v *UpdateGatewayOptionResponseBody) *UpdateGatewayOptionResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayOptionResponse) Validate() error {
	return dara.Validate(s)
}
