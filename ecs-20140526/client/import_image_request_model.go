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
	// - i386
	//
	// - x86_64
	//
	// - arm64
	//
	// Default value: x86_64.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// - BIOS: the BIOS boot mode.
	//
	// - UEFI: the UEFI boot mode.
	//
	// Default value: BIOS. If you set `Architecture` to `arm64`, the value of this parameter defaults to UEFI and can only be set to UEFI.
	//
	// 	Notice:
	//
	// To prevent startup failures, ensure the boot mode you specify is supported by the image. For more information, see [Image boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// A client-generated token that ensures the idempotence of a request. The token must be unique across requests. The token can contain only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the image. The description must be 2 to 256 characters long and cannot start with `http://` or `https://`. Both English and Chinese characters are supported.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The policy for checking the image. If you do not specify this parameter, the system does not check the image. This parameter supports only the standard detection mode. Set the value to `Standard`.
	//
	// > This feature is supported on most Linux and Windows versions. For more information about the check items and the operating systems that support this feature, see [Image detection overview](https://help.aliyun.com/document_detail/439819.html) and [Operating system limitations for image detection](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// A list of disk device mappings for the custom image.
	DiskDeviceMapping []*ImportImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run for the request. Valid values:
	//
	// - `true`: performs a check request without executing the actual operation. The system checks whether the request parameters are valid, the request format is correct, and the required permissions are granted. If the check fails, the system returns an error message. If the check succeeds, the system returns the `DryRunOperation` error code.
	//
	// - `false`: sends a normal request. After the request passes the check, the system returns a 2xx HTTP status code and performs the operation.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The properties of image features.
	Features *ImportImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image. The name must be 2 to 128 characters long and start with a letter or a Chinese character. It can contain digits, periods (.), colons (:), underscores (_), and hyphens (-). The name cannot start with `aliyun` or `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// ImageTestName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The license type. This parameter sets the licensing model for instances that are created from the image by calling the [RunInstances](https://help.aliyun.com/document_detail/2679677.html) operation. This parameter applies only to Windows Server images. Valid values:
	//
	// - Aliyun: Uses a license provided by Alibaba Cloud. When you start an instance created from this image, the system attempts to automatically connect to the Alibaba Cloud KMS server for activation. The fees for the instance include the cost of the Windows Server license.
	//
	// - BYOL: Bring Your Own License. When you start an instance created from this image, Alibaba Cloud does not provide activation. You must use your own license key to manually activate the operating system. The fees for the instance do not include the cost of the Windows Server license.
	//
	// Default value: Aliyun.
	//
	// example:
	//
	// BYOL
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type. Valid values:
	//
	// - `windows`: You must also set the `LicenseType` parameter.
	//
	// - `linux`
	//
	// Default value: linux.
	//
	// example:
	//
	// linux
	OSType  *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system distribution. Valid values:
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
	// The ID of the region where the source custom image is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the imported image belongs.
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
	// The Alibaba Cloud Resource Name (ARN) of the CloudBox, which uniquely identifies the cloud storage location.
	//
	// > You must specify this parameter only when you import an image file from OSS ON CloudBox. If you do not use OSS ON CloudBox, do not specify this parameter. For more information, see [What is OSS ON CloudBox?](https://help.aliyun.com/document_detail/430190.html).
	//
	// The ARN must be in the `arn:acs:cloudbox:{RegionId}:{AliUid}:cloudbox/{CloudBoxId}` format. Replace `{RegionId}` with the ID of the region where the CloudBox is located, `{AliUid}` with the ID of your Alibaba Cloud account, and `{CloudBoxId}` with the ID of the CloudBox.
	//
	// example:
	//
	// arn:acs:cloudbox:cn-hangzhou:123456:cloudbox/cb-xx***123
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags to add to the image.
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
	// The device name of the disk (`DiskDeviceMapping.N.Device`) in the custom image.
	//
	// > This parameter is being phased out. To ensure compatibility, we recommend that you avoid using this parameter.
	//
	// example:
	//
	// null
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the disk, in GiB.
	//
	// The system disk size must be greater than or equal to the size of the imported image file. Valid values:
	//
	// - For N=1, the disk is a system disk. The value must be in the range of 1 to 2,048.
	//
	// - For N=2 to 17, the disk is a data disk. The value must be in the range of 1 to 2,048.
	//
	// After you upload the source image file to an OSS bucket, you can view the size of the file in the bucket.
	//
	// > This parameter is being deprecated. For better compatibility, we recommend that you use the `DiskDeviceMapping.N.DiskImageSize` parameter.
	//
	// example:
	//
	// 80
	DiskImSize *int32 `json:"DiskImSize,omitempty" xml:"DiskImSize,omitempty"`
	// The size of the disk after the image is imported, in GiB.
	//
	// The value of this parameter for the system disk must be greater than or equal to the size of the image file. Valid values:
	//
	// - For N=1, the disk is a system disk. The value must be in the range of 1 to 2,048.
	//
	// - For N=2 to 17, the disk is a data disk. The value must be in the range of 1 to 2,048.
	//
	// After you upload the source image file to an OSS bucket, you can view the size of the file in the bucket.
	//
	// example:
	//
	// 80
	DiskImageSize *int32 `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	// The format of the image file. Valid values:
	//
	// - RAW
	//
	// - VHD
	//
	// - QCOW2
	//
	// - VMDK (This feature is in invitation-only preview.)
	//
	// Default value: None. If you leave this parameter empty, Alibaba Cloud automatically detects the image format and uses the detected format.
	//
	// example:
	//
	// QCOW2
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The OSS bucket where the image file is stored.
	//
	// > Before you import an image from an OSS bucket for the first time, you must add a RAM policy as described in the **Description*	- section of this topic. Otherwise, the API returns the `NoSetRoletoECSServiceAccount` error.
	//
	// example:
	//
	// ecsimageos
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The object key of the image file in the OSS bucket.
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
	// - v1: When you create an ECS instance from the image, you cannot set the metadata access mode to Security-Hardened Mode.
	//
	// - v2: When you create an ECS instance from the image, you can set the metadata access mode to Security-Hardened Mode.
	//
	// Default value: v1.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Specifies whether the image supports NVMe. Valid values:
	//
	// - supported: Instances created from the image support the NVMe protocol.
	//
	// - unsupported: Instances created from the image do not support the NVMe protocol.
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
	// The key of tag N. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters long and cannot start with `aliyun` or `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters long, cannot start with `acs:`, and cannot contain `http://` or `https://`.
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
