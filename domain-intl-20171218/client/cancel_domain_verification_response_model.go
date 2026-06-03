// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDomainVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDomainVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDomainVerificationResponse
	GetStatusCode() *int32
	SetBody(v *CancelDomainVerificationResponseBody) *CancelDomainVerificationResponse
	GetBody() *CancelDomainVerificationResponseBody
}

type CancelDomainVerificationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDomainVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDomainVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDomainVerificationResponse) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDomainVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDomainVerificationResponse) GetBody() *CancelDomainVerificationResponseBody {
	return s.Body
}

func (s *CancelDomainVerificationResponse) SetHeaders(v map[string]*string) *CancelDomainVerificationResponse {
	s.Headers = v
	return s
}

func (s *CancelDomainVerificationResponse) SetStatusCode(v int32) *CancelDomainVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDomainVerificationResponse) SetBody(v *CancelDomainVerificationResponseBody) *CancelDomainVerificationResponse {
	s.Body = v
	return s
}

func (s *CancelDomainVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
