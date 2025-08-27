// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceModifyResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceModifyResponseBody) *InvoiceModifyResponse
	GetBody() *InvoiceModifyResponseBody
}

type InvoiceModifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceModifyResponse) GoString() string {
	return s.String()
}

func (s *InvoiceModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceModifyResponse) GetBody() *InvoiceModifyResponseBody {
	return s.Body
}

func (s *InvoiceModifyResponse) SetHeaders(v map[string]*string) *InvoiceModifyResponse {
	s.Headers = v
	return s
}

func (s *InvoiceModifyResponse) SetStatusCode(v int32) *InvoiceModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceModifyResponse) SetBody(v *InvoiceModifyResponseBody) *InvoiceModifyResponse {
	s.Body = v
	return s
}

func (s *InvoiceModifyResponse) Validate() error {
	return dara.Validate(s)
}
