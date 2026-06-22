// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpaStrategyNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmDetail(v *UpdateOpaStrategyNewRequestAlarmDetail) *UpdateOpaStrategyNewRequest
	GetAlarmDetail() *UpdateOpaStrategyNewRequestAlarmDetail
	SetClusterId(v string) *UpdateOpaStrategyNewRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateOpaStrategyNewRequest
	GetClusterName() *string
	SetDescription(v string) *UpdateOpaStrategyNewRequest
	GetDescription() *string
	SetImageName(v []*string) *UpdateOpaStrategyNewRequest
	GetImageName() []*string
	SetLabel(v []*string) *UpdateOpaStrategyNewRequest
	GetLabel() []*string
	SetMaliciousImage(v bool) *UpdateOpaStrategyNewRequest
	GetMaliciousImage() *bool
	SetRuleAction(v int32) *UpdateOpaStrategyNewRequest
	GetRuleAction() *int32
	SetScopes(v []*UpdateOpaStrategyNewRequestScopes) *UpdateOpaStrategyNewRequest
	GetScopes() []*UpdateOpaStrategyNewRequestScopes
	SetStrategyId(v int64) *UpdateOpaStrategyNewRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *UpdateOpaStrategyNewRequest
	GetStrategyName() *string
	SetStrategyTemplateId(v int64) *UpdateOpaStrategyNewRequest
	GetStrategyTemplateId() *int64
	SetUnScanedImage(v bool) *UpdateOpaStrategyNewRequest
	GetUnScanedImage() *bool
	SetWhiteList(v []*string) *UpdateOpaStrategyNewRequest
	GetWhiteList() []*string
}

