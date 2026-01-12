// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstancePasswordRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The Supabase Dashboard password.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **ResetInstancePassword**.
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
