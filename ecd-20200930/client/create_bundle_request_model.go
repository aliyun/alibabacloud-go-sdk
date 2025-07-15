// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBundleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleName(v string) *CreateBundleRequest
	GetBundleName() *string
	SetDescription(v string) *CreateBundleRequest
	GetDescription() *string
	SetDesktopType(v string) *CreateBundleRequest
	GetDesktopType() *string
	SetImageId(v string) *CreateBundleRequest
	GetImageId() *string
	SetLanguage(v string) *CreateBundleRequest
	GetLanguage() *string
	SetRegionId(v string) *CreateBundleRequest
	GetRegionId() *string
	SetRootDiskPerformanceLevel(v string) *CreateBundleRequest
	GetRootDiskPerformanceLevel() *string
	SetRootDiskSizeGib(v int32) *CreateBundleRequest
	GetRootDiskSizeGib() *int32
	SetUserDiskPerformanceLevel(v string) *CreateBundleRequest
	GetUserDiskPerformanceLevel() *string
	SetUserDiskSizeGib(v []*int32) *CreateBundleRequest
	GetUserDiskSizeGib() []*int32
}

type CreateBundleRequest struct {
	// The name of the cloud computer template.
	//
	// example:
	//
	// testBundleName
	BundleName *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	// The description of the cloud computer template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance type of the cloud computers. You can call the [DescribeBundles](https://help.aliyun.com/document_detail/436974.html) operation to query cloud computer templates and obtain the instance types supported by the cloud computers from the `DesktopType` response parameter.
	//
	// >  If you want the template to use a non-GPU-accelerated image, you can only select a non-GPU-accelerated instance type. If you want the template to use a GPU-accelerated image, you can only select a GPU-accelerated instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd.basic.large
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The OS language. This parameter is available only for system images. Valid values:
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- zh-HK: Traditional Chinese (Hong Kong)
	//
	// 	- en-US: American English
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The performance level (PL) of the system disk. When the cloud computer instance type that is specified by the DesktopType parameter is set to a graphical instance type or instance type with a high clock speed, you can set the performance level of the disks. For more information about the differences among disks at different PLs, see [Enhanced SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// Valid values:
	//
	// 	- PL1
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL0
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL3
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL2
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PL1
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. The value of this parameter must be consistent with the system disk size supported by the cloud computer instance type. For more information, see [Overview](https://help.aliyun.com/document_detail/188609.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The PL of the data disk. When the cloud computer instance type that is specified by the DesktopType parameter is set to a graphical instance type or instance type with a high clock speed, you can set the performance level of the disks. For more information about the differences among disks at different PLs, see [Enhanced SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// Valid values:
	//
	// 	- PL1
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL0
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL3
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL2
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PL1
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The data disk sizes. You can configure only one data disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70
	UserDiskSizeGib []*int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty" type:"Repeated"`
}

func (s CreateBundleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateBundleRequest) GetBundleName() *string {
	return s.BundleName
}

func (s *CreateBundleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBundleRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *CreateBundleRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateBundleRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateBundleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBundleRequest) GetRootDiskPerformanceLevel() *string {
	return s.RootDiskPerformanceLevel
}

func (s *CreateBundleRequest) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *CreateBundleRequest) GetUserDiskPerformanceLevel() *string {
	return s.UserDiskPerformanceLevel
}

func (s *CreateBundleRequest) GetUserDiskSizeGib() []*int32 {
	return s.UserDiskSizeGib
}

func (s *CreateBundleRequest) SetBundleName(v string) *CreateBundleRequest {
	s.BundleName = &v
	return s
}

func (s *CreateBundleRequest) SetDescription(v string) *CreateBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateBundleRequest) SetDesktopType(v string) *CreateBundleRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateBundleRequest) SetImageId(v string) *CreateBundleRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBundleRequest) SetLanguage(v string) *CreateBundleRequest {
	s.Language = &v
	return s
}

func (s *CreateBundleRequest) SetRegionId(v string) *CreateBundleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskPerformanceLevel(v string) *CreateBundleRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskSizeGib(v int32) *CreateBundleRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskPerformanceLevel(v string) *CreateBundleRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskSizeGib(v []*int32) *CreateBundleRequest {
	s.UserDiskSizeGib = v
	return s
}

func (s *CreateBundleRequest) Validate() error {
	return dara.Validate(s)
}
