// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedFeatures(v string) *CreateDiskRequest
	GetAdvancedFeatures() *string
	SetArn(v []*CreateDiskRequestArn) *CreateDiskRequest
	GetArn() []*CreateDiskRequestArn
	SetBurstingEnabled(v bool) *CreateDiskRequest
	GetBurstingEnabled() *bool
	SetClientToken(v string) *CreateDiskRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDiskRequest
	GetDescription() *string
	SetDiskCategory(v string) *CreateDiskRequest
	GetDiskCategory() *string
	SetDiskName(v string) *CreateDiskRequest
	GetDiskName() *string
	SetEncryptAlgorithm(v string) *CreateDiskRequest
	GetEncryptAlgorithm() *string
	SetEncrypted(v bool) *CreateDiskRequest
	GetEncrypted() *bool
	SetInstanceId(v string) *CreateDiskRequest
	GetInstanceId() *string
	SetKMSKeyId(v string) *CreateDiskRequest
	GetKMSKeyId() *string
	SetMultiAttach(v string) *CreateDiskRequest
	GetMultiAttach() *string
	SetOwnerAccount(v string) *CreateDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDiskRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *CreateDiskRequest
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *CreateDiskRequest
	GetProvisionedIops() *int64
	SetRegionId(v string) *CreateDiskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDiskRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDiskRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *CreateDiskRequest
	GetSize() *int32
	SetSnapshotId(v string) *CreateDiskRequest
	GetSnapshotId() *string
	SetStorageClusterId(v string) *CreateDiskRequest
	GetStorageClusterId() *string
	SetStorageSetId(v string) *CreateDiskRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateDiskRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest
	GetTag() []*CreateDiskRequestTag
	SetZoneId(v string) *CreateDiskRequest
	GetZoneId() *string
}

type CreateDiskRequest struct {
	// This parameter is not available for use.
	//
	// example:
	//
	// hide
	AdvancedFeatures *string `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	// > This parameter is not available for use.
	Arn []*CreateDiskRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// > This parameter is supported only when `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The disk category of the data disk. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// - cloud_regional_disk_auto: regional Enterprise SSD (ESSD).
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - standard edition.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - premium edition.
	//
	// Default value: cloud.
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk name. The name must be 2 to 128 characters in length and can contain Unicode letters (including English and Chinese characters) and ASCII digits (0–9). It can also contain colons (:), underscores (_), periods (.), or hyphens (-). It must start with a Unicode letter.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// This parameter is not available for use.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// - true: The disk is encrypted.
	//
	// - false: The disk is not encrypted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// Creates a subscription disk and automatically attaches it to the specified subscription instance (InstanceId).
	//
	// - If you specify an instance ID, the ResourceGroupId, Tag.N.Key, Tag.N.Value, ClientToken, and KMSKeyId parameters are ignored.
	//
	// - You cannot specify both ZoneId and InstanceId at the same time.
	//
	// Default value: empty. An empty value indicates that a pay-as-you-go disk is created. The region of the disk is determined by RegionId and ZoneId.
	//
	// example:
	//
	// i-bp18pnlg1ds9rky4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key used for the disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption, and the KMSKeyId value is returned after the instance is created.
	//
	// > - - If the disk is created from an unshared encrypted snapshot, the encryption key used by that snapshot is used by default.
	//
	// > - - If the disk is created from a shared encrypted snapshot, the service key is used for encryption by default.
	//
	// > - - If the disk is created in a region where account-level default encryption for block storage is enabled, the specified account-level key is used for encryption by default.
	//
	// > - - In all other cases, the service key is used for encryption by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Specifies whether to enable the multi-attach attribute. Valid values:
	//
	// - Disabled: The feature is disabled.
	//
	// - Enabled: The feature is enabled. Currently, only enterprise SSDs support `Enabled`.
	//
	// Default value: Disabled.
	//
	// > Disks with the multi-attach attribute enabled support only the pay-as-you-go billing method. Therefore, when `MultiAttach=Enabled`, you cannot specify the `InstanceId` parameter at the same time. You can call [AttachDisk](https://help.aliyun.com/document_detail/25515.html) to attach the disk after it is created. Note that disks with multi-attach enabled can only be attached as data disks.
	//
	// example:
	//
	// Disabled
	MultiAttach  *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The performance level of the enterprise SSD to create. Valid values:
	//
	// - PL0: Maximum random read/write IOPS of 10,000 per disk.
	//
	// - PL1: Maximum random read/write IOPS of 50,000 per disk.
	//
	// - PL2: Maximum random read/write IOPS of 100,000 per disk.
	//
	// - PL3: Maximum random read/write IOPS of 1,000,000 per disk.
	//
	// Default value: PL1.
	//
	// For information about how to choose an ESSD performance level, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk.
	//
	// - Capacity (GiB) ≤ 3: Setting provisioned performance is not supported.
	//
	// - Capacity (GiB) ≥ 4: [0, min{(1,000 IOPS/GiB × capacity − baseline IOPS), 50,000}]
	//
	//
	// Baseline performance = max{min{1,800 + 50 × capacity, 50,000}, 3,000}.
	//
	//
	// > This parameter is supported only when `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of the disk. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the latest list of Alibaba Cloud regions.
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
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The disk capacity. Unit: GiB. This parameter is required. Valid values:
	//
	// -   cloud: 5 to 2,000.
	//
	// -   cloud_efficiency: 20 to 32,768.
	//
	// -   cloud_ssd: 20 to 32,768.
	//
	// -   cloud_essd: The valid range depends on the value of `PerformanceLevel`.
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
	// - elastic_ephemeral_disk_standard: 64 to 8,192.
	//
	// - elastic_ephemeral_disk_premium: 64 to 8,192.
	//
	// If you specify `SnapshotId`, the following limits apply:
	//
	// - If the snapshot size specified by `SnapshotId` is greater than the value of `Size`, the actual disk size equals the snapshot size.
	//
	// - If the snapshot size specified by `SnapshotId` is smaller than the value of `Size`, the actual disk size equals the value of `Size`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create the disk. Snapshots created on or before July 15, 2013 cannot be used to create disks.
	//
	// The following limits apply when you specify both `SnapshotId` and `Size`:
	//
	// - If the snapshot size specified by `SnapshotId` is greater than the value of `Size`, the actual disk size equals the snapshot size.
	//
	// - If the snapshot size specified by `SnapshotId` is smaller than the value of `Size`, the actual disk size equals the value of `Size`.
	//
	// - Creating elastic ephemeral disks from snapshots is not supported.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. Specify this parameter if you want to create a disk in a specific dedicated block storage cluster.
	//
	// > You can specify either the storage set parameters (`StorageSetId` and `StorageSetPartitionNumber`) or the dedicated block storage cluster parameter (`StorageClusterId`), but not both. If both are specified, the call fails.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The storage set ID.
	//
	// > You can specify either the storage set parameters (`StorageSetId` and `StorageSetPartitionNumber`) or the dedicated block storage cluster parameter (`StorageClusterId`), but not both. If both are specified, the call fails.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The number of partitions in the storage set. The value must be greater than or equal to 2 and cannot exceed the privilege quota limit returned by [DescribeAccountAttributes](https://help.aliyun.com/document_detail/73772.html).
	//
	// Default value: 2.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags to add to the disk.
	Tag []*CreateDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone in which to create a pay-as-you-go disk.
	//
	// - If you do not specify InstanceId, ZoneId is required.
	//
	// - You cannot specify both ZoneId and InstanceId at the same time.
	//
	//
	// > Disks of the `cloud_regional_disk_auto` type do not require a ZoneId.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskRequest) GetAdvancedFeatures() *string {
	return s.AdvancedFeatures
}

