// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchProductsResponse
	GetStatusCode() *int32
	SetBody(v *SearchProductsResponseBody) *SearchProductsResponse
	GetBody() *SearchProductsResponseBody
}

type SearchProductsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchProductsResponse) GoString() string {
	return s.String()
}

func (s *SearchProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchProductsResponse) GetBody() *SearchProductsResponseBody {
	return s.Body
}

func (s *SearchProductsResponse) SetHeaders(v map[string]*string) *SearchProductsResponse {
	s.Headers = v
	return s
}

func (s *SearchProductsResponse) SetStatusCode(v int32) *SearchProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchProductsResponse) SetBody(v *SearchProductsResponseBody) *SearchProductsResponse {
	s.Body = v
	return s
}

func (s *SearchProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
