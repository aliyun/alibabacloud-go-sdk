// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoThreeElementVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertNoThreeElementVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertNoThreeElementVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CertNoThreeElementVerificationResponseBody) *CertNoThreeElementVerificationResponse
	GetBody() *CertNoThreeElementVerificationResponseBody
}

type CertNoThreeElementVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertNoThreeElementVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertNoThreeElementVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CertNoThreeElementVerificationResponse) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertNoThreeElementVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertNoThreeElementVerificationResponse) GetBody() *CertNoThreeElementVerificationResponseBody {
	return s.Body
}

func (s *CertNoThreeElementVerificationResponse) SetHeaders(v map[string]*string) *CertNoThreeElementVerificationResponse {
	s.Headers = v
	return s
}

func (s *CertNoThreeElementVerificationResponse) SetStatusCode(v int32) *CertNoThreeElementVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CertNoThreeElementVerificationResponse) SetBody(v *CertNoThreeElementVerificationResponseBody) *CertNoThreeElementVerificationResponse {
	s.Body = v
	return s
}

func (s *CertNoThreeElementVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
