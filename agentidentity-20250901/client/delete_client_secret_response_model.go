// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientSecretResponseBody) *DeleteClientSecretResponse
	GetBody() *DeleteClientSecretResponseBody
}

type DeleteClientSecretResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientSecretResponse) GetBody() *DeleteClientSecretResponseBody {
	return s.Body
}

func (s *DeleteClientSecretResponse) SetHeaders(v map[string]*string) *DeleteClientSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientSecretResponse) SetStatusCode(v int32) *DeleteClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientSecretResponse) SetBody(v *DeleteClientSecretResponseBody) *DeleteClientSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteClientSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
