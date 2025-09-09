// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupplierInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSupplierInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSupplierInformationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSupplierInformationResponseBody) *UpdateSupplierInformationResponse
	GetBody() *UpdateSupplierInformationResponseBody
}

type UpdateSupplierInformationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSupplierInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSupplierInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupplierInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSupplierInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSupplierInformationResponse) GetBody() *UpdateSupplierInformationResponseBody {
	return s.Body
}

func (s *UpdateSupplierInformationResponse) SetHeaders(v map[string]*string) *UpdateSupplierInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateSupplierInformationResponse) SetStatusCode(v int32) *UpdateSupplierInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSupplierInformationResponse) SetBody(v *UpdateSupplierInformationResponseBody) *UpdateSupplierInformationResponse {
	s.Body = v
	return s
}

func (s *UpdateSupplierInformationResponse) Validate() error {
	return dara.Validate(s)
}
