// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplierInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupplierInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupplierInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetSupplierInformationResponseBody) *GetSupplierInformationResponse
	GetBody() *GetSupplierInformationResponseBody
}

type GetSupplierInformationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplierInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplierInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupplierInformationResponse) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupplierInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupplierInformationResponse) GetBody() *GetSupplierInformationResponseBody {
	return s.Body
}

func (s *GetSupplierInformationResponse) SetHeaders(v map[string]*string) *GetSupplierInformationResponse {
	s.Headers = v
	return s
}

func (s *GetSupplierInformationResponse) SetStatusCode(v int32) *GetSupplierInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplierInformationResponse) SetBody(v *GetSupplierInformationResponseBody) *GetSupplierInformationResponse {
	s.Body = v
	return s
}

func (s *GetSupplierInformationResponse) Validate() error {
	return dara.Validate(s)
}
