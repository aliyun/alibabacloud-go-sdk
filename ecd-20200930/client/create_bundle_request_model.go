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
	// The cloud computer specifications. You can call [DescribeBundles](https://help.aliyun.com/document_detail/436974.html) to query cloud computer templates and obtain the supported cloud computer specifications from the `DesktopType` parameter in the response.
	//
	// > Non-GPU images can only use non-GPU specifications, and GPU images can only use GPU specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// eds.enterprise_office.2c4g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The operating system language. Currently, only system images are supported. Valid values:
	//
	// - zh-CN: Simplified Chinese.
	//
	// - zh-HK: Traditional Chinese (Hong Kong (China)).
	//
	// - en-US: English.
	//
	// - ja-JP: Japanese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The performance level of the system cloud disk. When the cloud computer specifications are set to graphics or high frequency, you can configure the cloud disk performance level. For more information about the differences between performance levels, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html). Settings: standard SSD and ESSD cloud disks are supported.
	//
	// example:
	//
	// PL1
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The system disk size. Unit: GiB. The supported system disk sizes correspond to the specifications. For more information, see [Overview of cloud computer specifications](https://help.aliyun.com/document_detail/188609.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The performance level of the data cloud disk. When the cloud computer specifications are set to graphics or high frequency, you can configure the cloud disk performance level. For more information about the differences between performance levels, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html). Settings: standard SSD and ESSD cloud disks are supported.
	//
	// example:
	//
	// PL1
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The list of data disk sizes. Currently, only one data disk can be configured.
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
