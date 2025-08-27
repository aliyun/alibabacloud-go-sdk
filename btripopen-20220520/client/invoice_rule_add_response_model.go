// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceRuleAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceRuleAddResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceRuleAddResponseBody) *InvoiceRuleAddResponse
	GetBody() *InvoiceRuleAddResponseBody
}

type InvoiceRuleAddResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceRuleAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceRuleAddResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddResponse) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceRuleAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceRuleAddResponse) GetBody() *InvoiceRuleAddResponseBody {
	return s.Body
}

func (s *InvoiceRuleAddResponse) SetHeaders(v map[string]*string) *InvoiceRuleAddResponse {
	s.Headers = v
	return s
}

func (s *InvoiceRuleAddResponse) SetStatusCode(v int32) *InvoiceRuleAddResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceRuleAddResponse) SetBody(v *InvoiceRuleAddResponseBody) *InvoiceRuleAddResponse {
	s.Body = v
	return s
}

func (s *InvoiceRuleAddResponse) Validate() error {
	return dara.Validate(s)
}
