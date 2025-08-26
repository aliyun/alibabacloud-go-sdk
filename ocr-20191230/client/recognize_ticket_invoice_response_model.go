// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTicketInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeTicketInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeTicketInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeTicketInvoiceResponseBody) *RecognizeTicketInvoiceResponse
	GetBody() *RecognizeTicketInvoiceResponseBody
}

type RecognizeTicketInvoiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTicketInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTicketInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTicketInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeTicketInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeTicketInvoiceResponse) GetBody() *RecognizeTicketInvoiceResponseBody {
	return s.Body
}

func (s *RecognizeTicketInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTicketInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTicketInvoiceResponse) SetStatusCode(v int32) *RecognizeTicketInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponse) SetBody(v *RecognizeTicketInvoiceResponseBody) *RecognizeTicketInvoiceResponse {
	s.Body = v
	return s
}

func (s *RecognizeTicketInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
