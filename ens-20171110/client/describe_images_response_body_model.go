// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeImagesResponseBody
	GetCode() *int32
	SetImages(v *DescribeImagesResponseBodyImages) *DescribeImagesResponseBody
	GetImages() *DescribeImagesResponseBodyImages
	SetPageNumber(v int32) *DescribeImagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImagesResponseBody
	GetTotalCount() *int32
}

type DescribeImagesResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the images.
	Images *DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8331AA01-C16D-5481-BB47-D19CEBAA811E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of images.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeImagesResponseBody) GetImages() *DescribeImagesResponseBodyImages {
	return s.Images
}

func (s *DescribeImagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImagesResponseBody) SetCode(v int32) *DescribeImagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImagesResponseBody) SetImages(v *DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImagesResponseBody) SetPageNumber(v int32) *DescribeImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponseBody) SetPageSize(v int32) *DescribeImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetTotalCount(v int32) *DescribeImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagesResponseBody) Validate() error {
	if s.Images != nil {
		if err := s.Images.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagesResponseBodyImages struct {
	Image []*DescribeImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) GetImage() []*DescribeImagesResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeImagesResponseBodyImages) SetImage(v []*DescribeImagesResponseBodyImagesImage) *DescribeImagesResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeImagesResponseBodyImages) Validate() error {
	if s.Image != nil {
		for _, item := range s.Image {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImage struct {
	// The architecture of the image. Example: **x86_64**.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the image was created. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2017-12-08T12:10:03Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The mappings between the disk and the snapshot in the image.
	DiskDeviceMappings *DescribeImagesResponseBodyImagesImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_alibase_2017****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// Ubuntu_16.04
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The source of the image. Valid values:
	//
	// 	- system: Alibaba Cloud public images
	//
	// 	- self: your custom images
	//
	// 	- others: shared images from other Alibaba Cloud accounts, or community images published by other Alibaba Cloud accounts
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The size of the image. Unit: GiB.
	//
	// example:
	//
	// 40
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// example:
	//
	// centos
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// mock-clone_snapshot_id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImage) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeImagesResponseBodyImagesImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImagesResponseBodyImagesImage) GetDiskDeviceMappings() *DescribeImagesResponseBodyImagesImageDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageSize() *string {
	return s.ImageSize
}

func (s *DescribeImagesResponseBodyImagesImage) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImagesResponseBodyImagesImage) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagesResponseBodyImagesImage) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeImagesResponseBodyImagesImage) SetArchitecture(v string) *DescribeImagesResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetCreationTime(v string) *DescribeImagesResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetDiskDeviceMappings(v *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) *DescribeImagesResponseBodyImagesImage {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageId(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageName(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageSize(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetPlatform(v string) *DescribeImagesResponseBodyImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetRegionId(v string) *DescribeImagesResponseBodyImagesImage {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetSnapshotId(v string) *DescribeImagesResponseBodyImagesImage {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) Validate() error {
	if s.DiskDeviceMappings != nil {
		if err := s.DiskDeviceMappings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImageDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) *DescribeImagesResponseBodyImagesImageDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) Validate() error {
	if s.DiskDeviceMapping != nil {
		for _, item := range s.DiskDeviceMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping struct {
	// The format of the image.
	//
	// example:
	//
	// The format of the image.
	//
	// raw
	//
	// qcow2
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system: system disk.
	//
	// 	- data: data disk.
	//
	// example:
	//
	// Data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of image.
	//
	// example:
	//
	// i-test
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetSize() *string {
	return s.Size
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetType() *string {
	return s.Type
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetSize(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetType(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Type = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetImageId(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}
