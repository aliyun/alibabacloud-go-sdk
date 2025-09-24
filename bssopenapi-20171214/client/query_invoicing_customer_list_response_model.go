// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvoicingCustomerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInvoicingCustomerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInvoicingCustomerListResponse
	GetStatusCode() *int32
	SetBody(v *QueryInvoicingCustomerListResponseBody) *QueryInvoicingCustomerListResponse
	GetBody() *QueryInvoicingCustomerListResponseBody
}

type QueryInvoicingCustomerListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInvoicingCustomerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInvoicingCustomerListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListResponse) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInvoicingCustomerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInvoicingCustomerListResponse) GetBody() *QueryInvoicingCustomerListResponseBody {
	return s.Body
}

func (s *QueryInvoicingCustomerListResponse) SetHeaders(v map[string]*string) *QueryInvoicingCustomerListResponse {
	s.Headers = v
	return s
}

func (s *QueryInvoicingCustomerListResponse) SetStatusCode(v int32) *QueryInvoicingCustomerListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInvoicingCustomerListResponse) SetBody(v *QueryInvoicingCustomerListResponseBody) *QueryInvoicingCustomerListResponse {
	s.Body = v
	return s
}

func (s *QueryInvoicingCustomerListResponse) Validate() error {
	return dara.Validate(s)
}
