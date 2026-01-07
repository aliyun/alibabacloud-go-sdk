// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecretResponseBody) *DeleteSecretResponse
	GetBody() *DeleteSecretResponseBody
}

type DeleteSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecretResponse) GetBody() *DeleteSecretResponseBody {
	return s.Body
}

func (s *DeleteSecretResponse) SetHeaders(v map[string]*string) *DeleteSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretResponse) SetStatusCode(v int32) *DeleteSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretResponse) SetBody(v *DeleteSecretResponseBody) *DeleteSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
