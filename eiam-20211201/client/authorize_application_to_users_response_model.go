// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeApplicationToUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeApplicationToUsersResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeApplicationToUsersResponseBody) *AuthorizeApplicationToUsersResponse
	GetBody() *AuthorizeApplicationToUsersResponseBody
}

type AuthorizeApplicationToUsersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToUsersResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeApplicationToUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeApplicationToUsersResponse) GetBody() *AuthorizeApplicationToUsersResponseBody {
	return s.Body
}

func (s *AuthorizeApplicationToUsersResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToUsersResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToUsersResponse) SetStatusCode(v int32) *AuthorizeApplicationToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToUsersResponse) SetBody(v *AuthorizeApplicationToUsersResponseBody) *AuthorizeApplicationToUsersResponse {
	s.Body = v
	return s
}

func (s *AuthorizeApplicationToUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
