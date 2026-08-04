// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomerCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomerCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomerCategoryResponseBody) *UpdateCustomerCategoryResponse
	GetBody() *UpdateCustomerCategoryResponseBody
}

type UpdateCustomerCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomerCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomerCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomerCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomerCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomerCategoryResponse) GetBody() *UpdateCustomerCategoryResponseBody {
	return s.Body
}

func (s *UpdateCustomerCategoryResponse) SetHeaders(v map[string]*string) *UpdateCustomerCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomerCategoryResponse) SetStatusCode(v int32) *UpdateCustomerCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomerCategoryResponse) SetBody(v *UpdateCustomerCategoryResponseBody) *UpdateCustomerCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomerCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
