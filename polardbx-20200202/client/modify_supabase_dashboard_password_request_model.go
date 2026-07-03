// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseDashboardPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifySupabaseDashboardPasswordRequest
	GetDBInstanceName() *string
	SetNewPassword(v string) *ModifySupabaseDashboardPasswordRequest
	GetNewPassword() *string
	SetRegionId(v string) *ModifySupabaseDashboardPasswordRequest
	GetRegionId() *string
}

type ModifySupabaseDashboardPasswordRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The new password.
	//
	// This parameter is required.
	//
	// example:
	//
	// ********
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySupabaseDashboardPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseDashboardPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseDashboardPasswordRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySupabaseDashboardPasswordRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ModifySupabaseDashboardPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseDashboardPasswordRequest) SetDBInstanceName(v string) *ModifySupabaseDashboardPasswordRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordRequest) SetNewPassword(v string) *ModifySupabaseDashboardPasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordRequest) SetRegionId(v string) *ModifySupabaseDashboardPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordRequest) Validate() error {
	return dara.Validate(s)
}
