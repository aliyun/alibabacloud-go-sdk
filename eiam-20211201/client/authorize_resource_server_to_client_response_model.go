// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerToClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceServerToClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceServerToClientResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceServerToClientResponseBody) *AuthorizeResourceServerToClientResponse
	GetBody() *AuthorizeResourceServerToClientResponseBody
}

type AuthorizeResourceServerToClientResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceServerToClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceServerToClientResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerToClientResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerToClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceServerToClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceServerToClientResponse) GetBody() *AuthorizeResourceServerToClientResponseBody {
	return s.Body
}

func (s *AuthorizeResourceServerToClientResponse) SetHeaders(v map[string]*string) *AuthorizeResourceServerToClientResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceServerToClientResponse) SetStatusCode(v int32) *AuthorizeResourceServerToClientResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceServerToClientResponse) SetBody(v *AuthorizeResourceServerToClientResponseBody) *AuthorizeResourceServerToClientResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceServerToClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
