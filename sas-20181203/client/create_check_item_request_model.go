// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistInfo(v *CreateCheckItemRequestAssistInfo) *CreateCheckItemRequest
	GetAssistInfo() *CreateCheckItemRequestAssistInfo
	SetCheckRule(v string) *CreateCheckItemRequest
	GetCheckRule() *string
	SetCheckShowName(v string) *CreateCheckItemRequest
	GetCheckShowName() *string
	SetDescription(v *CreateCheckItemRequestDescription) *CreateCheckItemRequest
	GetDescription() *CreateCheckItemRequestDescription
	SetInstanceSubType(v string) *CreateCheckItemRequest
	GetInstanceSubType() *string
	SetInstanceType(v string) *CreateCheckItemRequest
	GetInstanceType() *string
	SetRemark(v string) *CreateCheckItemRequest
	GetRemark() *string
	SetRiskLevel(v string) *CreateCheckItemRequest
	GetRiskLevel() *string
	SetSectionIds(v []*int64) *CreateCheckItemRequest
	GetSectionIds() []*int64
	SetSolution(v *CreateCheckItemRequestSolution) *CreateCheckItemRequest
	GetSolution() *CreateCheckItemRequestSolution
	SetStatus(v string) *CreateCheckItemRequest
	GetStatus() *string
	SetVendor(v string) *CreateCheckItemRequest
	GetVendor() *string
}

type CreateCheckItemRequest struct {
	// Help information for the check item.
	AssistInfo *CreateCheckItemRequestAssistInfo `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty" type:"Struct"`
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
	Description *CreateCheckItemRequestDescription `json:"Description,omitempty" xml:"Description,omitempty" type:"Struct"`
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
	Solution *CreateCheckItemRequestSolution `json:"Solution,omitempty" xml:"Solution,omitempty" type:"Struct"`
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

func (s CreateCheckItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckItemRequest) GetAssistInfo() *CreateCheckItemRequestAssistInfo {
	return s.AssistInfo
}

func (s *CreateCheckItemRequest) GetCheckRule() *string {
	return s.CheckRule
}

func (s *CreateCheckItemRequest) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *CreateCheckItemRequest) GetDescription() *CreateCheckItemRequestDescription {
	return s.Description
}

func (s *CreateCheckItemRequest) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *CreateCheckItemRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCheckItemRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCheckItemRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *CreateCheckItemRequest) GetSectionIds() []*int64 {
	return s.SectionIds
}

func (s *CreateCheckItemRequest) GetSolution() *CreateCheckItemRequestSolution {
	return s.Solution
}

func (s *CreateCheckItemRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateCheckItemRequest) GetVendor() *string {
	return s.Vendor
}

func (s *CreateCheckItemRequest) SetAssistInfo(v *CreateCheckItemRequestAssistInfo) *CreateCheckItemRequest {
	s.AssistInfo = v
	return s
}

func (s *CreateCheckItemRequest) SetCheckRule(v string) *CreateCheckItemRequest {
	s.CheckRule = &v
	return s
}

func (s *CreateCheckItemRequest) SetCheckShowName(v string) *CreateCheckItemRequest {
	s.CheckShowName = &v
	return s
}

func (s *CreateCheckItemRequest) SetDescription(v *CreateCheckItemRequestDescription) *CreateCheckItemRequest {
	s.Description = v
	return s
}

func (s *CreateCheckItemRequest) SetInstanceSubType(v string) *CreateCheckItemRequest {
	s.InstanceSubType = &v
	return s
}

func (s *CreateCheckItemRequest) SetInstanceType(v string) *CreateCheckItemRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCheckItemRequest) SetRemark(v string) *CreateCheckItemRequest {
	s.Remark = &v
	return s
}

func (s *CreateCheckItemRequest) SetRiskLevel(v string) *CreateCheckItemRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateCheckItemRequest) SetSectionIds(v []*int64) *CreateCheckItemRequest {
	s.SectionIds = v
	return s
}

func (s *CreateCheckItemRequest) SetSolution(v *CreateCheckItemRequestSolution) *CreateCheckItemRequest {
	s.Solution = v
	return s
}

func (s *CreateCheckItemRequest) SetStatus(v string) *CreateCheckItemRequest {
	s.Status = &v
	return s
}

func (s *CreateCheckItemRequest) SetVendor(v string) *CreateCheckItemRequest {
	s.Vendor = &v
	return s
}

func (s *CreateCheckItemRequest) Validate() error {
	if s.AssistInfo != nil {
		if err := s.AssistInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Description != nil {
		if err := s.Description.Validate(); err != nil {
			return err
		}
	}
	if s.Solution != nil {
		if err := s.Solution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCheckItemRequestAssistInfo struct {
	// Type of the help information for the check item risk. Values:
	//
	// - **text**: Text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Content of the help information for the check item risk.
	//
	// example:
	//
	// custom assistInfo.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCheckItemRequestAssistInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemRequestAssistInfo) GoString() string {
	return s.String()
}

func (s *CreateCheckItemRequestAssistInfo) GetType() *string {
	return s.Type
}

func (s *CreateCheckItemRequestAssistInfo) GetValue() *string {
	return s.Value
}

func (s *CreateCheckItemRequestAssistInfo) SetType(v string) *CreateCheckItemRequestAssistInfo {
	s.Type = &v
	return s
}

func (s *CreateCheckItemRequestAssistInfo) SetValue(v string) *CreateCheckItemRequestAssistInfo {
	s.Value = &v
	return s
}

func (s *CreateCheckItemRequestAssistInfo) Validate() error {
	return dara.Validate(s)
}

type CreateCheckItemRequestDescription struct {
	// Type of the check item description information. Values:
	//
	// - **text**: Text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specific content of the description.
	//
	// example:
	//
	// custom description.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCheckItemRequestDescription) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemRequestDescription) GoString() string {
	return s.String()
}

func (s *CreateCheckItemRequestDescription) GetType() *string {
	return s.Type
}

func (s *CreateCheckItemRequestDescription) GetValue() *string {
	return s.Value
}

func (s *CreateCheckItemRequestDescription) SetType(v string) *CreateCheckItemRequestDescription {
	s.Type = &v
	return s
}

func (s *CreateCheckItemRequestDescription) SetValue(v string) *CreateCheckItemRequestDescription {
	s.Value = &v
	return s
}

func (s *CreateCheckItemRequestDescription) Validate() error {
	return dara.Validate(s)
}

type CreateCheckItemRequestSolution struct {
	// Type of the solution information for the check item. Values:
	//
	// - **text**: Text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Content of the solution for the check item risk.
	//
	// example:
	//
	// text
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCheckItemRequestSolution) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemRequestSolution) GoString() string {
	return s.String()
}

func (s *CreateCheckItemRequestSolution) GetType() *string {
	return s.Type
}

func (s *CreateCheckItemRequestSolution) GetValue() *string {
	return s.Value
}

func (s *CreateCheckItemRequestSolution) SetType(v string) *CreateCheckItemRequestSolution {
	s.Type = &v
	return s
}

func (s *CreateCheckItemRequestSolution) SetValue(v string) *CreateCheckItemRequestSolution {
	s.Value = &v
	return s
}

func (s *CreateCheckItemRequestSolution) Validate() error {
	return dara.Validate(s)
}
