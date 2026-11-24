// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecretValueResponseBody) *DeleteSecretValueResponse
	GetBody() *DeleteSecretValueResponseBody
}

type DeleteSecretValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecretValueResponse) GetBody() *DeleteSecretValueResponseBody {
	return s.Body
}

func (s *DeleteSecretValueResponse) SetHeaders(v map[string]*string) *DeleteSecretValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretValueResponse) SetStatusCode(v int32) *DeleteSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretValueResponse) SetBody(v *DeleteSecretValueResponseBody) *DeleteSecretValueResponse {
	s.Body = v
	return s
}

func (s *DeleteSecretValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
