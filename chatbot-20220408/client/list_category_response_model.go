// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListCategoryResponseBody) *ListCategoryResponse
	GetBody() *ListCategoryResponseBody
}

type ListCategoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCategoryResponse) GetBody() *ListCategoryResponseBody {
	return s.Body
}

func (s *ListCategoryResponse) SetHeaders(v map[string]*string) *ListCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListCategoryResponse) SetStatusCode(v int32) *ListCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoryResponse) SetBody(v *ListCategoryResponseBody) *ListCategoryResponse {
	s.Body = v
	return s
}

func (s *ListCategoryResponse) Validate() error {
	return dara.Validate(s)
}
