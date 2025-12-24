// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
	GetRequestId() *string
}

type SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) SetRequestId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
