// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTaxiInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeTaxiInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeTaxiInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse
	GetBody() *RecognizeTaxiInvoiceResponseBody
}

type RecognizeTaxiInvoiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxiInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeTaxiInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeTaxiInvoiceResponse) GetBody() *RecognizeTaxiInvoiceResponseBody {
	return s.Body
}

func (s *RecognizeTaxiInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTaxiInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetStatusCode(v int32) *RecognizeTaxiInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
