// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CheckCertificateResponseBody) *CheckCertificateResponse
	GetBody() *CheckCertificateResponseBody
}

type CheckCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCertificateResponse) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCertificateResponse) GetBody() *CheckCertificateResponseBody {
	return s.Body
}

func (s *CheckCertificateResponse) SetHeaders(v map[string]*string) *CheckCertificateResponse {
	s.Headers = v
	return s
}

func (s *CheckCertificateResponse) SetStatusCode(v int32) *CheckCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCertificateResponse) SetBody(v *CheckCertificateResponseBody) *CheckCertificateResponse {
	s.Body = v
	return s
}

func (s *CheckCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
