// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *DescribeImagesRequest
	GetActionType() *string
	SetArchitecture(v string) *DescribeImagesRequest
	GetArchitecture() *string
	SetDryRun(v bool) *DescribeImagesRequest
	GetDryRun() *bool
	SetFilter(v []*DescribeImagesRequestFilter) *DescribeImagesRequest
	GetFilter() []*DescribeImagesRequestFilter
	SetImageFamily(v string) *DescribeImagesRequest
	GetImageFamily() *string
	SetImageId(v string) *DescribeImagesRequest
	GetImageId() *string
	SetImageName(v string) *DescribeImagesRequest
	GetImageName() *string
	SetImageOwnerAlias(v string) *DescribeImagesRequest
	GetImageOwnerAlias() *string
	SetImageOwnerId(v int64) *DescribeImagesRequest
	GetImageOwnerId() *int64
	SetInstanceType(v string) *DescribeImagesRequest
	GetInstanceType() *string
	SetIsPublic(v bool) *DescribeImagesRequest
	GetIsPublic() *bool
	SetIsSupportCloudinit(v bool) *DescribeImagesRequest
	GetIsSupportCloudinit() *bool
	SetIsSupportIoOptimized(v bool) *DescribeImagesRequest
	GetIsSupportIoOptimized() *bool
	SetOSType(v string) *DescribeImagesRequest
	GetOSType() *string
	SetOwnerAccount(v string) *DescribeImagesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImagesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeImagesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeImagesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeImagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImagesRequest
	GetResourceOwnerId() *int64
	SetShowExpired(v bool) *DescribeImagesRequest
	GetShowExpired() *bool
	SetSnapshotId(v string) *DescribeImagesRequest
	GetSnapshotId() *string
	SetStatus(v string) *DescribeImagesRequest
	GetStatus() *string
	SetTag(v []*DescribeImagesRequestTag) *DescribeImagesRequest
	GetTag() []*DescribeImagesRequestTag
	SetUsage(v string) *DescribeImagesRequest
	GetUsage() *string
}

