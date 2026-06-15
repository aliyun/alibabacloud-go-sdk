// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *CreateImageRequest
	GetArchitecture() *string
	SetBootMode(v string) *CreateImageRequest
	GetBootMode() *string
	SetClientToken(v string) *CreateImageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateImageRequest
	GetDescription() *string
	SetDetectionStrategy(v string) *CreateImageRequest
	GetDetectionStrategy() *string
	SetDiskDeviceMapping(v []*CreateImageRequestDiskDeviceMapping) *CreateImageRequest
	GetDiskDeviceMapping() []*CreateImageRequestDiskDeviceMapping
	SetDryRun(v bool) *CreateImageRequest
	GetDryRun() *bool
	SetFeatures(v *CreateImageRequestFeatures) *CreateImageRequest
	GetFeatures() *CreateImageRequestFeatures
	SetImageFamily(v string) *CreateImageRequest
	GetImageFamily() *string
	SetImageName(v string) *CreateImageRequest
	GetImageName() *string
	SetImageVersion(v string) *CreateImageRequest
	GetImageVersion() *string
	SetInstanceId(v string) *CreateImageRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateImageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateImageRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateImageRequest
	GetPlatform() *string
	SetRegionId(v string) *CreateImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateImageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateImageRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *CreateImageRequest
	GetSnapshotId() *string
	SetTag(v []*CreateImageRequestTag) *CreateImageRequest
	GetTag() []*CreateImageRequestTag
}

type CreateImageRequest struct {
	// The system disk architecture. If you create the image\\"s system disk from a data disk snapshot, you must specify this parameter to identify the system disk architecture. Valid values:
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
	// - BIOS: BIOS boot mode.
	//
	// - UEFI: UEFI boot mode.
	//
	// - UEFI-Preferred: The image supports both BIOS and UEFI boot modes. The UEFI boot mode is preferred. This is the default value.
	//
	// 	Notice:
	//
	// If you specify a boot mode that the image does not support, instances created from the image may fail to start. Before you specify this parameter, ensure you know the boot modes that the image supports. For more information, see [Boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// A client-generated token to ensure the request is idempotent. You must ensure that the token is unique across different requests. The `ClientToken` value can contain only ASCII characters and cannot exceed 64 characters. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image description. It must be 2 to 256 characters long and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// ImageTestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image check policy. If you do not specify this parameter, no check is performed. Only the Standard mode is supported.
	//
	// > This feature is supported for most Linux and Windows versions. For more information about the check items and the operating systems that support this feature, see [Overview of image check](https://help.aliyun.com/document_detail/439819.html) and [Operating systems that support image check](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// The mappings between disks and snapshots used to create the custom image. If you need to create a custom image from a system disk snapshot and data disk snapshots, specify this parameter.
	DiskDeviceMapping []*CreateImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run to check the request. Valid values:
	//
	// - true: performs a dry run but does not create the image. The system checks whether your AccessKey pair is valid, whether RAM users are granted permissions, and whether the required parameters are specified. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - false: Sends the request to perform the operation. If the request is valid, a 2xx HTTP status code is returned and the image is created.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The image attributes.
	Features *CreateImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image family. The name must be 2 to 128 characters long and start with a letter or a Chinese character. It cannot start with aliyun or acs:, nor contain http\\:// or https\\://. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The name of the image. The name must be 2 to 128 characters long. It must start with a letter or a Chinese character and must not start with http\\:// or https\\://. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// TestCentOS
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The version of the image.
	//
	// > If you specify an instance ID (`InstanceId`) and the instance was created from an Alibaba Cloud Marketplace image (or a custom image based on a Marketplace image), this parameter must match the `ImageVersion` of the instance\\"s image or be left empty.
	//
	// example:
	//
	// 2017011017
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The ID of the instance. This parameter is required when you create a custom image from an instance.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system distribution. You must specify this parameter to identify the operating system distribution when you use a data disk snapshot to create the image\\"s system disk. Valid values:
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
	// Default value: Others Linux.
	//
	// example:
	//
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the region where the image will be created. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to add the custom image. If you do not specify this parameter, the image is added to the default resource group.
	//
	// > As a RAM user, you must have the required permissions to call this operation. If you leave `ResourceGroupId` empty, the `Forbidden: User not authorized to operate on the specified resource` error is returned if you lack permissions on the default resource group. To resolve this issue, specify the ID of a resource group for which you have permissions, or ask an administrator to grant you permissions on the default resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot used to create the custom image.
	//
	// > If you create a custom image from only a system disk snapshot, you can use either this parameter or the `DiskDeviceMapping.N.SnapshotId` parameter. If you want to include data disk snapshots, you must use the `DiskDeviceMapping.N.SnapshotId` parameter to specify the snapshots.
	//
	// example:
	//
	// s-bp17441ohwkdca0****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags to add to the image.
	Tag []*CreateImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *CreateImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageRequest) GetDetectionStrategy() *string {
	return s.DetectionStrategy
}

