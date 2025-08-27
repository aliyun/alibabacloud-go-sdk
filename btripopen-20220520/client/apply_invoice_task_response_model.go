// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyInvoiceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyInvoiceTaskResponse
	GetStatusCode() *int32
	SetBody(v *ApplyInvoiceTaskResponseBody) *ApplyInvoiceTaskResponse
	GetBody() *ApplyInvoiceTaskResponseBody
}

type ApplyInvoiceTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyInvoiceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyInvoiceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskResponse) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyInvoiceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyInvoiceTaskResponse) GetBody() *ApplyInvoiceTaskResponseBody {
	return s.Body
}

func (s *ApplyInvoiceTaskResponse) SetHeaders(v map[string]*string) *ApplyInvoiceTaskResponse {
	s.Headers = v
	return s
}

func (s *ApplyInvoiceTaskResponse) SetStatusCode(v int32) *ApplyInvoiceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyInvoiceTaskResponse) SetBody(v *ApplyInvoiceTaskResponseBody) *ApplyInvoiceTaskResponse {
	s.Body = v
	return s
}

func (s *ApplyInvoiceTaskResponse) Validate() error {
	return dara.Validate(s)
}
