// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyThreeElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompanyThreeElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompanyThreeElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CompanyThreeElementsVerificationResponseBody) *CompanyThreeElementsVerificationResponse
	GetBody() *CompanyThreeElementsVerificationResponseBody
}

type CompanyThreeElementsVerificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompanyThreeElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompanyThreeElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompanyThreeElementsVerificationResponse) GetBody() *CompanyThreeElementsVerificationResponseBody {
	return s.Body
}

func (s *CompanyThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyThreeElementsVerificationResponse) SetStatusCode(v int32) *CompanyThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponse) SetBody(v *CompanyThreeElementsVerificationResponseBody) *CompanyThreeElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *CompanyThreeElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
