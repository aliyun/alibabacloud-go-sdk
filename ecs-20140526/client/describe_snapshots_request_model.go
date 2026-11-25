// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeSnapshotsRequestFilter) *DescribeSnapshotsRequest
	GetFilter() []*DescribeSnapshotsRequestFilter
	SetCategory(v string) *DescribeSnapshotsRequest
	GetCategory() *string
	SetDiskId(v string) *DescribeSnapshotsRequest
	GetDiskId() *string
	SetDryRun(v bool) *DescribeSnapshotsRequest
	GetDryRun() *bool
	SetEncrypted(v bool) *DescribeSnapshotsRequest
	GetEncrypted() *bool
	SetInstanceId(v string) *DescribeSnapshotsRequest
	GetInstanceId() *string
	SetKMSKeyId(v string) *DescribeSnapshotsRequest
	GetKMSKeyId() *string
	SetMaxResults(v int32) *DescribeSnapshotsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnapshotsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSnapshotsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnapshotsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSnapshotsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotsRequest
	GetResourceOwnerId() *int64
	SetSnapshotIds(v string) *DescribeSnapshotsRequest
	GetSnapshotIds() *string
	SetSnapshotLinkId(v string) *DescribeSnapshotsRequest
	GetSnapshotLinkId() *string
	SetSnapshotName(v string) *DescribeSnapshotsRequest
	GetSnapshotName() *string
	SetSnapshotType(v string) *DescribeSnapshotsRequest
	GetSnapshotType() *string
	SetSourceDiskType(v string) *DescribeSnapshotsRequest
	GetSourceDiskType() *string
	SetStatus(v string) *DescribeSnapshotsRequest
	GetStatus() *string
	SetTag(v []*DescribeSnapshotsRequestTag) *DescribeSnapshotsRequest
	GetTag() []*DescribeSnapshotsRequestTag
	SetUsage(v string) *DescribeSnapshotsRequest
	GetUsage() *string
}

type DescribeSnapshotsRequest struct {
	Filter []*DescribeSnapshotsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The snapshot category. Valid values:
	//
	// - Standard: Standard snapshot.
	//
	// - Flash: Local snapshot. This value will soon be deprecated. Local snapshots are replaced by the snapshot instant access feature. The parameter description is as follows:
	//
	//   - If you used local snapshots before December 14, 2020, you can use this parameter normally.
	//
	//   - If you did not use local snapshots before December 14, 2020, you cannot use this parameter.
	//
	// - archive: Archive snapshot.
	//
	//
	//
	// <props="china">For more information, see [Notice on Alibaba Cloud Snapshot Service Upgrade and New Billing Items on December 14](https://help.aliyun.com/noticelist/articleid/1060755542.html).
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to only perform a dry run for this request.
	//
	// - true: The system sends a check request without querying resources. The check items include whether the AccessKey pair is valid, the authorization of Resource Access Management (RAM) users, and whether required parameters are specified. If the check fails, the system returns an error. If the check succeeds, the system returns the error code DryRunOperation.
	//
	// - false (default): The system sends a normal request. If the request passes the check, the system returns an HTTP status code 2xx and queries resources.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to filter encrypted snapshots. Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The specified instance ID. This parameter specifies the instance whose disk snapshots you want to query.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key that corresponds to the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. You can obtain the token from the response of the previous request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter will be deprecated soon. We recommend that you use the NextToken and MaxResults parameters to perform a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter will be deprecated soon. We recommend that you use the NextToken and MaxResults parameters to perform a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you use this parameter to filter resources, the resource count cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["s-bp67acfmxazb4p****", "s-bp67acfmxazb5p****", … "s-bp67acfmxazb6p****"]
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The ID of the snapshot chain.
	//
	// example:
	//
	// sl-bp1grgphbcc9brb5****
	SnapshotLinkId *string `json:"SnapshotLinkId,omitempty" xml:"SnapshotLinkId,omitempty"`
	// The snapshot name.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The automatic creation type. Valid values:
	//
	//
	//
	// - auto: The snapshot is automatically created.
	//
	// - user: The snapshot is manually created.
	//
	// - all (default): All automatic creation types.
	//
	// example:
	//
	// all
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The type of the source disk from which the snapshot is created. Valid values:
	//
	//
	//
	// - system: System disk.
	//
	// - data: Data disk.
	//
	// > The values are not case-sensitive.
	//
	// example:
	//
	// system
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The snapshot status. Valid values:
	//
	//
	//
	// - progressing: The snapshot is being created.
	//
	// - accomplished: The snapshot is created.
	//
	// - failed: The snapshot fails to be created.
	//
	// - all (default): All snapshot statuses.
	//
	// example:
	//
	// all
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tag []*DescribeSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The usage of the snapshot. Valid values:
	//
	// - image: The snapshot is used to create a custom image.
	//
	// - disk: The snapshot is used to create a disk.
	//
	// - image_disk: The snapshot is used to create a data disk and a custom image.
	//
	// - none: The snapshot is not used.
	//
	// example:
	//
	// none
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetFilter() []*DescribeSnapshotsRequestFilter {
	return s.Filter
}

func (s *DescribeSnapshotsRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeSnapshotsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeSnapshotsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeSnapshotsRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeSnapshotsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotsRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSnapshotsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotsRequest) GetSnapshotIds() *string {
	return s.SnapshotIds
}

func (s *DescribeSnapshotsRequest) GetSnapshotLinkId() *string {
	return s.SnapshotLinkId
}

func (s *DescribeSnapshotsRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsRequest) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsRequest) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsRequest) GetTag() []*DescribeSnapshotsRequestTag {
	return s.Tag
}

