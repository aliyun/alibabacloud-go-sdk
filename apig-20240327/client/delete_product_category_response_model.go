// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProductCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProductCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProductCategoryResponseBody) *DeleteProductCategoryResponse
	GetBody() *DeleteProductCategoryResponseBody
}

type DeleteProductCategoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProductCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProductCategoryResponse) GetBody() *DeleteProductCategoryResponseBody {
	return s.Body
}

func (s *DeleteProductCategoryResponse) SetHeaders(v map[string]*string) *DeleteProductCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductCategoryResponse) SetStatusCode(v int32) *DeleteProductCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductCategoryResponse) SetBody(v *DeleteProductCategoryResponseBody) *DeleteProductCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteProductCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
