// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthenticateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthenticateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthenticateResponse
	GetStatusCode() *int32
	SetBody(v *AuthenticateResponseBody) *AuthenticateResponse
	GetBody() *AuthenticateResponseBody
}

type AuthenticateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthenticateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthenticateResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateResponse) GoString() string {
	return s.String()
}

func (s *AuthenticateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthenticateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthenticateResponse) GetBody() *AuthenticateResponseBody {
	return s.Body
}

func (s *AuthenticateResponse) SetHeaders(v map[string]*string) *AuthenticateResponse {
	s.Headers = v
	return s
}

func (s *AuthenticateResponse) SetStatusCode(v int32) *AuthenticateResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthenticateResponse) SetBody(v *AuthenticateResponseBody) *AuthenticateResponse {
	s.Body = v
	return s
}

func (s *AuthenticateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
