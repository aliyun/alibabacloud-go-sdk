// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistInfoShrink(v string) *CreateCheckItemShrinkRequest
	GetAssistInfoShrink() *string
	SetCheckRule(v string) *CreateCheckItemShrinkRequest
	GetCheckRule() *string
	SetCheckShowName(v string) *CreateCheckItemShrinkRequest
	GetCheckShowName() *string
	SetDescriptionShrink(v string) *CreateCheckItemShrinkRequest
	GetDescriptionShrink() *string
	SetInstanceSubType(v string) *CreateCheckItemShrinkRequest
	GetInstanceSubType() *string
	SetInstanceType(v string) *CreateCheckItemShrinkRequest
	GetInstanceType() *string
	SetRemark(v string) *CreateCheckItemShrinkRequest
	GetRemark() *string
	SetRiskLevel(v string) *CreateCheckItemShrinkRequest
	GetRiskLevel() *string
	SetSectionIds(v []*int64) *CreateCheckItemShrinkRequest
	GetSectionIds() []*int64
	SetSolutionShrink(v string) *CreateCheckItemShrinkRequest
	GetSolutionShrink() *string
	SetStatus(v string) *CreateCheckItemShrinkRequest
	GetStatus() *string
	SetVendor(v string) *CreateCheckItemShrinkRequest
	GetVendor() *string
}

type CreateCheckItemShrinkRequest struct {
	// Help information for the check item.
	AssistInfoShrink *string `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty"`
	// Definition rule for the custom check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AssociatedData":{"ToDataList":[{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","FromPropertyPath":"InstanceId"}]},"MatchProperty":{"Operator":"AND","MatchProperties":[{"DataName":"ACS_ECS_Disk","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testId"},{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testInstanceId"}]}}
	CheckRule *string `json:"CheckRule,omitempty" xml:"CheckRule,omitempty"`
	// Name of the custom check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// testCheckItemName
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// Description information of the check item.
	DescriptionShrink *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Sub-asset type of the cloud product.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// Asset type of the cloud product.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Remark information.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Risk level of the check item. Values:
	//
	// - **HIGH**: High risk
	//
	// - **MEDIUM**: Medium risk
	//
	// - **LOW**: Low risk
	//
	// This parameter is required.
	//
	// example:
	//
	// LOW
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Array of section IDs associated with the check item.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// EDIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Cloud asset vendor.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) API to get the available vendors.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateCheckItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckItemShrinkRequest) GetAssistInfoShrink() *string {
	return s.AssistInfoShrink
}

func (s *CreateCheckItemShrinkRequest) GetCheckRule() *string {
	return s.CheckRule
}

func (s *CreateCheckItemShrinkRequest) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *CreateCheckItemShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *CreateCheckItemShrinkRequest) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *CreateCheckItemShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCheckItemShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCheckItemShrinkRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *CreateCheckItemShrinkRequest) GetSectionIds() []*int64 {
	return s.SectionIds
}

func (s *CreateCheckItemShrinkRequest) GetSolutionShrink() *string {
	return s.SolutionShrink
}

func (s *CreateCheckItemShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateCheckItemShrinkRequest) GetVendor() *string {
	return s.Vendor
}

func (s *CreateCheckItemShrinkRequest) SetAssistInfoShrink(v string) *CreateCheckItemShrinkRequest {
	s.AssistInfoShrink = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetCheckRule(v string) *CreateCheckItemShrinkRequest {
	s.CheckRule = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetCheckShowName(v string) *CreateCheckItemShrinkRequest {
	s.CheckShowName = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetDescriptionShrink(v string) *CreateCheckItemShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetInstanceSubType(v string) *CreateCheckItemShrinkRequest {
	s.InstanceSubType = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetInstanceType(v string) *CreateCheckItemShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetRemark(v string) *CreateCheckItemShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetRiskLevel(v string) *CreateCheckItemShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetSectionIds(v []*int64) *CreateCheckItemShrinkRequest {
	s.SectionIds = v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetSolutionShrink(v string) *CreateCheckItemShrinkRequest {
	s.SolutionShrink = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetStatus(v string) *CreateCheckItemShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) SetVendor(v string) *CreateCheckItemShrinkRequest {
	s.Vendor = &v
	return s
}

func (s *CreateCheckItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
