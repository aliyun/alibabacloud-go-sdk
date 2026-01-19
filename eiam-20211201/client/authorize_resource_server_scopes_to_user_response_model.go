// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceServerScopesToUserResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceServerScopesToUserResponseBody) *AuthorizeResourceServerScopesToUserResponse
	GetBody() *AuthorizeResourceServerScopesToUserResponseBody
}

type AuthorizeResourceServerScopesToUserResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceServerScopesToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceServerScopesToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToUserResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceServerScopesToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceServerScopesToUserResponse) GetBody() *AuthorizeResourceServerScopesToUserResponseBody {
	return s.Body
}

func (s *AuthorizeResourceServerScopesToUserResponse) SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToUserResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceServerScopesToUserResponse) SetStatusCode(v int32) *AuthorizeResourceServerScopesToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceServerScopesToUserResponse) SetBody(v *AuthorizeResourceServerScopesToUserResponseBody) *AuthorizeResourceServerScopesToUserResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceServerScopesToUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
