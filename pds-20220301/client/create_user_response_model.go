// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserResponseBody) *CreateUserResponse
	GetBody() *CreateUserResponseBody
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserResponse) GetBody() *CreateUserResponseBody {
	return s.Body
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

func (s *CreateUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
