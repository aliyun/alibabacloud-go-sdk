// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListCategoriesResponseBody) *ListCategoriesResponse
	GetBody() *ListCategoriesResponseBody
}

type ListCategoriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCategoriesResponse) GetBody() *ListCategoriesResponseBody {
	return s.Body
}

func (s *ListCategoriesResponse) SetHeaders(v map[string]*string) *ListCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListCategoriesResponse) SetStatusCode(v int32) *ListCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoriesResponse) SetBody(v *ListCategoriesResponseBody) *ListCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListCategoriesResponse) Validate() error {
	return dara.Validate(s)
}
