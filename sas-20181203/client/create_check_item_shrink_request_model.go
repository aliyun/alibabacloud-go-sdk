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
	// The help information of the check item.
	AssistInfoShrink *string `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty"`
	// The rule definition of the custom check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AssociatedData":{"ToDataList":[{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","FromPropertyPath":"InstanceId"}]},"MatchProperty":{"Operator":"AND","MatchProperties":[{"DataName":"ACS_ECS_Disk","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testId"},{"DataName":"ACS_ECS_Instance","PropertyPath":"InstanceId","MatchOperator":"EQ","MatchPropertyValue":"testInstanceId"}]}}
	CheckRule *string `json:"CheckRule,omitempty" xml:"CheckRule,omitempty"`
	// The name of the custom check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// testCheckItemName
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// The description of the check item.
	DescriptionShrink *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The asset subtype of the cloud service.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset type of the cloud service.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// - **HIGH**: High risk.
	//
	// - **MEDIUM**: Medium risk.
	//
	// - **LOW**: Low risk.
	//
	// This parameter is required.
	//
	// example:
	//
	// LOW
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The IDs of the sections associated with the check item.
	//
	// This parameter is required.
	SectionIds []*int64 `json:"SectionIds,omitempty" xml:"SectionIds,omitempty" type:"Repeated"`
	// The solution information of the check item.
	SolutionShrink *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The status of the check item. Valid values:
	//
	// - **EDIT**: Being edited.
	//
	// - **RELEASE**: Published.
	//
	// > - Changing the status from **Published*	- to **Being edited*	- will purge all historical records.
	//
	// > - Only check items in the **Published*	- status can be used for checks.
	//
	// This parameter is required.
	//
	// example:
	//
	// EDIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud asset vendor.
	//
	// > You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain the available vendors.
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
