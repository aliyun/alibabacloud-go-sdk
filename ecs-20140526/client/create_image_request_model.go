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
	// The system architecture. After a data disk snapshot is specified as the system disk of the image, use this parameter to specify the system architecture of the system disk. Valid values:
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
	// - (Default) UEFI-Preferred: dual boot mode.
	//
	// <notice>
	//
	// To prevent instances from failing to start due to an unsupported boot mode, make sure that you understand the boot modes supported by the target image before specifying this parameter. For more information about image boot modes, see [Image boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// </notice>
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the image. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// ImageTestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image detection strategy. If this parameter is not specified, detection is not triggered. Only the Standard detection mode is supported.
	//
	// > Most Linux and Windows versions are supported. For more information about image detection items and operating system limitations, see [Image detection overview](https://help.aliyun.com/document_detail/439819.html) and [Operating system limitations for image detection](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// The disk and snapshot information used to create the custom image. Use this parameter to specify snapshots when you want to create a custom image from system disk and data disk snapshots.
	DiskDeviceMapping []*CreateImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	DryRun            *bool                                  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The image feature properties.
	Features *CreateImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The image family name. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with aliyun or acs:. It cannot contain http:// or https://. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image name. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// TestCentOS
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image version.
	//
	// > If you specify an instance ID (`InstanceId`) and the image of the instance is an Alibaba Cloud Marketplace image or a custom image created from an Alibaba Cloud Marketplace image, this parameter must be the same as the `ImageVersion` of the current instance image or left empty.
	//
	// example:
	//
	// 2017011017
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The instance ID. This parameter is required when you create a custom image from an instance.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system distribution. After a data disk snapshot is specified as the system disk of the image, use this parameter to specify the operating system distribution of the system disk. Valid values:
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
	// The region ID of the image. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the custom image belongs. If this parameter is not set, the created image belongs to the default resource group.
	//
	// > If you invoke this operation as a Resource Access Management (RAM) user and `ResourceGroupId` is left empty, note that when the RAM user does not have permissions on the default resource group, the error message `Forbidden: User not authorized to operate on the specified resource` is returned. Set a resource group ID that the RAM user has permissions on, or grant the RAM user permissions on the default resource group before invoking this operation again.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot ID used to create the custom image.
	//
	// > If you want to create a custom image only from the system disk snapshot of an instance, you can use this parameter or the `DiskDeviceMapping.N.SnapshotId` parameter. To include data disk snapshots, use only the `DiskDeviceMapping.N.SnapshotId` parameter to specify snapshots.
	//
	// example:
	//
	// s-bp17441ohwkdca0****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags.
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
	// The device name in the custom image. Valid values:
	//
	// - The device name of the system disk must be /dev/xvda.
	//
	// - Data disk device names are sequentially ordered from /dev/xvdb to /dev/xvdz and cannot be duplicated.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The type of the disk in the new image. You can use this parameter to specify a data disk snapshot as the system disk of the image. If this parameter is not specified, the disk type defaults to the type of the disk corresponding to the snapshot. Valid values:
	//
	// - system: system disk. Only one system disk snapshot can be specified.
	//
	// - data: data disk. Up to 16 data disk snapshots can be specified.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The size of the disk, in GiB. The valid values and default value of DiskDeviceMapping.N.Size depend on DiskDeviceMapping.N.SnapshotId:
	//
	// - If SnapshotId is not specified, the valid values and default value of Size are:
	//
	//     - Basic disk: 5 to 2000 GiB. Default value: 5.
	//
	//     - Other disk types: 20 to 32768 GiB. Default value: 20.
	//
	// - If SnapshotId is specified, the value of Size must be greater than or equal to the size of the snapshot. Default value: the size of the snapshot.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID.
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
	// The metadata access mode of the image. Valid values:
	//
	// - v1: When you create an ECS instance from this image, you cannot set the metadata access mode to "security hardening mode only".
	//
	// - v2: When you create an ECS instance from this image, you can set the metadata access mode to "security hardening mode only".
	//
	// Default value: When creating an image from a snapshot, the default is v1. When creating an image from an instance, the default is the ImdsSupport value of the image used when the instance was created.
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
	// The tag key of the image. Valid values of N: 1 to 20. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the image. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
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
