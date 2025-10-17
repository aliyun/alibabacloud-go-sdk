// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserSayResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserSayResponseBody) *CreateUserSayResponse
	GetBody() *CreateUserSayResponseBody
}

type CreateUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSayResponse) GoString() string {
	return s.String()
}

func (s *CreateUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserSayResponse) GetBody() *CreateUserSayResponseBody {
	return s.Body
}

func (s *CreateUserSayResponse) SetHeaders(v map[string]*string) *CreateUserSayResponse {
	s.Headers = v
	return s
}

func (s *CreateUserSayResponse) SetStatusCode(v int32) *CreateUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserSayResponse) SetBody(v *CreateUserSayResponseBody) *CreateUserSayResponse {
	s.Body = v
	return s
}

func (s *CreateUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
