// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretValueResponseBody) *ListSecretValueResponse
	GetBody() *ListSecretValueResponseBody
}

type ListSecretValueResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretValueResponse) GoString() string {
	return s.String()
}

func (s *ListSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretValueResponse) GetBody() *ListSecretValueResponseBody {
	return s.Body
}

func (s *ListSecretValueResponse) SetHeaders(v map[string]*string) *ListSecretValueResponse {
	s.Headers = v
	return s
}

func (s *ListSecretValueResponse) SetStatusCode(v int32) *ListSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretValueResponse) SetBody(v *ListSecretValueResponseBody) *ListSecretValueResponse {
	s.Body = v
	return s
}

func (s *ListSecretValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
