// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateTemplateRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateTemplateRequest
	GetAutoRenew() *bool
	SetBizType(v string) *CreateTemplateRequest
	GetBizType() *string
	SetChargeType(v string) *CreateTemplateRequest
	GetChargeType() *string
	SetDataDiskList(v []*CreateTemplateRequestDataDiskList) *CreateTemplateRequest
	GetDataDiskList() []*CreateTemplateRequestDataDiskList
	SetDefaultLanguage(v string) *CreateTemplateRequest
	GetDefaultLanguage() *string
	SetDescription(v string) *CreateTemplateRequest
	GetDescription() *string
	SetImageId(v string) *CreateTemplateRequest
	GetImageId() *string
	SetPeriod(v int32) *CreateTemplateRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateTemplateRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateTemplateRequest
	GetPolicyGroupId() *string
	SetPostPaidAfterUsedUp(v bool) *CreateTemplateRequest
	GetPostPaidAfterUsedUp() *bool
	SetProductType(v string) *CreateTemplateRequest
	GetProductType() *string
	SetRegionConfigList(v []*CreateTemplateRequestRegionConfigList) *CreateTemplateRequest
	GetRegionConfigList() []*CreateTemplateRequestRegionConfigList
	SetResourceGroupId(v string) *CreateTemplateRequest
	GetResourceGroupId() *string
	SetResourceTagList(v []*CreateTemplateRequestResourceTagList) *CreateTemplateRequest
	GetResourceTagList() []*CreateTemplateRequestResourceTagList
	SetSiteConfigList(v []*CreateTemplateRequestSiteConfigList) *CreateTemplateRequest
	GetSiteConfigList() []*CreateTemplateRequestSiteConfigList
	SetSystemDiskPerformanceLevel(v string) *CreateTemplateRequest
	GetSystemDiskPerformanceLevel() *string
	SetSystemDiskSize(v int32) *CreateTemplateRequest
	GetSystemDiskSize() *int32
	SetTemplateName(v string) *CreateTemplateRequest
	GetTemplateName() *string
	SetTimerGroupId(v string) *CreateTemplateRequest
	GetTimerGroupId() *string
	SetUserDuration(v int32) *CreateTemplateRequest
	GetUserDuration() *int32
}

type CreateTemplateRequest struct {
	AutoPay   *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	BizType      *string                              `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ChargeType   *string                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DataDiskList []*CreateTemplateRequestDataDiskList `json:"DataDiskList,omitempty" xml:"DataDiskList,omitempty" type:"Repeated"`
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// desktopimage-windows-server-2022-64-asp
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Period     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// pg-8hlryfn331******
	PolicyGroupId       *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PostPaidAfterUsedUp *bool   `json:"PostPaidAfterUsedUp,omitempty" xml:"PostPaidAfterUsedUp,omitempty"`
	// example:
	//
	// CloudDesktop
	ProductType      *string                                  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionConfigList []*CreateTemplateRequestRegionConfigList `json:"RegionConfigList,omitempty" xml:"RegionConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// rg-4knxmfneq1e******
	ResourceGroupId *string                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTagList []*CreateTemplateRequestResourceTagList `json:"ResourceTagList,omitempty" xml:"ResourceTagList,omitempty" type:"Repeated"`
	SiteConfigList  []*CreateTemplateRequestSiteConfigList  `json:"SiteConfigList,omitempty" xml:"SiteConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// AutoPL
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	UserDuration *int32  `json:"UserDuration,omitempty" xml:"UserDuration,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateTemplateRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateTemplateRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateTemplateRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTemplateRequest) GetDataDiskList() []*CreateTemplateRequestDataDiskList {
	return s.DataDiskList
}

func (s *CreateTemplateRequest) GetDefaultLanguage() *string {
	return s.DefaultLanguage
}

func (s *CreateTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateTemplateRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateTemplateRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateTemplateRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateTemplateRequest) GetPostPaidAfterUsedUp() *bool {
	return s.PostPaidAfterUsedUp
}

func (s *CreateTemplateRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateTemplateRequest) GetRegionConfigList() []*CreateTemplateRequestRegionConfigList {
	return s.RegionConfigList
}

func (s *CreateTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateRequest) GetResourceTagList() []*CreateTemplateRequestResourceTagList {
	return s.ResourceTagList
}

func (s *CreateTemplateRequest) GetSiteConfigList() []*CreateTemplateRequestSiteConfigList {
	return s.SiteConfigList
}

func (s *CreateTemplateRequest) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateTemplateRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateRequest) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *CreateTemplateRequest) GetUserDuration() *int32 {
	return s.UserDuration
}

func (s *CreateTemplateRequest) SetAutoPay(v bool) *CreateTemplateRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTemplateRequest) SetAutoRenew(v bool) *CreateTemplateRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTemplateRequest) SetBizType(v string) *CreateTemplateRequest {
	s.BizType = &v
	return s
}

func (s *CreateTemplateRequest) SetChargeType(v string) *CreateTemplateRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTemplateRequest) SetDataDiskList(v []*CreateTemplateRequestDataDiskList) *CreateTemplateRequest {
	s.DataDiskList = v
	return s
}

