// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningUserPrimaryOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetBody() *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
}

type SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetBody() *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	return s.Body
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetBody(v *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
