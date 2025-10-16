// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTemplatesResponseBody
	GetCode() *string
	SetData(v []*DescribeTemplatesResponseBodyData) *DescribeTemplatesResponseBody
	GetData() []*DescribeTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *DescribeTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTemplatesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTemplatesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeTemplatesResponseBody
	GetTotalCount() *int32
}

type DescribeTemplatesResponseBody struct {
	// example:
	//
	// success
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1871984F-51F6-5588-BAF6-*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 94
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTemplatesResponseBody) GetData() []*DescribeTemplatesResponseBodyData {
	return s.Data
}

func (s *DescribeTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTemplatesResponseBody) SetCode(v string) *DescribeTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetData(v []*DescribeTemplatesResponseBodyData) *DescribeTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetHttpStatusCode(v int32) *DescribeTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetMessage(v string) *DescribeTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageNumber(v int32) *DescribeTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageSize(v int32) *DescribeTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetRequestId(v string) *DescribeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetSuccess(v bool) *DescribeTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTotalCount(v int32) *DescribeTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplatesResponseBodyData struct {
	AutoPay      *bool                                            `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew    *bool                                            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType   *string                                          `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DataDiskList []*DescribeTemplatesResponseBodyDataDataDiskList `json:"DataDiskList,omitempty" xml:"DataDiskList,omitempty" type:"Repeated"`
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2025-04-25T05:18:46.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-04-25T05:18:46.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// m-5q8ehbihx*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// User
	ImageType  *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Period     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// pg-0caoeogkhz*****
	PolicyGroupId       *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PostPaidAfterUsedUp *bool   `json:"PostPaidAfterUsedUp,omitempty" xml:"PostPaidAfterUsedUp,omitempty"`
	// example:
	//
	// CLOUD_DESKTOP
	ProductType      *string                                              `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionConfigList []*DescribeTemplatesResponseBodyDataRegionConfigList `json:"RegionConfigList,omitempty" xml:"RegionConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 1871984F-51F6-5588-BAF6-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-a5fqjjqaejt***
	ResourceGroupId *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTagList []*DescribeTemplatesResponseBodyDataResourceTagList `json:"ResourceTagList,omitempty" xml:"ResourceTagList,omitempty" type:"Repeated"`
	SiteConfigList  []*DescribeTemplatesResponseBodyDataSiteConfigList  `json:"SiteConfigList,omitempty" xml:"SiteConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// AutoPL
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// b-0caoeogs88y*****
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// USER_TEMPLATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// bcc-dweha*****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	UserDuration *string `json:"UserDuration,omitempty" xml:"UserDuration,omitempty"`
}

func (s DescribeTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyData) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *DescribeTemplatesResponseBodyData) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeTemplatesResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTemplatesResponseBodyData) GetDataDiskList() []*DescribeTemplatesResponseBodyDataDataDiskList {
	return s.DataDiskList
}

func (s *DescribeTemplatesResponseBodyData) GetDefaultLanguage() *string {
	return s.DefaultLanguage
}

func (s *DescribeTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplatesResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeTemplatesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeTemplatesResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeTemplatesResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeTemplatesResponseBodyData) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeTemplatesResponseBodyData) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeTemplatesResponseBodyData) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeTemplatesResponseBodyData) GetPostPaidAfterUsedUp() *bool {
	return s.PostPaidAfterUsedUp
}

func (s *DescribeTemplatesResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeTemplatesResponseBodyData) GetRegionConfigList() []*DescribeTemplatesResponseBodyDataRegionConfigList {
	return s.RegionConfigList
}

func (s *DescribeTemplatesResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplatesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTemplatesResponseBodyData) GetResourceTagList() []*DescribeTemplatesResponseBodyDataResourceTagList {
	return s.ResourceTagList
}

func (s *DescribeTemplatesResponseBodyData) GetSiteConfigList() []*DescribeTemplatesResponseBodyDataSiteConfigList {
	return s.SiteConfigList
}

func (s *DescribeTemplatesResponseBodyData) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeTemplatesResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeTemplatesResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeTemplatesResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeTemplatesResponseBodyData) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatesResponseBodyData) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *DescribeTemplatesResponseBodyData) GetUserDuration() *string {
	return s.UserDuration
}

func (s *DescribeTemplatesResponseBodyData) SetAutoPay(v bool) *DescribeTemplatesResponseBodyData {
	s.AutoPay = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetAutoRenew(v bool) *DescribeTemplatesResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetChargeType(v string) *DescribeTemplatesResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetDataDiskList(v []*DescribeTemplatesResponseBodyDataDataDiskList) *DescribeTemplatesResponseBodyData {
	s.DataDiskList = v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetDefaultLanguage(v string) *DescribeTemplatesResponseBodyData {
	s.DefaultLanguage = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetDescription(v string) *DescribeTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetGmtCreate(v string) *DescribeTemplatesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetGmtModified(v string) *DescribeTemplatesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetImageId(v string) *DescribeTemplatesResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetImageType(v string) *DescribeTemplatesResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetPeriod(v int32) *DescribeTemplatesResponseBodyData {
	s.Period = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetPeriodUnit(v string) *DescribeTemplatesResponseBodyData {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetPolicyGroupId(v string) *DescribeTemplatesResponseBodyData {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetPostPaidAfterUsedUp(v bool) *DescribeTemplatesResponseBodyData {
	s.PostPaidAfterUsedUp = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetProductType(v string) *DescribeTemplatesResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetRegionConfigList(v []*DescribeTemplatesResponseBodyDataRegionConfigList) *DescribeTemplatesResponseBodyData {
	s.RegionConfigList = v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetRequestId(v string) *DescribeTemplatesResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetResourceGroupId(v string) *DescribeTemplatesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetResourceTagList(v []*DescribeTemplatesResponseBodyDataResourceTagList) *DescribeTemplatesResponseBodyData {
	s.ResourceTagList = v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetSiteConfigList(v []*DescribeTemplatesResponseBodyDataSiteConfigList) *DescribeTemplatesResponseBodyData {
	s.SiteConfigList = v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetSystemDiskPerformanceLevel(v string) *DescribeTemplatesResponseBodyData {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetSystemDiskSize(v int32) *DescribeTemplatesResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetTemplateId(v string) *DescribeTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetTemplateName(v string) *DescribeTemplatesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetTemplateType(v string) *DescribeTemplatesResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetTimerGroupId(v string) *DescribeTemplatesResponseBodyData {
	s.TimerGroupId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) SetUserDuration(v string) *DescribeTemplatesResponseBodyData {
	s.UserDuration = &v
	return s
}

func (s *DescribeTemplatesResponseBodyData) Validate() error {
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

type DescribeTemplatesResponseBodyDataDataDiskList struct {
	// example:
	//
	// AutoPL
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTemplatesResponseBodyDataDataDiskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyDataDataDiskList) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyDataDataDiskList) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeTemplatesResponseBodyDataDataDiskList) GetSize() *string {
	return s.Size
}

func (s *DescribeTemplatesResponseBodyDataDataDiskList) SetPerformanceLevel(v string) *DescribeTemplatesResponseBodyDataDataDiskList {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataDataDiskList) SetSize(v string) *DescribeTemplatesResponseBodyDataDataDiskList {
	s.Size = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataDataDiskList) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplatesResponseBodyDataRegionConfigList struct {
	// example:
	//
	// 4
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// example:
	//
	// 4GiB
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// example:
	//
	// 8192
	MemorySize *int64 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// example:
	//
	// cn-beijing+dir-3040*****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// eds.enterprise_office.4c8g
	ResourceInstanceType *string `json:"ResourceInstanceType,omitempty" xml:"ResourceInstanceType,omitempty"`
	// example:
	//
	// sp-b9fasjuu0*****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// example:
	//
	// vsw-dgea1*****
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// example:
	//
	// false
	VolumeEncryptionEnable *bool `json:"VolumeEncryptionEnable,omitempty" xml:"VolumeEncryptionEnable,omitempty"`
	// example:
	//
	// 3bc77be0-cbce-4a29-b07b-13f16394****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s DescribeTemplatesResponseBodyDataRegionConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyDataRegionConfigList) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetMemorySize() *int64 {
	return s.MemorySize
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetResourceInstanceType() *string {
	return s.ResourceInstanceType
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetVolumeEncryptionEnable() *bool {
	return s.VolumeEncryptionEnable
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetCpuCount(v int32) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.CpuCount = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetGpuSpec(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.GpuSpec = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetMemorySize(v int64) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.MemorySize = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetOfficeSiteId(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetRegionId(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.RegionId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetResourceInstanceType(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.ResourceInstanceType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetSnapshotPolicyId(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.SnapshotPolicyId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetSubnetId(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.SubnetId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetVolumeEncryptionEnable(v bool) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.VolumeEncryptionEnable = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) SetVolumeEncryptionKey(v string) *DescribeTemplatesResponseBodyDataRegionConfigList {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataRegionConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplatesResponseBodyDataResourceTagList struct {
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTemplatesResponseBodyDataResourceTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyDataResourceTagList) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyDataResourceTagList) GetKey() *string {
	return s.Key
}

func (s *DescribeTemplatesResponseBodyDataResourceTagList) GetValue() *string {
	return s.Value
}

func (s *DescribeTemplatesResponseBodyDataResourceTagList) SetKey(v string) *DescribeTemplatesResponseBodyDataResourceTagList {
	s.Key = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataResourceTagList) SetValue(v string) *DescribeTemplatesResponseBodyDataResourceTagList {
	s.Value = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataResourceTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplatesResponseBodyDataSiteConfigList struct {
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	SiteId    *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeTemplatesResponseBodyDataSiteConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyDataSiteConfigList) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyDataSiteConfigList) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *DescribeTemplatesResponseBodyDataSiteConfigList) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeTemplatesResponseBodyDataSiteConfigList) SetAppRuleId(v string) *DescribeTemplatesResponseBodyDataSiteConfigList {
	s.AppRuleId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataSiteConfigList) SetSiteId(v string) *DescribeTemplatesResponseBodyDataSiteConfigList {
	s.SiteId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyDataSiteConfigList) Validate() error {
	return dara.Validate(s)
}
