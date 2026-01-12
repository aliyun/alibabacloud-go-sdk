// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoginResponse
	GetStatusCode() *int32
	SetBody(v *LoginResponseBody) *LoginResponse
	GetBody() *LoginResponseBody
}

type LoginResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginResponse) String() string {
	return dara.Prettify(s)
}

func (s LoginResponse) GoString() string {
	return s.String()
}

func (s *LoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoginResponse) GetBody() *LoginResponseBody {
	return s.Body
}

func (s *LoginResponse) SetHeaders(v map[string]*string) *LoginResponse {
	s.Headers = v
	return s
}

func (s *LoginResponse) SetStatusCode(v int32) *LoginResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginResponse) SetBody(v *LoginResponseBody) *LoginResponse {
	s.Body = v
	return s
}

func (s *LoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
