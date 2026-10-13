// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProductCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProductCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProductCategoryResponseBody) *UpdateProductCategoryResponse
	GetBody() *UpdateProductCategoryResponseBody
}

type UpdateProductCategoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProductCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProductCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProductCategoryResponse) GetBody() *UpdateProductCategoryResponseBody {
	return s.Body
}

func (s *UpdateProductCategoryResponse) SetHeaders(v map[string]*string) *UpdateProductCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductCategoryResponse) SetStatusCode(v int32) *UpdateProductCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProductCategoryResponse) SetBody(v *UpdateProductCategoryResponseBody) *UpdateProductCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateProductCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
