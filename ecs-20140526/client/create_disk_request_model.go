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
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	AdvancedFeatures *string `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	// > This parameter is not publicly available.
	Arn []*CreateDiskRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable burst (performance bursting). Valid values:
	//
	// - true: enables burst.
	//
	// - false: disables burst.
	//
	// > This parameter is supported only when `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Ensures the idempotency of the request. Generate a parameter value from your client to ensure uniqueness across different requests. **ClientToken*	- only supports ASCII characters and cannot exceed 64 characters. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
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
	// The category of the data disk. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// - cloud_regional_disk_auto: ESSD zone-redundant disk.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - standard.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - premium.
	//
	// Default value: cloud.
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk name. The name must be 2 to 128 characters in length and can contain Unicode letters (including English and Chinese characters) and ASCII digits (0-9). It can contain colons (:), underscores (_), periods (.), or hyphens (-). It must start with a Unicode letter.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// - true: encrypts the disk.
	//
	// - false: does not encrypt the disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// Creates a subscription disk and automatically attaches it to the specified subscription instance (InstanceId).
	//
	// - After you set the instance ID, the ResourceGroupId, Tag.N.Key, Tag.N.Value, ClientToken, and KMSKeyId parameters that you set are ignored.
	//
	// - You cannot specify both ZoneId and InstanceId.
	//
	// Default value: empty, which means a pay-as-you-go disk is created. The location of the disk is determined by RegionId and ZoneId.
	//
	// example:
	//
	// i-bp18pnlg1ds9rky4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The KMS key ID used by the disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption, and the KMSKeyId value is returned after the instance is successfully created.
	//
	// > - - Disk created from a non-shared encrypted snapshot: The encryption key used by the snapshot is used by default.
	//
	// > - - Disk created from a shared encrypted snapshot: The service key is used by default.
	//
	// > - - Disk created in a region where account-level default encryption for block storage is enabled: The specified account-level key is used by default.
	//
	// > - - Other cases: The service key is used by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Specifies whether to enable the multi-attach feature. Valid values:
	//
	// - Disabled: disables the feature.
	//
	// - Enabled: enables the feature. Currently, only ESSDs support setting this to `Enabled`.
	//
	// Default value: Disabled.
	//
	// > Disks with the multi-attach feature enabled only support the pay-as-you-go billing method. Therefore, when `MultiAttach=Enabled`, you cannot set the `InstanceId` parameter. You can call [AttachDisk](https://help.aliyun.com/document_detail/25515.html) to attach the disk after creation, but note that disks with multi-attach enabled can only be attached as data disks.
	//
	// example:
	//
	// Disabled
	MultiAttach  *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Sets the performance level of the disk when creating an ESSD. Valid values:
	//
	// - PL0: maximum random read/write IOPS of 10,000 per disk.
	//
	// - PL1: maximum random read/write IOPS of 50,000 per disk.
	//
	// - PL2: maximum random read/write IOPS of 100,000 per disk.
	//
	// - PL3: maximum random read/write IOPS of 1,000,000 per disk.
	//
	// Default value: PL1.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk (per disk). Possible values:
	//
	// - Capacity (GiB) <= 3: provisioned performance cannot be set.
	//
	// - Capacity (GiB) >= 4: [0, min{(1,000
	//
	//  IOPS/GiB 	- capacity - baseline IOPS), 50,000}]
	//
	//
	// Baseline performance = max{min{1,800 + 50 	- capacity, 50,000}, 3,000}.
	//
	//
	// > This parameter is supported only when `DiskCategory` = `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the disk belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The capacity size. Unit: GiB. You must specify a value for this parameter. Valid values:
	//
	// -   cloud: 5 to 2,000.
	//
	// -   cloud_efficiency: 20 to 32,768.
	//
	// -   cloud_ssd: 20 to 32,768.
	//
	// -   cloud_essd: The valid value range depends on the value of `PerformanceLevel`.
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
	// If you specify the `SnapshotId` parameter, the `SnapshotId` and `Size` parameters have the following restrictions:
	//
	// - If the snapshot capacity corresponding to the `SnapshotId` parameter is greater than the specified `Size` parameter value, the actual size of the created disk equals the size of the specified snapshot.
	//
	// - If the snapshot capacity corresponding to the `SnapshotId` parameter is less than the specified `Size` parameter value, the actual size of the created disk equals the specified `Size` parameter value.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID used to create the disk. Snapshots created on or before July 15, 2013 cannot be used to create disks.
	//
	// The `SnapshotId` and `Size` parameters have the following restrictions:
	//
	// - If the snapshot capacity corresponding to the `SnapshotId` parameter is greater than the specified `Size` parameter value, the actual size of the created disk equals the size of the specified snapshot.
	//
	// - If the snapshot capacity corresponding to the `SnapshotId` parameter is less than the specified `Size` parameter value, the actual size of the created disk equals the specified `Size` parameter value.
	//
	// - Creating elastic ephemeral disks from snapshots is not supported.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The dedicated block storage cluster ID. If you need to create a disk in a specified dedicated block storage cluster, specify this parameter.
	//
	// > Storage set parameters (`StorageSetId`, `StorageSetPartitionNumber`) and dedicated block storage cluster parameter (`StorageClusterId`) are mutually exclusive. If both are set, the API call will fail.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The storage set ID.
	//
	// > Storage set parameters (`StorageSetId`, `StorageSetPartitionNumber`) and dedicated block storage cluster parameter (`StorageClusterId`) are mutually exclusive. If both are set, the API call will fail.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The number of storage set partitions. Valid values: greater than or equal to 2, and cannot exceed the quota limit displayed after calling [DescribeAccountAttributes](https://help.aliyun.com/document_detail/73772.html).
	//
	// Default value: 2.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The list of tag information for the disk.
	Tag []*CreateDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Creates a pay-as-you-go disk in the specified zone.
	//
	// - If you do not set InstanceId, ZoneId is required.
	//
	// - You cannot specify both ZoneId and InstanceId.
	//
	//
	// > Disks of the `cloud_regional_disk_auto` type do not require ZoneId to be set.
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
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not publicly available.
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
	// The tag key of the disk. Valid values of N: 1 to 20. Once a Tag.N.Key value is specified, it cannot be an empty string. It supports up to 128 characters and cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk. Valid values of N: 1 to 20. Once a Tag.N.Value value is specified, it can be an empty string. It supports up to 128 characters and cannot contain `http://` or `https://`.
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
