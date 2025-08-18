// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetCertificateResponseBody) *SetCertificateResponse
	GetBody() *SetCertificateResponseBody
}

type SetCertificateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCertificateResponse) GetBody() *SetCertificateResponseBody {
	return s.Body
}

func (s *SetCertificateResponse) SetHeaders(v map[string]*string) *SetCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCertificateResponse) SetStatusCode(v int32) *SetCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCertificateResponse) SetBody(v *SetCertificateResponseBody) *SetCertificateResponse {
	s.Body = v
	return s
}

func (s *SetCertificateResponse) Validate() error {
	return dara.Validate(s)
}
