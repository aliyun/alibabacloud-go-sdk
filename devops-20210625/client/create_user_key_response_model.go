// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserKeyResponseBody) *CreateUserKeyResponse
	GetBody() *CreateUserKeyResponseBody
}

type CreateUserKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateUserKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserKeyResponse) GetBody() *CreateUserKeyResponseBody {
	return s.Body
}

func (s *CreateUserKeyResponse) SetHeaders(v map[string]*string) *CreateUserKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateUserKeyResponse) SetStatusCode(v int32) *CreateUserKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserKeyResponse) SetBody(v *CreateUserKeyResponseBody) *CreateUserKeyResponse {
	s.Body = v
	return s
}

func (s *CreateUserKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
