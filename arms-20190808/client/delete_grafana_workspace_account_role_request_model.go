// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *DeleteGrafanaWorkspaceAccountRoleRequest
	GetAccountId() *int64
	SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceAccountRoleRequest
	GetGrafanaWorkspaceId() *string
	SetOrgId(v int64) *DeleteGrafanaWorkspaceAccountRoleRequest
	GetOrgId() *int64
	SetRegionId(v string) *DeleteGrafanaWorkspaceAccountRoleRequest
	GetRegionId() *string
}

type DeleteGrafanaWorkspaceAccountRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1844567291383820
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-lbj3j3gfy04
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
}

func (s DeleteGrafanaWorkspaceAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) SetAccountId(v int64) *DeleteGrafanaWorkspaceAccountRoleRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceAccountRoleRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) SetOrgId(v int64) *DeleteGrafanaWorkspaceAccountRoleRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) SetRegionId(v string) *DeleteGrafanaWorkspaceAccountRoleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
