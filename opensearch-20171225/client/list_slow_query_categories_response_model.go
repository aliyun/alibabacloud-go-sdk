// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlowQueryCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSlowQueryCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSlowQueryCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListSlowQueryCategoriesResponseBody) *ListSlowQueryCategoriesResponse
	GetBody() *ListSlowQueryCategoriesResponseBody
}

type ListSlowQueryCategoriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlowQueryCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSlowQueryCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSlowQueryCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSlowQueryCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSlowQueryCategoriesResponse) GetBody() *ListSlowQueryCategoriesResponseBody {
	return s.Body
}

func (s *ListSlowQueryCategoriesResponse) SetHeaders(v map[string]*string) *ListSlowQueryCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListSlowQueryCategoriesResponse) SetStatusCode(v int32) *ListSlowQueryCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlowQueryCategoriesResponse) SetBody(v *ListSlowQueryCategoriesResponseBody) *ListSlowQueryCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListSlowQueryCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
