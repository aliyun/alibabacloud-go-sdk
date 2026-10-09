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
	SetBackupId(v string) *CreateSupabaseProjectRequest
	GetBackupId() *string
	SetClientToken(v string) *CreateSupabaseProjectRequest
	GetClientToken() *string
	SetCreateOptions(v string) *CreateSupabaseProjectRequest
	GetCreateOptions() *string
	SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest
	GetDiskPerformanceLevel() *string
	SetEngineVersion(v string) *CreateSupabaseProjectRequest
	GetEngineVersion() *string
	SetLightweight(v bool) *CreateSupabaseProjectRequest
	GetLightweight() *bool
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
	SetSrcProjectId(v string) *CreateSupabaseProjectRequest
	GetSrcProjectId() *string
	SetStorageSize(v int64) *CreateSupabaseProjectRequest
	GetStorageSize() *int64
	SetTags(v []*CreateSupabaseProjectRequestTags) *CreateSupabaseProjectRequest
	GetTags() []*CreateSupabaseProjectRequestTags
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
	// The initial account password.
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
	// Specifies whether to enable auto-start and auto-stop. If you do not specify this parameter, the default value is false.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The backup set ID.
	//
	// > You can call [ListSupabaseDataBackups](https://help.aliyun.com/document_detail/3064623.html) to view the IDs of all backup sets under the target Supabase project.
	//
	// example:
	//
	// 2176307784
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The client token. It is used to ensure idempotence and prevent duplicate requests from executing the same operation.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The optional creation parameters. The default value is empty.
	//
	// example:
	//
	// {}
	CreateOptions *string `json:"CreateOptions,omitempty" xml:"CreateOptions,omitempty"`
	// The performance level of the cloud disk. If you do not specify this parameter, the default value is PL0.
	//
	// Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// - PL2
	//
	// - PL3
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The DPI engine version. If you do not specify this parameter, the default value is PG15. PostgreSQL 17 and later versions support the data sandbox (branch) feature.
	//
	// Valid values:
	//
	// - PG15: PostgreSQL 15.
	//
	// - PG17: PostgreSQL 17, which supports the data sandbox feature.
	//
	// example:
	//
	// PG15
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether the project is the lightweight edition.
	//
	// example:
	//
	// false
	Lightweight *bool `json:"Lightweight,omitempty" xml:"Lightweight,omitempty"`
	// The billing method. If you do not specify this parameter, the default value is Free.
	//
	// Valid values:
	//
	// - Free: the free billing method.
	//
	// - Postpaid: pay-as-you-go.
	//
	// - Prepaid: subscription.
	//
	// example:
	//
	// Free
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. This parameter takes effect only when PayType is set to Prepaid. If you do not specify this parameter, the default value is Month.
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
	// - The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// - The name must start with a letter or an underscore (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// supabase_demo
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The specifications of the Supabase project. The free billing method uses the free specifications. For paid billing methods, the specifications must be consistent with those available in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2C4G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address whitelist. Separate multiple IP addresses or CIDR blocks with commas (,). If you do not specify this parameter, the default value 0.0.0.0/0 is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The ID of the Supabase project to which the backup set belongs.
	//
	// example:
	//
	// spb-xxxxxxxx
	SrcProjectId *string `json:"SrcProjectId,omitempty" xml:"SrcProjectId,omitempty"`
	// The storage capacity. Unit: GB. If you do not specify this parameter for a non-free billing method, the default value is 1.
	//
	// example:
	//
	// 50
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The list of tags.
	Tags []*CreateSupabaseProjectRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The subscription duration of the resource. This parameter takes effect only when PayType is set to Prepaid. If you do not specify this parameter, the default value is 1.
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

func (s *CreateSupabaseProjectRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateSupabaseProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSupabaseProjectRequest) GetCreateOptions() *string {
	return s.CreateOptions
}

func (s *CreateSupabaseProjectRequest) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *CreateSupabaseProjectRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateSupabaseProjectRequest) GetLightweight() *bool {
	return s.Lightweight
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

func (s *CreateSupabaseProjectRequest) GetSrcProjectId() *string {
	return s.SrcProjectId
}

func (s *CreateSupabaseProjectRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreateSupabaseProjectRequest) GetTags() []*CreateSupabaseProjectRequestTags {
	return s.Tags
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

func (s *CreateSupabaseProjectRequest) SetBackupId(v string) *CreateSupabaseProjectRequest {
	s.BackupId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetClientToken(v string) *CreateSupabaseProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetCreateOptions(v string) *CreateSupabaseProjectRequest {
	s.CreateOptions = &v
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

func (s *CreateSupabaseProjectRequest) SetLightweight(v bool) *CreateSupabaseProjectRequest {
	s.Lightweight = &v
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

func (s *CreateSupabaseProjectRequest) SetSrcProjectId(v string) *CreateSupabaseProjectRequest {
	s.SrcProjectId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetStorageSize(v int64) *CreateSupabaseProjectRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetTags(v []*CreateSupabaseProjectRequestTags) *CreateSupabaseProjectRequest {
	s.Tags = v
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSupabaseProjectRequestTags struct {
	// The tag key. Limits:
	//
	// - It cannot be an empty string.
	//
	// - It can be up to 128 characters in length.
	//
	// - It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be an empty string. It can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSupabaseProjectRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectRequestTags) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateSupabaseProjectRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateSupabaseProjectRequestTags) SetKey(v string) *CreateSupabaseProjectRequestTags {
	s.Key = &v
	return s
}

func (s *CreateSupabaseProjectRequestTags) SetValue(v string) *CreateSupabaseProjectRequestTags {
	s.Value = &v
	return s
}

func (s *CreateSupabaseProjectRequestTags) Validate() error {
	return dara.Validate(s)
}
