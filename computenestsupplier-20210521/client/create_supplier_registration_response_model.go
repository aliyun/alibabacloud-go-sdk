// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupplierRegistrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSupplierRegistrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSupplierRegistrationResponse
	GetStatusCode() *int32
	SetBody(v *CreateSupplierRegistrationResponseBody) *CreateSupplierRegistrationResponse
	GetBody() *CreateSupplierRegistrationResponseBody
}

type CreateSupplierRegistrationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSupplierRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSupplierRegistrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSupplierRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSupplierRegistrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSupplierRegistrationResponse) GetBody() *CreateSupplierRegistrationResponseBody {
	return s.Body
}

func (s *CreateSupplierRegistrationResponse) SetHeaders(v map[string]*string) *CreateSupplierRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CreateSupplierRegistrationResponse) SetStatusCode(v int32) *CreateSupplierRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSupplierRegistrationResponse) SetBody(v *CreateSupplierRegistrationResponseBody) *CreateSupplierRegistrationResponse {
	s.Body = v
	return s
}

func (s *CreateSupplierRegistrationResponse) Validate() error {
	return dara.Validate(s)
}
