// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RotateSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RotateSecretResponse
	GetStatusCode() *int32
	SetBody(v *RotateSecretResponseBody) *RotateSecretResponse
	GetBody() *RotateSecretResponseBody
}

type RotateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RotateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RotateSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s RotateSecretResponse) GoString() string {
	return s.String()
}

func (s *RotateSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RotateSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RotateSecretResponse) GetBody() *RotateSecretResponseBody {
	return s.Body
}

func (s *RotateSecretResponse) SetHeaders(v map[string]*string) *RotateSecretResponse {
	s.Headers = v
	return s
}

func (s *RotateSecretResponse) SetStatusCode(v int32) *RotateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *RotateSecretResponse) SetBody(v *RotateSecretResponseBody) *RotateSecretResponse {
	s.Body = v
	return s
}

func (s *RotateSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
