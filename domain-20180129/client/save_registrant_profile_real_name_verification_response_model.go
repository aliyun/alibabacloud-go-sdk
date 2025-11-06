// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileRealNameVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveRegistrantProfileRealNameVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveRegistrantProfileRealNameVerificationResponse
	GetStatusCode() *int32
	SetBody(v *SaveRegistrantProfileRealNameVerificationResponseBody) *SaveRegistrantProfileRealNameVerificationResponse
	GetBody() *SaveRegistrantProfileRealNameVerificationResponseBody
}

type SaveRegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveRegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveRegistrantProfileRealNameVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) GetBody() *SaveRegistrantProfileRealNameVerificationResponseBody {
	return s.Body
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *SaveRegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *SaveRegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetBody(v *SaveRegistrantProfileRealNameVerificationResponseBody) *SaveRegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
