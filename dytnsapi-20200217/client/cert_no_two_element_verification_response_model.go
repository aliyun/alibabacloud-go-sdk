// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoTwoElementVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CertNoTwoElementVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CertNoTwoElementVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CertNoTwoElementVerificationResponseBody) *CertNoTwoElementVerificationResponse
	GetBody() *CertNoTwoElementVerificationResponseBody
}

type CertNoTwoElementVerificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertNoTwoElementVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertNoTwoElementVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CertNoTwoElementVerificationResponse) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CertNoTwoElementVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CertNoTwoElementVerificationResponse) GetBody() *CertNoTwoElementVerificationResponseBody {
	return s.Body
}

func (s *CertNoTwoElementVerificationResponse) SetHeaders(v map[string]*string) *CertNoTwoElementVerificationResponse {
	s.Headers = v
	return s
}

func (s *CertNoTwoElementVerificationResponse) SetStatusCode(v int32) *CertNoTwoElementVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CertNoTwoElementVerificationResponse) SetBody(v *CertNoTwoElementVerificationResponseBody) *CertNoTwoElementVerificationResponse {
	s.Body = v
	return s
}

func (s *CertNoTwoElementVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
