// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceRuleSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceRuleSaveResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceRuleSaveResponseBody) *InvoiceRuleSaveResponse
	GetBody() *InvoiceRuleSaveResponseBody
}

type InvoiceRuleSaveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceRuleSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceRuleSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveResponse) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceRuleSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceRuleSaveResponse) GetBody() *InvoiceRuleSaveResponseBody {
	return s.Body
}

func (s *InvoiceRuleSaveResponse) SetHeaders(v map[string]*string) *InvoiceRuleSaveResponse {
	s.Headers = v
	return s
}

func (s *InvoiceRuleSaveResponse) SetStatusCode(v int32) *InvoiceRuleSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceRuleSaveResponse) SetBody(v *InvoiceRuleSaveResponseBody) *InvoiceRuleSaveResponse {
	s.Body = v
	return s
}

func (s *InvoiceRuleSaveResponse) Validate() error {
	return dara.Validate(s)
}
