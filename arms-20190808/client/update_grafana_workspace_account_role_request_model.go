// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *UpdateGrafanaWorkspaceAccountRoleRequest
	GetAccountId() *int64
	SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceAccountRoleRequest
	GetGrafanaWorkspaceId() *string
	SetOrgId(v int64) *UpdateGrafanaWorkspaceAccountRoleRequest
	GetOrgId() *int64
	SetRegionId(v string) *UpdateGrafanaWorkspaceAccountRoleRequest
	GetRegionId() *string
	SetRole(v string) *UpdateGrafanaWorkspaceAccountRoleRequest
	GetRole() *string
}

type UpdateGrafanaWorkspaceAccountRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1035730062732547
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-ipp4v54ms01
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	OrgId *int64 `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s UpdateGrafanaWorkspaceAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) GetRole() *string {
	return s.Role
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) SetAccountId(v int64) *UpdateGrafanaWorkspaceAccountRoleRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceAccountRoleRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) SetOrgId(v int64) *UpdateGrafanaWorkspaceAccountRoleRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) SetRegionId(v string) *UpdateGrafanaWorkspaceAccountRoleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) SetRole(v string) *UpdateGrafanaWorkspaceAccountRoleRequest {
	s.Role = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
