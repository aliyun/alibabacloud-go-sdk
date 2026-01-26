// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCustomHostnameCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCustomHostnameCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCustomHostnameCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCustomHostnameCertificateResponseBody) *ApplyCustomHostnameCertificateResponse
	GetBody() *ApplyCustomHostnameCertificateResponseBody
}

type ApplyCustomHostnameCertificateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCustomHostnameCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCustomHostnameCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCustomHostnameCertificateResponse) GoString() string {
	return s.String()
}

func (s *ApplyCustomHostnameCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCustomHostnameCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCustomHostnameCertificateResponse) GetBody() *ApplyCustomHostnameCertificateResponseBody {
	return s.Body
}

func (s *ApplyCustomHostnameCertificateResponse) SetHeaders(v map[string]*string) *ApplyCustomHostnameCertificateResponse {
	s.Headers = v
	return s
}

func (s *ApplyCustomHostnameCertificateResponse) SetStatusCode(v int32) *ApplyCustomHostnameCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCustomHostnameCertificateResponse) SetBody(v *ApplyCustomHostnameCertificateResponseBody) *ApplyCustomHostnameCertificateResponse {
	s.Body = v
	return s
}

func (s *ApplyCustomHostnameCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
