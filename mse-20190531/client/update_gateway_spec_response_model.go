// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewaySpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewaySpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewaySpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewaySpecResponseBody) *UpdateGatewaySpecResponse
	GetBody() *UpdateGatewaySpecResponseBody
}

type UpdateGatewaySpecResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewaySpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewaySpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewaySpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewaySpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewaySpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewaySpecResponse) GetBody() *UpdateGatewaySpecResponseBody {
	return s.Body
}

func (s *UpdateGatewaySpecResponse) SetHeaders(v map[string]*string) *UpdateGatewaySpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewaySpecResponse) SetStatusCode(v int32) *UpdateGatewaySpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewaySpecResponse) SetBody(v *UpdateGatewaySpecResponseBody) *UpdateGatewaySpecResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewaySpecResponse) Validate() error {
	return dara.Validate(s)
}
