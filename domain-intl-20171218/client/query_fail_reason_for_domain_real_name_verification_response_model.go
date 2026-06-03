// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForDomainRealNameVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFailReasonForDomainRealNameVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFailReasonForDomainRealNameVerificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryFailReasonForDomainRealNameVerificationResponseBody) *QueryFailReasonForDomainRealNameVerificationResponse
	GetBody() *QueryFailReasonForDomainRealNameVerificationResponseBody
}

type QueryFailReasonForDomainRealNameVerificationResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFailReasonForDomainRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) GetBody() *QueryFailReasonForDomainRealNameVerificationResponseBody {
	return s.Body
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetHeaders(v map[string]*string) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetStatusCode(v int32) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetBody(v *QueryFailReasonForDomainRealNameVerificationResponseBody) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.Body = v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