type UpdateOpaStrategyNewRequest struct {
	// The alert content details included in the policy.
	AlarmDetail *UpdateOpaStrategyNewRequestAlarmDetail `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// 	Notice: This parameter is deprecated. Use the Scopes parameter to specify the scope of clusters to which the policy applies..
	//
	// example:
	//
	// c870ec78ecbcb41d2a35c679823ef****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// 	Notice: This parameter is deprecated..
	//
	// example:
	//
	// docker-law
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The policy description.
	//
	// example:
	//
	// 4566
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of image names included in the policy.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The list of image tags included in the policy.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// Specifies whether the policy supports Internet malicious images. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	//
	// example:
	//
	// true
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The action to take when the policy is hit. Valid values:
	//
	// - **1**: Alert.
	//
	// - **2**: Block.
	//
	// - **3**: Allow.
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The scope of the policy.
	Scopes []*UpdateOpaStrategyNewRequestScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// >Call the [ListOpaClusterStrategyNew](https://help.aliyun.com/document_detail/2623574.html) operation to obtain this parameter.
	//
	// example:
	//
	// 1003
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The policy template ID.
	//
	// >Call the [GetOpaStrategyTemplateSummary](https://help.aliyun.com/document_detail/2539952.html) operation to obtain this parameter.
	//
	// example:
	//
	// 109
	StrategyTemplateId *int64 `json:"StrategyTemplateId,omitempty" xml:"StrategyTemplateId,omitempty"`
	// Specifies whether the policy supports unscanned images. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	//
	// example:
	//
	// true
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
	// The whitelist.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequest) GetAlarmDetail() *UpdateOpaStrategyNewRequestAlarmDetail {
	return s.AlarmDetail
}

func (s *UpdateOpaStrategyNewRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateOpaStrategyNewRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateOpaStrategyNewRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpaStrategyNewRequest) GetImageName() []*string {
	return s.ImageName
}

func (s *UpdateOpaStrategyNewRequest) GetLabel() []*string {
	return s.Label
}

func (s *UpdateOpaStrategyNewRequest) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *UpdateOpaStrategyNewRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *UpdateOpaStrategyNewRequest) GetScopes() []*UpdateOpaStrategyNewRequestScopes {
	return s.Scopes
}

func (s *UpdateOpaStrategyNewRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *UpdateOpaStrategyNewRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateOpaStrategyNewRequest) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *UpdateOpaStrategyNewRequest) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *UpdateOpaStrategyNewRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *UpdateOpaStrategyNewRequest) SetAlarmDetail(v *UpdateOpaStrategyNewRequestAlarmDetail) *UpdateOpaStrategyNewRequest {
	s.AlarmDetail = v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetClusterId(v string) *UpdateOpaStrategyNewRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetClusterName(v string) *UpdateOpaStrategyNewRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetDescription(v string) *UpdateOpaStrategyNewRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetImageName(v []*string) *UpdateOpaStrategyNewRequest {
	s.ImageName = v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetLabel(v []*string) *UpdateOpaStrategyNewRequest {
	s.Label = v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetMaliciousImage(v bool) *UpdateOpaStrategyNewRequest {
	s.MaliciousImage = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetRuleAction(v int32) *UpdateOpaStrategyNewRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetScopes(v []*UpdateOpaStrategyNewRequestScopes) *UpdateOpaStrategyNewRequest {
	s.Scopes = v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetStrategyId(v int64) *UpdateOpaStrategyNewRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetStrategyName(v string) *UpdateOpaStrategyNewRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetStrategyTemplateId(v int64) *UpdateOpaStrategyNewRequest {
	s.StrategyTemplateId = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetUnScanedImage(v bool) *UpdateOpaStrategyNewRequest {
	s.UnScanedImage = &v
	return s
}

func (s *UpdateOpaStrategyNewRequest) SetWhiteList(v []*string) *UpdateOpaStrategyNewRequest {
	s.WhiteList = v
	return s
}

func (s *UpdateOpaStrategyNewRequest) Validate() error {
	if s.AlarmDetail != nil {
		if err := s.AlarmDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Scopes != nil {
		for _, item := range s.Scopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetail struct {
	// The baseline risk information.
	Baseline *UpdateOpaStrategyNewRequestAlarmDetailBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	// The risky image build command information.
	BuildRisk *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk `json:"BuildRisk,omitempty" xml:"BuildRisk,omitempty" type:"Struct"`
	// The malicious sample information.
	MaliciousFile *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty" type:"Struct"`
	// The sensitive file information.
	SensitiveFile *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile `json:"SensitiveFile,omitempty" xml:"SensitiveFile,omitempty" type:"Struct"`
	// The vulnerability item information.
	Vul *UpdateOpaStrategyNewRequestAlarmDetailVul `json:"Vul,omitempty" xml:"Vul,omitempty" type:"Struct"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetail) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) GetBaseline() *UpdateOpaStrategyNewRequestAlarmDetailBaseline {
	return s.Baseline
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) GetBuildRisk() *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk {
	return s.BuildRisk
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) GetMaliciousFile() *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	return s.MaliciousFile
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) GetSensitiveFile() *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	return s.SensitiveFile
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) GetVul() *UpdateOpaStrategyNewRequestAlarmDetailVul {
	return s.Vul
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) SetBaseline(v *UpdateOpaStrategyNewRequestAlarmDetailBaseline) *UpdateOpaStrategyNewRequestAlarmDetail {
	s.Baseline = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) SetBuildRisk(v *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) *UpdateOpaStrategyNewRequestAlarmDetail {
	s.BuildRisk = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) SetMaliciousFile(v *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) *UpdateOpaStrategyNewRequestAlarmDetail {
	s.MaliciousFile = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) SetSensitiveFile(v *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) *UpdateOpaStrategyNewRequestAlarmDetail {
	s.SensitiveFile = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) SetVul(v *UpdateOpaStrategyNewRequestAlarmDetailVul) *UpdateOpaStrategyNewRequestAlarmDetail {
	s.Vul = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetail) Validate() error {
	if s.Baseline != nil {
		if err := s.Baseline.Validate(); err != nil {
			return err
		}
	}
	if s.BuildRisk != nil {
		if err := s.BuildRisk.Validate(); err != nil {
			return err
		}
	}
	if s.MaliciousFile != nil {
		if err := s.MaliciousFile.Validate(); err != nil {
			return err
		}
	}
	if s.SensitiveFile != nil {
		if err := s.SensitiveFile.Validate(); err != nil {
			return err
		}
	}
	if s.Vul != nil {
		if err := s.Vul.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailBaseline struct {
	// The baseline item information.
	Item []*UpdateOpaStrategyNewRequestAlarmDetailBaselineItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBaseline) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBaseline) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaseline) GetItem() []*UpdateOpaStrategyNewRequestAlarmDetailBaselineItem {
	return s.Item
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaseline) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaseline) SetItem(v []*UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) *UpdateOpaStrategyNewRequestAlarmDetailBaseline {
	s.Item = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaseline) SetRiskLevel(v []*string) *UpdateOpaStrategyNewRequestAlarmDetailBaseline {
	s.RiskLevel = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaseline) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailBaselineItem struct {
	// The baseline check item ID.
	//
	// > Call the [GetOpaClusterBaseLineList](https://help.aliyun.com/document_detail/2539883.html) operation to query this value.
	//
	// example:
	//
	// hc.image.checklist.identify.hc_exploit_couchdb_linux.item
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The baseline check item name.
	//
	// > Call the [GetOpaClusterBaseLineList](https://help.aliyun.com/document_detail/2539883.html) operation to query this value.
	//
	// example:
	//
	// passwd
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBaselineItem) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestAlarmDetailBuildRisk struct {
	// The list of risk items.
	Item []*UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) GetItem() []*UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	return s.Item
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) SetItem(v []*UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk {
	s.Item = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) SetRiskLevel(v []*string) *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk {
	s.RiskLevel = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRisk) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem struct {
	// The risky image build command type ID.
	//
	// > Call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query this value.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The risky image build command type name.
	//
	// > Call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query this value.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailBuildRiskItem) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile struct {
	// The malicious sample list.
	Item []*UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) GetItem() []*UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	return s.Item
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) SetItem(v []*UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	s.Item = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) SetRiskLevel(v []*string) *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	s.RiskLevel = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFile) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem struct {
	// The malicious sample ID.
	//
	// > Call the [DescribeMatchedMaliciousNames](~~DescribeMatchedMaliciousNames~~) operation to query this value.
	//
	// example:
	//
	// 65201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The malicious sample name.
	//
	// > Call the [DescribeMatchedMaliciousNames](~~DescribeMatchedMaliciousNames~~) operation to query this value.
	//
	// example:
	//
	// abnormal binary file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile struct {
	// The list of sensitive file check items.
	Item []*UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) GetItem() []*UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	return s.Item
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) SetItem(v []*UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	s.Item = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) SetRiskLevel(v []*string) *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	s.RiskLevel = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFile) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem struct {
	// The sensitive file type ID.
	//
	// > Call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query this value.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sensitive file type name.
	//
	// > Call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query this value.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestAlarmDetailVul struct {
	// The vulnerability list.
	Item []*UpdateOpaStrategyNewRequestAlarmDetailVulItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The list of vulnerability types.
	RiskClass []*UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass `json:"RiskClass,omitempty" xml:"RiskClass,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVul) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVul) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) GetItem() []*UpdateOpaStrategyNewRequestAlarmDetailVulItem {
	return s.Item
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) GetRiskClass() []*UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	return s.RiskClass
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) SetItem(v []*UpdateOpaStrategyNewRequestAlarmDetailVulItem) *UpdateOpaStrategyNewRequestAlarmDetailVul {
	s.Item = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) SetRiskClass(v []*UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) *UpdateOpaStrategyNewRequestAlarmDetailVul {
	s.RiskClass = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) SetRiskLevel(v []*string) *UpdateOpaStrategyNewRequestAlarmDetailVul {
	s.RiskLevel = v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVul) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskClass != nil {
		for _, item := range s.RiskClass {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewRequestAlarmDetailVulItem struct {
	// The vulnerability ID.
	//
	// > Call the [DescribeVulListPage](https://help.aliyun.com/document_detail/471928.html) operation to query this value.
	//
	// example:
	//
	// CVE-2023-36034
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The vulnerability name.
	//
	// > Call the [DescribeVulListPage](https://help.aliyun.com/document_detail/471928.html) operation to query this value.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20227002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVulItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVulItem) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulItem) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulItem) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulItem) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailVulItem {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulItem) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailVulItem {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulItem) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass struct {
	// The vulnerability type ID. Valid values:
	//
	// - **cve**: system vulnerability
	//
	// - **app**: application vulnerability.
	//
	// example:
	//
	// cve
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The vulnerability type name. Valid values:
	//
	// - **system vulnerability**
	//
	// - **application vulnerability**
	//
	// example:
	//
	// 系统漏洞
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) GetId() *string {
	return s.Id
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) GetName() *string {
	return s.Name
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) SetId(v string) *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	s.Id = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) SetName(v string) *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	s.Name = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestAlarmDetailVulRiskClass) Validate() error {
	return dara.Validate(s)
}

