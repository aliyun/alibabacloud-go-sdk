// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *ImportImageRequest
	GetArchitecture() *string
	SetComputeType(v string) *ImportImageRequest
	GetComputeType() *string
	SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest
	GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping
	SetImageFormat(v string) *ImportImageRequest
	GetImageFormat() *string
	SetImageName(v string) *ImportImageRequest
	GetImageName() *string
	SetLicenseType(v string) *ImportImageRequest
	GetLicenseType() *string
	SetOSSBucket(v string) *ImportImageRequest
	GetOSSBucket() *string
	SetOSSObject(v string) *ImportImageRequest
	GetOSSObject() *string
	SetOSSRegion(v string) *ImportImageRequest
	GetOSSRegion() *string
	SetOSType(v string) *ImportImageRequest
	GetOSType() *string
	SetOSVersion(v string) *ImportImageRequest
	GetOSVersion() *string
	SetPlatform(v string) *ImportImageRequest
	GetPlatform() *string
	SetTargetOSSRegionId(v string) *ImportImageRequest
	GetTargetOSSRegionId() *string
}

type ImportImageRequest struct {
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
	ComputeType       *string                                `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	DiskDeviceMapping []*ImportImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
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

func (s ImportImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequest) GoString() string {
	return s.String()
}

func (s *ImportImageRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ImportImageRequest) GetComputeType() *string {
	return s.ComputeType
}

func (s *ImportImageRequest) GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *ImportImageRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *ImportImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ImportImageRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ImportImageRequest) GetOSSObject() *string {
	return s.OSSObject
}

func (s *ImportImageRequest) GetOSSRegion() *string {
	return s.OSSRegion
}

func (s *ImportImageRequest) GetOSType() *string {
	return s.OSType
}

func (s *ImportImageRequest) GetOSVersion() *string {
	return s.OSVersion
}

func (s *ImportImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ImportImageRequest) GetTargetOSSRegionId() *string {
	return s.TargetOSSRegionId
}

func (s *ImportImageRequest) SetArchitecture(v string) *ImportImageRequest {
	s.Architecture = &v
	return s
}

func (s *ImportImageRequest) SetComputeType(v string) *ImportImageRequest {
	s.ComputeType = &v
	return s
}

func (s *ImportImageRequest) SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest {
	s.DiskDeviceMapping = v
	return s
}

func (s *ImportImageRequest) SetImageFormat(v string) *ImportImageRequest {
	s.ImageFormat = &v
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

func (s *ImportImageRequest) SetOSSBucket(v string) *ImportImageRequest {
	s.OSSBucket = &v
	return s
}

func (s *ImportImageRequest) SetOSSObject(v string) *ImportImageRequest {
	s.OSSObject = &v
	return s
}

func (s *ImportImageRequest) SetOSSRegion(v string) *ImportImageRequest {
	s.OSSRegion = &v
	return s
}

func (s *ImportImageRequest) SetOSType(v string) *ImportImageRequest {
	s.OSType = &v
	return s
}

func (s *ImportImageRequest) SetOSVersion(v string) *ImportImageRequest {
	s.OSVersion = &v
	return s
}

func (s *ImportImageRequest) SetPlatform(v string) *ImportImageRequest {
	s.Platform = &v
	return s
}

func (s *ImportImageRequest) SetTargetOSSRegionId(v string) *ImportImageRequest {
	s.TargetOSSRegionId = &v
	return s
}

func (s *ImportImageRequest) Validate() error {
	return dara.Validate(s)
}

type ImportImageRequestDiskDeviceMapping struct {
	// example:
	//
	// www-cn
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// example:
	//
	// image-bucket
	OSSObject *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
	// example:
	//
	// cn-beijing
	OSSRegion *string `json:"OSSRegion,omitempty" xml:"OSSRegion,omitempty"`
}

func (s ImportImageRequestDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSObject() *string {
	return s.OSSObject
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSRegion() *string {
	return s.OSSRegion
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSBucket(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSBucket = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSObject(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSObject = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSRegion(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSRegion = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}
