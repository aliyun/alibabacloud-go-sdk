// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningUserPrimaryOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetApplicationId() *string
	SetInstanceId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetInstanceId() *string
	SetUserPrimaryOrganizationalUnitId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetUserPrimaryOrganizationalUnitId() *string
}

type SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest struct {
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

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetUserPrimaryOrganizationalUnitId() *string {
	return s.UserPrimaryOrganizationalUnitId
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetApplicationId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetInstanceId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetUserPrimaryOrganizationalUnitId(v string) *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.UserPrimaryOrganizationalUnitId = &v
	return s
}

func (s *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
