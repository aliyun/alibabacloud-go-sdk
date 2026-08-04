// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerCategoryResponseBody) *GetCustomerCategoryResponse
	GetBody() *GetCustomerCategoryResponseBody
}

type GetCustomerCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerCategoryResponse) GetBody() *GetCustomerCategoryResponseBody {
	return s.Body
}

func (s *GetCustomerCategoryResponse) SetHeaders(v map[string]*string) *GetCustomerCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerCategoryResponse) SetStatusCode(v int32) *GetCustomerCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerCategoryResponse) SetBody(v *GetCustomerCategoryResponseBody) *GetCustomerCategoryResponse {
	s.Body = v
	return s
}

func (s *GetCustomerCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
