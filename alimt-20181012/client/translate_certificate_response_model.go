// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateCertificateResponse
	GetStatusCode() *int32
	SetBody(v *TranslateCertificateResponseBody) *TranslateCertificateResponse
	GetBody() *TranslateCertificateResponseBody
}

type TranslateCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateResponse) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateCertificateResponse) GetBody() *TranslateCertificateResponseBody {
	return s.Body
}

func (s *TranslateCertificateResponse) SetHeaders(v map[string]*string) *TranslateCertificateResponse {
	s.Headers = v
	return s
}

func (s *TranslateCertificateResponse) SetStatusCode(v int32) *TranslateCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateCertificateResponse) SetBody(v *TranslateCertificateResponseBody) *TranslateCertificateResponse {
	s.Body = v
	return s
}

func (s *TranslateCertificateResponse) Validate() error {
	return dara.Validate(s)
}
