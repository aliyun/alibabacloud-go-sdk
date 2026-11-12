// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ImportImageRequest
	GetBizRegionId() *string
	SetBizType(v int32) *ImportImageRequest
	GetBizType() *int32
	SetBootMode(v string) *ImportImageRequest
	GetBootMode() *string
	SetDescription(v string) *ImportImageRequest
	GetDescription() *string
	SetEnableSecurityCheck(v bool) *ImportImageRequest
	GetEnableSecurityCheck() *bool
	SetGpuCategory(v bool) *ImportImageRequest
	GetGpuCategory() *bool
	SetGpuDriverType(v string) *ImportImageRequest
	GetGpuDriverType() *string
	SetImageName(v string) *ImportImageRequest
	GetImageName() *string
	SetLicenseType(v string) *ImportImageRequest
	GetLicenseType() *string
	SetOsType(v string) *ImportImageRequest
	GetOsType() *string
	SetOssObjectPath(v string) *ImportImageRequest
	GetOssObjectPath() *string
	SetProductType(v string) *ImportImageRequest
	GetProductType() *string
	SetProtocolType(v string) *ImportImageRequest
	GetProtocolType() *string
	SetSystemDiskSize(v int32) *ImportImageRequest
	GetSystemDiskSize() *int32
}

type ImportImageRequest struct {
	// The business region ID. Example: cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The business line to which the image belongs. Default value: 1.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// - BIOS: legacy BIOS boot.
	//
	// - UEFI: UEFI boot.
	//
	// example:
	//
	// UEFI
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// Custom cloud desktop image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the image security check. Valid values:
	//
	// - true: enables the security check.
	//
	// - false: disables the security check.
	EnableSecurityCheck *bool `json:"EnableSecurityCheck,omitempty" xml:"EnableSecurityCheck,omitempty"`
	// Specifies whether the image is a GPU image. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	GpuCategory *bool `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The GPU driver type. Valid values:
	//
	// - gpu_grid9: GRID 9 driver.
	//
	// - gpu_grid12: GRID 12 driver.
	//
	// - gpu_custom: custom driver.
	//
	// example:
	//
	// gpu_grid12
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	// The name of the image after it is imported. This parameter cannot be empty.
	//
	// example:
	//
	// DemoImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The operating system license type. Valid values:
	//
	// - Auto: automatically selected.
	//
	// - Aliyun: provided by Alibaba Cloud.
	//
	// - BYOL: Bring Your Own License.
	//
	// Default value: Auto.
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type of the image file. It must be consistent with the actual operating system of the image, such as Windows or Linux.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The complete OSS object URL of the image file to be imported, including the bucket domain name and file path. The file must be uploaded before you call this operation. Example: https://BucketName.oss-cn-hangzhou.aliyuncs.com/FilePath.
	OssObjectPath *string `json:"OssObjectPath,omitempty" xml:"OssObjectPath,omitempty"`
	// The product type. Currently, only CloudDesktop (cloud desktop) is supported.
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The connection protocol of the image. Valid values: ASP and HDX. Default value: ASP.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s ImportImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequest) GoString() string {
	return s.String()
}

func (s *ImportImageRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ImportImageRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ImportImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *ImportImageRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportImageRequest) GetEnableSecurityCheck() *bool {
	return s.EnableSecurityCheck
}

func (s *ImportImageRequest) GetGpuCategory() *bool {
	return s.GpuCategory
}

func (s *ImportImageRequest) GetGpuDriverType() *string {
	return s.GpuDriverType
}

func (s *ImportImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ImportImageRequest) GetOsType() *string {
	return s.OsType
}

func (s *ImportImageRequest) GetOssObjectPath() *string {
	return s.OssObjectPath
}

func (s *ImportImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ImportImageRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ImportImageRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ImportImageRequest) SetBizRegionId(v string) *ImportImageRequest {
	s.BizRegionId = &v
	return s
}

func (s *ImportImageRequest) SetBizType(v int32) *ImportImageRequest {
	s.BizType = &v
	return s
}

func (s *ImportImageRequest) SetBootMode(v string) *ImportImageRequest {
	s.BootMode = &v
	return s
}

func (s *ImportImageRequest) SetDescription(v string) *ImportImageRequest {
	s.Description = &v
	return s
}

func (s *ImportImageRequest) SetEnableSecurityCheck(v bool) *ImportImageRequest {
	s.EnableSecurityCheck = &v
	return s
}

func (s *ImportImageRequest) SetGpuCategory(v bool) *ImportImageRequest {
	s.GpuCategory = &v
	return s
}

func (s *ImportImageRequest) SetGpuDriverType(v string) *ImportImageRequest {
	s.GpuDriverType = &v
	return s
}

func (s *ImportImageRequest) SetImageName(v string) *ImportImageRequest {
	s.ImageName = &v
	return s
}

func (s *ImportImageRequest) SetLicenseType(v string) *ImportImageRequest {
	s.LicenseType = &v
	return s
}

func (s *ImportImageRequest) SetOsType(v string) *ImportImageRequest {
	s.OsType = &v
	return s
}

func (s *ImportImageRequest) SetOssObjectPath(v string) *ImportImageRequest {
	s.OssObjectPath = &v
	return s
}

func (s *ImportImageRequest) SetProductType(v string) *ImportImageRequest {
	s.ProductType = &v
	return s
}

func (s *ImportImageRequest) SetProtocolType(v string) *ImportImageRequest {
	s.ProtocolType = &v
	return s
}

func (s *ImportImageRequest) SetSystemDiskSize(v int32) *ImportImageRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ImportImageRequest) Validate() error {
	return dara.Validate(s)
}
