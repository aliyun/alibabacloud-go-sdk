// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeImageInfosResponseBody
	GetCode() *int32
	SetImages(v *DescribeImageInfosResponseBodyImages) *DescribeImageInfosResponseBody
	GetImages() *DescribeImageInfosResponseBodyImages
	SetRequestId(v string) *DescribeImageInfosResponseBody
	GetRequestId() *string
}

type DescribeImageInfosResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 0
	Code   *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Images *DescribeImageInfosResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5568A08C-10A9-47F3-902F-647298B463FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeImageInfosResponseBody) GetImages() *DescribeImageInfosResponseBodyImages {
	return s.Images
}

func (s *DescribeImageInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageInfosResponseBody) SetCode(v int32) *DescribeImageInfosResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageInfosResponseBody) SetImages(v *DescribeImageInfosResponseBodyImages) *DescribeImageInfosResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImageInfosResponseBody) SetRequestId(v string) *DescribeImageInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageInfosResponseBody) Validate() error {
	if s.Images != nil {
		if err := s.Images.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageInfosResponseBodyImages struct {
	Image []*DescribeImageInfosResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeImageInfosResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImages) GetImage() []*DescribeImageInfosResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeImageInfosResponseBodyImages) SetImage(v []*DescribeImageInfosResponseBodyImagesImage) *DescribeImageInfosResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeImageInfosResponseBodyImages) Validate() error {
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

type DescribeImageInfosResponseBodyImagesImage struct {
	ComputeType        *string                                                      `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	Description        *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskDeviceMappings *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	ImageId            *string                                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageSize          *string                                                      `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	ImageVersion       *string                                                      `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	OSName             *string                                                      `json:"OSName,omitempty" xml:"OSName,omitempty"`
	OSType             *string                                                      `json:"OSType,omitempty" xml:"OSType,omitempty"`
	RegionId           *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeImageInfosResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetComputeType() *string {
	return s.ComputeType
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetDiskDeviceMappings() *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetImageSize() *string {
	return s.ImageSize
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetOSName() *string {
	return s.OSName
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetOSType() *string {
	return s.OSType
}

func (s *DescribeImageInfosResponseBodyImagesImage) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetComputeType(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ComputeType = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetDescription(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.Description = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetDiskDeviceMappings(v *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) *DescribeImageInfosResponseBodyImagesImage {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageId(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageSize(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageSize = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetImageVersion(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetOSName(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.OSName = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetOSType(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.OSType = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) SetRegionId(v string) *DescribeImageInfosResponseBodyImagesImage {
	s.RegionId = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImage) Validate() error {
	if s.DiskDeviceMappings != nil {
		if err := s.DiskDeviceMappings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappings) Validate() error {
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

type DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping struct {
	Format  *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Size    *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetSize() *string {
	return s.Size
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetType() *string {
	return s.Type
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetSize(v string) *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetType(v string) *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Type = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetImageId(v string) *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInfosResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}
