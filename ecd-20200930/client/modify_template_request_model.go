// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyTemplateRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *ModifyTemplateRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *ModifyTemplateRequest
	GetChargeType() *string
	SetDataDiskList(v []*ModifyTemplateRequestDataDiskList) *ModifyTemplateRequest
	GetDataDiskList() []*ModifyTemplateRequestDataDiskList
	SetDefaultLanguage(v string) *ModifyTemplateRequest
	GetDefaultLanguage() *string
	SetDescription(v string) *ModifyTemplateRequest
	GetDescription() *string
	SetImageId(v string) *ModifyTemplateRequest
	GetImageId() *string
	SetPeriod(v int32) *ModifyTemplateRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyTemplateRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *ModifyTemplateRequest
	GetPolicyGroupId() *string
	SetPostPaidAfterUsedUp(v bool) *ModifyTemplateRequest
	GetPostPaidAfterUsedUp() *bool
	SetRegionConfigList(v []*ModifyTemplateRequestRegionConfigList) *ModifyTemplateRequest
	GetRegionConfigList() []*ModifyTemplateRequestRegionConfigList
	SetResourceGroupId(v string) *ModifyTemplateRequest
	GetResourceGroupId() *string
	SetResourceTagList(v []*ModifyTemplateRequestResourceTagList) *ModifyTemplateRequest
	GetResourceTagList() []*ModifyTemplateRequestResourceTagList
	SetSiteConfigList(v []*ModifyTemplateRequestSiteConfigList) *ModifyTemplateRequest
	GetSiteConfigList() []*ModifyTemplateRequestSiteConfigList
	SetSystemDiskPerformanceLevel(v string) *ModifyTemplateRequest
	GetSystemDiskPerformanceLevel() *string
	SetSystemDiskSize(v int32) *ModifyTemplateRequest
	GetSystemDiskSize() *int32
	SetTemplateId(v string) *ModifyTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *ModifyTemplateRequest
	GetTemplateName() *string
	SetTimerGroupId(v string) *ModifyTemplateRequest
	GetTimerGroupId() *string
	SetUserDuration(v int32) *ModifyTemplateRequest
	GetUserDuration() *int32
}

type ModifyTemplateRequest struct {
	AutoPay      *bool                                `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew    *bool                                `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType   *string                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DataDiskList []*ModifyTemplateRequestDataDiskList `json:"DataDiskList,omitempty" xml:"DataDiskList,omitempty" type:"Repeated"`
	// The default language to set when the WUYING Workspace starts. This parameter is valid only when you create a WUYING Workspace from an OS image.
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	// The description of the template. The description must meet the following requirements:
	//
	// - It must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// - It can contain Chinese characters, letters, digits, spaces, and special characters. Use line breaks to start a new line.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WUYING Workspace image. You can find the ID on the Image Management page. OS images and custom images are supported.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Period     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the global policy.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId       *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PostPaidAfterUsedUp *bool   `json:"PostPaidAfterUsedUp,omitempty" xml:"PostPaidAfterUsedUp,omitempty"`
	// The region-specific template configurations. You can specify configurations for multiple regions. The system matches the configuration based on the specific region.
	//
	// > You can specify configurations for up to 20 regions.
	RegionConfigList []*ModifyTemplateRequestRegionConfigList `json:"RegionConfigList,omitempty" xml:"RegionConfigList,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-a5fqjjqaejt***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Tags for the cloud computer, in key-value format. You can specify up to 20 tags.
	ResourceTagList []*ModifyTemplateRequestResourceTagList `json:"ResourceTagList,omitempty" xml:"ResourceTagList,omitempty" type:"Repeated"`
	SiteConfigList  []*ModifyTemplateRequestSiteConfigList  `json:"SiteConfigList,omitempty" xml:"SiteConfigList,omitempty" type:"Repeated"`
	// The type of the system disk.
	//
	// > Enhanced SSD (ESSD) disks are supported only by cloud computers with high clock speeds and powerful graphics capabilities.
	//
	// example:
	//
	// AutoPL
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. The value must be between 40 and 500, inclusive. The step size is 10 GiB.
	//
	// > The system disk size cannot be smaller than the size of the image.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-0caoeogs88y*****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template. The name must meet the following requirements:
	//
	// - It must be 2 to 126 characters in length.
	//
	// - It must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - It can contain letters, digits, Chinese characters, colons (:), underscores (_), and hyphens (-). It cannot contain periods (.).
	//
	// example:
	//
	// My cloud desktop template 001
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the configuration group.
	//
	// example:
	//
	// bcc-dweha*****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	UserDuration *int32  `json:"UserDuration,omitempty" xml:"UserDuration,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyTemplateRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyTemplateRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyTemplateRequest) GetDataDiskList() []*ModifyTemplateRequestDataDiskList {
	return s.DataDiskList
}

