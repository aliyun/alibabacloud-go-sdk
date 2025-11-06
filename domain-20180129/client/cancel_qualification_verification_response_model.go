// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelQualificationVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelQualificationVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelQualificationVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CancelQualificationVerificationResponseBody) *CancelQualificationVerificationResponse
	GetBody() *CancelQualificationVerificationResponseBody
}

type CancelQualificationVerificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelQualificationVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelQualificationVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelQualificationVerificationResponse) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelQualificationVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelQualificationVerificationResponse) GetBody() *CancelQualificationVerificationResponseBody {
	return s.Body
}

func (s *CancelQualificationVerificationResponse) SetHeaders(v map[string]*string) *CancelQualificationVerificationResponse {
	s.Headers = v
	return s
}

func (s *CancelQualificationVerificationResponse) SetStatusCode(v int32) *CancelQualificationVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelQualificationVerificationResponse) SetBody(v *CancelQualificationVerificationResponseBody) *CancelQualificationVerificationResponse {
	s.Body = v
	return s
}

func (s *CancelQualificationVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
