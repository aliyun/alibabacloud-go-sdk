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
	SetBootMode(v string) *ImportImageRequest
	GetBootMode() *string
	SetClientToken(v string) *ImportImageRequest
	GetClientToken() *string
	SetDescription(v string) *ImportImageRequest
	GetDescription() *string
	SetDetectionStrategy(v string) *ImportImageRequest
	GetDetectionStrategy() *string
	SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest
	GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping
	SetDryRun(v bool) *ImportImageRequest
	GetDryRun() *bool
	SetFeatures(v *ImportImageRequestFeatures) *ImportImageRequest
	GetFeatures() *ImportImageRequestFeatures
	SetImageName(v string) *ImportImageRequest
	GetImageName() *string
	SetLicenseType(v string) *ImportImageRequest
	GetLicenseType() *string
	SetOSType(v string) *ImportImageRequest
	GetOSType() *string
	SetOwnerId(v int64) *ImportImageRequest
	GetOwnerId() *int64
	SetPlatform(v string) *ImportImageRequest
	GetPlatform() *string
	SetRegionId(v string) *ImportImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ImportImageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ImportImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportImageRequest
	GetResourceOwnerId() *int64
	SetRoleName(v string) *ImportImageRequest
	GetRoleName() *string
	SetStorageLocationArn(v string) *ImportImageRequest
	GetStorageLocationArn() *string
	SetTag(v []*ImportImageRequestTag) *ImportImageRequest
	GetTag() []*ImportImageRequestTag
}

type ImportImageRequest struct {
	// The system architecture. Valid values:
	//
	// - i386.
	//
	// - x86_64.
	//
	// - arm64.
	//
	// Default value: x86_64.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// - BIOS: BIOS boot mode.
	//
	// - UEFI: UEFI boot mode.
	//
	// Default value: BIOS. If `Architecture=arm64`, the default value is UEFI, and only UEFI is supported.
	//
	// <notice>
	//
	// To prevent instances from failing to start due to an unsupported boot mode, make sure that you understand the boot mode supported by the target image before you set this parameter. For more information about image boot modes, see [Image boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// </notice>.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image detection strategy. If this parameter is not specified, detection is not triggered. Only the Standard detection mode is supported.
	//
	// > Most Linux and Windows versions are supported. For more information about image detection items and operating system limitations, see [Image detection overview](https://help.aliyun.com/document_detail/439819.html) and [Operating system limitations for image detection](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// The information list of the custom image to create.
	DiskDeviceMapping []*ImportImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - false: performs a dry run and sends the request. If the request passes the dry run, a 2XX HTTP status code is returned and the resource status is queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The image feature-related properties.
	Features *ImportImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// ImageTestName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The license type. This parameter specifies the authorization mode when instances are created by calling [RunInstances](https://help.aliyun.com/document_detail/2679677.html) with the image. This value takes effect only for Windows Server images. Valid values:
	//
	// - Aliyun: Use the Alibaba Cloud official license. After the instance starts, the system attempts to automatically connect to the Alibaba Cloud KMS server for activation. The billing for the instance includes the Windows Server license fee.
	//
	// - BYOL: Bring Your Own License. After the instance starts, Alibaba Cloud does not automatically activate it. You must manually activate it by using your own valid license key. The billing for the instance does not include the Windows Server license fee.
	//
	// Default value: Aliyun.
	//
	// example:
	//
	// BYOL
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type. Valid values:
	//
	// - windows. You must also set `LicenseType`.
	//
	// - linux.
	//
	// Default value: linux.
	//
	// example:
	//
	// linux
	OSType  *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system version. Valid values:
	//
	// - Aliyun
	//
	// - Anolis
	//
	// - CentOS
	//
	// - Ubuntu
	//
	// - CoreOS
	//
	// - SUSE
	//
	// - Debian
	//
	// - OpenSUSE
	//
	// - FreeBSD
	//
	// - RedHat
	//
	// - Kylin
	//
	// - UOS
	//
	// - Fedora
	//
	// - Fedora CoreOS
	//
	// - CentOS Stream
	//
	// - AlmaLinux
	//
	// - Rocky Linux
	//
	// - Gentoo
	//
	// - Customized Linux
	//
	// - Others Linux
	//
	// - Windows Server 2022
	//
	// - Windows Server 2019
	//
	// - Windows Server 2016
	//
	// - Windows Server 2012
	//
	// - Windows Server 2008
	//
	// - Windows Server 2003
	//
	// - Other Windows
	//
	// Default value: Others Linux.
	//
	// example:
	//
	// Aliyun
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the source custom image. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the imported image belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the RAM role used to import the image.
	//
	// example:
	//
	// AliyunECSImageImportDefaultRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the CloudBox, which is used to uniquely identify the cloud storage location.
	//
	// > You need to specify this parameter only when you import an image file from OSS on CloudBox. If you do not use OSS on CloudBox, do not set this parameter. For more information, see [What is OSS on CloudBox](https://help.aliyun.com/document_detail/430190.html).
	//
	// The ARN must follow this format: `arn:acs:cloudbox:{RegionId}:{AliUid}:cloudbox/{CloudBoxId}`, where `{RegionId}` is the region ID of the CloudBox, `{AliUid}` is the Alibaba Cloud account ID, and `{CloudBoxId}` is the CloudBox ID.
	//
	// example:
	//
	// arn:acs:cloudbox:cn-hangzhou:123456:cloudbox/cb-xx***123
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags of the image.
	Tag []*ImportImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ImportImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *ImportImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ImportImageRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportImageRequest) GetDetectionStrategy() *string {
	return s.DetectionStrategy
}

