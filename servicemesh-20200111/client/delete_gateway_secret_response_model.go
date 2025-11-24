// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewaySecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewaySecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewaySecretResponseBody) *DeleteGatewaySecretResponse
	GetBody() *DeleteGatewaySecretResponseBody
}

type DeleteGatewaySecretResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewaySecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewaySecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewaySecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewaySecretResponse) GetBody() *DeleteGatewaySecretResponseBody {
	return s.Body
}

func (s *DeleteGatewaySecretResponse) SetHeaders(v map[string]*string) *DeleteGatewaySecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySecretResponse) SetStatusCode(v int32) *DeleteGatewaySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySecretResponse) SetBody(v *DeleteGatewaySecretResponseBody) *DeleteGatewaySecretResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewaySecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
