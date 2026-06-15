// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeDisksRequestFilter) *DescribeDisksRequest
	GetFilter() []*DescribeDisksRequestFilter
	SetAdditionalAttributes(v []*string) *DescribeDisksRequest
	GetAdditionalAttributes() []*string
	SetAutoSnapshotPolicyId(v string) *DescribeDisksRequest
	GetAutoSnapshotPolicyId() *string
	SetCategory(v string) *DescribeDisksRequest
	GetCategory() *string
	SetDeleteAutoSnapshot(v bool) *DescribeDisksRequest
	GetDeleteAutoSnapshot() *bool
	SetDeleteWithInstance(v bool) *DescribeDisksRequest
	GetDeleteWithInstance() *bool
	SetDiskChargeType(v string) *DescribeDisksRequest
	GetDiskChargeType() *string
	SetDiskIds(v string) *DescribeDisksRequest
	GetDiskIds() *string
	SetDiskName(v string) *DescribeDisksRequest
	GetDiskName() *string
	SetDiskType(v string) *DescribeDisksRequest
	GetDiskType() *string
	SetDryRun(v bool) *DescribeDisksRequest
	GetDryRun() *bool
	SetEnableAutoSnapshot(v bool) *DescribeDisksRequest
	GetEnableAutoSnapshot() *bool
	SetEnableAutomatedSnapshotPolicy(v bool) *DescribeDisksRequest
	GetEnableAutomatedSnapshotPolicy() *bool
	SetEnableShared(v bool) *DescribeDisksRequest
	GetEnableShared() *bool
	SetEncrypted(v bool) *DescribeDisksRequest
	GetEncrypted() *bool
	SetInstanceId(v string) *DescribeDisksRequest
	GetInstanceId() *string
	SetKMSKeyId(v string) *DescribeDisksRequest
	GetKMSKeyId() *string
	SetLockReason(v string) *DescribeDisksRequest
	GetLockReason() *string
	SetMaxResults(v int32) *DescribeDisksRequest
	GetMaxResults() *int32
	SetMultiAttach(v string) *DescribeDisksRequest
	GetMultiAttach() *string
	SetNextToken(v string) *DescribeDisksRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDisksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDisksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDisksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDisksRequest
	GetPageSize() *int32
	SetPortable(v bool) *DescribeDisksRequest
	GetPortable() *bool
	SetRegionId(v string) *DescribeDisksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDisksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDisksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDisksRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *DescribeDisksRequest
	GetSnapshotId() *string
	SetStatus(v string) *DescribeDisksRequest
	GetStatus() *string
	SetTag(v []*DescribeDisksRequestTag) *DescribeDisksRequest
	GetTag() []*DescribeDisksRequestTag
	SetZoneId(v string) *DescribeDisksRequest
	GetZoneId() *string
}

