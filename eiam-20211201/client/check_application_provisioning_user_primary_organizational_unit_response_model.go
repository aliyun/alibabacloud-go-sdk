// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetBody() *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
}

type CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string                                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetBody() *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	return s.Body
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetBody(v *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
