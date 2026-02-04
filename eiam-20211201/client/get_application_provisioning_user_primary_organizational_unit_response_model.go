// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningUserPrimaryOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
	GetBody() *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
}

type GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) GetBody() *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	return s.Body
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) SetBody(v *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
