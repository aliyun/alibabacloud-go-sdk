// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootMode(v string) *ModifyImageAttributeRequest
	GetBootMode() *string
	SetDescription(v string) *ModifyImageAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyImageAttributeRequest
	GetDryRun() *bool
	SetFeatures(v *ModifyImageAttributeRequestFeatures) *ModifyImageAttributeRequest
	GetFeatures() *ModifyImageAttributeRequestFeatures
	SetImageFamily(v string) *ModifyImageAttributeRequest
	GetImageFamily() *string
	SetImageId(v string) *ModifyImageAttributeRequest
	GetImageId() *string
	SetImageName(v string) *ModifyImageAttributeRequest
	GetImageName() *string
	SetLicenseType(v string) *ModifyImageAttributeRequest
	GetLicenseType() *string
	SetOwnerAccount(v string) *ModifyImageAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyImageAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImageAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyImageAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyImageAttributeRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ModifyImageAttributeRequest
	GetStatus() *string
}

type ModifyImageAttributeRequest struct {
	// The boot mode of the image. Valid values:
	//
	// - BIOS: Basic Input/Output System (BIOS) boot mode.
	//
	// - UEFI: Unified Extensible Firmware Interface (UEFI) boot mode.
	//
	// - UEFI-Preferred: dual boot mode.
	//
	//
	// <notice>
	//
	//    To prevent instances from failing to start due to an unsupported boot mode, make sure that you understand the boot modes supported by the image before you modify this parameter. For more information about image boot modes, see [Image boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// </notice>
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The description of the custom image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: null, which indicates that the original description is retained.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The image feature attributes.
	//
	// if can be null:
	// true
	Features *ModifyImageAttributeRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), or hyphens (-).
	//
	// Default value: null.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp18ygjuqnwhechc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), or hyphens (-).
	//
	// Default value: null, which indicates that the original name is retained.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The license type used to activate the operating system after the image is imported. Currently, only BYOL is supported.
	//
	// BYOL: the license that comes with the source operating system. When you use BYOL, make sure that your license key supports use on Alibaba Cloud.
	//
	// example:
	//
	// BYOL
	LicenseType  *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the custom image. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The image status. Valid values:
	//
	// - Deprecated: sets the image to the deprecated state. If you have shared the custom image, you must unshare it before you can set it to the deprecated state. You cannot share or copy a deprecated image. However, you can use the image to create instances or replace system disks.
	//
	// - Available: sets the image to the available state. You can restore a deprecated image to the available state.
	//
	// > To roll back a custom image in an image family to the previous version, you can set the latest available custom image to the deprecated state. However, if the image is the only available custom image in the image family, the image family will have no available custom image for creating instances after the image is deprecated. Proceed with caution.
	//
	// example:
	//
	// Deprecated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *ModifyImageAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyImageAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyImageAttributeRequest) GetFeatures() *ModifyImageAttributeRequestFeatures {
	return s.Features
}

func (s *ModifyImageAttributeRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *ModifyImageAttributeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageAttributeRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyImageAttributeRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ModifyImageAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyImageAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImageAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyImageAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyImageAttributeRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyImageAttributeRequest) SetBootMode(v string) *ModifyImageAttributeRequest {
	s.BootMode = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDryRun(v bool) *ModifyImageAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetFeatures(v *ModifyImageAttributeRequestFeatures) *ModifyImageAttributeRequest {
	s.Features = v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageFamily(v string) *ModifyImageAttributeRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageName(v string) *ModifyImageAttributeRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetLicenseType(v string) *ModifyImageAttributeRequest {
	s.LicenseType = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetOwnerAccount(v string) *ModifyImageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetOwnerId(v int64) *ModifyImageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetResourceOwnerAccount(v string) *ModifyImageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetResourceOwnerId(v int64) *ModifyImageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetStatus(v string) *ModifyImageAttributeRequest {
	s.Status = &v
	return s
}

func (s *ModifyImageAttributeRequest) Validate() error {
	if s.Features != nil {
		if err := s.Features.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyImageAttributeRequestFeatures struct {
	// The metadata access mode of the image. Valid values:
	//
	// - v1: when you create an ECS instance from this image, you cannot set the metadata access mode to IMDSv2 only (hardened mode).
	//
	// - v2: when you create an ECS instance from this image, you can set the metadata access mode to IMDSv2 only (hardened mode).
	//
	// <notice>
	//
	//   ImdsSupport cannot be changed from v2 to v1. If you need to change it, create a new image from the snapshot associated with this image and set the value to v1.
	//
	// </notice>
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Modifies the NVMe support attribute of the image. If this parameter is not specified, the current value is retained.
	//
	// 	Notice: Before enabling this feature, make sure that the NVMe driver is pre-installed in the operating system. Recommended procedure: install the driver on an instance, create a custom image, and then call this operation. Forcibly enabling this feature without the driver will cause instance startup failures.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// supported
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s ModifyImageAttributeRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeRequestFeatures) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *ModifyImageAttributeRequestFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *ModifyImageAttributeRequestFeatures) SetImdsSupport(v string) *ModifyImageAttributeRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *ModifyImageAttributeRequestFeatures) SetNvmeSupport(v string) *ModifyImageAttributeRequestFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *ModifyImageAttributeRequestFeatures) Validate() error {
	return dara.Validate(s)
}
