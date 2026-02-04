// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckApplicationProvisioningUserPrimaryOuResult(v *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
	GetCheckApplicationProvisioningUserPrimaryOuResult() *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult
	SetRequestId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody
	GetRequestId() *string
}

type CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody struct {
	CheckApplicationProvisioningUserPrimaryOuResult *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult `json:"CheckApplicationProvisioningUserPrimaryOuResult,omitempty" xml:"CheckApplicationProvisioningUserPrimaryOuResult,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GetCheckApplicationProvisioningUserPrimaryOuResult() *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult {
	return s.CheckApplicationProvisioningUserPrimaryOuResult
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) SetCheckApplicationProvisioningUserPrimaryOuResult(v *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	s.CheckApplicationProvisioningUserPrimaryOuResult = v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) SetRequestId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBody) Validate() error {
	if s.CheckApplicationProvisioningUserPrimaryOuResult != nil {
		if err := s.CheckApplicationProvisioningUserPrimaryOuResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult struct {
	// 是否授权
	AllowAuthorization *bool `json:"AllowAuthorization,omitempty" xml:"AllowAuthorization,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) String() string {
	return dara.Prettify(s)
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) GoString() string {
	return s.String()
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) GetAllowAuthorization() *bool {
	return s.AllowAuthorization
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) SetAllowAuthorization(v bool) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult {
	s.AllowAuthorization = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) SetApplicationId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult {
	s.ApplicationId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) SetInstanceId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult {
	s.InstanceId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponseBodyCheckApplicationProvisioningUserPrimaryOuResult) Validate() error {
	return dara.Validate(s)
}