type DescribeImagesRequest struct {
	// The scenario in which the image will be used. Valid values:
	//
	// - CreateEcs (default): Create an instance.
	//
	// - ChangeOS: Replace the system disk or change the operating system.
	//
	// example:
	//
	// CreateEcs
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The architecture of the image. Valid values:
	//
	// - i386
	//
	// - x86_64
	//
	// - arm64
	//
	// example:
	//
	// i386
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Specifies whether to perform a dry run of the request.
	//
	// - true: Sends a dry run request without querying resource status. The system checks whether the AccessKey is valid, whether the Resource Access Management (RAM) user is authorized, and whether all required parameters are specified. If the check fails, an error is returned. If the check passes, the error code DryRunOperation is returned.
	//
	// - false: Sends a normal request. After the check passes, an HTTP 2XX status code is returned and the resource status is queried directly.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A list of filter conditions for querying resources.
	Filter []*DescribeImagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The name of the image family. When querying images, you can use this parameter to filter images belonging to the specified image family.
	//
	// Default value: empty.
	//
	// > For information about image families associated with official Alibaba Cloud images, see [Overview of public images](https://help.aliyun.com/document_detail/108393.html).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image ID.
	//
	// <details>
	//
	// <summary>
	//
	// Naming convention for image IDs
	//
	// </summary>
	//
	// - Public images: Named based on the operating system version, architecture, language, and published date. For example, the image ID for Windows Server 2008 R2 Enterprise Edition, 64-bit English system is `win2008r2_64_ent_sp1_en-us_40G_alibase_20190318.vhd`.
	//
	// - Custom images, shared images, Alibaba Cloud Marketplace images, and community images: Start with the letter `m`.
	//
	// </details>
	//
	// example:
	//
	// m-bp1g7004ksh0oeuc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Fuzzy search is supported.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The source of the image. Valid values:
	//
	// - system: Images provided by Alibaba Cloud that are not published on Alibaba Cloud Marketplace. This differs from the "public image" concept in the console.
	//
	// - self: Your custom images.
	//
	// - others: Includes shared images (images directly shared with you by other Alibaba Cloud users) and community images (custom images fully publicly shared by any Alibaba Cloud user). Note the following:
	//
	//   - To find community images, IsPublic must be true.
	//
	//   - To find shared images, IsPublic must be set to false or omitted.
	//
	// - marketplace: Images published on Alibaba Cloud Marketplace by Alibaba Cloud or third-party ISVs, which must be purchased together with ECS instances. Please review the pricing details of Alibaba Cloud Marketplace images yourself.
	//
	// Default value: empty.
	//
	// > An empty value returns results with ImageOwnerAlias values of system, self, and others.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The Alibaba Cloud account ID to which the image belongs. This parameter takes effect only when you query shared images and community images.
	//
	// example:
	//
	// 20169351435666****
	ImageOwnerId *int64 `json:"ImageOwnerId,omitempty" xml:"ImageOwnerId,omitempty"`
	// Queries images that can be used with the specified instance type.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to query published community images. Valid values:
	//
	// - true: Queries published community images. When this parameter is set to true, ImageOwnerAlias must be set to others.
	//
	// - false: Queries other image types excluding community images, depending on the value of the ImageOwnerAlias parameter.
	//
	// Default Value: false.
	//
	// example:
	//
	// false
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// Indicates whether the image supports cloud-init.
	//
	// example:
	//
	// true
	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitempty" xml:"IsSupportCloudinit,omitempty"`
	// Indicates whether the image can run on I/O optimized instances.
	//
	// example:
	//
	// true
	IsSupportIoOptimized *bool `json:"IsSupportIoOptimized,omitempty" xml:"IsSupportIoOptimized,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// - windows
	//
	// - linux
	//
	// example:
	//
	// linux
	OSType       *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the image resource list.
	//
	// Starting value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID to which the image belongs. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the custom image belongs. When using this parameter to filter resources, the number of resources cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Indicates whether subscription-based images have exceeded their usage period.
	//
	// example:
	//
	// false
	ShowExpired *bool `json:"ShowExpired,omitempty" xml:"ShowExpired,omitempty"`
	// The custom image created from a specific snapshot ID.
	//
	// example:
	//
	// s-bp17ot2q7x72ggtw****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// Queries images in the specified status. If this parameter is not configured, only images in the Available status are returned by default. Valid values:
	//
	// - Creating: The image is being created.
	//
	// - Waiting: The image is queued for multitasking.
	//
	// - Available (default): The image is available for your use.
	//
	// - UnAvailable: The image is unavailable for your use.
	//
	// - CreateFailed: The image creation failed.
	//
	// - Deprecated: The image has been deprecated.
	//
	// Default value: Available. This parameter supports multiple values separated by commas (,).
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tag []*DescribeImagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Indicates whether the image is already running on an ECS instance. Valid values:
	//
	// - instance: The image is in use by one or more ECS instances.
	//
	// - none: The image is idle and not used by any ECS instance.
	//
	// example:
	//
	// instance
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeImagesRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeImagesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeImagesRequest) GetFilter() []*DescribeImagesRequestFilter {
	return s.Filter
}

func (s *DescribeImagesRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImagesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagesRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeImagesRequest) GetImageOwnerId() *int64 {
	return s.ImageOwnerId
}

func (s *DescribeImagesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeImagesRequest) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *DescribeImagesRequest) GetIsSupportCloudinit() *bool {
	return s.IsSupportCloudinit
}

func (s *DescribeImagesRequest) GetIsSupportIoOptimized() *bool {
	return s.IsSupportIoOptimized
}

func (s *DescribeImagesRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeImagesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImagesRequest) GetShowExpired() *bool {
	return s.ShowExpired
}

func (s *DescribeImagesRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeImagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagesRequest) GetTag() []*DescribeImagesRequestTag {
	return s.Tag
}

func (s *DescribeImagesRequest) GetUsage() *string {
	return s.Usage
}

func (s *DescribeImagesRequest) SetActionType(v string) *DescribeImagesRequest {
	s.ActionType = &v
	return s
}

func (s *DescribeImagesRequest) SetArchitecture(v string) *DescribeImagesRequest {
	s.Architecture = &v
	return s
}

func (s *DescribeImagesRequest) SetDryRun(v bool) *DescribeImagesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeImagesRequest) SetFilter(v []*DescribeImagesRequestFilter) *DescribeImagesRequest {
	s.Filter = v
	return s
}

func (s *DescribeImagesRequest) SetImageFamily(v string) *DescribeImagesRequest {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v string) *DescribeImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesRequest) SetImageName(v string) *DescribeImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesRequest) SetImageOwnerAlias(v string) *DescribeImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImagesRequest) SetImageOwnerId(v int64) *DescribeImagesRequest {
	s.ImageOwnerId = &v
	return s
}

func (s *DescribeImagesRequest) SetInstanceType(v string) *DescribeImagesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeImagesRequest) SetIsPublic(v bool) *DescribeImagesRequest {
	s.IsPublic = &v
	return s
}

func (s *DescribeImagesRequest) SetIsSupportCloudinit(v bool) *DescribeImagesRequest {
	s.IsSupportCloudinit = &v
	return s
}

func (s *DescribeImagesRequest) SetIsSupportIoOptimized(v bool) *DescribeImagesRequest {
	s.IsSupportIoOptimized = &v
	return s
}

func (s *DescribeImagesRequest) SetOSType(v string) *DescribeImagesRequest {
	s.OSType = &v
	return s
}

func (s *DescribeImagesRequest) SetOwnerAccount(v string) *DescribeImagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImagesRequest) SetOwnerId(v int64) *DescribeImagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImagesRequest) SetPageNumber(v int32) *DescribeImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesRequest) SetPageSize(v int32) *DescribeImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetResourceGroupId(v string) *DescribeImagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImagesRequest) SetResourceOwnerAccount(v string) *DescribeImagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImagesRequest) SetResourceOwnerId(v int64) *DescribeImagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImagesRequest) SetShowExpired(v bool) *DescribeImagesRequest {
	s.ShowExpired = &v
	return s
}

func (s *DescribeImagesRequest) SetSnapshotId(v string) *DescribeImagesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImagesRequest) SetStatus(v string) *DescribeImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeImagesRequest) SetTag(v []*DescribeImagesRequestTag) *DescribeImagesRequest {
	s.Tag = v
	return s
}

func (s *DescribeImagesRequest) SetUsage(v string) *DescribeImagesRequest {
	s.Usage = &v
	return s
}

func (s *DescribeImagesRequest) Validate() error {
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

type DescribeImagesRequestFilter struct {
	// The filter key used when querying resources. Valid values:
	//
	// - If this parameter is set to `CreationStartTime`, you can query resources created after the specified time point (`Filter.N.Value`).
	//
	// - If this parameter is set to `CreationEndTime`, you can query resources created before the specified time point (`Filter.N.Value`).
	//
	// - If this parameter is set to `NetworkType`, you can query resources of the specified network type.
	//
	// - If this parameter is set to any of `CpuOnlineUpgrade`, `CpuOnlineDowngrade`, `MemoryOnlineUpgrade`, or `MemoryOnlineDowngrade`, you can query the hot-swapping support status of CPU or memory for the specified image.
	//
	// Default Value: null.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used when querying resources.
	//
	// - When (`Filter.N.Key`) is `CreationStartTime` or `CreationEndTime`, the format is `yyyy-MM-ddTHH:mmZ` in UTC+0 time zone.
	//
	// - When (`Filter.N.Key`) is `NetworkType`, valid values include `vpc`, `classic`, etc.
	//
	// - When (`Filter.N.Key`) is `CpuOnlineUpgrade`, `CpuOnlineDowngrade`, `MemoryOnlineUpgrade`, or `MemoryOnlineDowngrade`, the value can be `supported` or `unsupported`.
	//
	// Default Value: null.
	//
	// example:
	//
	// 2017-12-05T22:40Z
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeImagesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeImagesRequestFilter) SetKey(v string) *DescribeImagesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeImagesRequestFilter) SetValue(v string) *DescribeImagesRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeImagesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeImagesRequestTag struct {
	// The tag key of the image. Valid values of N: 1 to 20.
	//
	// When you use one tag to filter resources, the number of resources retrieved under this tag cannot exceed 1,000. When you use multiple tags to filter resources, the number of resources that are attached with all specified tags cannot exceed 1,000. If the resource count exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) API to query the resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the image. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeImagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeImagesRequestTag) SetKey(v string) *DescribeImagesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImagesRequestTag) SetValue(v string) *DescribeImagesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeImagesRequestTag) Validate() error {
	return dara.Validate(s)
}
