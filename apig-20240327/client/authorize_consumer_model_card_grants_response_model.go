// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeConsumerModelCardGrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeConsumerModelCardGrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeConsumerModelCardGrantsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeConsumerModelCardGrantsResponseBody) *AuthorizeConsumerModelCardGrantsResponse
	GetBody() *AuthorizeConsumerModelCardGrantsResponseBody
}

type AuthorizeConsumerModelCardGrantsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeConsumerModelCardGrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeConsumerModelCardGrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeConsumerModelCardGrantsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeConsumerModelCardGrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeConsumerModelCardGrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeConsumerModelCardGrantsResponse) GetBody() *AuthorizeConsumerModelCardGrantsResponseBody {
	return s.Body
}

func (s *AuthorizeConsumerModelCardGrantsResponse) SetHeaders(v map[string]*string) *AuthorizeConsumerModelCardGrantsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsResponse) SetStatusCode(v int32) *AuthorizeConsumerModelCardGrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsResponse) SetBody(v *AuthorizeConsumerModelCardGrantsResponseBody) *AuthorizeConsumerModelCardGrantsResponse {
	s.Body = v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
