// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetProductCategoryResponseBody) *GetProductCategoryResponse
	GetBody() *GetProductCategoryResponseBody
}

type GetProductCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetProductCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductCategoryResponse) GetBody() *GetProductCategoryResponseBody {
	return s.Body
}

func (s *GetProductCategoryResponse) SetHeaders(v map[string]*string) *GetProductCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetProductCategoryResponse) SetStatusCode(v int32) *GetProductCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductCategoryResponse) SetBody(v *GetProductCategoryResponseBody) *GetProductCategoryResponse {
	s.Body = v
	return s
}

func (s *GetProductCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