func (s *ModifyTemplateRequest) GetDefaultLanguage() *string {
	return s.DefaultLanguage
}

func (s *ModifyTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTemplateRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyTemplateRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyTemplateRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyTemplateRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyTemplateRequest) GetPostPaidAfterUsedUp() *bool {
	return s.PostPaidAfterUsedUp
}

func (s *ModifyTemplateRequest) GetRegionConfigList() []*ModifyTemplateRequestRegionConfigList {
	return s.RegionConfigList
}

func (s *ModifyTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyTemplateRequest) GetResourceTagList() []*ModifyTemplateRequestResourceTagList {
	return s.ResourceTagList
}

func (s *ModifyTemplateRequest) GetSiteConfigList() []*ModifyTemplateRequestSiteConfigList {
	return s.SiteConfigList
}

func (s *ModifyTemplateRequest) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *ModifyTemplateRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ModifyTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyTemplateRequest) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *ModifyTemplateRequest) GetUserDuration() *int32 {
	return s.UserDuration
}

func (s *ModifyTemplateRequest) SetAutoPay(v bool) *ModifyTemplateRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyTemplateRequest) SetAutoRenew(v bool) *ModifyTemplateRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyTemplateRequest) SetChargeType(v string) *ModifyTemplateRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyTemplateRequest) SetDataDiskList(v []*ModifyTemplateRequestDataDiskList) *ModifyTemplateRequest {
	s.DataDiskList = v
	return s
}

func (s *ModifyTemplateRequest) SetDefaultLanguage(v string) *ModifyTemplateRequest {
	s.DefaultLanguage = &v
	return s
}

func (s *ModifyTemplateRequest) SetDescription(v string) *ModifyTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyTemplateRequest) SetImageId(v string) *ModifyTemplateRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyTemplateRequest) SetPeriod(v int32) *ModifyTemplateRequest {
	s.Period = &v
	return s
}

func (s *ModifyTemplateRequest) SetPeriodUnit(v string) *ModifyTemplateRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyTemplateRequest) SetPolicyGroupId(v string) *ModifyTemplateRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyTemplateRequest) SetPostPaidAfterUsedUp(v bool) *ModifyTemplateRequest {
	s.PostPaidAfterUsedUp = &v
	return s
}

func (s *ModifyTemplateRequest) SetRegionConfigList(v []*ModifyTemplateRequestRegionConfigList) *ModifyTemplateRequest {
	s.RegionConfigList = v
	return s
}

func (s *ModifyTemplateRequest) SetResourceGroupId(v string) *ModifyTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyTemplateRequest) SetResourceTagList(v []*ModifyTemplateRequestResourceTagList) *ModifyTemplateRequest {
	s.ResourceTagList = v
	return s
}

func (s *ModifyTemplateRequest) SetSiteConfigList(v []*ModifyTemplateRequestSiteConfigList) *ModifyTemplateRequest {
	s.SiteConfigList = v
	return s
}

func (s *ModifyTemplateRequest) SetSystemDiskPerformanceLevel(v string) *ModifyTemplateRequest {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *ModifyTemplateRequest) SetSystemDiskSize(v int32) *ModifyTemplateRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateId(v string) *ModifyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateName(v string) *ModifyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTimerGroupId(v string) *ModifyTemplateRequest {
	s.TimerGroupId = &v
	return s
}

func (s *ModifyTemplateRequest) SetUserDuration(v int32) *ModifyTemplateRequest {
	s.UserDuration = &v
	return s
}

