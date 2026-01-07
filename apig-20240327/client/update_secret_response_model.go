// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecretResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecretResponseBody) *UpdateSecretResponse
	GetBody() *UpdateSecretResponseBody
}

type UpdateSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecretResponse) GetBody() *UpdateSecretResponseBody {
	return s.Body
}

func (s *UpdateSecretResponse) SetHeaders(v map[string]*string) *UpdateSecretResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretResponse) SetStatusCode(v int32) *UpdateSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretResponse) SetBody(v *UpdateSecretResponseBody) *UpdateSecretResponse {
	s.Body = v
	return s
}

func (s *UpdateSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
