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
	// This parameter is required.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ens_vm
	ComputeType             *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	DiskDeviceMappingShrink *string `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qcow2
	ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	// This parameter is required.
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// example:
	//
	// tmp-hybrid
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// example:
	//
	// image-test
	OSSObject *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
	// example:
	//
	// cn-beijing
	OSSRegion *string `json:"OSSRegion,omitempty" xml:"OSSRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// example:
	//
	// 6.8
	OSVersion *string `json:"OSVersion,omitempty" xml:"OSVersion,omitempty"`
	// example:
	//
	// centos
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
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