func (s *ModifyTemplateRequest) Validate() error {
	if s.DataDiskList != nil {
		for _, item := range s.DataDiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionConfigList != nil {
		for _, item := range s.RegionConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceTagList != nil {
		for _, item := range s.ResourceTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SiteConfigList != nil {
		for _, item := range s.SiteConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyTemplateRequestDataDiskList struct {
	// The performance level of the data disk. The default value is `AutoPL`.
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB. The value must be between 40 and 2040, inclusive. The step size is 10 GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyTemplateRequestDataDiskList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequestDataDiskList) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequestDataDiskList) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyTemplateRequestDataDiskList) GetSize() *int32 {
	return s.Size
}

func (s *ModifyTemplateRequestDataDiskList) SetPerformanceLevel(v string) *ModifyTemplateRequestDataDiskList {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyTemplateRequestDataDiskList) SetSize(v int32) *ModifyTemplateRequestDataDiskList {
	s.Size = &v
	return s
}

func (s *ModifyTemplateRequestDataDiskList) Validate() error {
	return dara.Validate(s)
}

type ModifyTemplateRequestRegionConfigList struct {
	// The ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou+dir-709****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to obtain a list of regions that WUYING Workspace supports.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud desktop instance type.
	//
	// example:
	//
	// eds.enterprise_office.8c16g
	ResourceInstanceType *string `json:"ResourceInstanceType,omitempty" xml:"ResourceInstanceType,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-35fvn8m2*****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-adjrehad1****
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnable *bool `json:"VolumeEncryptionEnable,omitempty" xml:"VolumeEncryptionEnable,omitempty"`
	// The ID of the KMS key to use when disk encryption is enabled.
	//
	// example:
	//
	// a7b3c0c8-b3a2-4876-b1cc-116dddc9****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s ModifyTemplateRequestRegionConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequestRegionConfigList) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequestRegionConfigList) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyTemplateRequestRegionConfigList) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTemplateRequestRegionConfigList) GetResourceInstanceType() *string {
	return s.ResourceInstanceType
}

func (s *ModifyTemplateRequestRegionConfigList) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *ModifyTemplateRequestRegionConfigList) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ModifyTemplateRequestRegionConfigList) GetVolumeEncryptionEnable() *bool {
	return s.VolumeEncryptionEnable
}

func (s *ModifyTemplateRequestRegionConfigList) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *ModifyTemplateRequestRegionConfigList) SetOfficeSiteId(v string) *ModifyTemplateRequestRegionConfigList {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetRegionId(v string) *ModifyTemplateRequestRegionConfigList {
	s.RegionId = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetResourceInstanceType(v string) *ModifyTemplateRequestRegionConfigList {
	s.ResourceInstanceType = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetSnapshotPolicyId(v string) *ModifyTemplateRequestRegionConfigList {
	s.SnapshotPolicyId = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetSubnetId(v string) *ModifyTemplateRequestRegionConfigList {
	s.SubnetId = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetVolumeEncryptionEnable(v bool) *ModifyTemplateRequestRegionConfigList {
	s.VolumeEncryptionEnable = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) SetVolumeEncryptionKey(v string) *ModifyTemplateRequestRegionConfigList {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *ModifyTemplateRequestRegionConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyTemplateRequestResourceTagList struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// design
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyTemplateRequestResourceTagList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequestResourceTagList) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequestResourceTagList) GetKey() *string {
	return s.Key
}

func (s *ModifyTemplateRequestResourceTagList) GetValue() *string {
	return s.Value
}

func (s *ModifyTemplateRequestResourceTagList) SetKey(v string) *ModifyTemplateRequestResourceTagList {
	s.Key = &v
	return s
}

func (s *ModifyTemplateRequestResourceTagList) SetValue(v string) *ModifyTemplateRequestResourceTagList {
	s.Value = &v
	return s
}

func (s *ModifyTemplateRequestResourceTagList) Validate() error {
	return dara.Validate(s)
}

type ModifyTemplateRequestSiteConfigList struct {
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	SiteId    *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ModifyTemplateRequestSiteConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequestSiteConfigList) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequestSiteConfigList) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *ModifyTemplateRequestSiteConfigList) GetSiteId() *string {
	return s.SiteId
}

func (s *ModifyTemplateRequestSiteConfigList) SetAppRuleId(v string) *ModifyTemplateRequestSiteConfigList {
	s.AppRuleId = &v
	return s
}

func (s *ModifyTemplateRequestSiteConfigList) SetSiteId(v string) *ModifyTemplateRequestSiteConfigList {
	s.SiteId = &v
	return s
}

func (s *ModifyTemplateRequestSiteConfigList) Validate() error {
	return dara.Validate(s)
}
