// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetSupabaseProjectResponseBody
	GetCreateTime() *string
	SetDashboardPassword(v string) *GetSupabaseProjectResponseBody
	GetDashboardPassword() *string
	SetDashboardUserName(v string) *GetSupabaseProjectResponseBody
	GetDashboardUserName() *string
	SetDiskPerformanceLevel(v string) *GetSupabaseProjectResponseBody
	GetDiskPerformanceLevel() *string
	SetEngine(v string) *GetSupabaseProjectResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *GetSupabaseProjectResponseBody
	GetEngineVersion() *string
	SetPrivateConnectUrl(v string) *GetSupabaseProjectResponseBody
	GetPrivateConnectUrl() *string
	SetProjectId(v string) *GetSupabaseProjectResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetSupabaseProjectResponseBody
	GetProjectName() *string
	SetProjectSpec(v string) *GetSupabaseProjectResponseBody
	GetProjectSpec() *string
	SetPublicConnectUrl(v string) *GetSupabaseProjectResponseBody
	GetPublicConnectUrl() *string
	SetRegionId(v string) *GetSupabaseProjectResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetSupabaseProjectResponseBody
	GetRequestId() *string
	SetSecurityIpList(v string) *GetSupabaseProjectResponseBody
	GetSecurityIpList() *string
	SetStatus(v string) *GetSupabaseProjectResponseBody
	GetStatus() *string
	SetStorageSize(v int64) *GetSupabaseProjectResponseBody
	GetStorageSize() *int64
	SetVSwitchId(v string) *GetSupabaseProjectResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *GetSupabaseProjectResponseBody
	GetVpcId() *string
	SetZoneId(v string) *GetSupabaseProjectResponseBody
	GetZoneId() *string
}

type GetSupabaseProjectResponseBody struct {
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxpassword
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// example:
	//
	// username
	DashboardUserName *string `json:"DashboardUserName,omitempty" xml:"DashboardUserName,omitempty"`
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// example:
	//
	// postgres
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 15
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 192.168.0.11
	PrivateConnectUrl *string `json:"PrivateConnectUrl,omitempty" xml:"PrivateConnectUrl,omitempty"`
	// example:
	//
	// sbp-545434
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// supabase_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1C1G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// example:
	//
	// 10.154.11.10
	PublicConnectUrl *string `json:"PublicConnectUrl,omitempty" xml:"PublicConnectUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSupabaseProjectResponseBody) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *GetSupabaseProjectResponseBody) GetDashboardUserName() *string {
	return s.DashboardUserName
}

func (s *GetSupabaseProjectResponseBody) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *GetSupabaseProjectResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *GetSupabaseProjectResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GetSupabaseProjectResponseBody) GetPrivateConnectUrl() *string {
	return s.PrivateConnectUrl
}

func (s *GetSupabaseProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSupabaseProjectResponseBody) GetProjectSpec() *string {
	return s.ProjectSpec
}

func (s *GetSupabaseProjectResponseBody) GetPublicConnectUrl() *string {
	return s.PublicConnectUrl
}

func (s *GetSupabaseProjectResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupabaseProjectResponseBody) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *GetSupabaseProjectResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSupabaseProjectResponseBody) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *GetSupabaseProjectResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetSupabaseProjectResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetSupabaseProjectResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetSupabaseProjectResponseBody) SetCreateTime(v string) *GetSupabaseProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetDashboardPassword(v string) *GetSupabaseProjectResponseBody {
	s.DashboardPassword = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetDashboardUserName(v string) *GetSupabaseProjectResponseBody {
	s.DashboardUserName = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetDiskPerformanceLevel(v string) *GetSupabaseProjectResponseBody {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetEngine(v string) *GetSupabaseProjectResponseBody {
	s.Engine = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetEngineVersion(v string) *GetSupabaseProjectResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetPrivateConnectUrl(v string) *GetSupabaseProjectResponseBody {
	s.PrivateConnectUrl = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetProjectId(v string) *GetSupabaseProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetProjectName(v string) *GetSupabaseProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetProjectSpec(v string) *GetSupabaseProjectResponseBody {
	s.ProjectSpec = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetPublicConnectUrl(v string) *GetSupabaseProjectResponseBody {
	s.PublicConnectUrl = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetRegionId(v string) *GetSupabaseProjectResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetRequestId(v string) *GetSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetSecurityIpList(v string) *GetSupabaseProjectResponseBody {
	s.SecurityIpList = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetStatus(v string) *GetSupabaseProjectResponseBody {
	s.Status = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetStorageSize(v int64) *GetSupabaseProjectResponseBody {
	s.StorageSize = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetVSwitchId(v string) *GetSupabaseProjectResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetVpcId(v string) *GetSupabaseProjectResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetZoneId(v string) *GetSupabaseProjectResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
