// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecretResponseBody) *CreateSecretResponse
	GetBody() *CreateSecretResponseBody
}

type CreateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecretResponse) GetBody() *CreateSecretResponseBody {
	return s.Body
}

func (s *CreateSecretResponse) SetHeaders(v map[string]*string) *CreateSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretResponse) SetStatusCode(v int32) *CreateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretResponse) SetBody(v *CreateSecretResponseBody) *CreateSecretResponse {
	s.Body = v
	return s
}

func (s *CreateSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
