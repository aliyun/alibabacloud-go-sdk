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
	Code   *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Architecture       *string                                                  `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	CreationTime       *string                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DiskDeviceMappings *DescribeImagesResponseBodyImagesImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	ImageId            *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName          *string                                                  `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias    *string                                                  `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageSize          *string                                                  `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	Platform           *string                                                  `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegionId           *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId         *string                                                  `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
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
	Format  *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Size    *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
