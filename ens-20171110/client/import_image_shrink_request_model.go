// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *ImportImageShrinkRequest
	GetArchitecture() *string
	SetComputeType(v string) *ImportImageShrinkRequest
	GetComputeType() *string
	SetDiskDeviceMappingShrink(v string) *ImportImageShrinkRequest
	GetDiskDeviceMappingShrink() *string
	SetImageFormat(v string) *ImportImageShrinkRequest
	GetImageFormat() *string
	SetImageName(v string) *ImportImageShrinkRequest
	GetImageName() *string
	SetLicenseType(v string) *ImportImageShrinkRequest
	GetLicenseType() *string
	SetOSSBucket(v string) *ImportImageShrinkRequest
	GetOSSBucket() *string
	SetOSSObject(v string) *ImportImageShrinkRequest
	GetOSSObject() *string
	SetOSSRegion(v string) *ImportImageShrinkRequest
	GetOSSRegion() *string
	SetOSType(v string) *ImportImageShrinkRequest
	GetOSType() *string
	SetOSVersion(v string) *ImportImageShrinkRequest
	GetOSVersion() *string
	SetPlatform(v string) *ImportImageShrinkRequest
	GetPlatform() *string
	SetTargetOSSRegionId(v string) *ImportImageShrinkRequest
	GetTargetOSSRegionId() *string
}

type ImportImageShrinkRequest struct {
	// The system architecture. Valid values:</br>
	//
	// - x86_64
	//
	// Currently, only x86_64 is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image type. Valid values:
	//
	// ens_vm: ENS virtual machine image (default).
	//
	// This parameter is required.
	//
	// example:
	//
	// ens_vm
	ComputeType *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	// The list of custom image information.
	DiskDeviceMappingShrink *string `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty"`
	// The image format. Valid values:</br>
	//
	// qcow2.</br>
	//
	// Currently, only qcow2 is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcow2
	ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	// The image name. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image license type. Valid values: Aliyun (Alibaba Cloud license) and BYOL (Bring Your Own Key).
	//
	// example:
	//
	// Aliyun
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The OSS bucket that stores the image file.
	//
	// example:
	//
	// tmp-hybrid
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The name of the image file.
	//
	// example:
	//
	// image-test
	OSSObject *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
	// The region where the image resides. Currently, only cn-beijing is supported.
	//
	// example:
	//
	// cn-beijing
	OSSRegion *string `json:"OSSRegion,omitempty" xml:"OSSRegion,omitempty"`
	// The type of the operating system platform. Valid values:
	//
	// - windows
	//
	// - linux
	//
	// Currently, only linux is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The version of the operating system distribution.
	//
	// example:
	//
	// 6.8
	OSVersion *string `json:"OSVersion,omitempty" xml:"OSVersion,omitempty"`
	// The operating system distribution. Valid values:
	//
	// 	- centos
	//
	// 	- ubuntu
	//
	// example:
	//
	// centos
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region of the destination OSS bucket where the image is stored.</br>
	//
	// > Currently, only cn-beijing and ap-southeast-1 are supported.
	//
	// example:
	//
	// cn-beijing
	TargetOSSRegionId *string `json:"TargetOSSRegionId,omitempty" xml:"TargetOSSRegionId,omitempty"`
}

func (s ImportImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportImageShrinkRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ImportImageShrinkRequest) GetComputeType() *string {
	return s.ComputeType
}

func (s *ImportImageShrinkRequest) GetDiskDeviceMappingShrink() *string {
	return s.DiskDeviceMappingShrink
}

func (s *ImportImageShrinkRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *ImportImageShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageShrinkRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ImportImageShrinkRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ImportImageShrinkRequest) GetOSSObject() *string {
	return s.OSSObject
}

func (s *ImportImageShrinkRequest) GetOSSRegion() *string {
	return s.OSSRegion
}

func (s *ImportImageShrinkRequest) GetOSType() *string {
	return s.OSType
}

func (s *ImportImageShrinkRequest) GetOSVersion() *string {
	return s.OSVersion
}

func (s *ImportImageShrinkRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ImportImageShrinkRequest) GetTargetOSSRegionId() *string {
	return s.TargetOSSRegionId
}

func (s *ImportImageShrinkRequest) SetArchitecture(v string) *ImportImageShrinkRequest {
	s.Architecture = &v
	return s
}

func (s *ImportImageShrinkRequest) SetComputeType(v string) *ImportImageShrinkRequest {
	s.ComputeType = &v
	return s
}

func (s *ImportImageShrinkRequest) SetDiskDeviceMappingShrink(v string) *ImportImageShrinkRequest {
	s.DiskDeviceMappingShrink = &v
	return s
}

func (s *ImportImageShrinkRequest) SetImageFormat(v string) *ImportImageShrinkRequest {
	s.ImageFormat = &v
	return s
}

func (s *ImportImageShrinkRequest) SetImageName(v string) *ImportImageShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *ImportImageShrinkRequest) SetLicenseType(v string) *ImportImageShrinkRequest {
	s.LicenseType = &v
	return s
}

func (s *ImportImageShrinkRequest) SetOSSBucket(v string) *ImportImageShrinkRequest {
	s.OSSBucket = &v
	return s
}

func (s *ImportImageShrinkRequest) SetOSSObject(v string) *ImportImageShrinkRequest {
	s.OSSObject = &v
	return s
}

func (s *ImportImageShrinkRequest) SetOSSRegion(v string) *ImportImageShrinkRequest {
	s.OSSRegion = &v
	return s
}

func (s *ImportImageShrinkRequest) SetOSType(v string) *ImportImageShrinkRequest {
	s.OSType = &v
	return s
}

func (s *ImportImageShrinkRequest) SetOSVersion(v string) *ImportImageShrinkRequest {
	s.OSVersion = &v
	return s
}

func (s *ImportImageShrinkRequest) SetPlatform(v string) *ImportImageShrinkRequest {
	s.Platform = &v
	return s
}

func (s *ImportImageShrinkRequest) SetTargetOSSRegionId(v string) *ImportImageShrinkRequest {
	s.TargetOSSRegionId = &v
	return s
}

func (s *ImportImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
