// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportedProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportedProductsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportedProductsResponseBody) *ListSupportedProductsResponse
	GetBody() *ListSupportedProductsResponseBody
}

type ListSupportedProductsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportedProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportedProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedProductsResponse) GoString() string {
	return s.String()
}

func (s *ListSupportedProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportedProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportedProductsResponse) GetBody() *ListSupportedProductsResponseBody {
	return s.Body
}

func (s *ListSupportedProductsResponse) SetHeaders(v map[string]*string) *ListSupportedProductsResponse {
	s.Headers = v
	return s
}

func (s *ListSupportedProductsResponse) SetStatusCode(v int32) *ListSupportedProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportedProductsResponse) SetBody(v *ListSupportedProductsResponseBody) *ListSupportedProductsResponse {
	s.Body = v
	return s
}

func (s *ListSupportedProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
