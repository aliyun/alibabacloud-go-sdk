// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeUsersForAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeUsersForAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeUsersForAppResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeUsersForAppResponseBody) *AuthorizeUsersForAppResponse
	GetBody() *AuthorizeUsersForAppResponseBody
}

type AuthorizeUsersForAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeUsersForAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeUsersForAppResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeUsersForAppResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeUsersForAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeUsersForAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeUsersForAppResponse) GetBody() *AuthorizeUsersForAppResponseBody {
	return s.Body
}

func (s *AuthorizeUsersForAppResponse) SetHeaders(v map[string]*string) *AuthorizeUsersForAppResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeUsersForAppResponse) SetStatusCode(v int32) *AuthorizeUsersForAppResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeUsersForAppResponse) SetBody(v *AuthorizeUsersForAppResponseBody) *AuthorizeUsersForAppResponse {
	s.Body = v
	return s
}

func (s *AuthorizeUsersForAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
