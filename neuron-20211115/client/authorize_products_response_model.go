// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeProductsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeProductsResponseBody) *AuthorizeProductsResponse
	GetBody() *AuthorizeProductsResponseBody
}

type AuthorizeProductsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeProductsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeProductsResponse) GetBody() *AuthorizeProductsResponseBody {
	return s.Body
}

func (s *AuthorizeProductsResponse) SetHeaders(v map[string]*string) *AuthorizeProductsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeProductsResponse) SetStatusCode(v int32) *AuthorizeProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeProductsResponse) SetBody(v *AuthorizeProductsResponseBody) *AuthorizeProductsResponse {
	s.Body = v
	return s
}

func (s *AuthorizeProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