func (s *ImportImageRequest) GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *ImportImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ImportImageRequest) GetFeatures() *ImportImageRequestFeatures {
	return s.Features
}

func (s *ImportImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ImportImageRequest) GetOSType() *string {
	return s.OSType
}

func (s *ImportImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ImportImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ImportImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportImageRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ImportImageRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *ImportImageRequest) GetTag() []*ImportImageRequestTag {
	return s.Tag
}

func (s *ImportImageRequest) SetArchitecture(v string) *ImportImageRequest {
	s.Architecture = &v
	return s
}

func (s *ImportImageRequest) SetBootMode(v string) *ImportImageRequest {
	s.BootMode = &v
	return s
}

func (s *ImportImageRequest) SetClientToken(v string) *ImportImageRequest {
	s.ClientToken = &v
	return s
}

func (s *ImportImageRequest) SetDescription(v string) *ImportImageRequest {
	s.Description = &v
	return s
}

func (s *ImportImageRequest) SetDetectionStrategy(v string) *ImportImageRequest {
	s.DetectionStrategy = &v
	return s
}

func (s *ImportImageRequest) SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest {
	s.DiskDeviceMapping = v
	return s
}

func (s *ImportImageRequest) SetDryRun(v bool) *ImportImageRequest {
	s.DryRun = &v
	return s
}

func (s *ImportImageRequest) SetFeatures(v *ImportImageRequestFeatures) *ImportImageRequest {
	s.Features = v
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

func (s *ImportImageRequest) SetOSType(v string) *ImportImageRequest {
	s.OSType = &v
	return s
}

func (s *ImportImageRequest) SetOwnerId(v int64) *ImportImageRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportImageRequest) SetPlatform(v string) *ImportImageRequest {
	s.Platform = &v
	return s
}

func (s *ImportImageRequest) SetRegionId(v string) *ImportImageRequest {
	s.RegionId = &v
	return s
}

func (s *ImportImageRequest) SetResourceGroupId(v string) *ImportImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerAccount(v string) *ImportImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerId(v int64) *ImportImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportImageRequest) SetRoleName(v string) *ImportImageRequest {
	s.RoleName = &v
	return s
}

func (s *ImportImageRequest) SetStorageLocationArn(v string) *ImportImageRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *ImportImageRequest) SetTag(v []*ImportImageRequestTag) *ImportImageRequest {
	s.Tag = v
	return s
}

