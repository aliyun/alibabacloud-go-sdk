// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadImageRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The size of the data disk. Valid values: 80 to 500. Unit: GiB.
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
	// Specifies whether the image is a GPU-accelerated image.
	//
	// example:
	//
	// true
	GpuCategory *bool `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The type of the pre-installed GPU driver.
	//
	// Valid values:
	//
	// 	- gpu_grid9: This GPU driver is used on cloud computers of the following two specifications: graphics – 4 vCPUs, 23 GiB memory, 4 GiB GPU memory, and graphics – 10 vCPUs, 46 GiB memory, 8 GiB GPU memory.
	//
	// 	- gpu_custom: You can install the driver later.
	//
	// 	- gpu_grid12: This GPU driver is used on graphical cloud computers of specifications other than the following two specifications: graphics – 4 vCPUs, 23 GiB memory, & 4 GiB GPU memory, and graphics – 10 vCPUs, 46 GiB memory, & 8 GiB GPU memory.
	//
	// example:
	//
	// gpu_grid9
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	// The name of the image. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// Win10_Test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the license that is used to activate the operating system after the image is imported. Valid values:
	//
	// 	- Auto: Elastic Desktop Service detects the operating system of the image and allocates a license to the operating system. In this mode, the system first checks whether a license allocated by an official Alibaba Cloud channel is specified in the `Platform`. If a license allocated by an official Alibaba Cloud channel is specified, the system allocates the license to the imported image. If no such license is specified, the BYOL (Bring Your Own License) mode is used.
	//
	// 	- Aliyun: The license that is allocated by an official Alibaba Cloud channel and is specified by `Platform` is used for the operating system distribution.
	//
	// 	- BYOL: The license that comes with the source operating system is used. When you use the BYOL mode, make sure that your license key is supported by Alibaba Cloud.
	//
	// Default value: Auto.
	//
	// >  Windows 10 cannot be activated by Alibaba Cloud. Set the `LicenseType` to BYOL for Windows 10.
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The type of the operating system.
	//
	// Valid values:
	//
	// 	- Linux
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The object path of the image file in Object Storage Service (OSS).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ossbucket:endpoint/object
	OssObjectPath *string `json:"OssObjectPath,omitempty" xml:"OssObjectPath,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- ASP: in-house Adaptive Streaming Protocol (ASP)
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// >  The system disk must be at least as large as the image.
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
