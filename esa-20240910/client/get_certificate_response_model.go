// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetCertificateResponseBody) *GetCertificateResponse
	GetBody() *GetCertificateResponseBody
}

type GetCertificateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertificateResponse) GetBody() *GetCertificateResponseBody {
	return s.Body
}

func (s *GetCertificateResponse) SetHeaders(v map[string]*string) *GetCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetCertificateResponse) SetStatusCode(v int32) *GetCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificateResponse) SetBody(v *GetCertificateResponseBody) *GetCertificateResponse {
	s.Body = v
	return s
}

func (s *GetCertificateResponse) Validate() error {
	return dara.Validate(s)
}
