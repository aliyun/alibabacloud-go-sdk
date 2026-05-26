// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientSecretResponseBody) *CreateClientSecretResponse
	GetBody() *CreateClientSecretResponseBody
}

type CreateClientSecretResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientSecretResponse) GetBody() *CreateClientSecretResponseBody {
	return s.Body
}

func (s *CreateClientSecretResponse) SetHeaders(v map[string]*string) *CreateClientSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateClientSecretResponse) SetStatusCode(v int32) *CreateClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientSecretResponse) SetBody(v *CreateClientSecretResponseBody) *CreateClientSecretResponse {
	s.Body = v
	return s
}

func (s *CreateClientSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
