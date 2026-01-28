// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceUserOrg interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v int64) *GrafanaWorkspaceUserOrg
	GetOrgId() *int64
	SetOrgName(v string) *GrafanaWorkspaceUserOrg
	GetOrgName() *string
	SetRole(v string) *GrafanaWorkspaceUserOrg
	GetRole() *string
}

type GrafanaWorkspaceUserOrg struct {
	// example:
	//
	// 1
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// main org
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// admin
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GrafanaWorkspaceUserOrg) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceUserOrg) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceUserOrg) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GrafanaWorkspaceUserOrg) GetOrgName() *string {
	return s.OrgName
}

func (s *GrafanaWorkspaceUserOrg) GetRole() *string {
	return s.Role
}

func (s *GrafanaWorkspaceUserOrg) SetOrgId(v int64) *GrafanaWorkspaceUserOrg {
	s.OrgId = &v
	return s
}

func (s *GrafanaWorkspaceUserOrg) SetOrgName(v string) *GrafanaWorkspaceUserOrg {
	s.OrgName = &v
	return s
}

func (s *GrafanaWorkspaceUserOrg) SetRole(v string) *GrafanaWorkspaceUserOrg {
	s.Role = &v
	return s
}

func (s *GrafanaWorkspaceUserOrg) Validate() error {
	return dara.Validate(s)
}
