// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *ApplyInvoiceResponseBody) *ApplyInvoiceResponse
	GetBody() *ApplyInvoiceResponseBody
}

type ApplyInvoiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceResponse) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyInvoiceResponse) GetBody() *ApplyInvoiceResponseBody {
	return s.Body
}

func (s *ApplyInvoiceResponse) SetHeaders(v map[string]*string) *ApplyInvoiceResponse {
	s.Headers = v
	return s
}

func (s *ApplyInvoiceResponse) SetStatusCode(v int32) *ApplyInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyInvoiceResponse) SetBody(v *ApplyInvoiceResponseBody) *ApplyInvoiceResponse {
	s.Body = v
	return s
}

func (s *ApplyInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
