// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListAvailableImagesResponseBodyImages) *ListAvailableImagesResponseBody
	GetImages() []*ListAvailableImagesResponseBodyImages
	SetPageNumber(v string) *ListAvailableImagesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListAvailableImagesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListAvailableImagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvailableImagesResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *ListAvailableImagesResponseBody
	GetTotalCount() *string
}

type ListAvailableImagesResponseBody struct {
	// The information about the images.
	Images []*ListAvailableImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvailableImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponseBody) GetImages() []*ListAvailableImagesResponseBodyImages {
	return s.Images
}

func (s *ListAvailableImagesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAvailableImagesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAvailableImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableImagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvailableImagesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListAvailableImagesResponseBody) SetImages(v []*ListAvailableImagesResponseBodyImages) *ListAvailableImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListAvailableImagesResponseBody) SetPageNumber(v string) *ListAvailableImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetPageSize(v string) *ListAvailableImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetRequestId(v string) *ListAvailableImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetSuccess(v bool) *ListAvailableImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetTotalCount(v string) *ListAvailableImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAvailableImagesResponseBody) Validate() error {
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

type ListAvailableImagesResponseBodyImages struct {
	// The OS architecture of the image. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// 	- BIOS: Basic Input/Output System (BIOS).
	//
	// 	- UEFI: Unified Extensible Firmware Interface (UEFI).
	//
	// >  When you change the OS boot mode of an instance, you must make sure that the boot mode matches the boot mode of the associated image. Otherwise, the instance fails to be booted.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The image description.
	//
	// example:
	//
	// ExampleDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// centos_7_06_64_20G_alibase_2019071****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CHESS5V5.0.27
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: system image.
	//
	// 	- self: custom image.
	//
	// 	- others: shared image.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The OS name in Chinese.
	OSName *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	// The OS name in English.
	//
	// example:
	//
	// CentOS  7.9 64 bit
	OSNameEn *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	// The OS. Valid values:
	//
	// 	- CentOS
	//
	// 	- windows
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The image size. Unit: GiB
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListAvailableImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponseBodyImages) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListAvailableImagesResponseBodyImages) GetBootMode() *string {
	return s.BootMode
}

func (s *ListAvailableImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListAvailableImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListAvailableImagesResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *ListAvailableImagesResponseBodyImages) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *ListAvailableImagesResponseBodyImages) GetOSName() *string {
	return s.OSName
}

func (s *ListAvailableImagesResponseBodyImages) GetOSNameEn() *string {
	return s.OSNameEn
}

func (s *ListAvailableImagesResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *ListAvailableImagesResponseBodyImages) GetSize() *string {
	return s.Size
}

func (s *ListAvailableImagesResponseBodyImages) SetArchitecture(v string) *ListAvailableImagesResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetBootMode(v string) *ListAvailableImagesResponseBodyImages {
	s.BootMode = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetDescription(v string) *ListAvailableImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageId(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageName(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageOwnerAlias(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetOSName(v string) *ListAvailableImagesResponseBodyImages {
	s.OSName = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetOSNameEn(v string) *ListAvailableImagesResponseBodyImages {
	s.OSNameEn = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetPlatform(v string) *ListAvailableImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetSize(v string) *ListAvailableImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
