// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayAuthResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayAuthResponseBody) *DeleteGatewayAuthResponse
	GetBody() *DeleteGatewayAuthResponseBody
}

type DeleteGatewayAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayAuthResponse) GetBody() *DeleteGatewayAuthResponseBody {
	return s.Body
}

func (s *DeleteGatewayAuthResponse) SetHeaders(v map[string]*string) *DeleteGatewayAuthResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayAuthResponse) SetStatusCode(v int32) *DeleteGatewayAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayAuthResponse) SetBody(v *DeleteGatewayAuthResponseBody) *DeleteGatewayAuthResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
