// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v []*CloneDisksRequestArn) *CloneDisksRequest
	GetArn() []*CloneDisksRequestArn
	SetBurstingEnabled(v bool) *CloneDisksRequest
	GetBurstingEnabled() *bool
	SetClientToken(v string) *CloneDisksRequest
	GetClientToken() *string
	SetDiskCategory(v string) *CloneDisksRequest
	GetDiskCategory() *string
	SetDiskName(v string) *CloneDisksRequest
	GetDiskName() *string
	SetDryRun(v string) *CloneDisksRequest
	GetDryRun() *string
	SetEncrypted(v bool) *CloneDisksRequest
	GetEncrypted() *bool
	SetKmsKeyId(v string) *CloneDisksRequest
	GetKmsKeyId() *string
	SetMultiAttach(v string) *CloneDisksRequest
	GetMultiAttach() *string
	SetOwnerId(v int64) *CloneDisksRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *CloneDisksRequest
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *CloneDisksRequest
	GetProvisionedIops() *int64
	SetRegionId(v string) *CloneDisksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CloneDisksRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CloneDisksRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *CloneDisksRequest
	GetSize() *int32
	SetSourceDiskId(v string) *CloneDisksRequest
	GetSourceDiskId() *string
	SetTag(v []*CloneDisksRequestTag) *CloneDisksRequest
	GetTag() []*CloneDisksRequestTag
}

type CloneDisksRequest struct {
	// > This parameter is not publicly available.
	Arn []*CloneDisksRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature for the new disk. Valid values:
	//
	// - true: enables the performance burst feature.
	//
	// - false: does not enable the performance burst feature.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://www.alibabacloud.com/help/en/ecs/user-guide/essd-autopl-disks).
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The category of the new disk. Valid values:
	//
	// - cloud_essd: standard SSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// - cloud_regional_disk_auto: regional Standard SSD (ESSD).
	//
	// > Disk category restrictions for disk cloning:
	//
	// > - Non-regional disks can only be cloned to non-regional disk types.
	//
	// > - Regional disks can only be cloned to regional disk types.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). The name must start with a letter.
	//
	// Default value: empty.
	//
	// example:
	//
	// MyDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - true: sends a check request without querying the filing status. The system checks whether your AccessKey pair is valid, whether RAM user authorization is granted, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check passes, the DryRunOperation error code is returned.
	//
	// - false (default): sends a normal request. After the check passes, a 2XX HTTP status code is returned and the filing status is queried.
	//
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the new disk is encrypted. Valid values:
	//
	// - true: The new disk is encrypted.
	//
	// - false: The new disk is not encrypted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The key ID of the KMS key used by the new disk.
	//
	// example:
	//
	// key-szz67b2f696f4wh9yeg5d
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// Specifies whether to enable the multi-attach attribute for the new disk. Settings for this parameter. Valid values:
	//
	// - Disabled: disables the multi-attach attribute.
	//
	// - Enabled: enables the multi-attach attribute. Only standard SSDs support the value `Enabled`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The performance level (PL) of the enterprise SSD to create. Settings for this parameter vary based on the standard SSD type. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// > If DiskCategory is set to cloud_essd, PerformanceLevel is required.
	//
	// For more information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values:
	//
	// - Capacity (GiB) <= 3: provisioned performance is not supported.
	//
	// - Capacity (GiB) >= 4: [0, min{(1,000 IOPS/GiB × Capacity - Baseline IOPS), 50,000}]
	//
	// Baseline performance = max{min{1,800 + 50 × Capacity, 50,000}, 3,000}.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://www.alibabacloud.com/help/en/ecs/user-guide/essd-autopl-disks).
	//
	// example:
	//
	// 10
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID. You can call [DescribeRegions](https://www.alibabacloud.com/help/en/ecs/api-regions-describeregions) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-bp199lyny9b3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The capacity of the new disk. Unit: GiB. You must specify a value for this parameter. Valid values:
	//
	// - cloud_essd: the valid value range varies based on the performance level.
	//
	//     - PL0: 1 to 65,536.
	//
	//     - PL1: 20 to 65,536.
	//
	//     - PL2: 461 to 65,536.
	//
	//     - PL3: 1,261 to 65,536.
	//
	// - cloud_auto: 1 to 65,536.
	//
	// - cloud_essd_entry: 10 to 32,768.
	//
	// - cloud_regional_disk_auto: 10 to 65,536.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the source disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1d6tsvznfghy7y****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The list of tags for the disk.
	Tag []*CloneDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CloneDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequest) GoString() string {
	return s.String()
}

func (s *CloneDisksRequest) GetArn() []*CloneDisksRequestArn {
	return s.Arn
}

func (s *CloneDisksRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CloneDisksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloneDisksRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CloneDisksRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CloneDisksRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CloneDisksRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CloneDisksRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CloneDisksRequest) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *CloneDisksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloneDisksRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CloneDisksRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CloneDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneDisksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CloneDisksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloneDisksRequest) GetSize() *int32 {
	return s.Size
}

