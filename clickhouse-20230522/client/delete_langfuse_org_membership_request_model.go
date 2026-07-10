// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgMembershipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteLangfuseOrgMembershipRequest
	GetDBInstanceId() *string
	SetEmail(v string) *DeleteLangfuseOrgMembershipRequest
	GetEmail() *string
	SetOrganizationId(v string) *DeleteLangfuseOrgMembershipRequest
	GetOrganizationId() *string
	SetRegionId(v string) *DeleteLangfuseOrgMembershipRequest
	GetRegionId() *string
}

type DeleteLangfuseOrgMembershipRequest struct {
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
}

func (s DeleteLangfuseOrgMembershipRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgMembershipRequest) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgMembershipRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseOrgMembershipRequest) GetEmail() *string {
	return s.Email
}

func (s *DeleteLangfuseOrgMembershipRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteLangfuseOrgMembershipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLangfuseOrgMembershipRequest) SetDBInstanceId(v string) *DeleteLangfuseOrgMembershipRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipRequest) SetEmail(v string) *DeleteLangfuseOrgMembershipRequest {
	s.Email = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipRequest) SetOrganizationId(v string) *DeleteLangfuseOrgMembershipRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipRequest) SetRegionId(v string) *DeleteLangfuseOrgMembershipRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipRequest) Validate() error {
	return dara.Validate(s)
}
