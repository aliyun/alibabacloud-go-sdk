// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeApplicationResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeApplicationResponseBody) *AuthorizeApplicationResponse
	GetBody() *AuthorizeApplicationResponseBody
}

type AuthorizeApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeApplicationResponse) GetBody() *AuthorizeApplicationResponseBody {
	return s.Body
}

func (s *AuthorizeApplicationResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationResponse) SetStatusCode(v int32) *AuthorizeApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationResponse) SetBody(v *AuthorizeApplicationResponseBody) *AuthorizeApplicationResponse {
	s.Body = v
	return s
}

func (s *AuthorizeApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
