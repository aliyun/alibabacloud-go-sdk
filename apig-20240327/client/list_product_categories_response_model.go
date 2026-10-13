// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListProductCategoriesResponseBody) *ListProductCategoriesResponse
	GetBody() *ListProductCategoriesResponseBody
}

type ListProductCategoriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListProductCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductCategoriesResponse) GetBody() *ListProductCategoriesResponseBody {
	return s.Body
}

func (s *ListProductCategoriesResponse) SetHeaders(v map[string]*string) *ListProductCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListProductCategoriesResponse) SetStatusCode(v int32) *ListProductCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductCategoriesResponse) SetBody(v *ListProductCategoriesResponseBody) *ListProductCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListProductCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
