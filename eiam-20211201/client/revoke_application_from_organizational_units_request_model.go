// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeApplicationFromOrganizationalUnitsRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *RevokeApplicationFromOrganizationalUnitsRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *RevokeApplicationFromOrganizationalUnitsRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *RevokeApplicationFromOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
}

type RevokeApplicationFromOrganizationalUnitsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色ID。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations. You can revoke the access permissions from a maximum of 100 organizations at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s RevokeApplicationFromOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetApplicationId(v string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetApplicationRoleId(v string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetInstanceId(v string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
