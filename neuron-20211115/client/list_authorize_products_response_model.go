// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizeProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizeProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizeProductsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeProductListResult) *ListAuthorizeProductsResponse
	GetBody() *AuthorizeProductListResult
}

type ListAuthorizeProductsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeProductListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizeProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizeProductsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizeProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizeProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizeProductsResponse) GetBody() *AuthorizeProductListResult {
	return s.Body
}

func (s *ListAuthorizeProductsResponse) SetHeaders(v map[string]*string) *ListAuthorizeProductsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizeProductsResponse) SetStatusCode(v int32) *ListAuthorizeProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizeProductsResponse) SetBody(v *AuthorizeProductListResult) *ListAuthorizeProductsResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizeProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
