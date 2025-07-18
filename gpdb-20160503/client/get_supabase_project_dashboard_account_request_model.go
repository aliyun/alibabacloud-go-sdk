// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectDashboardAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetSupabaseProjectDashboardAccountRequest
	GetProjectId() *string
	SetRegionId(v string) *GetSupabaseProjectDashboardAccountRequest
	GetRegionId() *string
}

type GetSupabaseProjectDashboardAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sbp-4ptxp5mp****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupabaseProjectDashboardAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectDashboardAccountRequest) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectDashboardAccountRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseProjectDashboardAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseProjectDashboardAccountRequest) SetProjectId(v string) *GetSupabaseProjectDashboardAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountRequest) SetRegionId(v string) *GetSupabaseProjectDashboardAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountRequest) Validate() error {
	return dara.Validate(s)
}
