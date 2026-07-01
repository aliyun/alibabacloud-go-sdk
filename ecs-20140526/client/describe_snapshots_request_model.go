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
	// The category of the snapshot. Valid values:
	//
	// - Standard: standard snapshot.
	//
	// - Flash: local snapshot. This value is about to be deprecated. Local snapshots have been replaced by the snapshot instant access feature. The metric description is as follows:
	//
	//   - If you used local snapshots before December 14, 2020, you can use this parameter. The parameter is active.
	//
	//   - If you did not use local snapshots before December 14, 2020, you cannot use this parameter.
	//
	// - archive: archive snapshot.
	//
	//
	//
	// <props="china">For more information, see [Chinese site notice on snapshot service upgrade and new billing items on December 14](https://help.aliyun.com/noticelist/articleid/1060755542.html).
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
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// - true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
	// The instance ID. Specify this parameter to query snapshot information of disks attached to the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Key Management Service (KMS) key used by the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The maximum number of entries per page for paging. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Obtained from the response of the previous request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults for paging instead.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults for paging instead.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// > Filtering by default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of snapshots. The value is a JSON array that consists of up to 100 snapshot IDs. Separate the IDs with commas (,).
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
	// The name of the snapshot.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The type of automatic creation. Valid values:
	//
	//
	//
	// - auto: automatic snapshot.
	//
	// - user: manual snapshot.
	//
	// - all (default): All automatic creation types.
	//
	// example:
	//
	// all
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The type of the source disk. Valid values:
	//
	//
	//
	// - system: system disk.
	//
	// - data: data disk.
	//
	// > The value is case-insensitive.
	//
	// example:
	//
	// system
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The status of the snapshot. Valid values:
	//
	//
	//
	// - progressing: The snapshot is being created.
	//
	// - accomplished: The snapshot is created.
	//
	// - failed: The snapshot failed to be created.
	//
	// - all (default): All snapshot statuses.
	//
	// example:
	//
	// all
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether the snapshot has been used to create images or disks. Valid values:
	//
	// - image: The snapshot has been used to create custom images.
	//
	// - disk: The snapshot has been used to create disks.
	//
	// - image_disk: The snapshot has been used to create both data disks and custom images.
	//
	// - none: The snapshot has not been used.
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
	// The filter key used to query resources. Set the value to `CreationStartTime`. When you specify both `Filter.1.Key` and `Filter.1.Value`, you can query resources created after the specified point in time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used to query resources. When you specify this parameter, you must also specify `Filter.1.Key`. Specify the time in the `yyyy-MM-ddTHH:mmZ` format in UTC.
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
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1000. If you use multiple tags to filter resources, the resource count with all specified tags attached cannot exceed 1000. If the resource count exceeds 1000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
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
