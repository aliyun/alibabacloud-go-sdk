// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealNameVerificationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainRealNameVerificationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainRealNameVerificationInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainRealNameVerificationInfoResponseBody) *QueryDomainRealNameVerificationInfoResponse
	GetBody() *QueryDomainRealNameVerificationInfoResponseBody
}

type QueryDomainRealNameVerificationInfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainRealNameVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainRealNameVerificationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainRealNameVerificationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainRealNameVerificationInfoResponse) GetBody() *QueryDomainRealNameVerificationInfoResponseBody {
	return s.Body
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetHeaders(v map[string]*string) *QueryDomainRealNameVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetStatusCode(v int32) *QueryDomainRealNameVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetBody(v *QueryDomainRealNameVerificationInfoResponseBody) *QueryDomainRealNameVerificationInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
