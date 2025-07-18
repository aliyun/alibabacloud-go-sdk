// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *CreateSupabaseProjectRequest
	GetAccountPassword() *string
	SetClientToken(v string) *CreateSupabaseProjectRequest
	GetClientToken() *string
	SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest
	GetDiskPerformanceLevel() *string
	SetProjectName(v string) *CreateSupabaseProjectRequest
	GetProjectName() *string
	SetProjectSpec(v string) *CreateSupabaseProjectRequest
	GetProjectSpec() *string
	SetRegionId(v string) *CreateSupabaseProjectRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *CreateSupabaseProjectRequest
	GetSecurityIPList() *string
	SetStorageSize(v int64) *CreateSupabaseProjectRequest
	GetStorageSize() *int64
	SetVSwitchId(v string) *CreateSupabaseProjectRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateSupabaseProjectRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateSupabaseProjectRequest
	GetZoneId() *string
}

type CreateSupabaseProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88888888****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// saas_iot_x86_modbustcp_lqt01
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1C1G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// 2
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateSupabaseProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSupabaseProjectRequest) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *CreateSupabaseProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSupabaseProjectRequest) GetProjectSpec() *string {
	return s.ProjectSpec
}

func (s *CreateSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSupabaseProjectRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateSupabaseProjectRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreateSupabaseProjectRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateSupabaseProjectRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSupabaseProjectRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateSupabaseProjectRequest) SetAccountPassword(v string) *CreateSupabaseProjectRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetClientToken(v string) *CreateSupabaseProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetProjectName(v string) *CreateSupabaseProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetProjectSpec(v string) *CreateSupabaseProjectRequest {
	s.ProjectSpec = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetRegionId(v string) *CreateSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetSecurityIPList(v string) *CreateSupabaseProjectRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetStorageSize(v int64) *CreateSupabaseProjectRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetVSwitchId(v string) *CreateSupabaseProjectRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetVpcId(v string) *CreateSupabaseProjectRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetZoneId(v string) *CreateSupabaseProjectRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
