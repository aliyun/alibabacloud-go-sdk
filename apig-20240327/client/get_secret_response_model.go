// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretResponseBody) *GetSecretResponse
	GetBody() *GetSecretResponseBody
}

type GetSecretResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretResponse) GoString() string {
	return s.String()
}

func (s *GetSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretResponse) GetBody() *GetSecretResponseBody {
	return s.Body
}

func (s *GetSecretResponse) SetHeaders(v map[string]*string) *GetSecretResponse {
	s.Headers = v
	return s
}

func (s *GetSecretResponse) SetStatusCode(v int32) *GetSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretResponse) SetBody(v *GetSecretResponseBody) *GetSecretResponse {
	s.Body = v
	return s
}

func (s *GetSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
