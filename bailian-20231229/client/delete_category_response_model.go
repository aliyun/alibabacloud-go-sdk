// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse
	GetBody() *DeleteCategoryResponseBody
}

type DeleteCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCategoryResponse) GetBody() *DeleteCategoryResponseBody {
	return s.Body
}

func (s *DeleteCategoryResponse) SetHeaders(v map[string]*string) *DeleteCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategoryResponse) SetStatusCode(v int32) *DeleteCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteCategoryResponse) Validate() error {
	return dara.Validate(s)
}
