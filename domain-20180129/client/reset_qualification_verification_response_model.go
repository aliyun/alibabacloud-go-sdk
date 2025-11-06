// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetQualificationVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetQualificationVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetQualificationVerificationResponse
	GetStatusCode() *int32
	SetBody(v *ResetQualificationVerificationResponseBody) *ResetQualificationVerificationResponse
	GetBody() *ResetQualificationVerificationResponseBody
}

type ResetQualificationVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetQualificationVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetQualificationVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetQualificationVerificationResponse) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetQualificationVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetQualificationVerificationResponse) GetBody() *ResetQualificationVerificationResponseBody {
	return s.Body
}

func (s *ResetQualificationVerificationResponse) SetHeaders(v map[string]*string) *ResetQualificationVerificationResponse {
	s.Headers = v
	return s
}

func (s *ResetQualificationVerificationResponse) SetStatusCode(v int32) *ResetQualificationVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetQualificationVerificationResponse) SetBody(v *ResetQualificationVerificationResponseBody) *ResetQualificationVerificationResponse {
	s.Body = v
	return s
}

func (s *ResetQualificationVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