type UpdateOpaStrategyNewRequestScopes struct {
	// The policy instance ID in the cluster.
	//
	// > Call the [GetOpaStrategyDetailNew](~~GetOpaStrategyDetailNew~~) operation to obtain this parameter.
	//
	// example:
	//
	// ack-1
	AckPolicyInstanceId *string `json:"AckPolicyInstanceId,omitempty" xml:"AckPolicyInstanceId,omitempty"`
	// Specifies whether all namespaces are included. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the container cluster.
	//
	// > Call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// cdcb56a931c**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of namespaces.
	//
	// 	Notice: This parameter takes effect only when AllNamespace is set to 0..
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewRequestScopes) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewRequestScopes) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewRequestScopes) GetAckPolicyInstanceId() *string {
	return s.AckPolicyInstanceId
}

func (s *UpdateOpaStrategyNewRequestScopes) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *UpdateOpaStrategyNewRequestScopes) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateOpaStrategyNewRequestScopes) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *UpdateOpaStrategyNewRequestScopes) SetAckPolicyInstanceId(v string) *UpdateOpaStrategyNewRequestScopes {
	s.AckPolicyInstanceId = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestScopes) SetAllNamespace(v int32) *UpdateOpaStrategyNewRequestScopes {
	s.AllNamespace = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestScopes) SetClusterId(v string) *UpdateOpaStrategyNewRequestScopes {
	s.ClusterId = &v
	return s
}

func (s *UpdateOpaStrategyNewRequestScopes) SetNamespaceList(v []*string) *UpdateOpaStrategyNewRequestScopes {
	s.NamespaceList = v
	return s
}

func (s *UpdateOpaStrategyNewRequestScopes) Validate() error {
	return dara.Validate(s)
}
