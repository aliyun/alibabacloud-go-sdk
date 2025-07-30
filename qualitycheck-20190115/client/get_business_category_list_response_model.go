// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessCategoryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBusinessCategoryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBusinessCategoryListResponse
	GetStatusCode() *int32
	SetBody(v *GetBusinessCategoryListResponseBody) *GetBusinessCategoryListResponse
	GetBody() *GetBusinessCategoryListResponseBody
}

type GetBusinessCategoryListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBusinessCategoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBusinessCategoryListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessCategoryListResponse) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBusinessCategoryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBusinessCategoryListResponse) GetBody() *GetBusinessCategoryListResponseBody {
	return s.Body
}

func (s *GetBusinessCategoryListResponse) SetHeaders(v map[string]*string) *GetBusinessCategoryListResponse {
	s.Headers = v
	return s
}

func (s *GetBusinessCategoryListResponse) SetStatusCode(v int32) *GetBusinessCategoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusinessCategoryListResponse) SetBody(v *GetBusinessCategoryListResponseBody) *GetBusinessCategoryListResponse {
	s.Body = v
	return s
}

func (s *GetBusinessCategoryListResponse) Validate() error {
	return dara.Validate(s)
}
