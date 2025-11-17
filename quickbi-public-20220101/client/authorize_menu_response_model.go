// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMenuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeMenuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeMenuResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeMenuResponseBody) *AuthorizeMenuResponse
	GetBody() *AuthorizeMenuResponseBody
}

type AuthorizeMenuResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeMenuResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMenuResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeMenuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeMenuResponse) GetBody() *AuthorizeMenuResponseBody {
	return s.Body
}

func (s *AuthorizeMenuResponse) SetHeaders(v map[string]*string) *AuthorizeMenuResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeMenuResponse) SetStatusCode(v int32) *AuthorizeMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeMenuResponse) SetBody(v *AuthorizeMenuResponseBody) *AuthorizeMenuResponse {
	s.Body = v
	return s
}

func (s *AuthorizeMenuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
