// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistInfoShrink(v string) *UpdateCheckItemShrinkRequest
	GetAssistInfoShrink() *string
	SetCheckId(v int64) *UpdateCheckItemShrinkRequest
	GetCheckId() *int64
	SetCheckRule(v string) *UpdateCheckItemShrinkRequest
	GetCheckRule() *string
	SetCheckShowName(v string) *UpdateCheckItemShrinkRequest
	GetCheckShowName() *string
	SetDescriptionShrink(v string) *UpdateCheckItemShrinkRequest
	GetDescriptionShrink() *string
	SetInstanceSubType(v string) *UpdateCheckItemShrinkRequest
	GetInstanceSubType() *string
	SetInstanceType(v string) *UpdateCheckItemShrinkRequest
	GetInstanceType() *string
	SetRemark(v string) *UpdateCheckItemShrinkRequest
	GetRemark() *string
	SetRiskLevel(v string) *UpdateCheckItemShrinkRequest
	GetRiskLevel() *string
	SetSectionIds(v []*int64) *UpdateCheckItemShrinkRequest
	GetSectionIds() []*int64
	SetSolutionShrink(v string) *UpdateCheckItemShrinkRequest
	GetSolutionShrink() *string
	SetStatus(v string) *UpdateCheckItemShrinkRequest
	GetStatus() *string
	SetVendor(v string) *UpdateCheckItemShrinkRequest
	GetVendor() *string
}

type UpdateCheckItemShrinkRequest struct {
	// Help information for the check item.
	AssistInfoShrink *string `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty"`
	// ID of the custom check item to be updated.
	//
	// > You can call the [ListCheckItems](~~ListCheckItems~~) API to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000000001
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// Definition rule for the custom check item.
	//
	// example:
	//
	// {"AssociatedData":{"ToDataList":[{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","FromPropertyPath":"InstanceId"}]},"MatchProperty":{"Operator":"AND","MatchProperties":[{"DataName":"ACS_ECS_Disk","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testId"},{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testInstanceId"}]}}
	CheckRule *string `json:"CheckRule,omitempty" xml:"CheckRule,omitempty"`
	// Name of the custom check item.
	//
	// example:
	//
	// testCheckItemName
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// Description of the check item.
	DescriptionShrink *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Sub-asset type of the cloud product.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get this parameter.
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// Asset type of the cloud product.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get this parameter.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Remark information
	//
	// example:
	//
	// remark.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Risk level of the check item. Values:
	//
	// - **HIGH**: High risk
	//
	// - **MEDIUM**: Medium risk
	//
	// - **LOW**: Low risk
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Array of section IDs associated with the check item.
	SectionIds []*int64 `json:"SectionIds,omitempty" xml:"SectionIds,omitempty" type:"Repeated"`
	// Solution information for the check item.
	SolutionShrink *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// Status of the check item. Values:
	//
	// - **EDIT**: In editing
	//
	// - **RELEASE**: Released
	//
	// > - Changing from **Released*	- to **In editing*	- will clear all historical records
	//
	// > - Only the **Released*	- status allows the use of the check item for inspection.
	//
	// example:
	//
	// RELEASE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Cloud asset vendor.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get the available vendors.
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s UpdateCheckItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemShrinkRequest) GetAssistInfoShrink() *string {
	return s.AssistInfoShrink
}

func (s *UpdateCheckItemShrinkRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *UpdateCheckItemShrinkRequest) GetCheckRule() *string {
	return s.CheckRule
}

func (s *UpdateCheckItemShrinkRequest) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *UpdateCheckItemShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *UpdateCheckItemShrinkRequest) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *UpdateCheckItemShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateCheckItemShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCheckItemShrinkRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *UpdateCheckItemShrinkRequest) GetSectionIds() []*int64 {
	return s.SectionIds
}

func (s *UpdateCheckItemShrinkRequest) GetSolutionShrink() *string {
	return s.SolutionShrink
}

func (s *UpdateCheckItemShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCheckItemShrinkRequest) GetVendor() *string {
	return s.Vendor
}

func (s *UpdateCheckItemShrinkRequest) SetAssistInfoShrink(v string) *UpdateCheckItemShrinkRequest {
	s.AssistInfoShrink = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetCheckId(v int64) *UpdateCheckItemShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetCheckRule(v string) *UpdateCheckItemShrinkRequest {
	s.CheckRule = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetCheckShowName(v string) *UpdateCheckItemShrinkRequest {
	s.CheckShowName = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetDescriptionShrink(v string) *UpdateCheckItemShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetInstanceSubType(v string) *UpdateCheckItemShrinkRequest {
	s.InstanceSubType = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetInstanceType(v string) *UpdateCheckItemShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetRemark(v string) *UpdateCheckItemShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetRiskLevel(v string) *UpdateCheckItemShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetSectionIds(v []*int64) *UpdateCheckItemShrinkRequest {
	s.SectionIds = v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetSolutionShrink(v string) *UpdateCheckItemShrinkRequest {
	s.SolutionShrink = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetStatus(v string) *UpdateCheckItemShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) SetVendor(v string) *UpdateCheckItemShrinkRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateCheckItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
