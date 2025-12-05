// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserPublicKeyResponseBody) *CreateUserPublicKeyResponse
	GetBody() *CreateUserPublicKeyResponseBody
}

type CreateUserPublicKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserPublicKeyResponse) GetBody() *CreateUserPublicKeyResponseBody {
	return s.Body
}

func (s *CreateUserPublicKeyResponse) SetHeaders(v map[string]*string) *CreateUserPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateUserPublicKeyResponse) SetStatusCode(v int32) *CreateUserPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserPublicKeyResponse) SetBody(v *CreateUserPublicKeyResponseBody) *CreateUserPublicKeyResponse {
	s.Body = v
	return s
}

func (s *CreateUserPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
