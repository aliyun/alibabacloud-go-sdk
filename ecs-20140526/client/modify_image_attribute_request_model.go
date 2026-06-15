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
	// - `BIOS`: BIOS boot mode.
	//
	// - `UEFI`: UEFI boot mode.
	//
	// - `UEFI-Preferred`: UEFI-preferred boot mode.
	//
	// 	Notice:
	//
	// To prevent startup failures, verify the boot modes that the image supports before you change its boot mode. For more information, see [Boot modes](~~2244655#b9caa9b8bb1wf~~).
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The new description of the custom image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// If you do not specify this parameter, the original description is retained.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run to check whether the request is valid. Valid values:
	//
	// - `true`: performs a dry run to check the request for validity, syntax, and required permissions. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - `false` (default): sends the request. If the request passes the validation checks, the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The features of the image.
	//
	// if can be null:
	// true
	Features *ModifyImageAttributeRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character. The name cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// By default, this parameter is empty.
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
	// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character. The name cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// If you do not specify this parameter, the original name is retained.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The license type for activating the operating system after you import the image. The only valid value is `BYOL`.
	//
	// `BYOL`: Bring Your Own License. If you use the BYOL license type, you must ensure that your license key is supported for use on Alibaba Cloud.
	//
	// example:
	//
	// BYOL
	LicenseType  *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the custom image is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to view the latest list of Alibaba Cloud regions.
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
	// - `Deprecated`: Deprecates the image. If a custom image that you want to deprecate is shared, you must unshare it first. You cannot share or copy a deprecated image. However, you can use the image to create an instance or replace a system disk.
	//
	// - `Available`: Makes the image available. You can change the status of a deprecated image to `Available`.
	//
	// > However, if this is the only available custom image in the image family, deprecating it prevents the creation of instances from any image in that family. Use this option with caution.
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
	// - `v1`: When you create an ECS instance from this image, you cannot set the metadata access mode to `enforced mode`.
	//
	// - `v2`: When you create an ECS instance from this image, you can set the metadata access mode to `enforced mode`.
	//
	//   	Notice:
	//
	//   You cannot change the value of `ImdsSupport` from `v2` to `v1`. To use the `v1` mode, create a new image from a snapshot that is associated with the image and set `ImdsSupport` to `v1`.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Specifies whether the image supports NVMe. Valid values:
	//
	// - `supported`: The image supports NVMe. Instances that you create from this image support the NVMe protocol.
	//
	// - `unsupported`: The image does not support NVMe. Instances that you create from this image do not support the NVMe protocol.
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
