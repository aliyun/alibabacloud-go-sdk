// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesInCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTablesInCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTablesInCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListTablesInCategoryResponseBody) *ListTablesInCategoryResponse
	GetBody() *ListTablesInCategoryResponseBody
}

type ListTablesInCategoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTablesInCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTablesInCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTablesInCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListTablesInCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTablesInCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTablesInCategoryResponse) GetBody() *ListTablesInCategoryResponseBody {
	return s.Body
}

func (s *ListTablesInCategoryResponse) SetHeaders(v map[string]*string) *ListTablesInCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListTablesInCategoryResponse) SetStatusCode(v int32) *ListTablesInCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesInCategoryResponse) SetBody(v *ListTablesInCategoryResponseBody) *ListTablesInCategoryResponse {
	s.Body = v
	return s
}

func (s *ListTablesInCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
