// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvoiceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInvoiceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInvoiceEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInvoiceEntityResponseBody) *DeleteInvoiceEntityResponse
	GetBody() *DeleteInvoiceEntityResponseBody
}

type DeleteInvoiceEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInvoiceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInvoiceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInvoiceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInvoiceEntityResponse) GetBody() *DeleteInvoiceEntityResponseBody {
	return s.Body
}

func (s *DeleteInvoiceEntityResponse) SetHeaders(v map[string]*string) *DeleteInvoiceEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteInvoiceEntityResponse) SetStatusCode(v int32) *DeleteInvoiceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInvoiceEntityResponse) SetBody(v *DeleteInvoiceEntityResponseBody) *DeleteInvoiceEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteInvoiceEntityResponse) Validate() error {
	return dara.Validate(s)
}
