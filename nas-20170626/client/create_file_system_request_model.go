// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateFileSystemRequest
	GetBandwidth() *int64
	SetCapacity(v int64) *CreateFileSystemRequest
	GetCapacity() *int64
	SetChargeType(v string) *CreateFileSystemRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateFileSystemRequest
	GetClientToken() *string
	SetDescription(v string) *CreateFileSystemRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFileSystemRequest
	GetDryRun() *bool
	SetDuration(v int32) *CreateFileSystemRequest
	GetDuration() *int32
	SetEncryptType(v int32) *CreateFileSystemRequest
	GetEncryptType() *int32
	SetFileSystemType(v string) *CreateFileSystemRequest
	GetFileSystemType() *string
	SetHpnZone(v string) *CreateFileSystemRequest
	GetHpnZone() *string
	SetKmsKeyId(v string) *CreateFileSystemRequest
	GetKmsKeyId() *string
	SetProtocolType(v string) *CreateFileSystemRequest
	GetProtocolType() *string
	SetRedundancyType(v string) *CreateFileSystemRequest
	GetRedundancyType() *string
	SetRedundancyVSwitchIds(v []*string) *CreateFileSystemRequest
	GetRedundancyVSwitchIds() []*string
	SetResourceGroupId(v string) *CreateFileSystemRequest
	GetResourceGroupId() *string
	SetSnapshotId(v string) *CreateFileSystemRequest
	GetSnapshotId() *string
	SetStorageType(v string) *CreateFileSystemRequest
	GetStorageType() *string
	SetTag(v []*CreateFileSystemRequestTag) *CreateFileSystemRequest
	GetTag() []*CreateFileSystemRequestTag
	SetVSwitchId(v string) *CreateFileSystemRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateFileSystemRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateFileSystemRequest
	GetZoneId() *string
}

