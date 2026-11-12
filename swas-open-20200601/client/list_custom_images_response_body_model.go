// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomImages(v []*ListCustomImagesResponseBodyCustomImages) *ListCustomImagesResponseBody
	GetCustomImages() []*ListCustomImagesResponseBodyCustomImages
	SetPageNumber(v string) *ListCustomImagesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListCustomImagesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListCustomImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListCustomImagesResponseBody
	GetTotalCount() *string
}

type ListCustomImagesResponseBody struct {
	// The custom images.
	CustomImages []*ListCustomImagesResponseBodyCustomImages `json:"CustomImages,omitempty" xml:"CustomImages,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in Settings for the paged query.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) GetCustomImages() []*ListCustomImagesResponseBodyCustomImages {
	return s.CustomImages
}

func (s *ListCustomImagesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListCustomImagesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListCustomImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomImagesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListCustomImagesResponseBody) SetCustomImages(v []*ListCustomImagesResponseBodyCustomImages) *ListCustomImagesResponseBody {
	s.CustomImages = v
	return s
}

func (s *ListCustomImagesResponseBody) SetPageNumber(v string) *ListCustomImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetPageSize(v string) *ListCustomImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetTotalCount(v string) *ListCustomImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomImagesResponseBody) Validate() error {
	if s.CustomImages != nil {
		for _, item := range s.CustomImages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomImagesResponseBodyCustomImages struct {
	// The information about instances created from the image.
	CreateInstances []*string `json:"CreateInstances,omitempty" xml:"CreateInstances,omitempty" type:"Repeated"`
	// The time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-09T02:28:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data cloud disk snapshot ID.
	//
	// example:
	//
	// s-bp19rn9u8eqzlfb***
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The name of the data cloud disk snapshot.
	//
	// example:
	//
	// DataDiskSnapshot
	DataSnapshotName *string `json:"DataSnapshotName,omitempty" xml:"DataSnapshotName,omitempty"`
	// The description.
	//
	// example:
	//
	// CustomImage-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The custom image ID.
	//
	// example:
	//
	// m-bp1e79zktg26n2b***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Indicates whether the custom image is shared to ECS.
	//
	// example:
	//
	// false
	InShare *bool `json:"InShare,omitempty" xml:"InShare,omitempty"`
	// Indicates whether the custom image is shared across accounts.
	//
	// example:
	//
	// False
	InShareUser *bool `json:"InShareUser,omitempty" xml:"InShareUser,omitempty"`
	// The instance ID of the simple application server from which the image was created.
	//
	// example:
	//
	// 2d06ee0520b44de1ae88d4be****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server instance.
	//
	// example:
	//
	// swas-asdf23***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the custom image.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system type of the image.
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data cloud disk size of the image. Unit: GiB.
	//
	// example:
	//
	// 20
	RequiredDataDiskSize *int32 `json:"RequiredDataDiskSize,omitempty" xml:"RequiredDataDiskSize,omitempty"`
	// The system cloud disk size of the image. Unit: GiB.
	//
	// example:
	//
	// 20
	RequiredSystemDiskSize *int32 `json:"RequiredSystemDiskSize,omitempty" xml:"RequiredSystemDiskSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm2h2lvp3ublq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the source image.
	//
	// example:
	//
	// CentOS
	SourceImageName *string `json:"SourceImageName,omitempty" xml:"SourceImageName,omitempty"`
	// The image version of the source image.
	//
	// example:
	//
	// 8.0
	SourceImageVersion *string `json:"SourceImageVersion,omitempty" xml:"SourceImageVersion,omitempty"`
	// The status of the custom image. Valid values:
	//
	// - 0: copying.
	//
	// - 1: active.
	//
	// - 2: unavailable.
	//
	// - 3: creation failed.
	//
	// - 4: creating.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The system cloud disk snapshot ID.
	//
	// example:
	//
	// s-bp1h173hj21puxb***
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The name of the system cloud disk snapshot.
	//
	// example:
	//
	// SystemDiskSnapshot
	SystemSnapshotName *string `json:"SystemSnapshotName,omitempty" xml:"SystemSnapshotName,omitempty"`
	// The tags of the custom images.
	Tags []*ListCustomImagesResponseBodyCustomImagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the image owner.
	//
	// example:
	//
	// 180185828710****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCustomImagesResponseBodyCustomImages) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesResponseBodyCustomImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyCustomImages) GetCreateInstances() []*string {
	return s.CreateInstances
}

