// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceServerScopesToClientResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceServerScopesToClientResponseBody) *AuthorizeResourceServerScopesToClientResponse
	GetBody() *AuthorizeResourceServerScopesToClientResponseBody
}

type AuthorizeResourceServerScopesToClientResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceServerScopesToClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceServerScopesToClientResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToClientResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceServerScopesToClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceServerScopesToClientResponse) GetBody() *AuthorizeResourceServerScopesToClientResponseBody {
	return s.Body
}

func (s *AuthorizeResourceServerScopesToClientResponse) SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToClientResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceServerScopesToClientResponse) SetStatusCode(v int32) *AuthorizeResourceServerScopesToClientResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceServerScopesToClientResponse) SetBody(v *AuthorizeResourceServerScopesToClientResponseBody) *AuthorizeResourceServerScopesToClientResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceServerScopesToClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
