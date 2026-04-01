// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCImageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*DescribeRCImageListResponseBodyImages) *DescribeRCImageListResponseBody
	GetImages() []*DescribeRCImageListResponseBodyImages
	SetPageNumber(v int32) *DescribeRCImageListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCImageListResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCImageListResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeRCImageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCImageListResponseBody
	GetTotalCount() *int32
}

type DescribeRCImageListResponseBody struct {
	// The information about the images.
	Images []*DescribeRCImageListResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2553A660-E4EB-4AF4-A402-8AFF70A49143
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of images.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCImageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCImageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCImageListResponseBody) GetImages() []*DescribeRCImageListResponseBodyImages {
	return s.Images
}

func (s *DescribeRCImageListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCImageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCImageListResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCImageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCImageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCImageListResponseBody) SetImages(v []*DescribeRCImageListResponseBodyImages) *DescribeRCImageListResponseBody {
	s.Images = v
	return s
}

func (s *DescribeRCImageListResponseBody) SetPageNumber(v int32) *DescribeRCImageListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCImageListResponseBody) SetPageSize(v int32) *DescribeRCImageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCImageListResponseBody) SetRegionId(v string) *DescribeRCImageListResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeRCImageListResponseBody) SetRequestId(v string) *DescribeRCImageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCImageListResponseBody) SetTotalCount(v int32) *DescribeRCImageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCImageListResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCImageListResponseBodyImages struct {
	// The image architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2024-04-25T02:17:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// test
	Description        *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskDeviceMappings []*DescribeRCImageListResponseBodyImagesDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Repeated"`
	// The image ID.
	//
	// example:
	//
	// m-2oqiu973jwcxe****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Created_from_i-2zeh17y17sz677x****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image version.
	//
	// example:
	//
	// 2
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// Indicates whether the image is a public image. Public images include public images provided by Alibaba Cloud and custom images published as community images.
	//
	// 	- **true**: The image is a public image.
	//
	// 	- **false**: The image is not a public image.
	//
	// example:
	//
	// false
	IsPublic           *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	IsSupportRdsCustom *bool `json:"IsSupportRdsCustom,omitempty" xml:"IsSupportRdsCustom,omitempty"`
	// The display name of the operating system in Chinese.
	OSName *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	// The display name of the operating system in English.
	//
	// example:
	//
	// Alibaba Cloud Linux  2.1903 LTS 64 bit Quick Boot
	OSNameEn *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// linux
	OSType   *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The image size. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The image status. Valid values:
	//
	// 	- **Unavailable**
	//
	// 	- **Available**
	//
	// 	- **Creating**
	//
	// 	- **CreateFailed**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the image is used by the RDS Custom instance. Valid values:
	//
	// 	- **instance**: The image is used to create one or more RDS Custom instances.
	//
	// 	- **none**: The image is not used to create RDS Custom instances.
	//
	// example:
	//
	// instance
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeRCImageListResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCImageListResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeRCImageListResponseBodyImages) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeRCImageListResponseBodyImages) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCImageListResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCImageListResponseBodyImages) GetDiskDeviceMappings() []*DescribeRCImageListResponseBodyImagesDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeRCImageListResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCImageListResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeRCImageListResponseBodyImages) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeRCImageListResponseBodyImages) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *DescribeRCImageListResponseBodyImages) GetIsSupportRdsCustom() *bool {
	return s.IsSupportRdsCustom
}

func (s *DescribeRCImageListResponseBodyImages) GetOSName() *string {
	return s.OSName
}

func (s *DescribeRCImageListResponseBodyImages) GetOSNameEn() *string {
	return s.OSNameEn
}

func (s *DescribeRCImageListResponseBodyImages) GetOSType() *string {
	return s.OSType
}

func (s *DescribeRCImageListResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeRCImageListResponseBodyImages) GetSize() *int64 {
	return s.Size
}

func (s *DescribeRCImageListResponseBodyImages) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCImageListResponseBodyImages) GetUsage() *string {
	return s.Usage
}

func (s *DescribeRCImageListResponseBodyImages) SetArchitecture(v string) *DescribeRCImageListResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetCreationTime(v string) *DescribeRCImageListResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetDescription(v string) *DescribeRCImageListResponseBodyImages {
	s.Description = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetDiskDeviceMappings(v []*DescribeRCImageListResponseBodyImagesDiskDeviceMappings) *DescribeRCImageListResponseBodyImages {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetImageId(v string) *DescribeRCImageListResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetImageName(v string) *DescribeRCImageListResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetImageVersion(v string) *DescribeRCImageListResponseBodyImages {
	s.ImageVersion = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetIsPublic(v bool) *DescribeRCImageListResponseBodyImages {
	s.IsPublic = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetIsSupportRdsCustom(v bool) *DescribeRCImageListResponseBodyImages {
	s.IsSupportRdsCustom = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetOSName(v string) *DescribeRCImageListResponseBodyImages {
	s.OSName = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetOSNameEn(v string) *DescribeRCImageListResponseBodyImages {
	s.OSNameEn = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetOSType(v string) *DescribeRCImageListResponseBodyImages {
	s.OSType = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetPlatform(v string) *DescribeRCImageListResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetSize(v int64) *DescribeRCImageListResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetStatus(v string) *DescribeRCImageListResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) SetUsage(v string) *DescribeRCImageListResponseBodyImages {
	s.Usage = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImages) Validate() error {
	if s.DiskDeviceMappings != nil {
		for _, item := range s.DiskDeviceMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCImageListResponseBodyImagesDiskDeviceMappings struct {
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRCImageListResponseBodyImagesDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCImageListResponseBodyImagesDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) GetDevice() *string {
	return s.Device
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) GetSize() *string {
	return s.Size
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) GetType() *string {
	return s.Type
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) SetDevice(v string) *DescribeRCImageListResponseBodyImagesDiskDeviceMappings {
	s.Device = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) SetSize(v string) *DescribeRCImageListResponseBodyImagesDiskDeviceMappings {
	s.Size = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) SetType(v string) *DescribeRCImageListResponseBodyImagesDiskDeviceMappings {
	s.Type = &v
	return s
}

func (s *DescribeRCImageListResponseBodyImagesDiskDeviceMappings) Validate() error {
	return dara.Validate(s)
}
