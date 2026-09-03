// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootMode(v string) *UploadImageRequest
	GetBootMode() *string
	SetDataDiskSize(v int32) *UploadImageRequest
	GetDataDiskSize() *int32
	SetDescription(v string) *UploadImageRequest
	GetDescription() *string
	SetEnableSecurityCheck(v bool) *UploadImageRequest
	GetEnableSecurityCheck() *bool
	SetGpuCategory(v bool) *UploadImageRequest
	GetGpuCategory() *bool
	SetGpuDriverType(v string) *UploadImageRequest
	GetGpuDriverType() *string
	SetImageName(v string) *UploadImageRequest
	GetImageName() *string
	SetLicenseType(v string) *UploadImageRequest
	GetLicenseType() *string
	SetOsType(v string) *UploadImageRequest
	GetOsType() *string
	SetOssObjectPath(v string) *UploadImageRequest
	GetOssObjectPath() *string
	SetProtocolType(v string) *UploadImageRequest
	GetProtocolType() *string
	SetRegionId(v string) *UploadImageRequest
	GetRegionId() *string
	SetSystemDiskSize(v string) *UploadImageRequest
	GetSystemDiskSize() *string
}

type UploadImageRequest struct {
	// The boot mode of the image.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The data cloud disk size. Valid values: 80 to 500. Unit: GiB.
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The description of the image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable security check.
	//
	// example:
	//
	// true
	EnableSecurityCheck *bool `json:"EnableSecurityCheck,omitempty" xml:"EnableSecurityCheck,omitempty"`
	// Specifies whether the image is a GPU-type image.
	//
	// example:
	//
	// true
	GpuCategory *bool `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The type of the pre-installed GPU driver.
	//
	// example:
	//
	// gpu_grid9
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	// The image name. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// Win10_Test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The license type used to activate the operating system after the image is imported. Valid values:
	//
	// - Auto: Alibaba Cloud detects the source operating system and assigns a license. In automatic mode, the system first checks whether an Alibaba Cloud official license is available for the `Platform` you specified and assigns it to the imported image. If no such license is available, the system switches to BYOL (Bring Your Own License) mode.
	//
	// - Aliyun: Uses an Alibaba Cloud official license based on the `Platform` you specified.
	//
	// - BYOL: Uses the license that comes with the source operating system. When using BYOL, ensure that your license key supports use on Alibaba Cloud.
	//
	// Default value: Auto
	//
	// > Systems such as Windows 10 cannot be activated through Alibaba Cloud. Set `LicenseType` to custom activation (BYOL).
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The OSS object path of the image file.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ossbucket:endpoint/object
	OssObjectPath *string `json:"OssObjectPath,omitempty" xml:"OssObjectPath,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The system cloud disk size. Unit: GiB.
	//
	// > The system cloud disk size cannot be smaller than the image file.
	//
	// example:
	//
	// 80
	SystemDiskSize *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s UploadImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadImageRequest) GoString() string {
	return s.String()
}

func (s *UploadImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *UploadImageRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *UploadImageRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadImageRequest) GetEnableSecurityCheck() *bool {
	return s.EnableSecurityCheck
}

func (s *UploadImageRequest) GetGpuCategory() *bool {
	return s.GpuCategory
}

func (s *UploadImageRequest) GetGpuDriverType() *string {
	return s.GpuDriverType
}

func (s *UploadImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *UploadImageRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *UploadImageRequest) GetOsType() *string {
	return s.OsType
}

func (s *UploadImageRequest) GetOssObjectPath() *string {
	return s.OssObjectPath
}

func (s *UploadImageRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *UploadImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadImageRequest) GetSystemDiskSize() *string {
	return s.SystemDiskSize
}

func (s *UploadImageRequest) SetBootMode(v string) *UploadImageRequest {
	s.BootMode = &v
	return s
}

func (s *UploadImageRequest) SetDataDiskSize(v int32) *UploadImageRequest {
	s.DataDiskSize = &v
	return s
}

func (s *UploadImageRequest) SetDescription(v string) *UploadImageRequest {
	s.Description = &v
	return s
}

func (s *UploadImageRequest) SetEnableSecurityCheck(v bool) *UploadImageRequest {
	s.EnableSecurityCheck = &v
	return s
}

func (s *UploadImageRequest) SetGpuCategory(v bool) *UploadImageRequest {
	s.GpuCategory = &v
	return s
}

func (s *UploadImageRequest) SetGpuDriverType(v string) *UploadImageRequest {
	s.GpuDriverType = &v
	return s
}

func (s *UploadImageRequest) SetImageName(v string) *UploadImageRequest {
	s.ImageName = &v
	return s
}

func (s *UploadImageRequest) SetLicenseType(v string) *UploadImageRequest {
	s.LicenseType = &v
	return s
}

func (s *UploadImageRequest) SetOsType(v string) *UploadImageRequest {
	s.OsType = &v
	return s
}

func (s *UploadImageRequest) SetOssObjectPath(v string) *UploadImageRequest {
	s.OssObjectPath = &v
	return s
}

func (s *UploadImageRequest) SetProtocolType(v string) *UploadImageRequest {
	s.ProtocolType = &v
	return s
}

func (s *UploadImageRequest) SetRegionId(v string) *UploadImageRequest {
	s.RegionId = &v
	return s
}

func (s *UploadImageRequest) SetSystemDiskSize(v string) *UploadImageRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *UploadImageRequest) Validate() error {
	return dara.Validate(s)
}
