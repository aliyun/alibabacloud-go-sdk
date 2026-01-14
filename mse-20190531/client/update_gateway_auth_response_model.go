// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayAuthResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayAuthResponseBody) *UpdateGatewayAuthResponse
	GetBody() *UpdateGatewayAuthResponseBody
}

type UpdateGatewayAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayAuthResponse) GetBody() *UpdateGatewayAuthResponseBody {
	return s.Body
}

func (s *UpdateGatewayAuthResponse) SetHeaders(v map[string]*string) *UpdateGatewayAuthResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayAuthResponse) SetStatusCode(v int32) *UpdateGatewayAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayAuthResponse) SetBody(v *UpdateGatewayAuthResponseBody) *UpdateGatewayAuthResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
