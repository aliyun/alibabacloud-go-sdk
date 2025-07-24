// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceCategoriesResponseBody) *ListServiceCategoriesResponse
	GetBody() *ListServiceCategoriesResponseBody
}

type ListServiceCategoriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceCategoriesResponse) GetBody() *ListServiceCategoriesResponseBody {
	return s.Body
}

func (s *ListServiceCategoriesResponse) SetHeaders(v map[string]*string) *ListServiceCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceCategoriesResponse) SetStatusCode(v int32) *ListServiceCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceCategoriesResponse) SetBody(v *ListServiceCategoriesResponseBody) *ListServiceCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListServiceCategoriesResponse) Validate() error {
	return dara.Validate(s)
}
