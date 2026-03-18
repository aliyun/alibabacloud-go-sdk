// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthenticationTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthenticationTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthenticationTokensResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthenticationTokensResponseBody) *ListAuthenticationTokensResponse
	GetBody() *ListAuthenticationTokensResponseBody
}

type ListAuthenticationTokensResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthenticationTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthenticationTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthenticationTokensResponse) GoString() string {
	return s.String()
}

func (s *ListAuthenticationTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthenticationTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthenticationTokensResponse) GetBody() *ListAuthenticationTokensResponseBody {
	return s.Body
}

func (s *ListAuthenticationTokensResponse) SetHeaders(v map[string]*string) *ListAuthenticationTokensResponse {
	s.Headers = v
	return s
}

func (s *ListAuthenticationTokensResponse) SetStatusCode(v int32) *ListAuthenticationTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthenticationTokensResponse) SetBody(v *ListAuthenticationTokensResponseBody) *ListAuthenticationTokensResponse {
	s.Body = v
	return s
}

func (s *ListAuthenticationTokensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