type CreateFileSystemRequest struct {
	// The maximum throughput of the file system. Unit: MB/s.
	//
	// For valid values, refer to the specifications on the buy page.
	//
	// <props="china">
	//
	// [Buy page for pay-as-you-go Parallel file systems](https://common-buy.aliyun.com/?commodityCode=nas_cpfs_post#/buy)
	//
	//
	//
	// <props="intl">
	//
	// [Buy page for pay-as-you-go Parallel file systems](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 150
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The capacity of the file system. Unit: GiB.
	//
	// This parameter is required and valid only when FileSystemType is set to extreme, cpfs, or cpfsse.
	//
	// For valid values, refer to the specifications on the buy page:
	//
	// <props="china">
	//
	// - [Buy page for pay-as-you-go Extreme NAS file systems](https://common-buy.aliyun.com/?commodityCode=nas_extreme_post#/buy)
	//
	// - [Buy page for pay-as-you-go Parallel file systems](https://common-buy.aliyun.com/?commodityCode=nas_cpfs_post#/buy)
	//
	//
	//
	// <props="intl">
	//
	// - [Buy page for pay-as-you-go Extreme NAS file systems](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
	//
	// - [Buy page for pay-as-you-go Parallel file systems](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method. Valid values:
	//
	// - PayAsYouGo: pay-as-you-go. This is the default value.
	//
	// - Subscription: subscription.
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	//
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the file system.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - It must start with a letter and cannot start with `http://` or `https://`.
	//
	// - It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// The description of this file system
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the creation request.
	//
	// A dry run checks whether the parameters are valid and whether resources are sufficient, but does not actually create the instance or incur fees.
	//
	// Valid values:
	//
	// - true: performs a dry run without creating the instance. The system checks whether required parameters are specified, the request format is valid, the service limits are met, and NAS resources are sufficient. If the request fails the dry run, an error message is returned. If the request passes the dry run, the 200 HTTP status code is returned, but FileSystemId is empty.
	//
	// - false: performs a normal request and creates the instance directly after the request passes the check. This is the default value.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration. Unit: months. This parameter is required and valid only when ChargeType is set to Subscription.
	//
	// If a subscription instance is not renewed upon expiration, it is automatically released.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to encrypt the file system.
	//
	// Uses KMS-managed keys to encrypt data at rest in the file system. You do not need to perform decryption when reading or writing the encrypted data.
	//
	// Valid values:
	//
	// - 0: Not encrypted. This is the default value.
	//
	// - 1: NAS-managed key. Supported when FileSystemType is set to standard or extreme.
	//
	// - 2: Custom key. Supported when FileSystemType is set to standard or extreme.
	//
	// > - Extreme NAS file systems: The custom key (KMS) feature is supported in all regions except China East 1 Finance Cloud.
	//
	// > - General-purpose NAS file systems: The custom key (KMS) feature is supported in all regions.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The file system type. Valid values:
	//
	// - standard: General-purpose NAS file system. This is the default value.
	//
	// - extreme: Extreme NAS file system.
	//
	// - cpfs: Cloud Parallel File Storage (CPFS) with locally redundant storage (LRS).
	//
	// - cpfsse: Cloud Parallel File Storage SE with zone-redundant storage (ZRS).
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The cluster ID of the high-performance network zone. This parameter is required when FileSystemType is set to bmcpfs and cannot be modified after creation. You can call the DescribeSpecifications operation to query valid values. Example: B1.
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The ID of the KMS key.
	//
	// This parameter is required only when EncryptType is set to 2.
	//
	// example:
	//
	// fcbd****-62**-4a**-b605-c58cc1d5****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The file transfer protocol type.
	//
	// - When FileSystemType is set to standard, valid values are NFS and SMB.
	//
	// - When FileSystemType is set to extreme, the valid value is NFS.
	//
	// - When FileSystemType is set to cpfs, the valid value is cpfs.
	//
	// - When FileSystemType is set to cpfsse, the valid value is cpfs.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage redundancy type. This parameter takes effect only for CPFS SE. Valid value: ZRS.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ZRS
	RedundancyType *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	// The zone-redundant vSwitch IDs.
	//
	// This parameter is required when RedundancyType is set to ZRS. You must specify three vSwitch IDs, and they must belong to three different zones.
	//
	// if can be null:
	// true
	RedundancyVSwitchIds []*string `json:"RedundancyVSwitchIds,omitempty" xml:"RedundancyVSwitchIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// You can view the resource group ID in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?).
	//
	// example:
	//
	// rg-acfmwavnfdf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is supported only for Extreme NAS file systems with the Advanced storage type.
	//
	// > When you create a file system from a snapshot, the created file system has the same version as the source file system of the snapshot. For example, if the source file system of the snapshot is version 1 and you want to create a version 2 file system, you can first create File System A from the snapshot, then create File System B that meets the version 2 configurations, copy the data from File System A to File System B, and migrate your business to File System B after the copy is complete.
	//
	// example:
	//
	// s-extreme-snapsho****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The storage type.
	//
	// - When FileSystemType is set to standard, valid values are Performance, Capacity, Premium, and Agentic.
	//
	// - When FileSystemType is set to extreme, valid values are standard and advance.
	//
	// - When FileSystemType is set to cpfs, valid values are advance_100 (100 MB/s/TiB baseline), advance_200 (200 MB/s/TiB baseline), and economic (Economy).
	//
	// - When FileSystemType is set to cpfsse, the valid value is advance_100 (100 MB/s/TiB baseline).
	//
	// This parameter is required.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags.
	//
	// Array length: 1 to 20. If the array contains multiple tag objects, the tag keys must be unique.
	Tag []*CreateFileSystemRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// - This parameter is required when FileSystemType is set to cpfs.
	//
	// - This parameter is reserved and does not take effect when FileSystemType is not set to cpfs. You do not need to configure it.
	//
	// example:
	//
	// vsw-bp131dkqilvw5pnlt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// - This parameter is required when FileSystemType is set to cpfs or cpfsse.
	//
	// - This parameter is reserved and does not take effect when FileSystemType is set to standard or extreme. You do not need to configure it.
	//
	// example:
	//
	// vpc-bp18cx9a7zoh0h9b4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// A zone is a physical area with independent power grids and networks in the same region.
	//
	// When FileSystemType is set to standard, this parameter is optional. By default, a zone that meets the conditional requirements is randomly selected based on the ProtocolType and StorageType settings.
	//
	// When FileSystemType is set to extreme or cpfs, this parameter is required.
	//
	// > - File systems and Elastic Compute Service (ECS) instances in different zones of the same region can communicate with each other.
	//
	// > - Deploy the file system and the ECS instance in the same zone to avoid cross-zone latency.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateFileSystemRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateFileSystemRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateFileSystemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFileSystemRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFileSystemRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateFileSystemRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *CreateFileSystemRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *CreateFileSystemRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *CreateFileSystemRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateFileSystemRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateFileSystemRequest) GetRedundancyType() *string {
	return s.RedundancyType
}

func (s *CreateFileSystemRequest) GetRedundancyVSwitchIds() []*string {
	return s.RedundancyVSwitchIds
}

func (s *CreateFileSystemRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFileSystemRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateFileSystemRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateFileSystemRequest) GetTag() []*CreateFileSystemRequestTag {
	return s.Tag
}

func (s *CreateFileSystemRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateFileSystemRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateFileSystemRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateFileSystemRequest) SetBandwidth(v int64) *CreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateFileSystemRequest) SetCapacity(v int64) *CreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetChargeType(v string) *CreateFileSystemRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequest) SetClientToken(v string) *CreateFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetDryRun(v bool) *CreateFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFileSystemRequest) SetDuration(v int32) *CreateFileSystemRequest {
	s.Duration = &v
	return s
}

func (s *CreateFileSystemRequest) SetEncryptType(v int32) *CreateFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemType(v string) *CreateFileSystemRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateFileSystemRequest) SetHpnZone(v string) *CreateFileSystemRequest {
	s.HpnZone = &v
	return s
}

func (s *CreateFileSystemRequest) SetKmsKeyId(v string) *CreateFileSystemRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetRedundancyType(v string) *CreateFileSystemRequest {
	s.RedundancyType = &v
	return s
}

func (s *CreateFileSystemRequest) SetRedundancyVSwitchIds(v []*string) *CreateFileSystemRequest {
	s.RedundancyVSwitchIds = v
	return s
}

func (s *CreateFileSystemRequest) SetResourceGroupId(v string) *CreateFileSystemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFileSystemRequest) SetSnapshotId(v string) *CreateFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetTag(v []*CreateFileSystemRequestTag) *CreateFileSystemRequest {
	s.Tag = v
	return s
}

func (s *CreateFileSystemRequest) SetVSwitchId(v string) *CreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVpcId(v string) *CreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateFileSystemRequest) Validate() error {
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

type CreateFileSystemRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// - It cannot be empty or an empty string.
	//
	// - It can be up to 128 characters in length.
	//
	// - It cannot start with `aliyun` or `acs:`.
	//
	// - It cannot contain `http://` or `https://`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// nastest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// - It cannot be empty or an empty string.
	//
	// - It can be up to 128 characters in length.
	//
	// - It cannot contain `http://` or `https://`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFileSystemRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFileSystemRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFileSystemRequestTag) SetKey(v string) *CreateFileSystemRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFileSystemRequestTag) SetValue(v string) *CreateFileSystemRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFileSystemRequestTag) Validate() error {
	return dara.Validate(s)
}
