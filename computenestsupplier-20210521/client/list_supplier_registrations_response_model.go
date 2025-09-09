// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupplierRegistrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupplierRegistrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupplierRegistrationsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupplierRegistrationsResponseBody) *ListSupplierRegistrationsResponse
	GetBody() *ListSupplierRegistrationsResponseBody
}

type ListSupplierRegistrationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupplierRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupplierRegistrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupplierRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupplierRegistrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupplierRegistrationsResponse) GetBody() *ListSupplierRegistrationsResponseBody {
	return s.Body
}

func (s *ListSupplierRegistrationsResponse) SetHeaders(v map[string]*string) *ListSupplierRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListSupplierRegistrationsResponse) SetStatusCode(v int32) *ListSupplierRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupplierRegistrationsResponse) SetBody(v *ListSupplierRegistrationsResponseBody) *ListSupplierRegistrationsResponse {
	s.Body = v
	return s
}

func (s *ListSupplierRegistrationsResponse) Validate() error {
	return dara.Validate(s)
}