func (s *DescribeSnapshotsRequest) GetUsage() *string {
	return s.Usage
}

func (s *DescribeSnapshotsRequest) SetFilter(v []*DescribeSnapshotsRequestFilter) *DescribeSnapshotsRequest {
	s.Filter = v
	return s
}

func (s *DescribeSnapshotsRequest) SetCategory(v string) *DescribeSnapshotsRequest {
	s.Category = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDiskId(v string) *DescribeSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDryRun(v bool) *DescribeSnapshotsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetEncrypted(v bool) *DescribeSnapshotsRequest {
	s.Encrypted = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetInstanceId(v string) *DescribeSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetKMSKeyId(v string) *DescribeSnapshotsRequest {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetOwnerAccount(v string) *DescribeSnapshotsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetOwnerId(v int64) *DescribeSnapshotsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetResourceGroupId(v string) *DescribeSnapshotsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetResourceOwnerId(v int64) *DescribeSnapshotsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotIds(v string) *DescribeSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotLinkId(v string) *DescribeSnapshotsRequest {
	s.SnapshotLinkId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotName(v string) *DescribeSnapshotsRequest {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotType(v string) *DescribeSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSourceDiskType(v string) *DescribeSnapshotsRequest {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStatus(v string) *DescribeSnapshotsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetTag(v []*DescribeSnapshotsRequestTag) *DescribeSnapshotsRequest {
	s.Tag = v
	return s
}

func (s *DescribeSnapshotsRequest) SetUsage(v string) *DescribeSnapshotsRequest {
	s.Usage = &v
	return s
}

func (s *DescribeSnapshotsRequest) Validate() error {
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

type DescribeSnapshotsRequestFilter struct {
	// The filter key used when querying resources. The value must be `CreationStartTime`. If you specify both `Filter.1.Key` and `Filter.1.Value`, you can query the resources created after the specified point in time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used when querying resources. If you specify this parameter, you must also specify the `Filter.1.Key` parameter. The value must be in the `yyyy-MM-ddTHH:mmZ` format and use the UTC+0 time zone.
	//
	// example:
	//
	// 2019-12-13T17:00Z
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSnapshotsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeSnapshotsRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeSnapshotsRequestFilter) SetKey(v string) *DescribeSnapshotsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeSnapshotsRequestFilter) SetValue(v string) *DescribeSnapshotsRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeSnapshotsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotsRequestTag struct {
	// The tag key of the snapshot. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the resource count under this tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that are attached with multiple tags cannot exceed 1,000. If the resource count exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSnapshotsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSnapshotsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSnapshotsRequestTag) SetKey(v string) *DescribeSnapshotsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSnapshotsRequestTag) SetValue(v string) *DescribeSnapshotsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSnapshotsRequestTag) Validate() error {
	return dara.Validate(s)
}
