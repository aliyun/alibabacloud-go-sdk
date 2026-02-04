// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningUserPrimaryOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
	GetInstanceId() *string
}

type GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest struct {
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
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetApplicationId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) SetInstanceId(v string) *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
