// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScale(v string) *GetSupabaseProjectResponseBody
	GetAutoScale() *string
	SetCreateTime(v string) *GetSupabaseProjectResponseBody
	GetCreateTime() *string
	SetDBSecurityIpList(v string) *GetSupabaseProjectResponseBody
	GetDBSecurityIpList() *string
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
	SetEni(v string) *GetSupabaseProjectResponseBody
	GetEni() *string
	SetInstanceVersion(v string) *GetSupabaseProjectResponseBody
	GetInstanceVersion() *string
	SetPayType(v string) *GetSupabaseProjectResponseBody
	GetPayType() *string
	SetPrivateConnectUrl(v string) *GetSupabaseProjectResponseBody
	GetPrivateConnectUrl() *string
	SetProjectDescription(v string) *GetSupabaseProjectResponseBody
	GetProjectDescription() *string
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
	SetStorageType(v string) *GetSupabaseProjectResponseBody
	GetStorageType() *string
	SetVSwitchId(v string) *GetSupabaseProjectResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *GetSupabaseProjectResponseBody
	GetVpcId() *string
	SetZoneId(v string) *GetSupabaseProjectResponseBody
	GetZoneId() *string
}

type GetSupabaseProjectResponseBody struct {
	// Indicates whether the **auto pause and resume*	- feature is enabled.
	//
	// Valid values:
	//
	// - `true`: The feature is enabled. The project automatically pauses and resumes based on traffic.
	//
	// - `false`: The feature is disabled.
	//
	// example:
	//
	// false
	AutoScale *string `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The creation time of the project.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database IP address whitelist, specified as a comma-separated string.
	//
	// example:
	//
	// 127.0.0.1,100.64.XX.XX/10
	DBSecurityIpList *string `json:"DBSecurityIpList,omitempty" xml:"DBSecurityIpList,omitempty"`
	// The password for the Supabase Dashboard. This parameter is not used.
	//
	// example:
	//
	// xxpassword
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The username for the Supabase Dashboard. This parameter is not used.
	//
	// example:
	//
	// username
	DashboardUserName *string `json:"DashboardUserName,omitempty" xml:"DashboardUserName,omitempty"`
	// The performance level (PL) of the cloud disk. Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The database engine.
	//
	// example:
	//
	// postgres
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 15
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The elastic network interface (ENI) ID.
	//
	// example:
	//
	// eni-xxxxxx
	Eni *string `json:"Eni,omitempty" xml:"Eni,omitempty"`
	// The current instance version.
	//
	// example:
	//
	// v1.0.3
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The billing method. Valid values:
	//
	// - `POSTPAY`: pay-as-you-go
	//
	// - `PREPAY`: subscription
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The private connection URL for the Supabase Dashboard.
	//
	// example:
	//
	// 192.168.0.11
	PrivateConnectUrl *string `json:"PrivateConnectUrl,omitempty" xml:"PrivateConnectUrl,omitempty"`
	// The description of the Supabase project.
	//
	// example:
	//
	// for-test-project
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The Supabase project ID.
	//
	// example:
	//
	// sbp-545434
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The Supabase project name.
	//
	// example:
	//
	// supabase_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Supabase instance specification.
	//
	// example:
	//
	// 1C1G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// The public connection URL for the Supabase Dashboard.
	//
	// example:
	//
	// 10.154.11.10
	PublicConnectUrl *string `json:"PublicConnectUrl,omitempty" xml:"PublicConnectUrl,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address whitelist, specified as a comma-separated string.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	// The Supabase instance status.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage space, in GB.
	//
	// example:
	//
	// 2
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud_essd_pl0**
	//
	// - **cloud_essd_pl1**
	//
	// - **cloud_essd_pl2**
	//
	// - **cloud_essd_pl3**
	//
	// example:
	//
	// cloud_essd_pl0
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// > - You can call the [DescribeRdsVpcs](https://help.aliyun.com/document_detail/208327.html) operation to query the available VPCs.
	//
	// >
	//
	// > - This parameter is required.
	//
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available zones.
	//
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

func (s *GetSupabaseProjectResponseBody) GetAutoScale() *string {
	return s.AutoScale
}

func (s *GetSupabaseProjectResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSupabaseProjectResponseBody) GetDBSecurityIpList() *string {
	return s.DBSecurityIpList
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

func (s *GetSupabaseProjectResponseBody) GetEni() *string {
	return s.Eni
}

func (s *GetSupabaseProjectResponseBody) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *GetSupabaseProjectResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetSupabaseProjectResponseBody) GetPrivateConnectUrl() *string {
	return s.PrivateConnectUrl
}

func (s *GetSupabaseProjectResponseBody) GetProjectDescription() *string {
	return s.ProjectDescription
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

func (s *GetSupabaseProjectResponseBody) GetStorageType() *string {
	return s.StorageType
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

func (s *GetSupabaseProjectResponseBody) SetAutoScale(v string) *GetSupabaseProjectResponseBody {
	s.AutoScale = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetCreateTime(v string) *GetSupabaseProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetDBSecurityIpList(v string) *GetSupabaseProjectResponseBody {
	s.DBSecurityIpList = &v
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

func (s *GetSupabaseProjectResponseBody) SetEni(v string) *GetSupabaseProjectResponseBody {
	s.Eni = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetInstanceVersion(v string) *GetSupabaseProjectResponseBody {
	s.InstanceVersion = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetPayType(v string) *GetSupabaseProjectResponseBody {
	s.PayType = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetPrivateConnectUrl(v string) *GetSupabaseProjectResponseBody {
	s.PrivateConnectUrl = &v
	return s
}

func (s *GetSupabaseProjectResponseBody) SetProjectDescription(v string) *GetSupabaseProjectResponseBody {
	s.ProjectDescription = &v
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

func (s *GetSupabaseProjectResponseBody) SetStorageType(v string) *GetSupabaseProjectResponseBody {
	s.StorageType = &v
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
