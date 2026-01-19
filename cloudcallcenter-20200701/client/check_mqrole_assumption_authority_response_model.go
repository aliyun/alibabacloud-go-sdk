// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMQRoleAssumptionAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMQRoleAssumptionAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *CheckMQRoleAssumptionAuthorityResponseBody) *CheckMQRoleAssumptionAuthorityResponse
	GetBody() *CheckMQRoleAssumptionAuthorityResponseBody
}

type CheckMQRoleAssumptionAuthorityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMQRoleAssumptionAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMQRoleAssumptionAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMQRoleAssumptionAuthorityResponse) GoString() string {
	return s.String()
}

func (s *CheckMQRoleAssumptionAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMQRoleAssumptionAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMQRoleAssumptionAuthorityResponse) GetBody() *CheckMQRoleAssumptionAuthorityResponseBody {
	return s.Body
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetHeaders(v map[string]*string) *CheckMQRoleAssumptionAuthorityResponse {
	s.Headers = v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetBody(v *CheckMQRoleAssumptionAuthorityResponseBody) *CheckMQRoleAssumptionAuthorityResponse {
	s.Body = v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
