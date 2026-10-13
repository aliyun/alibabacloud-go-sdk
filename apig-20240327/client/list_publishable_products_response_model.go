// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishableProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishableProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishableProductsResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishableProductsResponseBody) *ListPublishableProductsResponse
	GetBody() *ListPublishableProductsResponseBody
}

type ListPublishableProductsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishableProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishableProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishableProductsResponse) GoString() string {
	return s.String()
}

func (s *ListPublishableProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishableProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishableProductsResponse) GetBody() *ListPublishableProductsResponseBody {
	return s.Body
}

func (s *ListPublishableProductsResponse) SetHeaders(v map[string]*string) *ListPublishableProductsResponse {
	s.Headers = v
	return s
}

func (s *ListPublishableProductsResponse) SetStatusCode(v int32) *ListPublishableProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishableProductsResponse) SetBody(v *ListPublishableProductsResponseBody) *ListPublishableProductsResponse {
	s.Body = v
	return s
}

func (s *ListPublishableProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
