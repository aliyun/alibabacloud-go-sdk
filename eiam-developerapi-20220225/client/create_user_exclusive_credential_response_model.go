// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserExclusiveCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserExclusiveCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserExclusiveCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserExclusiveCredentialResponseBody) *CreateUserExclusiveCredentialResponse
	GetBody() *CreateUserExclusiveCredentialResponseBody
}

type CreateUserExclusiveCredentialResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserExclusiveCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserExclusiveCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserExclusiveCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserExclusiveCredentialResponse) GetBody() *CreateUserExclusiveCredentialResponseBody {
	return s.Body
}

func (s *CreateUserExclusiveCredentialResponse) SetHeaders(v map[string]*string) *CreateUserExclusiveCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateUserExclusiveCredentialResponse) SetStatusCode(v int32) *CreateUserExclusiveCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserExclusiveCredentialResponse) SetBody(v *CreateUserExclusiveCredentialResponseBody) *CreateUserExclusiveCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateUserExclusiveCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
