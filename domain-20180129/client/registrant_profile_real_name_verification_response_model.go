// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistrantProfileRealNameVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegistrantProfileRealNameVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegistrantProfileRealNameVerificationResponse
	GetStatusCode() *int32
	SetBody(v *RegistrantProfileRealNameVerificationResponseBody) *RegistrantProfileRealNameVerificationResponse
	GetBody() *RegistrantProfileRealNameVerificationResponseBody
}

type RegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegistrantProfileRealNameVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegistrantProfileRealNameVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegistrantProfileRealNameVerificationResponse) GetBody() *RegistrantProfileRealNameVerificationResponseBody {
	return s.Body
}

func (s *RegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *RegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *RegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponse) SetBody(v *RegistrantProfileRealNameVerificationResponseBody) *RegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
