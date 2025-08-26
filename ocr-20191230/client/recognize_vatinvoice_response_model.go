// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVATInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeVATInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeVATInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeVATInvoiceResponseBody) *RecognizeVATInvoiceResponse
	GetBody() *RecognizeVATInvoiceResponseBody
}

type RecognizeVATInvoiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVATInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVATInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeVATInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeVATInvoiceResponse) GetBody() *RecognizeVATInvoiceResponseBody {
	return s.Body
}

func (s *RecognizeVATInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeVATInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVATInvoiceResponse) SetStatusCode(v int32) *RecognizeVATInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponse) SetBody(v *RecognizeVATInvoiceResponseBody) *RecognizeVATInvoiceResponse {
	s.Body = v
	return s
}

func (s *RecognizeVATInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
