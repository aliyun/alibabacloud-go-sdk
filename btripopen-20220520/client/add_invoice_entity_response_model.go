// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInvoiceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInvoiceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInvoiceEntityResponse
	GetStatusCode() *int32
	SetBody(v *AddInvoiceEntityResponseBody) *AddInvoiceEntityResponse
	GetBody() *AddInvoiceEntityResponseBody
}

type AddInvoiceEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInvoiceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInvoiceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityResponse) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInvoiceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInvoiceEntityResponse) GetBody() *AddInvoiceEntityResponseBody {
	return s.Body
}

func (s *AddInvoiceEntityResponse) SetHeaders(v map[string]*string) *AddInvoiceEntityResponse {
	s.Headers = v
	return s
}

func (s *AddInvoiceEntityResponse) SetStatusCode(v int32) *AddInvoiceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInvoiceEntityResponse) SetBody(v *AddInvoiceEntityResponseBody) *AddInvoiceEntityResponse {
	s.Body = v
	return s
}

func (s *AddInvoiceEntityResponse) Validate() error {
	return dara.Validate(s)
}