func (s *ListCustomImagesResponseBodyCustomImages) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListCustomImagesResponseBodyCustomImages) GetDataSnapshotId() *string {
	return s.DataSnapshotId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetDataSnapshotName() *string {
	return s.DataSnapshotName
}

func (s *ListCustomImagesResponseBodyCustomImages) GetDescription() *string {
	return s.Description
}

func (s *ListCustomImagesResponseBodyCustomImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetInShare() *bool {
	return s.InShare
}

func (s *ListCustomImagesResponseBodyCustomImages) GetInShareUser() *bool {
	return s.InShareUser
}

func (s *ListCustomImagesResponseBodyCustomImages) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCustomImagesResponseBodyCustomImages) GetName() *string {
	return s.Name
}

func (s *ListCustomImagesResponseBodyCustomImages) GetOsType() *string {
	return s.OsType
}

func (s *ListCustomImagesResponseBodyCustomImages) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetRequiredDataDiskSize() *int32 {
	return s.RequiredDataDiskSize
}

func (s *ListCustomImagesResponseBodyCustomImages) GetRequiredSystemDiskSize() *int32 {
	return s.RequiredSystemDiskSize
}

func (s *ListCustomImagesResponseBodyCustomImages) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetSourceImageName() *string {
	return s.SourceImageName
}

func (s *ListCustomImagesResponseBodyCustomImages) GetSourceImageVersion() *string {
	return s.SourceImageVersion
}

func (s *ListCustomImagesResponseBodyCustomImages) GetStatus() *string {
	return s.Status
}

func (s *ListCustomImagesResponseBodyCustomImages) GetSystemSnapshotId() *string {
	return s.SystemSnapshotId
}

func (s *ListCustomImagesResponseBodyCustomImages) GetSystemSnapshotName() *string {
	return s.SystemSnapshotName
}

func (s *ListCustomImagesResponseBodyCustomImages) GetTags() []*ListCustomImagesResponseBodyCustomImagesTags {
	return s.Tags
}

func (s *ListCustomImagesResponseBodyCustomImages) GetUserId() *int64 {
	return s.UserId
}

func (s *ListCustomImagesResponseBodyCustomImages) SetCreateInstances(v []*string) *ListCustomImagesResponseBodyCustomImages {
	s.CreateInstances = v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetCreationTime(v string) *ListCustomImagesResponseBodyCustomImages {
	s.CreationTime = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDataSnapshotId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.DataSnapshotId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDataSnapshotName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.DataSnapshotName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDescription(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetImageId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInShare(v bool) *ListCustomImagesResponseBodyCustomImages {
	s.InShare = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInShareUser(v bool) *ListCustomImagesResponseBodyCustomImages {
	s.InShareUser = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInstanceId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.InstanceId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInstanceName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.InstanceName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Name = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetOsType(v string) *ListCustomImagesResponseBodyCustomImages {
	s.OsType = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetRegionId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.RegionId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetRequiredDataDiskSize(v int32) *ListCustomImagesResponseBodyCustomImages {
	s.RequiredDataDiskSize = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetRequiredSystemDiskSize(v int32) *ListCustomImagesResponseBodyCustomImages {
	s.RequiredSystemDiskSize = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetResourceGroupId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSourceImageName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SourceImageName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSourceImageVersion(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SourceImageVersion = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetStatus(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Status = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSystemSnapshotId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SystemSnapshotId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSystemSnapshotName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SystemSnapshotName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetTags(v []*ListCustomImagesResponseBodyCustomImagesTags) *ListCustomImagesResponseBodyCustomImages {
	s.Tags = v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetUserId(v int64) *ListCustomImagesResponseBodyCustomImages {
	s.UserId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomImagesResponseBodyCustomImagesTags struct {
	// The tag key of the custom image.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom image.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCustomImagesResponseBodyCustomImagesTags) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesResponseBodyCustomImagesTags) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) GetKey() *string {
	return s.Key
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) GetValue() *string {
	return s.Value
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) SetKey(v string) *ListCustomImagesResponseBodyCustomImagesTags {
	s.Key = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) SetValue(v string) *ListCustomImagesResponseBodyCustomImagesTags {
	s.Value = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) Validate() error {
	return dara.Validate(s)
}
