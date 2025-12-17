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
	// The default language of the cloud computer during startup. This parameter takes effect only when the cloud computer is created from a system image.
	//
	// Valid values:
	//
	// 	- en-US: English.
	//
	// 	- zh-HK: Chinese, Traditional (Hong Kong, China).
	//
	// 	- zh-CN: Simplified Chinese.
	//
	// 	- ja-JP: Japanese.
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	// The template description. It must meet the following criteria:
	//
	// 	- It can be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// 	- It can contain letters, digits, and special characters, including spaces. Note: You can use carriage returns to break lines.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud computer image. You can query image IDs on the Images page. System images and custom images are supported.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Period     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy group.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId       *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PostPaidAfterUsedUp *bool   `json:"PostPaidAfterUsedUp,omitempty" xml:"PostPaidAfterUsedUp,omitempty"`
	// The regions by which you can search for cloud computer templates. When this parameter takes effect, cloud computer templates are matched based on the specified regions.
	//
	// >  You can specify up to 20 regions.
	RegionConfigList []*ModifyTemplateRequestRegionConfigList `json:"RegionConfigList,omitempty" xml:"RegionConfigList,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-a5fqjjqaejt***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to cloud computers. Specify tags in key-value pairs. You can specify up to 20 tags.
	ResourceTagList []*ModifyTemplateRequestResourceTagList `json:"ResourceTagList,omitempty" xml:"ResourceTagList,omitempty" type:"Repeated"`
	SiteConfigList  []*ModifyTemplateRequestSiteConfigList  `json:"SiteConfigList,omitempty" xml:"SiteConfigList,omitempty" type:"Repeated"`
	// The performance level (PL) of the system disk.
	//
	// >  Only cloud computers of the Graphics or High Frequency type support Enterprise SSDs (ESSDs).
	//
	// Valid values:
	//
	// 	- PL1: a PL1 ESSD.
	//
	// 	- PL0: a PL0 ESSD.
	//
	// 	- AutoPL: an AutoPL ESSD.
	//
	// example:
	//
	// AutoPL
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values: 40 to 500. Increments: 10 GiB.
	//
	// >  The system disk size must be at least as large as the configured image size.
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
	// The template name. It must meet the following criteria:
	//
	// 	- It can be 2 to 126 characters in length.
	//
	// 	- It must begin with a letter and cannot start with `http://` or `https://`.
	//
	// 	- It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Note: Periods (.) are not supported in the name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the scheduled task group.
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
	// The PL of the data disk. Default value: `AutoPL`.
	//
	// Valid values:
	//
	// 	- PL1: a PL1 ESSD
	//
	// 	- PL0: a PL0 ESSD
	//
	// 	- AutoPL: an AutoPL ESSD
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB.Valid range: 40 to 2040 GiB with an increment of 10 GiB.
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
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-709****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud computer type.
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
	// The subnet ID.
	//
	// example:
	//
	// vsw-adjrehad1****
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// Valid values:
	//
	// 	- false (default): disables disk encryption.
	//
	// 	- true: enables disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnable *bool `json:"VolumeEncryptionEnable,omitempty" xml:"VolumeEncryptionEnable,omitempty"`
	// The ID of the Key Management Service (KMS) key that you want to use to encrypt disks.
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
