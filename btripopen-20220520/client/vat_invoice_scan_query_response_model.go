// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVatInvoiceScanQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VatInvoiceScanQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VatInvoiceScanQueryResponse
	GetStatusCode() *int32
	SetBody(v *VatInvoiceScanQueryResponseBody) *VatInvoiceScanQueryResponse
	GetBody() *VatInvoiceScanQueryResponseBody
}

type VatInvoiceScanQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VatInvoiceScanQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VatInvoiceScanQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryResponse) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VatInvoiceScanQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VatInvoiceScanQueryResponse) GetBody() *VatInvoiceScanQueryResponseBody {
	return s.Body
}

func (s *VatInvoiceScanQueryResponse) SetHeaders(v map[string]*string) *VatInvoiceScanQueryResponse {
	s.Headers = v
	return s
}

func (s *VatInvoiceScanQueryResponse) SetStatusCode(v int32) *VatInvoiceScanQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VatInvoiceScanQueryResponse) SetBody(v *VatInvoiceScanQueryResponseBody) *VatInvoiceScanQueryResponse {
	s.Body = v
	return s
}

func (s *VatInvoiceScanQueryResponse) Validate() error {
	return dara.Validate(s)
}
