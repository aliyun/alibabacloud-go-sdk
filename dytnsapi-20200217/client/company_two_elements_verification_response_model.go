// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyTwoElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompanyTwoElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompanyTwoElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CompanyTwoElementsVerificationResponseBody) *CompanyTwoElementsVerificationResponse
	GetBody() *CompanyTwoElementsVerificationResponseBody
}

type CompanyTwoElementsVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyTwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompanyTwoElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompanyTwoElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompanyTwoElementsVerificationResponse) GetBody() *CompanyTwoElementsVerificationResponseBody {
	return s.Body
}

func (s *CompanyTwoElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyTwoElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyTwoElementsVerificationResponse) SetStatusCode(v int32) *CompanyTwoElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponse) SetBody(v *CompanyTwoElementsVerificationResponseBody) *CompanyTwoElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *CompanyTwoElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