func (s *CreateDiskRequest) GetArn() []*CreateDiskRequestArn {
	return s.Arn
}

func (s *CreateDiskRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDiskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDiskRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateDiskRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateDiskRequest) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateDiskRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDiskRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateDiskRequest) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *CreateDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDiskRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateDiskRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDiskRequest) GetSize() *int32 {
	return s.Size
}

func (s *CreateDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateDiskRequest) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateDiskRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateDiskRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateDiskRequest) GetTag() []*CreateDiskRequestTag {
	return s.Tag
}

func (s *CreateDiskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDiskRequest) SetAdvancedFeatures(v string) *CreateDiskRequest {
	s.AdvancedFeatures = &v
	return s
}

func (s *CreateDiskRequest) SetArn(v []*CreateDiskRequestArn) *CreateDiskRequest {
	s.Arn = v
	return s
}

func (s *CreateDiskRequest) SetBurstingEnabled(v bool) *CreateDiskRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateDiskRequest) SetClientToken(v string) *CreateDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskRequest) SetDescription(v string) *CreateDiskRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskRequest) SetDiskCategory(v string) *CreateDiskRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateDiskRequest) SetDiskName(v string) *CreateDiskRequest {
	s.DiskName = &v
	return s
}

func (s *CreateDiskRequest) SetEncryptAlgorithm(v string) *CreateDiskRequest {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateDiskRequest) SetEncrypted(v bool) *CreateDiskRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateDiskRequest) SetInstanceId(v string) *CreateDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDiskRequest) SetKMSKeyId(v string) *CreateDiskRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CreateDiskRequest) SetMultiAttach(v string) *CreateDiskRequest {
	s.MultiAttach = &v
	return s
}

func (s *CreateDiskRequest) SetOwnerAccount(v string) *CreateDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDiskRequest) SetOwnerId(v int64) *CreateDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDiskRequest) SetPerformanceLevel(v string) *CreateDiskRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateDiskRequest) SetProvisionedIops(v int64) *CreateDiskRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateDiskRequest) SetRegionId(v string) *CreateDiskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskRequest) SetResourceGroupId(v string) *CreateDiskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskRequest) SetResourceOwnerAccount(v string) *CreateDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDiskRequest) SetResourceOwnerId(v int64) *CreateDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDiskRequest) SetSize(v int32) *CreateDiskRequest {
	s.Size = &v
	return s
}

func (s *CreateDiskRequest) SetSnapshotId(v string) *CreateDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageClusterId(v string) *CreateDiskRequest {
	s.StorageClusterId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageSetId(v string) *CreateDiskRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageSetPartitionNumber(v int32) *CreateDiskRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateDiskRequest) SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest {
	s.Tag = v
	return s
}

func (s *CreateDiskRequest) SetZoneId(v string) *CreateDiskRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDiskRequest) Validate() error {
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

type CreateDiskRequestArn struct {
	// > This parameter is not available for use.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// hide
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateDiskRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequestArn) GoString() string {
	return s.String()
}

func (s *CreateDiskRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateDiskRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateDiskRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateDiskRequestArn) SetAssumeRoleFor(v int64) *CreateDiskRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateDiskRequestArn) SetRoleType(v string) *CreateDiskRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateDiskRequestArn) SetRolearn(v string) *CreateDiskRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CreateDiskRequestArn) Validate() error {
	return dara.Validate(s)
}

type CreateDiskRequestTag struct {
	// The tag key of the disk. Valid values of N: 1 to 20. The tag key cannot be an empty string once specified. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk. Valid values of N: 1 to 20. The tag value can be an empty string once specified. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDiskRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDiskRequestTag) SetKey(v string) *CreateDiskRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskRequestTag) SetValue(v string) *CreateDiskRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDiskRequestTag) Validate() error {
	return dara.Validate(s)
}
