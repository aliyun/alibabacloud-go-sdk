// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetApplicationId() *string
	SetInstanceId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetInstanceId() *string
	SetUserPrimaryOrganizationalUnitId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetUserPrimaryOrganizationalUnitId() *string
}

type CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 组织ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	UserPrimaryOrganizationalUnitId *string `json:"UserPrimaryOrganizationalUnitId,omitempty" xml:"UserPrimaryOrganizationalUnitId,omitempty"`
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetUserPrimaryOrganizationalUnitId() *string {
	return s.UserPrimaryOrganizationalUnitId
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetApplicationId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.ApplicationId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetInstanceId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetUserPrimaryOrganizationalUnitId(v string) *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.UserPrimaryOrganizationalUnitId = &v
	return s
}

func (s *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
