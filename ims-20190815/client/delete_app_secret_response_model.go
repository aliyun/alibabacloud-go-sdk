// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppSecretResponseBody) *DeleteAppSecretResponse
	GetBody() *DeleteAppSecretResponseBody
}

type DeleteAppSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppSecretResponse) GetBody() *DeleteAppSecretResponseBody {
	return s.Body
}

func (s *DeleteAppSecretResponse) SetHeaders(v map[string]*string) *DeleteAppSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppSecretResponse) SetStatusCode(v int32) *DeleteAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppSecretResponse) SetBody(v *DeleteAppSecretResponseBody) *DeleteAppSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteAppSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
