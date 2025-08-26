// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQuotaInvoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeQuotaInvoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeQuotaInvoiceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse
	GetBody() *RecognizeQuotaInvoiceResponseBody
}

type RecognizeQuotaInvoiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeQuotaInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeQuotaInvoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeQuotaInvoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeQuotaInvoiceResponse) GetBody() *RecognizeQuotaInvoiceResponseBody {
	return s.Body
}

func (s *RecognizeQuotaInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeQuotaInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetStatusCode(v int32) *RecognizeQuotaInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse {
	s.Body = v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) Validate() error {
	return dara.Validate(s)
}
