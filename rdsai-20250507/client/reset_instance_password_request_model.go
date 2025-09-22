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
	SetInstanceName(v string) *ResetInstancePasswordRequest
	GetInstanceName() *string
	SetRegionId(v string) *ResetInstancePasswordRequest
	GetRegionId() *string
}

type ResetInstancePasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
