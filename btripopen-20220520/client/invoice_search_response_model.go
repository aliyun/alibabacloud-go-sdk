// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvoiceSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvoiceSearchResponse
	GetStatusCode() *int32
	SetBody(v *InvoiceSearchResponseBody) *InvoiceSearchResponse
	GetBody() *InvoiceSearchResponseBody
}

type InvoiceSearchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvoiceSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvoiceSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s InvoiceSearchResponse) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvoiceSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvoiceSearchResponse) GetBody() *InvoiceSearchResponseBody {
	return s.Body
}

func (s *InvoiceSearchResponse) SetHeaders(v map[string]*string) *InvoiceSearchResponse {
	s.Headers = v
	return s
}

func (s *InvoiceSearchResponse) SetStatusCode(v int32) *InvoiceSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *InvoiceSearchResponse) SetBody(v *InvoiceSearchResponseBody) *InvoiceSearchResponse {
	s.Body = v
	return s
}

func (s *InvoiceSearchResponse) Validate() error {
	return dara.Validate(s)
}