func (s *CreateTemplateRequest) SetDefaultLanguage(v string) *CreateTemplateRequest {
	s.DefaultLanguage = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetImageId(v string) *CreateTemplateRequest {
	s.ImageId = &v
	return s
}

func (s *CreateTemplateRequest) SetPeriod(v int32) *CreateTemplateRequest {
	s.Period = &v
	return s
}

func (s *CreateTemplateRequest) SetPeriodUnit(v string) *CreateTemplateRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateTemplateRequest) SetPolicyGroupId(v string) *CreateTemplateRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetPostPaidAfterUsedUp(v bool) *CreateTemplateRequest {
	s.PostPaidAfterUsedUp = &v
	return s
}

func (s *CreateTemplateRequest) SetProductType(v string) *CreateTemplateRequest {
	s.ProductType = &v
	return s
}

func (s *CreateTemplateRequest) SetRegionConfigList(v []*CreateTemplateRequestRegionConfigList) *CreateTemplateRequest {
	s.RegionConfigList = v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceTagList(v []*CreateTemplateRequestResourceTagList) *CreateTemplateRequest {
	s.ResourceTagList = v
	return s
}

func (s *CreateTemplateRequest) SetSiteConfigList(v []*CreateTemplateRequestSiteConfigList) *CreateTemplateRequest {
	s.SiteConfigList = v
	return s
}

func (s *CreateTemplateRequest) SetSystemDiskPerformanceLevel(v string) *CreateTemplateRequest {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateTemplateRequest) SetSystemDiskSize(v int32) *CreateTemplateRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTimerGroupId(v string) *CreateTemplateRequest {
	s.TimerGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetUserDuration(v int32) *CreateTemplateRequest {
	s.UserDuration = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
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

type CreateTemplateRequestDataDiskList struct {
	// example:
	//
	// AutoPL
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateTemplateRequestDataDiskList) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestDataDiskList) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestDataDiskList) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateTemplateRequestDataDiskList) GetSize() *int32 {
	return s.Size
}

func (s *CreateTemplateRequestDataDiskList) SetPerformanceLevel(v string) *CreateTemplateRequestDataDiskList {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateTemplateRequestDataDiskList) SetSize(v int32) *CreateTemplateRequestDataDiskList {
	s.Size = &v
	return s
}

func (s *CreateTemplateRequestDataDiskList) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateRequestRegionConfigList struct {
	// example:
	//
	// cn-hangzhou+dir-709******
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// eds.enterprise_office.8c16g
	ResourceInstanceType *string `json:"ResourceInstanceType,omitempty" xml:"ResourceInstanceType,omitempty"`
	// example:
	//
	// sp-35fvn8m21pnx2****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// example:
	//
	// vsw-bp1yiu**********
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// example:
	//
	// false
	VolumeEncryptionEnable *bool `json:"VolumeEncryptionEnable,omitempty" xml:"VolumeEncryptionEnable,omitempty"`
	// example:
	//
	// a7b3c0c8-b3a2-4876-b1cc-*********
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s CreateTemplateRequestRegionConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestRegionConfigList) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestRegionConfigList) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateTemplateRequestRegionConfigList) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateRequestRegionConfigList) GetResourceInstanceType() *string {
	return s.ResourceInstanceType
}

func (s *CreateTemplateRequestRegionConfigList) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *CreateTemplateRequestRegionConfigList) GetSubnetId() *string {
	return s.SubnetId
}

func (s *CreateTemplateRequestRegionConfigList) GetVolumeEncryptionEnable() *bool {
	return s.VolumeEncryptionEnable
}

func (s *CreateTemplateRequestRegionConfigList) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateTemplateRequestRegionConfigList) SetOfficeSiteId(v string) *CreateTemplateRequestRegionConfigList {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetRegionId(v string) *CreateTemplateRequestRegionConfigList {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetResourceInstanceType(v string) *CreateTemplateRequestRegionConfigList {
	s.ResourceInstanceType = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetSnapshotPolicyId(v string) *CreateTemplateRequestRegionConfigList {
	s.SnapshotPolicyId = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetSubnetId(v string) *CreateTemplateRequestRegionConfigList {
	s.SubnetId = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetVolumeEncryptionEnable(v bool) *CreateTemplateRequestRegionConfigList {
	s.VolumeEncryptionEnable = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) SetVolumeEncryptionKey(v string) *CreateTemplateRequestRegionConfigList {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateTemplateRequestRegionConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateRequestResourceTagList struct {
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// design
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateRequestResourceTagList) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestResourceTagList) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestResourceTagList) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateRequestResourceTagList) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateRequestResourceTagList) SetKey(v string) *CreateTemplateRequestResourceTagList {
	s.Key = &v
	return s
}

func (s *CreateTemplateRequestResourceTagList) SetValue(v string) *CreateTemplateRequestResourceTagList {
	s.Value = &v
	return s
}

func (s *CreateTemplateRequestResourceTagList) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateRequestSiteConfigList struct {
	// example:
	//
	// bwr-5a5371e0db954d********
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	// example:
	//
	// mainland
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateTemplateRequestSiteConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestSiteConfigList) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestSiteConfigList) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *CreateTemplateRequestSiteConfigList) GetSiteId() *string {
	return s.SiteId
}

func (s *CreateTemplateRequestSiteConfigList) SetAppRuleId(v string) *CreateTemplateRequestSiteConfigList {
	s.AppRuleId = &v
	return s
}

func (s *CreateTemplateRequestSiteConfigList) SetSiteId(v string) *CreateTemplateRequestSiteConfigList {
	s.SiteId = &v
	return s
}

func (s *CreateTemplateRequestSiteConfigList) Validate() error {
	return dara.Validate(s)
}
