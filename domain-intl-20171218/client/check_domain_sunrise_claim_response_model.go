// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainSunriseClaimResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDomainSunriseClaimResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDomainSunriseClaimResponse
	GetStatusCode() *int32
	SetBody(v *CheckDomainSunriseClaimResponseBody) *CheckDomainSunriseClaimResponse
	GetBody() *CheckDomainSunriseClaimResponseBody
}

type CheckDomainSunriseClaimResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainSunriseClaimResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainSunriseClaimResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainSunriseClaimResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDomainSunriseClaimResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDomainSunriseClaimResponse) GetBody() *CheckDomainSunriseClaimResponseBody {
	return s.Body
}

func (s *CheckDomainSunriseClaimResponse) SetHeaders(v map[string]*string) *CheckDomainSunriseClaimResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainSunriseClaimResponse) SetStatusCode(v int32) *CheckDomainSunriseClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainSunriseClaimResponse) SetBody(v *CheckDomainSunriseClaimResponseBody) *CheckDomainSunriseClaimResponse {
	s.Body = v
	return s
}

func (s *CheckDomainSunriseClaimResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
