// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelfImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeSelfImagesResponseBody
	GetCode() *int32
	SetImages(v *DescribeSelfImagesResponseBodyImages) *DescribeSelfImagesResponseBody
	GetImages() *DescribeSelfImagesResponseBodyImages
	SetPageNumber(v string) *DescribeSelfImagesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSelfImagesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeSelfImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSelfImagesResponseBody
	GetTotalCount() *string
}

type DescribeSelfImagesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image information.
	Images *DescribeSelfImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8B8EB73-B4FD-4262-8EF6-680DF39C9BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSelfImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeSelfImagesResponseBody) GetImages() *DescribeSelfImagesResponseBodyImages {
	return s.Images
}

func (s *DescribeSelfImagesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSelfImagesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSelfImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSelfImagesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSelfImagesResponseBody) SetCode(v int32) *DescribeSelfImagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSelfImagesResponseBody) SetImages(v *DescribeSelfImagesResponseBodyImages) *DescribeSelfImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeSelfImagesResponseBody) SetPageNumber(v string) *DescribeSelfImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSelfImagesResponseBody) SetPageSize(v string) *DescribeSelfImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSelfImagesResponseBody) SetRequestId(v string) *DescribeSelfImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSelfImagesResponseBody) SetTotalCount(v string) *DescribeSelfImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSelfImagesResponseBody) Validate() error {
	if s.Images != nil {
		if err := s.Images.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSelfImagesResponseBodyImages struct {
	Image []*DescribeSelfImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeSelfImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponseBodyImages) GetImage() []*DescribeSelfImagesResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeSelfImagesResponseBodyImages) SetImage(v []*DescribeSelfImagesResponseBodyImagesImage) *DescribeSelfImagesResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeSelfImagesResponseBodyImages) Validate() error {
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

type DescribeSelfImagesResponseBodyImagesImage struct {
	// The image architecture. Valid values:
	//
	// 	- i386
	//
	// 	- x86_64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Computing type. ens_vm/ens: x86 computing. bare_metal: x86 bare machine or x86 bare metal. arm_vm: ARM computing. arm_bare_metal: ARM bare machine or ARM bare metal. pcfarm: heterogeneous computing.
	//
	// example:
	//
	// ens_vm
	ComputeType *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	// The image creation time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-08T12:10:03Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The mappings between the disk and the snapshot in the image.
	DiskDeviceMappings *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	// The ID of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The source of the image. Valid values:
	//
	// 	- **others**: a custom image that is shared by other Alibaba Cloud accounts.
	//
	// 	- **self**: your own custom image.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The size of the image. Unit: GiB.
	//
	// example:
	//
	// 20
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The size of the image storage.
	//
	// example:
	//
	// 400
	ImageStorageSize *string `json:"ImageStorageSize,omitempty" xml:"ImageStorageSize,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5iqczfxps7csjrxeca****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operating system version.
	//
	// example:
	//
	// 7.2
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// The platform.
	//
	// 	- centos
	//
	// 	- ubuntu
	//
	// 	- alios
	//
	// 	- debian
	//
	// 	- rhel
	//
	// 	- windows
	//
	// example:
	//
	// centos
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// sp-5yt3bdedxzdz6t6uuw****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The state of the image.
	//
	// 	- Creating.
	//
	// 	- Packing.
	//
	// 	- Uploading.
	//
	// 	- Pack_failed.
	//
	// 	- Upload_failed.
	//
	// 	- Available: Only images in the Available state can be used and operated.
	//
	// 	- Unavailable.
	//
	// 	- Copying.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSelfImagesResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetComputeType() *string {
	return s.ComputeType
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetDiskDeviceMappings() *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetImageSize() *string {
	return s.ImageSize
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetImageStorageSize() *string {
	return s.ImageStorageSize
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSelfImagesResponseBodyImagesImage) GetStatus() *string {
	return s.Status
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetArchitecture(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetComputeType(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ComputeType = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetCreationTime(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetDiskDeviceMappings(v *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) *DescribeSelfImagesResponseBodyImagesImage {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetImageId(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetImageName(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetImageSize(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetImageStorageSize(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.ImageStorageSize = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetInstanceId(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.InstanceId = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetOsVersion(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.OsVersion = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetPlatform(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetRegionId(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.RegionId = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetSnapshotId(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) SetStatus(v string) *DescribeSelfImagesResponseBodyImagesImage {
	s.Status = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImage) Validate() error {
	if s.DiskDeviceMappings != nil {
		if err := s.DiskDeviceMappings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappings) Validate() error {
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

type DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping struct {
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

func (s DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetSize() *string {
	return s.Size
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetType() *string {
	return s.Type
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetSize(v string) *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetType(v string) *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Type = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetImageId(v string) *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImageId = &v
	return s
}

func (s *DescribeSelfImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}
