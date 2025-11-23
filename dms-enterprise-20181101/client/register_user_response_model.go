// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterUserResponse
	GetStatusCode() *int32
	SetBody(v *RegisterUserResponseBody) *RegisterUserResponse
	GetBody() *RegisterUserResponseBody
}

type RegisterUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterUserResponse) GoString() string {
	return s.String()
}

func (s *RegisterUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterUserResponse) GetBody() *RegisterUserResponseBody {
	return s.Body
}

func (s *RegisterUserResponse) SetHeaders(v map[string]*string) *RegisterUserResponse {
	s.Headers = v
	return s
}

func (s *RegisterUserResponse) SetStatusCode(v int32) *RegisterUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterUserResponse) SetBody(v *RegisterUserResponseBody) *RegisterUserResponse {
	s.Body = v
	return s
}

func (s *RegisterUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