func (s *ImportImageRequest) Validate() error {
	if s.DiskDeviceMapping != nil {
		for _, item := range s.DiskDeviceMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Features != nil {
		if err := s.Features.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportImageRequestDiskDeviceMapping struct {
	// The device name of DiskDeviceMapping.N.Device in the custom image.
	//
	// > This parameter will be deprecated. For better compatibility, do not use this parameter.
	//
	// example:
	//
	// null
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the custom image. Unit: GiB.
	//
	// The size includes the system disk and data disks. Make sure that the system disk space is greater than or equal to the size of the imported image file. Valid values:
	//
	// - When N is 1, the value specifies the system disk size. Valid values: 1 to 2048.
	//
	// - When N is 2 to 17, the value specifies the data disk size. Valid values: 1 to 2048.
	//
	// After you upload the source image file to OSS, you can view the image file size in the OSS bucket.
	//
	// > This parameter will be deprecated. For better compatibility, use `DiskDeviceMapping.N.DiskImageSize`.
	//
	// example:
	//
	// 80
	DiskImSize *int32 `json:"DiskImSize,omitempty" xml:"DiskImSize,omitempty"`
	// The size of the custom image after the image is imported.
	//
	// The size includes the system disk and data disks. Make sure that the system disk space is greater than or equal to the size of the imported image file. Valid values:
	//
	// - When N is 1, the value specifies the system disk size. Valid values: 1 to 2048. Unit: GiB.
	//
	// - When N is 2 to 17, the value specifies the data disk size. Valid values: 1 to 2048. Unit: GiB.
	//
	// After you upload the source image file to OSS, you can view the image file size in the OSS bucket.
	//
	// example:
	//
	// 80
	DiskImageSize *int32 `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	// The image format. Valid values:
	//
	// - RAW.
	//
	// - VHD.
	//
	// - QCOW2.
	//
	// - VMDK (in invitational preview).
	//
	// Default value: null, which indicates that Alibaba Cloud automatically detects the image format. The detected format prevails.
	//
	// example:
	//
	// QCOW2
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The OSS bucket where the image file is stored.
	//
	// > Before you import an image to this OSS bucket for the first time, add the RAM authorization policy as described in the **Operation description*	- section of this topic. Otherwise, the `NoSetRoletoECSServiceAccount` error is returned.
	//
	// example:
	//
	// ecsimageos
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The file name (key) of the image file stored in the OSS bucket after the image is uploaded to OSS.
	//
	// example:
	//
	// CentOS_5.4_32.raw
	OSSObject *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
}

func (s ImportImageRequestDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *ImportImageRequestDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *ImportImageRequestDiskDeviceMapping) GetDiskImSize() *int32 {
	return s.DiskImSize
}

func (s *ImportImageRequestDiskDeviceMapping) GetDiskImageSize() *int32 {
	return s.DiskImageSize
}

func (s *ImportImageRequestDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSObject() *string {
	return s.OSSObject
}

func (s *ImportImageRequestDiskDeviceMapping) SetDevice(v string) *ImportImageRequestDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetDiskImSize(v int32) *ImportImageRequestDiskDeviceMapping {
	s.DiskImSize = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetDiskImageSize(v int32) *ImportImageRequestDiskDeviceMapping {
	s.DiskImageSize = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetFormat(v string) *ImportImageRequestDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSBucket(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSBucket = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSObject(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSObject = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type ImportImageRequestFeatures struct {
	// The metadata access mode of the image. Valid values:
	//
	// - v1: When you create an ECS instance from this image, you cannot set the metadata access mode to hardened mode only.
	//
	// - v2: When you create an ECS instance from this image, you can set the metadata access mode to hardened mode only.
	//
	// Default value: v1.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Specifies whether the image supports NVMe. Valid values:
	//
	//  - supported: Instances created from this image support NVMe.
	//
	//  - unsupported: Instances created from this image do not support NVMe.
	//
	// example:
	//
	// supported
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s ImportImageRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestFeatures) GoString() string {
	return s.String()
}

func (s *ImportImageRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *ImportImageRequestFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *ImportImageRequestFeatures) SetImdsSupport(v string) *ImportImageRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *ImportImageRequestFeatures) SetNvmeSupport(v string) *ImportImageRequestFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *ImportImageRequestFeatures) Validate() error {
	return dara.Validate(s)
}

type ImportImageRequestTag struct {
	// The key of the image tag. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the image tag. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImportImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestTag) GoString() string {
	return s.String()
}

func (s *ImportImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *ImportImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *ImportImageRequestTag) SetKey(v string) *ImportImageRequestTag {
	s.Key = &v
	return s
}

func (s *ImportImageRequestTag) SetValue(v string) *ImportImageRequestTag {
	s.Value = &v
	return s
}

func (s *ImportImageRequestTag) Validate() error {
	return dara.Validate(s)
}
