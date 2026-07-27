// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstancePasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ResetInstancePasswordRequest
	GetBranchName() *string
	SetDashboardPassword(v string) *ResetInstancePasswordRequest
	GetDashboardPassword() *string
	SetDatabasePassword(v string) *ResetInstancePasswordRequest
	GetDatabasePassword() *string
	SetInstanceName(v string) *ResetInstancePasswordRequest
	GetInstanceName() *string
	SetRegionId(v string) *ResetInstancePasswordRequest
	GetRegionId() *string
}

type ResetInstancePasswordRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The Supabase Dashboard password.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The RDS database access password.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// 	Notice: This password change also updates the access passwords of the following accounts on the associated PostgreSQL instance. These accounts are required by Supabase: postgres, supabase_admin, supabase_auth_admin, supabase_functions_admin, supabase_storage_admin, authenticator, pgbouncer.
	//
	// </notice>
	//
	// example:
	//
	// test_Password
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetInstancePasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetInstancePasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetInstancePasswordRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ResetInstancePasswordRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *ResetInstancePasswordRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *ResetInstancePasswordRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ResetInstancePasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetInstancePasswordRequest) SetBranchName(v string) *ResetInstancePasswordRequest {
	s.BranchName = &v
	return s
}

func (s *ResetInstancePasswordRequest) SetDashboardPassword(v string) *ResetInstancePasswordRequest {
	s.DashboardPassword = &v
	return s
}

func (s *ResetInstancePasswordRequest) SetDatabasePassword(v string) *ResetInstancePasswordRequest {
	s.DatabasePassword = &v
	return s
}

func (s *ResetInstancePasswordRequest) SetInstanceName(v string) *ResetInstancePasswordRequest {
	s.InstanceName = &v
	return s
}

func (s *ResetInstancePasswordRequest) SetRegionId(v string) *ResetInstancePasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetInstancePasswordRequest) Validate() error {
	return dara.Validate(s)
}
