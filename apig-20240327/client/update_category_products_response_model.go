// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCategoryProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCategoryProductsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCategoryProductsResponseBody) *UpdateCategoryProductsResponse
	GetBody() *UpdateCategoryProductsResponseBody
}

type UpdateCategoryProductsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategoryProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCategoryProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryProductsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCategoryProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCategoryProductsResponse) GetBody() *UpdateCategoryProductsResponseBody {
	return s.Body
}

func (s *UpdateCategoryProductsResponse) SetHeaders(v map[string]*string) *UpdateCategoryProductsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryProductsResponse) SetStatusCode(v int32) *UpdateCategoryProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategoryProductsResponse) SetBody(v *UpdateCategoryProductsResponseBody) *UpdateCategoryProductsResponse {
	s.Body = v
	return s
}

func (s *UpdateCategoryProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
