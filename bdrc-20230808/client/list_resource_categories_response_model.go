// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceCategoriesResponseBody) *ListResourceCategoriesResponse
	GetBody() *ListResourceCategoriesResponseBody
}

type ListResourceCategoriesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceCategoriesResponse) GetBody() *ListResourceCategoriesResponseBody {
	return s.Body
}

func (s *ListResourceCategoriesResponse) SetHeaders(v map[string]*string) *ListResourceCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceCategoriesResponse) SetStatusCode(v int32) *ListResourceCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceCategoriesResponse) SetBody(v *ListResourceCategoriesResponseBody) *ListResourceCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListResourceCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
