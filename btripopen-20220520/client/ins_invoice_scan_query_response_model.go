// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsInvoiceScanQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsInvoiceScanQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsInvoiceScanQueryResponse
	GetStatusCode() *int32
	SetBody(v *InsInvoiceScanQueryResponseBody) *InsInvoiceScanQueryResponse
	GetBody() *InsInvoiceScanQueryResponseBody
}

type InsInvoiceScanQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsInvoiceScanQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsInvoiceScanQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryResponse) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsInvoiceScanQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsInvoiceScanQueryResponse) GetBody() *InsInvoiceScanQueryResponseBody {
	return s.Body
}

func (s *InsInvoiceScanQueryResponse) SetHeaders(v map[string]*string) *InsInvoiceScanQueryResponse {
	s.Headers = v
	return s
}

func (s *InsInvoiceScanQueryResponse) SetStatusCode(v int32) *InsInvoiceScanQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *InsInvoiceScanQueryResponse) SetBody(v *InsInvoiceScanQueryResponseBody) *InsInvoiceScanQueryResponse {
	s.Body = v
	return s
}

func (s *InsInvoiceScanQueryResponse) Validate() error {
	return dara.Validate(s)
}
