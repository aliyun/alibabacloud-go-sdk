// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretValueResponseBody) *GetSecretValueResponse
	GetBody() *GetSecretValueResponseBody
}

type GetSecretValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponse) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretValueResponse) GetBody() *GetSecretValueResponseBody {
	return s.Body
}

func (s *GetSecretValueResponse) SetHeaders(v map[string]*string) *GetSecretValueResponse {
	s.Headers = v
	return s
}

func (s *GetSecretValueResponse) SetStatusCode(v int32) *GetSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretValueResponse) SetBody(v *GetSecretValueResponseBody) *GetSecretValueResponse {
	s.Body = v
	return s
}

func (s *GetSecretValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
