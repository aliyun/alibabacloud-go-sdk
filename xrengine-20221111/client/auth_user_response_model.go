// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthUserResponse
	GetStatusCode() *int32
	SetBody(v *AuthUserResponseBody) *AuthUserResponse
	GetBody() *AuthUserResponseBody
}

type AuthUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthUserResponse) GoString() string {
	return s.String()
}

func (s *AuthUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthUserResponse) GetBody() *AuthUserResponseBody {
	return s.Body
}

func (s *AuthUserResponse) SetHeaders(v map[string]*string) *AuthUserResponse {
	s.Headers = v
	return s
}

func (s *AuthUserResponse) SetStatusCode(v int32) *AuthUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthUserResponse) SetBody(v *AuthUserResponseBody) *AuthUserResponse {
	s.Body = v
	return s
}

func (s *AuthUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
