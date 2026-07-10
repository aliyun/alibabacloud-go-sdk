// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseOrgMembershipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyLangfuseOrgMembershipRequest
	GetDBInstanceId() *string
	SetEmail(v string) *ModifyLangfuseOrgMembershipRequest
	GetEmail() *string
	SetOrganizationId(v string) *ModifyLangfuseOrgMembershipRequest
	GetOrganizationId() *string
	SetRegionId(v string) *ModifyLangfuseOrgMembershipRequest
	GetRegionId() *string
	SetRole(v string) *ModifyLangfuseOrgMembershipRequest
	GetRole() *string
}

type ModifyLangfuseOrgMembershipRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The Langfuse organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the user in the organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ModifyLangfuseOrgMembershipRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseOrgMembershipRequest) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseOrgMembershipRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyLangfuseOrgMembershipRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseOrgMembershipRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ModifyLangfuseOrgMembershipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLangfuseOrgMembershipRequest) GetRole() *string {
	return s.Role
}

func (s *ModifyLangfuseOrgMembershipRequest) SetDBInstanceId(v string) *ModifyLangfuseOrgMembershipRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipRequest) SetEmail(v string) *ModifyLangfuseOrgMembershipRequest {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipRequest) SetOrganizationId(v string) *ModifyLangfuseOrgMembershipRequest {
	s.OrganizationId = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipRequest) SetRegionId(v string) *ModifyLangfuseOrgMembershipRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipRequest) SetRole(v string) *ModifyLangfuseOrgMembershipRequest {
	s.Role = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipRequest) Validate() error {
	return dara.Validate(s)
}
