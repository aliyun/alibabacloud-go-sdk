// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSupabaseProjectPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *ResetSupabaseProjectPasswordRequest
	GetAccountPassword() *string
	SetDashboardPassword(v string) *ResetSupabaseProjectPasswordRequest
	GetDashboardPassword() *string
	SetProjectId(v string) *ResetSupabaseProjectPasswordRequest
	GetProjectId() *string
	SetRegionId(v string) *ResetSupabaseProjectPasswordRequest
	GetRegionId() *string
}

type ResetSupabaseProjectPasswordRequest struct {
	// The password of the database account.
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The supported special characters are: `!@#$%^&*()_+-=`
	//
	// - The password must be 8 to 32 characters in length.
	//
	// example:
	//
	// Pw123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The Supabase Dashboard password.
	//
	// example:
	//
	// zaQBjed6d4ubV4o
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The Supabase instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-tyarplz****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetSupabaseProjectPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSupabaseProjectPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetSupabaseProjectPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetSupabaseProjectPasswordRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *ResetSupabaseProjectPasswordRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ResetSupabaseProjectPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSupabaseProjectPasswordRequest) SetAccountPassword(v string) *ResetSupabaseProjectPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) SetDashboardPassword(v string) *ResetSupabaseProjectPasswordRequest {
	s.DashboardPassword = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) SetProjectId(v string) *ResetSupabaseProjectPasswordRequest {
	s.ProjectId = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) SetRegionId(v string) *ResetSupabaseProjectPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) Validate() error {
	return dara.Validate(s)
}
