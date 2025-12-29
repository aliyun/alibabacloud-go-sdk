// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyFourElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompanyFourElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompanyFourElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CompanyFourElementsVerificationResponseBody) *CompanyFourElementsVerificationResponse
	GetBody() *CompanyFourElementsVerificationResponseBody
}

type CompanyFourElementsVerificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyFourElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompanyFourElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CompanyFourElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompanyFourElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompanyFourElementsVerificationResponse) GetBody() *CompanyFourElementsVerificationResponseBody {
	return s.Body
}

func (s *CompanyFourElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyFourElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyFourElementsVerificationResponse) SetStatusCode(v int32) *CompanyFourElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyFourElementsVerificationResponse) SetBody(v *CompanyFourElementsVerificationResponseBody) *CompanyFourElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *CompanyFourElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
