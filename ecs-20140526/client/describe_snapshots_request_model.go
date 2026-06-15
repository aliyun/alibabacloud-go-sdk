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
	// - `Standard`: A standard snapshot.
	//
	// - `Flash`: A local snapshot. This value is deprecated because the local snapshot feature has been replaced by the instant access feature.
	//
	//   - If you have used local snapshots before December 14, 2020, you can continue to use this value.
	//
	//   - If you have not used local snapshots before December 14, 2020, you cannot use this value.
	//
	// - `archive`: An archive snapshot.
	//
	// <props="china">
	//
	// For more information, see [December 14: Alibaba Cloud snapshot service upgrade and new billing items notice](https://help.aliyun.com/noticelist/articleid/1060755542.html).
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the cloud disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// - `true`: Performs a dry run but does not query resources. The system checks the request for potential issues, including missing required parameters, invalid parameter values, and insufficient permissions. If the request is invalid, an error is returned. If the request is valid, the `DryRunOperation` error code is returned.
	//
	// - `false` (Default): Sends a normal request. If the request is valid, the system returns a 2xx HTTP status code and the query results.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to return only encrypted snapshots. Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the instance. When you specify this ID, the operation returns snapshots of cloud disks attached to the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key used to encrypt the snapshot.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to start the next page of results. You can obtain this token from the response to a previous query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is deprecated. We recommend that you use the `NextToken` and `MaxResults` parameters for paged queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is deprecated. We recommend that you use the `NextToken` and `MaxResults` parameters for paged queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the snapshot belongs. When you filter by this parameter, the query can return a maximum of 1,000 snapshots.
	//
	// > You cannot filter resources that are in the default resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A JSON array that contains the IDs of up to 100 snapshots to query.
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
	// The snapshot creation type. Valid values:
	//
	// - `auto`: An automatically created snapshot.
	//
	// - `user`: A manually created snapshot.
	//
	// - `all` (Default): All snapshot creation types.
	//
	// example:
	//
	// all
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The type of the source disk of the snapshot. Valid values:
	//
	// - `system`: The snapshot was created from a system disk.
	//
	// - `data`: The snapshot was created from a data disk.
	//
	// > The value is case-insensitive.
	//
	// example:
	//
	// system
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The status of the snapshot. Valid values:
	//
	// - `progressing`: The snapshot is being created.
	//
	// - `accomplished`: The snapshot is complete.
	//
	// - `failed`: Snapshot creation failed.
	//
	// - `all` (Default): All snapshot statuses.
	//
	// example:
	//
	// all
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags by which to filter snapshots.
	Tag []*DescribeSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The usage of the snapshot. Valid values:
	//
	// - `image`: The snapshot is used to create a custom image.
	//
	// - `disk`: The snapshot is used to create a cloud disk.
	//
	// - `image_disk`: The snapshot is used to create a custom image and a data disk.
	//
	// - `none`: The snapshot is not used.
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
	// The filter key for querying resources. The value must be `CreationStartTime`. If you specify `Filter.1.Key` and `Filter.1.Value`, you can query for resources that were created after the specified point in time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. If you specify this parameter, you must also specify `Filter.1.Key`. The value must be in the `yyyy-MM-ddTHH:mmZ` format and in UTC.
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
	// The tag key.
	//
	// > For better compatibility, use the `Tag.N.Key` parameter.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
