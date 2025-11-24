// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewaySecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewaySecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewaySecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewaySecretResponseBody) *CreateGatewaySecretResponse
	GetBody() *CreateGatewaySecretResponseBody
}

type CreateGatewaySecretResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewaySecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewaySecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewaySecretResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewaySecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewaySecretResponse) GetBody() *CreateGatewaySecretResponseBody {
	return s.Body
}

func (s *CreateGatewaySecretResponse) SetHeaders(v map[string]*string) *CreateGatewaySecretResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewaySecretResponse) SetStatusCode(v int32) *CreateGatewaySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewaySecretResponse) SetBody(v *CreateGatewaySecretResponseBody) *CreateGatewaySecretResponse {
	s.Body = v
	return s
}

func (s *CreateGatewaySecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
