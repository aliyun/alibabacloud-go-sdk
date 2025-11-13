// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistInfo(v *UpdateCheckItemRequestAssistInfo) *UpdateCheckItemRequest
	GetAssistInfo() *UpdateCheckItemRequestAssistInfo
	SetCheckId(v int64) *UpdateCheckItemRequest
	GetCheckId() *int64
	SetCheckRule(v string) *UpdateCheckItemRequest
	GetCheckRule() *string
	SetCheckShowName(v string) *UpdateCheckItemRequest
	GetCheckShowName() *string
	SetDescription(v *UpdateCheckItemRequestDescription) *UpdateCheckItemRequest
	GetDescription() *UpdateCheckItemRequestDescription
	SetInstanceSubType(v string) *UpdateCheckItemRequest
	GetInstanceSubType() *string
	SetInstanceType(v string) *UpdateCheckItemRequest
	GetInstanceType() *string
	SetRemark(v string) *UpdateCheckItemRequest
	GetRemark() *string
	SetRiskLevel(v string) *UpdateCheckItemRequest
	GetRiskLevel() *string
	SetSectionIds(v []*int64) *UpdateCheckItemRequest
	GetSectionIds() []*int64
	SetSolution(v *UpdateCheckItemRequestSolution) *UpdateCheckItemRequest
	GetSolution() *UpdateCheckItemRequestSolution
	SetStatus(v string) *UpdateCheckItemRequest
	GetStatus() *string
	SetVendor(v string) *UpdateCheckItemRequest
	GetVendor() *string
}

type UpdateCheckItemRequest struct {
	// Help information for the check item.
	AssistInfo *UpdateCheckItemRequestAssistInfo `json:"AssistInfo,omitempty" xml:"AssistInfo,omitempty" type:"Struct"`
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
	Description *UpdateCheckItemRequestDescription `json:"Description,omitempty" xml:"Description,omitempty" type:"Struct"`
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
	Solution *UpdateCheckItemRequestSolution `json:"Solution,omitempty" xml:"Solution,omitempty" type:"Struct"`
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

func (s UpdateCheckItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemRequest) GetAssistInfo() *UpdateCheckItemRequestAssistInfo {
	return s.AssistInfo
}

func (s *UpdateCheckItemRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *UpdateCheckItemRequest) GetCheckRule() *string {
	return s.CheckRule
}

func (s *UpdateCheckItemRequest) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *UpdateCheckItemRequest) GetDescription() *UpdateCheckItemRequestDescription {
	return s.Description
}

func (s *UpdateCheckItemRequest) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *UpdateCheckItemRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateCheckItemRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCheckItemRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *UpdateCheckItemRequest) GetSectionIds() []*int64 {
	return s.SectionIds
}

func (s *UpdateCheckItemRequest) GetSolution() *UpdateCheckItemRequestSolution {
	return s.Solution
}

func (s *UpdateCheckItemRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCheckItemRequest) GetVendor() *string {
	return s.Vendor
}

func (s *UpdateCheckItemRequest) SetAssistInfo(v *UpdateCheckItemRequestAssistInfo) *UpdateCheckItemRequest {
	s.AssistInfo = v
	return s
}

func (s *UpdateCheckItemRequest) SetCheckId(v int64) *UpdateCheckItemRequest {
	s.CheckId = &v
	return s
}

func (s *UpdateCheckItemRequest) SetCheckRule(v string) *UpdateCheckItemRequest {
	s.CheckRule = &v
	return s
}

func (s *UpdateCheckItemRequest) SetCheckShowName(v string) *UpdateCheckItemRequest {
	s.CheckShowName = &v
	return s
}

func (s *UpdateCheckItemRequest) SetDescription(v *UpdateCheckItemRequestDescription) *UpdateCheckItemRequest {
	s.Description = v
	return s
}

func (s *UpdateCheckItemRequest) SetInstanceSubType(v string) *UpdateCheckItemRequest {
	s.InstanceSubType = &v
	return s
}

func (s *UpdateCheckItemRequest) SetInstanceType(v string) *UpdateCheckItemRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateCheckItemRequest) SetRemark(v string) *UpdateCheckItemRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCheckItemRequest) SetRiskLevel(v string) *UpdateCheckItemRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateCheckItemRequest) SetSectionIds(v []*int64) *UpdateCheckItemRequest {
	s.SectionIds = v
	return s
}

func (s *UpdateCheckItemRequest) SetSolution(v *UpdateCheckItemRequestSolution) *UpdateCheckItemRequest {
	s.Solution = v
	return s
}

func (s *UpdateCheckItemRequest) SetStatus(v string) *UpdateCheckItemRequest {
	s.Status = &v
	return s
}

func (s *UpdateCheckItemRequest) SetVendor(v string) *UpdateCheckItemRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateCheckItemRequest) Validate() error {
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

type UpdateCheckItemRequestAssistInfo struct {
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

func (s UpdateCheckItemRequestAssistInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemRequestAssistInfo) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemRequestAssistInfo) GetType() *string {
	return s.Type
}

func (s *UpdateCheckItemRequestAssistInfo) GetValue() *string {
	return s.Value
}

func (s *UpdateCheckItemRequestAssistInfo) SetType(v string) *UpdateCheckItemRequestAssistInfo {
	s.Type = &v
	return s
}

func (s *UpdateCheckItemRequestAssistInfo) SetValue(v string) *UpdateCheckItemRequestAssistInfo {
	s.Value = &v
	return s
}

func (s *UpdateCheckItemRequestAssistInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckItemRequestDescription struct {
	// Type of the check item description. Values:
	//
	// - **text**: Text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Description of the check item.
	//
	// example:
	//
	// custom description.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateCheckItemRequestDescription) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemRequestDescription) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemRequestDescription) GetType() *string {
	return s.Type
}

func (s *UpdateCheckItemRequestDescription) GetValue() *string {
	return s.Value
}

func (s *UpdateCheckItemRequestDescription) SetType(v string) *UpdateCheckItemRequestDescription {
	s.Type = &v
	return s
}

func (s *UpdateCheckItemRequestDescription) SetValue(v string) *UpdateCheckItemRequestDescription {
	s.Value = &v
	return s
}

func (s *UpdateCheckItemRequestDescription) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckItemRequestSolution struct {
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
	// custom solution.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateCheckItemRequestSolution) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemRequestSolution) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemRequestSolution) GetType() *string {
	return s.Type
}

func (s *UpdateCheckItemRequestSolution) GetValue() *string {
	return s.Value
}

func (s *UpdateCheckItemRequestSolution) SetType(v string) *UpdateCheckItemRequestSolution {
	s.Type = &v
	return s
}

func (s *UpdateCheckItemRequestSolution) SetValue(v string) *UpdateCheckItemRequestSolution {
	s.Value = &v
	return s
}

func (s *UpdateCheckItemRequestSolution) Validate() error {
	return dara.Validate(s)
}
