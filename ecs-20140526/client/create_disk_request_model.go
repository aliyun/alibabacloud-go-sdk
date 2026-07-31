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
	// >This parameter is not publicly available.
	Arn []*CreateDiskRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The category of the data disk. Valid values:
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the disk. The name must be 2 to 128 characters in length and can contain characters that are categorized as letter in Unicode, including Chinese and English characters, and ASCII digits (0-9). The name can contain colons (:), underscores (_), periods (.), or hyphens (-). The name must start with a character that is categorized as letter in Unicode.
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
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the subscription instance to which the subscription disk is automatically attached after the disk is created.
	//
	// example:
	//
	// i-bp18pnlg1ds9rky4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Key Management Service (KMS) key used by the disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Specifies whether to enable the multi-attach feature. Valid values:
	//
	// example:
	//
	// Disabled
	MultiAttach  *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The performance level of the enterprise SSD (ESSD) to create. Valid values:
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned performance read/write IOPS of the ESSD AutoPL disk. Valid values:
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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
	// The size of the disk. Unit: GiB. You must specify a value for this parameter. Valid values:
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create the disk. Snapshots created on or before July 15, 2013 cannot be used to create disks.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. To create a disk in a specific dedicated block storage cluster, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The number of partitions in the storage set. Valid values: greater than or equal to 2, up to the privilege quota limit returned by calling [DescribeAccountAttributes](https://help.aliyun.com/document_detail/73772.html).
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The list of tags for the disk.
	Tag []*CreateDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone in which to create a pay-as-you-go disk.
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
	// >This parameter is not publicly available.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// >This parameter is not publicly available.
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
	// The tag key of the disk. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
