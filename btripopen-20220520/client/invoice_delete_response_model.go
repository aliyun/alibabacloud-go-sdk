// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceDeleteResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceDeleteResponseBody) *InvoiceDeleteResponse
	GetBody() *InvoiceDeleteResponseBody
}

type InvoiceDeleteResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceDeleteResponse) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceDeleteResponse) GetBody() *InvoiceDeleteResponseBody {
	return s.Body
}

func (s *InvoiceDeleteResponse) SetHeaders(v map[string]*string) *InvoiceDeleteResponse {
	s.Headers = v
	return s
}

func (s *InvoiceDeleteResponse) SetStatusCode(v int32) *InvoiceDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceDeleteResponse) SetBody(v *InvoiceDeleteResponseBody) *InvoiceDeleteResponse {
	s.Body = v
	return s
}

func (s *InvoiceDeleteResponse) Validate() error {
	return dara.Validate(s)
}
