// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
	GetRequestId() *string
	SetUserPrimaryOrganizationalUnitId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
	GetUserPrimaryOrganizationalUnitId() *string
}

type GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ou_12121xxxxxx
	UserPrimaryOrganizationalUnitId *string `json:"UserPrimaryOrganizationalUnitId,omitempty" xml:"UserPrimaryOrganizationalUnitId,omitempty"`
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GetUserPrimaryOrganizationalUnitId() *string {
	return s.UserPrimaryOrganizationalUnitId
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) SetRequestId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) SetUserPrimaryOrganizationalUnitId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	s.UserPrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