type DescribeDisksRequest struct {
	Filter []*DescribeDisksRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The additional attributes. Currently, the only valid value is `Placement`, which you can specify to query the data placement location of the disk.
	//
	// > Currently, only regional redundant disks have a data placement location.
	//
	// example:
	//
	// IOPS
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// The ID of the automatic snapshot policy that is applied to the cloud disk.
	//
	// example:
	//
	// sp-m5e2w2jutw8bv31****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The category of the disk. Valid values:
	//
	// - all: all disk categories, including all cloud disks, local disks, and elastic ephemeral disks.
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: SSD disk.
	//
	// - cloud_essd: ESSD
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional redundant ESSD disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// - elastic_ephemeral_disk_standard: Standard elastic ephemeral disk.
	//
	// - elastic_ephemeral_disk_premium: Premium elastic ephemeral disk.
	//
	// - local_ssd_pro: I/O-intensive local disk.
	//
	// - local_hdd_pro: throughput-intensive local disk.
	//
	// - ephemeral: (Discontinued) local disk.
	//
	// - ephemeral_ssd: (Discontinued) local SSD disk.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the automatic snapshots of the disk when the disk is released. Valid values:
	//
	// - true: The automatic snapshots are released.
	//
	// - false: The automatic snapshots are retained.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Specifies whether the disk is released when the instance is released. Valid values:
	//
	// - true: The disk is released with the instance.
	//
	// - false: The disk is retained as a pay-as-you-go data disk and is not released.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The billing method of the disk. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The IDs of the cloud disks, local disks, or elastic ephemeral disks. The value is a JSON array of up to 100 disk IDs.
	//
	// example:
	//
	// ["d-bp67acfmxazb4p****", "d-bp67acfmxazb4g****", … "d-bp67acfmxazb4d****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The name of the disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The type of the disk. Valid values:
	//
	// - all: both system disks and data disks.
	//
	// - system: only system disks.
	//
	// - data: only data disks.
	//
	// Default value: all.
	//
	// > Elastic ephemeral disks cannot be used as system disks.
	//
	// example:
	//
	// all
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// Specifies whether to perform a dry run for the request. Valid values:
	//
	// - true: Performs a dry run to check the request\\"s validity, including AccessKey authentication, permissions, and required parameters. If the request is valid, the `DryRunOperation` error code is returned. Otherwise, an error message is returned.
	//
	// - false: Sends the request. If the request is valid, a 2xx HTTP status code is returned and the resources are queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the automatic snapshot policy feature is enabled for the disk. Valid values:
	//
	// - true: The feature is enabled.
	//
	// - false: The feature is disabled.
	//
	// > This parameter is deprecated. The automatic snapshot policy feature is enabled by default for all disks. You only need to apply an automatic snapshot policy to a disk.
	//
	// example:
	//
	// true
	EnableAutoSnapshot *bool `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	// Specifies whether an automatic snapshot policy is applied to the disk. Valid values:
	//
	// - true: An automatic snapshot policy is applied.
	//
	// - false: No automatic snapshot policy is applied to the disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableAutomatedSnapshotPolicy *bool `json:"EnableAutomatedSnapshotPolicy,omitempty" xml:"EnableAutomatedSnapshotPolicy,omitempty"`
	// Specifies whether the disk is a shared block storage device.
	//
	// example:
	//
	// false
	EnableShared *bool `json:"EnableShared,omitempty" xml:"EnableShared,omitempty"`
	// Specifies whether to query only encrypted disks. Valid values:
	//
	// - true: Only encrypted disks are queried.
	//
	// - false: The query includes both encrypted and unencrypted disks.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the instance to which the disk is attached.
	//
	// example:
	//
	// i-bp67acfmxazb4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key that is used to encrypt the cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The reason why the disk is locked. Valid values:
	//
	// - financial: The disk is locked due to overdue payments.
	//
	// - security: The disk is locked for security reasons.
	//
	// example:
	//
	// security
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 500.
	//
	// Default value:
	//
	// - If this parameter is not specified or is set to a value less than 10, the default value is 10.
	//
	// - If this parameter is set to a value greater than 500, the value 500 is used.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether the multi-attach feature is enabled for the disk. Valid values:
	//
	// - Disabled: The multi-attach feature is disabled.
	//
	// - Enabled: The multi-attach feature is enabled.
	//
	// - LegacyShared: Used to query shared block storage devices.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	// The token used to start the next query. To retrieve the next page of results, set this parameter to the `NextToken` value returned from the previous call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is deprecated. Use the `NextToken` and `MaxResults` parameters for paged queries instead.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is deprecated. Use the `NextToken` and `MaxResults` parameters for paged queries instead.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the disk is detachable. Valid values:
	//
	// - true: The disk is detachable. A detachable disk can exist independently and can be attached to or detached from an instance within the same availability zone.
	//
	// - false: The disk is not detachable. Its lifecycle is bound to the instance to which it is attached.
	//
	// The `Portable` attribute for the following disk types is always `false`, and their lifecycles are bound to the instances to which they are attached:
	//
	// - Local disks
	//
	// - Local SSD disks
	//
	// - Subscription data disks
	//
	// example:
	//
	// false
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The ID of the region where the disk resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs. When you use this parameter to filter resources, the number of resources cannot exceed 1,000.
	//
	// > You cannot filter resources in the default resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot that was used to create the cloud disk.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The state of the disk. For more information, see [Disk states](https://help.aliyun.com/document_detail/25689.html). Valid values:
	//
	// - In_use: The disk is in use.
	//
	// - Available: The disk is ready for attachment.
	//
	// - Attaching: The disk is being attached.
	//
	// - Detaching: The disk is being detached.
	//
	// - Creating: The disk is being created.
	//
	// - ReIniting: The disk is being re-initialized.
	//
	// - All: all disk states.
	//
	// Default value: All.
	//
	// example:
	//
	// All
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the disk.
	Tag []*DescribeDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the availability zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequest) GetFilter() []*DescribeDisksRequestFilter {
	return s.Filter
}

func (s *DescribeDisksRequest) GetAdditionalAttributes() []*string {
	return s.AdditionalAttributes
}

func (s *DescribeDisksRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeDisksRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeDisksRequest) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *DescribeDisksRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeDisksRequest) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeDisksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeDisksRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeDisksRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDisksRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeDisksRequest) GetEnableAutoSnapshot() *bool {
	return s.EnableAutoSnapshot
}

func (s *DescribeDisksRequest) GetEnableAutomatedSnapshotPolicy() *bool {
	return s.EnableAutomatedSnapshotPolicy
}

func (s *DescribeDisksRequest) GetEnableShared() *bool {
	return s.EnableShared
}

func (s *DescribeDisksRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDisksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeDisksRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDisksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDisksRequest) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *DescribeDisksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDisksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDisksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDisksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDisksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisksRequest) GetPortable() *bool {
	return s.Portable
}

func (s *DescribeDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDisksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDisksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDisksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDisksRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeDisksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDisksRequest) GetTag() []*DescribeDisksRequestTag {
	return s.Tag
}

func (s *DescribeDisksRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDisksRequest) SetFilter(v []*DescribeDisksRequestFilter) *DescribeDisksRequest {
	s.Filter = v
	return s
}

func (s *DescribeDisksRequest) SetAdditionalAttributes(v []*string) *DescribeDisksRequest {
	s.AdditionalAttributes = v
	return s
}

func (s *DescribeDisksRequest) SetAutoSnapshotPolicyId(v string) *DescribeDisksRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeDisksRequest) SetCategory(v string) *DescribeDisksRequest {
	s.Category = &v
	return s
}

func (s *DescribeDisksRequest) SetDeleteAutoSnapshot(v bool) *DescribeDisksRequest {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeDisksRequest) SetDeleteWithInstance(v bool) *DescribeDisksRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskChargeType(v string) *DescribeDisksRequest {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskIds(v string) *DescribeDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskName(v string) *DescribeDisksRequest {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskType(v string) *DescribeDisksRequest {
	s.DiskType = &v
	return s
}

func (s *DescribeDisksRequest) SetDryRun(v bool) *DescribeDisksRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeDisksRequest) SetEnableAutoSnapshot(v bool) *DescribeDisksRequest {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeDisksRequest) SetEnableAutomatedSnapshotPolicy(v bool) *DescribeDisksRequest {
	s.EnableAutomatedSnapshotPolicy = &v
	return s
}

func (s *DescribeDisksRequest) SetEnableShared(v bool) *DescribeDisksRequest {
	s.EnableShared = &v
	return s
}

func (s *DescribeDisksRequest) SetEncrypted(v bool) *DescribeDisksRequest {
	s.Encrypted = &v
	return s
}

func (s *DescribeDisksRequest) SetInstanceId(v string) *DescribeDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksRequest) SetKMSKeyId(v string) *DescribeDisksRequest {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDisksRequest) SetLockReason(v string) *DescribeDisksRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeDisksRequest) SetMaxResults(v int32) *DescribeDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisksRequest) SetMultiAttach(v string) *DescribeDisksRequest {
	s.MultiAttach = &v
	return s
}

func (s *DescribeDisksRequest) SetNextToken(v string) *DescribeDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDisksRequest) SetOwnerAccount(v string) *DescribeDisksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDisksRequest) SetOwnerId(v int64) *DescribeDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDisksRequest) SetPageNumber(v int32) *DescribeDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksRequest) SetPageSize(v int32) *DescribeDisksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksRequest) SetPortable(v bool) *DescribeDisksRequest {
	s.Portable = &v
	return s
}

func (s *DescribeDisksRequest) SetRegionId(v string) *DescribeDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisksRequest) SetResourceGroupId(v string) *DescribeDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDisksRequest) SetResourceOwnerAccount(v string) *DescribeDisksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDisksRequest) SetResourceOwnerId(v int64) *DescribeDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDisksRequest) SetSnapshotId(v string) *DescribeDisksRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeDisksRequest) SetStatus(v string) *DescribeDisksRequest {
	s.Status = &v
	return s
}

func (s *DescribeDisksRequest) SetTag(v []*DescribeDisksRequestTag) *DescribeDisksRequest {
	s.Tag = v
	return s
}

func (s *DescribeDisksRequest) SetZoneId(v string) *DescribeDisksRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDisksRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
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

type DescribeDisksRequestFilter struct {
	// The key of the filter to use for the query. Set the value to `CreationStartTime`. You can use `Filter.1.Key` and `Filter.1.Value` to query resources that were created after a specific point in time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter to use for the query. When you specify this parameter, you must also specify the `Filter.1.Key` parameter. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-05T22:40Z
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDisksRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeDisksRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeDisksRequestFilter) SetKey(v string) *DescribeDisksRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeDisksRequestFilter) SetValue(v string) *DescribeDisksRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeDisksRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksRequestTag struct {
	// The tag key of the disk.
	//
	// > For better compatibility, we recommend that you use the `Tag.N.Key` parameter.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDisksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDisksRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDisksRequestTag) SetKey(v string) *DescribeDisksRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDisksRequestTag) SetValue(v string) *DescribeDisksRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDisksRequestTag) Validate() error {
	return dara.Validate(s)
}
