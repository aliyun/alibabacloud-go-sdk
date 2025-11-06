// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForRegistrantProfileRealNameVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse
	GetBody() *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) GetBody() *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody {
	return s.Body
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetBody(v *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
