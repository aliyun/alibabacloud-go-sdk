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
	SetAutoScale(v bool) *CreateSupabaseProjectRequest
	GetAutoScale() *bool
	SetClientToken(v string) *CreateSupabaseProjectRequest
	GetClientToken() *string
	SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest
	GetDiskPerformanceLevel() *string
	SetEngineVersion(v string) *CreateSupabaseProjectRequest
	GetEngineVersion() *string
	SetPayType(v string) *CreateSupabaseProjectRequest
	GetPayType() *string
	SetPeriod(v string) *CreateSupabaseProjectRequest
	GetPeriod() *string
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
	SetUsedTime(v string) *CreateSupabaseProjectRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateSupabaseProjectRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateSupabaseProjectRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateSupabaseProjectRequest
	GetZoneId() *string
}

type CreateSupabaseProjectRequest struct {
	// The password of the initial account.
	//
	// Password rules:
	//
	// - The password must be 8 to 32 characters in length.
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Supported special characters include !@#$%^&*()_+-=.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestPassword123!
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto start/stop. If this parameter is not specified, the default value is false.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The idempotency token. This token ensures that duplicate requests do not trigger the same operation more than once.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The performance level (PL) of the cloud disk. If this parameter is not specified, the default value PL0 is used.
	//
	// Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// - PL2
	//
	// - PL3.
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The DPI engine version. If this parameter is not specified, the default value PG15 is used.
	//
	// Valid values:
	//
	// - PG15: PostgreSQL 15.
	//
	// - PG17: PostgreSQL 17.
	//
	// example:
	//
	// PG15
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The billing method. If this parameter is not specified, the default value Free is used.
	//
	// Valid values:
	//
	// - Free: free tier.
	//
	// - Postpaid: pay-as-you-go.
	//
	// - Prepaid: subscription.
	//
	// example:
	//
	// Free
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. This parameter takes effect only when PayType is set to PrePay. If this parameter is not specified, the default value Month is used.
	//
	// Valid values:
	//
	// - Month: month.
	//
	// - Year: year.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the Supabase project.
	//
	// Naming rules:
	//
	// - The name must be 1 to 128 characters in length.
	//
	// - The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// - The name must start with a letter or an underscore (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// supabase_demo
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The specifications of the Supabase project. The Free billing type uses free-tier specifications. For paid billing types, the specifications must match those available on the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2C4G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// The region ID. Specifies the region in which to create the project.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address whitelist. Separate multiple IP addresses or CIDR blocks with commas (,). If this parameter is not specified, the default value 0.0.0.0/0 is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The storage size. Unit: GB. If this parameter is not specified for non-Free billing types, the default value is 1 GB.
	//
	// example:
	//
	// 50
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The subscription duration of the resource. This parameter takes effect only when PayType is set to PrePay. If this parameter is not specified, the default value is 1.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID. This parameter is required. The zone of the vSwitch must be the same as the value of ZoneId.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1234567890
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC). This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID. The zone of the vSwitch specified by VSwitchId must be the same as the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
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

func (s *CreateSupabaseProjectRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *CreateSupabaseProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSupabaseProjectRequest) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *CreateSupabaseProjectRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateSupabaseProjectRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateSupabaseProjectRequest) GetPeriod() *string {
	return s.Period
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

func (s *CreateSupabaseProjectRequest) GetUsedTime() *string {
	return s.UsedTime
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

func (s *CreateSupabaseProjectRequest) SetAutoScale(v bool) *CreateSupabaseProjectRequest {
	s.AutoScale = &v
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

func (s *CreateSupabaseProjectRequest) SetEngineVersion(v string) *CreateSupabaseProjectRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetPayType(v string) *CreateSupabaseProjectRequest {
	s.PayType = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetPeriod(v string) *CreateSupabaseProjectRequest {
	s.Period = &v
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

func (s *CreateSupabaseProjectRequest) SetUsedTime(v string) *CreateSupabaseProjectRequest {
	s.UsedTime = &v
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
