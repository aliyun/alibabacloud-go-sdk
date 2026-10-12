// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateProductCategoryResponseBody) *CreateProductCategoryResponse
	GetBody() *CreateProductCategoryResponseBody
}

type CreateProductCategoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateProductCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductCategoryResponse) GetBody() *CreateProductCategoryResponseBody {
	return s.Body
}

func (s *CreateProductCategoryResponse) SetHeaders(v map[string]*string) *CreateProductCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateProductCategoryResponse) SetStatusCode(v int32) *CreateProductCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductCategoryResponse) SetBody(v *CreateProductCategoryResponseBody) *CreateProductCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateProductCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