func (s *CreateImageRequest) GetDiskDeviceMapping() []*CreateImageRequestDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *CreateImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateImageRequest) GetFeatures() *CreateImageRequestFeatures {
	return s.Features
}

func (s *CreateImageRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImageRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *CreateImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateImageRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequest) GetTag() []*CreateImageRequestTag {
	return s.Tag
}

func (s *CreateImageRequest) SetArchitecture(v string) *CreateImageRequest {
	s.Architecture = &v
	return s
}

func (s *CreateImageRequest) SetBootMode(v string) *CreateImageRequest {
	s.BootMode = &v
	return s
}

func (s *CreateImageRequest) SetClientToken(v string) *CreateImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetDetectionStrategy(v string) *CreateImageRequest {
	s.DetectionStrategy = &v
	return s
}

func (s *CreateImageRequest) SetDiskDeviceMapping(v []*CreateImageRequestDiskDeviceMapping) *CreateImageRequest {
	s.DiskDeviceMapping = v
	return s
}

func (s *CreateImageRequest) SetDryRun(v bool) *CreateImageRequest {
	s.DryRun = &v
	return s
}

func (s *CreateImageRequest) SetFeatures(v *CreateImageRequestFeatures) *CreateImageRequest {
	s.Features = v
	return s
}

func (s *CreateImageRequest) SetImageFamily(v string) *CreateImageRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetImageVersion(v string) *CreateImageRequest {
	s.ImageVersion = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetOwnerAccount(v string) *CreateImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetOwnerId(v int64) *CreateImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageRequest) SetPlatform(v string) *CreateImageRequest {
	s.Platform = &v
	return s
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetResourceGroupId(v string) *CreateImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerAccount(v string) *CreateImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerId(v int64) *CreateImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetTag(v []*CreateImageRequestTag) *CreateImageRequest {
	s.Tag = v
	return s
}

func (s *CreateImageRequest) Validate() error {
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

type CreateImageRequestDiskDeviceMapping struct {
	// The device name of the disk in the custom image. Valid values:
	//
	// - The device name of the system disk must be /dev/xvda.
	//
	// - The device names of data disks are assigned in sequence from /dev/xvdb to /dev/xvdz and cannot be repeated.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The type of the disk in the image. You can specify this parameter to use a data disk snapshot as the system disk of the image. If you do not specify this parameter, the disk type is determined by the type of the source snapshot. Valid values:
	//
	// - system: system disk. You can specify only one system disk snapshot.
	//
	// - data: data disk. You can specify a maximum of 16 data disk snapshots.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The size of the cloud disk, in GiB. The valid values and default value of `DiskDeviceMapping.N.Size` vary based on whether `DiskDeviceMapping.N.SnapshotId` is specified.
	//
	// - If `DiskDeviceMapping.N.SnapshotId` is not specified, the value of this parameter depends on the disk type:
	//
	//   - For basic disks, the value range is 5 to 2,000 and the default value is 5.
	//
	//   - For other disk types, the value range is 20 to 32,768 and the default value is 20.
	//
	// - If `DiskDeviceMapping.N.SnapshotId` is specified, the value of `DiskDeviceMapping.N.Size` must be greater than or equal to the snapshot\\"s size. The default value is the snapshot\\"s size.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-bp17441ohwkdca0****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateImageRequestDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *CreateImageRequestDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *CreateImageRequestDiskDeviceMapping) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateImageRequestDiskDeviceMapping) GetSize() *int32 {
	return s.Size
}

func (s *CreateImageRequestDiskDeviceMapping) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequestDiskDeviceMapping) SetDevice(v string) *CreateImageRequestDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetDiskType(v string) *CreateImageRequestDiskDeviceMapping {
	s.DiskType = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetSize(v int32) *CreateImageRequestDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetSnapshotId(v string) *CreateImageRequestDiskDeviceMapping {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type CreateImageRequestFeatures struct {
	// The instance metadata access mode. Valid values:
	//
	// - v1: The normal mode. When you create an ECS instance from an image that has the metadata access mode set to this value, you cannot configure the instance metadata access mode as Enforced.
	//
	// - v2: The enforced mode. When you create an ECS instance from an image that has the metadata access mode set to this value, you can configure the instance metadata access mode as Enforced.
	//
	// Default value: v1 if you create the image from a snapshot. If you create the image from an instance, the value is inherited from the source instance\\"s image.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
}

func (s CreateImageRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestFeatures) GoString() string {
	return s.String()
}

func (s *CreateImageRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *CreateImageRequestFeatures) SetImdsSupport(v string) *CreateImageRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *CreateImageRequestFeatures) Validate() error {
	return dara.Validate(s)
}

type CreateImageRequestTag struct {
	// The key of tag N to add to the image.
	//
	// > For compatibility, we recommend that you use the `Tag.N.Key` parameter.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the image. Valid values of N: 1 to 20. The tag value can be an empty string, up to 128 characters long, and cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateImageRequestTag) SetKey(v string) *CreateImageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageRequestTag) SetValue(v string) *CreateImageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateImageRequestTag) Validate() error {
	return dara.Validate(s)
}
