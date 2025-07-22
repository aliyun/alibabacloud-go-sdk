// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceTitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInvoiceTitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInvoiceTitleResponse
	GetStatusCode() *int32
	SetBody(v *ListInvoiceTitleResponseBody) *ListInvoiceTitleResponse
	GetBody() *ListInvoiceTitleResponseBody
}

type ListInvoiceTitleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInvoiceTitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInvoiceTitleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceTitleResponse) GoString() string {
	return s.String()
}

func (s *ListInvoiceTitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInvoiceTitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInvoiceTitleResponse) GetBody() *ListInvoiceTitleResponseBody {
	return s.Body
}

func (s *ListInvoiceTitleResponse) SetHeaders(v map[string]*string) *ListInvoiceTitleResponse {
	s.Headers = v
	return s
}

func (s *ListInvoiceTitleResponse) SetStatusCode(v int32) *ListInvoiceTitleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInvoiceTitleResponse) SetBody(v *ListInvoiceTitleResponseBody) *ListInvoiceTitleResponse {
	s.Body = v
	return s
}

func (s *ListInvoiceTitleResponse) Validate() error {
	return dara.Validate(s)
}