func (s *CloneDisksRequest) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *CloneDisksRequest) GetTag() []*CloneDisksRequestTag {
	return s.Tag
}

func (s *CloneDisksRequest) SetArn(v []*CloneDisksRequestArn) *CloneDisksRequest {
	s.Arn = v
	return s
}

func (s *CloneDisksRequest) SetBurstingEnabled(v bool) *CloneDisksRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CloneDisksRequest) SetClientToken(v string) *CloneDisksRequest {
	s.ClientToken = &v
	return s
}

func (s *CloneDisksRequest) SetDiskCategory(v string) *CloneDisksRequest {
	s.DiskCategory = &v
	return s
}

func (s *CloneDisksRequest) SetDiskName(v string) *CloneDisksRequest {
	s.DiskName = &v
	return s
}

func (s *CloneDisksRequest) SetDryRun(v string) *CloneDisksRequest {
	s.DryRun = &v
	return s
}

func (s *CloneDisksRequest) SetEncrypted(v bool) *CloneDisksRequest {
	s.Encrypted = &v
	return s
}

func (s *CloneDisksRequest) SetKmsKeyId(v string) *CloneDisksRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CloneDisksRequest) SetMultiAttach(v string) *CloneDisksRequest {
	s.MultiAttach = &v
	return s
}

func (s *CloneDisksRequest) SetOwnerId(v int64) *CloneDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *CloneDisksRequest) SetPerformanceLevel(v string) *CloneDisksRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CloneDisksRequest) SetProvisionedIops(v int64) *CloneDisksRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CloneDisksRequest) SetRegionId(v string) *CloneDisksRequest {
	s.RegionId = &v
	return s
}

func (s *CloneDisksRequest) SetResourceGroupId(v string) *CloneDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloneDisksRequest) SetResourceOwnerId(v int64) *CloneDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloneDisksRequest) SetSize(v int32) *CloneDisksRequest {
	s.Size = &v
	return s
}

func (s *CloneDisksRequest) SetSourceDiskId(v string) *CloneDisksRequest {
	s.SourceDiskId = &v
	return s
}

func (s *CloneDisksRequest) SetTag(v []*CloneDisksRequestTag) *CloneDisksRequest {
	s.Tag = v
	return s
}

func (s *CloneDisksRequest) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloneDisksRequestArn struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CloneDisksRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequestArn) GoString() string {
	return s.String()
}

func (s *CloneDisksRequestArn) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CloneDisksRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CloneDisksRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CloneDisksRequestArn) SetAssumeRoleFor(v string) *CloneDisksRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CloneDisksRequestArn) SetRoleType(v string) *CloneDisksRequestArn {
	s.RoleType = &v
	return s
}

func (s *CloneDisksRequestArn) SetRolearn(v string) *CloneDisksRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CloneDisksRequestArn) Validate() error {
	return dara.Validate(s)
}

type CloneDisksRequestTag struct {
	// The tag key of the disk. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CloneDisksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequestTag) GoString() string {
	return s.String()
}

func (s *CloneDisksRequestTag) GetKey() *string {
	return s.Key
}

func (s *CloneDisksRequestTag) GetValue() *string {
	return s.Value
}

func (s *CloneDisksRequestTag) SetKey(v string) *CloneDisksRequestTag {
	s.Key = &v
	return s
}

func (s *CloneDisksRequestTag) SetValue(v string) *CloneDisksRequestTag {
	s.Value = &v
	return s
}

func (s *CloneDisksRequestTag) Validate() error {
	return dara.Validate(s)
}
