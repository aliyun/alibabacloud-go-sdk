// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceAddResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceAddResponseBody) *InvoiceAddResponse
	GetBody() *InvoiceAddResponseBody
}

type InvoiceAddResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceAddResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceAddResponse) GoString() string {
	return s.String()
}

func (s *InvoiceAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceAddResponse) GetBody() *InvoiceAddResponseBody {
	return s.Body
}

func (s *InvoiceAddResponse) SetHeaders(v map[string]*string) *InvoiceAddResponse {
	s.Headers = v
	return s
}

func (s *InvoiceAddResponse) SetStatusCode(v int32) *InvoiceAddResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceAddResponse) SetBody(v *InvoiceAddResponseBody) *InvoiceAddResponse {
	s.Body = v
	return s
}

func (s *InvoiceAddResponse) Validate() error {
	return dara.Validate(s)
}
