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
	SetPerformanceLevel(v string) *DescribeDisksRequest
	GetPerformanceLevel() *string
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
	// The list of additional attributes. Currently, only `Placement` is supported, which indicates the data storage location of the disk.
	//
	// > Currently, only regional disks have a data storage location.
	//
	// example:
	//
	// IOPS
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-m5e2w2jutw8bv31****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The disk category. Valid values:
	//
	// - all: all cloud disks, local disks, and elastic ephemeral disks.
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
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// - elastic_ephemeral_disk_standard: standard elastic ephemeral disk.
	//
	// - elastic_ephemeral_disk_premium: premium elastic ephemeral disk.
	//
	// - local_ssd_pro: I/O-intensive local disk.
	//
	// - local_hdd_pro: throughput-intensive local disk.
	//
	// - ephemeral: (retired) local disk.
	//
	// - ephemeral_ssd: (retired) local SSD.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether automatic snapshots are released when the cloud disk is released. Valid values:
	//
	// - true: Automatic snapshots are released.
	//
	// - false: Automatic snapshots are not released.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Specifies whether the disk is released together with the instance. Valid values:
	//
	// - true: The disk is released together with the instance.
	//
	// - false: The disk is retained and not released. It is converted to a pay-as-you-go data disk.
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
	// The IDs of the cloud disks, local disks, or elastic ephemeral disks. The value is a JSON array that contains up to 100 IDs separated by commas (,).
	//
	// example:
	//
	// ["d-bp67acfmxazb4p****", "d-bp67acfmxazb4g****", … "d-bp67acfmxazb4d****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The name of the disk. The name must be 2 to 128 characters in length and can contain Unicode letters, including English letters, Chinese characters, and digits. It can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The type of the cloud disk, local disk, or elastic ephemeral disk to query. Valid values:
	//
	// - all: queries both system disks and data disks.
	//
	// - system: queries only system disks.
	//
	// - data: queries only data disks.
	//
	// Default value: all.
	//
	// > Elastic ephemeral disks cannot be used as system disks.
	//
	// example:
	//
	// all
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// Specifies whether to only perform a dry run for this request. Valid values:
	//
	// - true: send a check request and do not query resources. The check items include whether the AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - false: send a Normal request. If the request passes the check, a 2xx HTTP status code is returned and resources are queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the automatic snapshot policy feature is enabled for the cloud disk. Valid values:
	//
	// - true: The feature is enabled.
	//
	// - false: The feature is not enabled.
	//
	// > This parameter is deprecated. The automatic snapshot policy feature is enabled by default for newly created cloud disks. You only need to apply an automatic snapshot policy to the cloud disk to use the feature.
	//
	// example:
	//
	// true
	EnableAutoSnapshot *bool `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	// Specifies whether an automatic snapshot policy is applied to the cloud disk. Valid values:
	//
	// - true: An automatic snapshot policy is applied.
	//
	// - false: No automatic snapshot policy is applied.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableAutomatedSnapshotPolicy *bool `json:"EnableAutomatedSnapshotPolicy,omitempty" xml:"EnableAutomatedSnapshotPolicy,omitempty"`
	// Specifies whether the block storage device is Shared Block Storage.
	//
	// example:
	//
	// false
	EnableShared *bool `json:"EnableShared,omitempty" xml:"EnableShared,omitempty"`
	// Specifies whether to query only encrypted cloud disks. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The instance ID to which the disk, local disk, or elastic ephemeral disk is attached.
	//
	// example:
	//
	// i-bp67acfmxazb4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Key Management Service (KMS) key used by the cloud disk.
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
	// The maximum number of entries to return. Valid values: 10 to 500.
	//
	// Default value:
	//
	// - If this parameter is not specified or the specified value is less than 10, the default value is 10.
	//
	// - If the specified value is greater than 500, the default value is 500.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether the multi-attach feature is enabled. Valid values:
	//
	// - Disabled: The feature is disabled.
	//
	// - Enabled: The feature is enabled.
	//
	// - LegacyShared: Used to query shared block storage devices.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	// The query token. The value is the `NextToken` value returned from the previous API call.
	//
	// For more information about how to configure the returned data for this operation, see the operation description section above.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is about to go offline. We recommend that you use the `NextToken` and `MaxResults` parameters to perform paging and complete the paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to go offline. We recommend that you use the `NextToken` and `MaxResults` parameters to perform paging and complete the paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The performance level (PL) of the ESSD disk. Valid values: PL0, PL1, PL2, and PL3.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Specifies whether the disk can be detached. Valid values:
	//
	// - true: The disk can be detached. It can exist independently and can be freely attached and detached within a zone.
	//
	// - false: The disk cannot be detached. It cannot exist independently and cannot be freely attached or detached within a zone.
	//
	// The `Portable` attribute is `false` for the following types of block storage devices, and their lifecycles are the same as those of the instances:
	//
	// - Local disks.
	//
	// - Local SSDs.
	//
	// - Subscription data disks.
	//
	// example:
	//
	// false
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The region ID of the block storage device. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the disk belongs. When you use this parameter to filter resources, the resource count cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot used to create the cloud disk.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The disk status. For more information, see [Disk status](https://help.aliyun.com/document_detail/25689.html). Valid values:
	//
	// - In_use: The disk is in use.
	//
	// - Available: The disk is available for attachment.
	//
	// - Attaching: The disk is being attached.
	//
	// - Detaching: The disk is being detached.
	//
	// - Creating: The disk is being created.
	//
	// - ReIniting: The disk is being initialized.
	//
	// - All: all statuses.
	//
	// Default value: All.
	//
	// example:
	//
	// All
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the disk.
	Tag []*DescribeDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID.
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

func (s *DescribeDisksRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
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

func (s *DescribeDisksRequest) SetPerformanceLevel(v string) *DescribeDisksRequest {
	s.PerformanceLevel = &v
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
	// The filter key used to query resources. Set the value to `CreationStartTime`. If you specify both `Filter.1.Key` and `Filter.1.Value`, you can query resources created after the specified time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used to query resources. If you specify this parameter, you must also specify the `Filter.1.Key` parameter. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
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
	// The tag key of the disk. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the resource count returned cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that attach multiple tags cannot exceed 1,000. If the resource count exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the disk. Valid values of N: 1 to 20.
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
