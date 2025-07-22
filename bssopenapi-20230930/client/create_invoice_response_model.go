// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateInvoiceResponseBody) *CreateInvoiceResponse
	GetBody() *CreateInvoiceResponseBody
}

type CreateInvoiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceResponse) GoString() string {
	return s.String()
}

func (s *CreateInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInvoiceResponse) GetBody() *CreateInvoiceResponseBody {
	return s.Body
}

func (s *CreateInvoiceResponse) SetHeaders(v map[string]*string) *CreateInvoiceResponse {
	s.Headers = v
	return s
}

func (s *CreateInvoiceResponse) SetStatusCode(v int32) *CreateInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInvoiceResponse) SetBody(v *CreateInvoiceResponseBody) *CreateInvoiceResponse {
	s.Body = v
	return s
}

func (s *CreateInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
