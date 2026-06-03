// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfileRealNameVerificationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRegistrantProfileRealNameVerificationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRegistrantProfileRealNameVerificationInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryRegistrantProfileRealNameVerificationInfoResponseBody) *QueryRegistrantProfileRealNameVerificationInfoResponse
	GetBody() *QueryRegistrantProfileRealNameVerificationInfoResponseBody
}

type QueryRegistrantProfileRealNameVerificationInfoResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRegistrantProfileRealNameVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) GetBody() *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	return s.Body
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetHeaders(v map[string]*string) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetStatusCode(v int32) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetBody(v *QueryRegistrantProfileRealNameVerificationInfoResponseBody) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.Body = v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
