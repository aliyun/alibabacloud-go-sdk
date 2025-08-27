// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceRuleDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceRuleDeleteResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceRuleDeleteResponseBody) *InvoiceRuleDeleteResponse
	GetBody() *InvoiceRuleDeleteResponseBody
}

type InvoiceRuleDeleteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceRuleDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceRuleDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteResponse) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceRuleDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceRuleDeleteResponse) GetBody() *InvoiceRuleDeleteResponseBody {
	return s.Body
}

func (s *InvoiceRuleDeleteResponse) SetHeaders(v map[string]*string) *InvoiceRuleDeleteResponse {
	s.Headers = v
	return s
}

func (s *InvoiceRuleDeleteResponse) SetStatusCode(v int32) *InvoiceRuleDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceRuleDeleteResponse) SetBody(v *InvoiceRuleDeleteResponseBody) *InvoiceRuleDeleteResponse {
	s.Body = v
	return s
}

func (s *InvoiceRuleDeleteResponse) Validate() error {
	return dara.Validate(s)
}
