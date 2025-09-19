// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaStrategyNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmDetail(v *CreateOpaStrategyNewRequestAlarmDetail) *CreateOpaStrategyNewRequest
	GetAlarmDetail() *CreateOpaStrategyNewRequestAlarmDetail
	SetClusterId(v string) *CreateOpaStrategyNewRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateOpaStrategyNewRequest
	GetClusterName() *string
	SetDescription(v string) *CreateOpaStrategyNewRequest
	GetDescription() *string
	SetImageName(v []*string) *CreateOpaStrategyNewRequest
	GetImageName() []*string
	SetLabel(v []*string) *CreateOpaStrategyNewRequest
	GetLabel() []*string
	SetMaliciousImage(v bool) *CreateOpaStrategyNewRequest
	GetMaliciousImage() *bool
	SetRuleAction(v int32) *CreateOpaStrategyNewRequest
	GetRuleAction() *int32
	SetScopes(v []*CreateOpaStrategyNewRequestScopes) *CreateOpaStrategyNewRequest
	GetScopes() []*CreateOpaStrategyNewRequestScopes
	SetStrategyId(v int64) *CreateOpaStrategyNewRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *CreateOpaStrategyNewRequest
	GetStrategyName() *string
	SetStrategyTemplateId(v int64) *CreateOpaStrategyNewRequest
	GetStrategyTemplateId() *int64
	SetUnScanedImage(v bool) *CreateOpaStrategyNewRequest
	GetUnScanedImage() *bool
	SetWhiteList(v []*string) *CreateOpaStrategyNewRequest
	GetWhiteList() []*string
}

type CreateOpaStrategyNewRequest struct {
	// The risks that you want to detect by using the rule.
	AlarmDetail *CreateOpaStrategyNewRequestAlarmDetail `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// cfa7e2fb8c221483ba59e098c34c6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// *
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The rule description.
	//
	// example:
	//
	// default policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The container tags.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// Specifies whether the rule supports malicious Internet images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The action that is performed when the rule is hit. Valid values:
	//
	// 	- **1**: trigger alerts
	//
	// 	- **2**: block
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The application scope of the rule.
	Scopes []*CreateOpaStrategyNewRequestScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// >  You can call the [ListOpaClusterStrategyNew](https://help.aliyun.com/document_detail/2623574.html) operation to query the rule ID.
	//
	// > This parameter is invalid when you create a rule.
	//
	// example:
	//
	// 16
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// default
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The ID of the rule template.
	//
	// >  You can call the [GetOpaStrategyTemplateSummary](https://help.aliyun.com/document_detail/2539952.html) operation to query the ID of the rule template.
	//
	// example:
	//
	// 109
	StrategyTemplateId *int64 `json:"StrategyTemplateId,omitempty" xml:"StrategyTemplateId,omitempty"`
	// Specifies whether the rule supports unscanned images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
	// The whitelist.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequest) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequest) GetAlarmDetail() *CreateOpaStrategyNewRequestAlarmDetail {
	return s.AlarmDetail
}

func (s *CreateOpaStrategyNewRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOpaStrategyNewRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateOpaStrategyNewRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOpaStrategyNewRequest) GetImageName() []*string {
	return s.ImageName
}

func (s *CreateOpaStrategyNewRequest) GetLabel() []*string {
	return s.Label
}

func (s *CreateOpaStrategyNewRequest) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *CreateOpaStrategyNewRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *CreateOpaStrategyNewRequest) GetScopes() []*CreateOpaStrategyNewRequestScopes {
	return s.Scopes
}

func (s *CreateOpaStrategyNewRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *CreateOpaStrategyNewRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateOpaStrategyNewRequest) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *CreateOpaStrategyNewRequest) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *CreateOpaStrategyNewRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *CreateOpaStrategyNewRequest) SetAlarmDetail(v *CreateOpaStrategyNewRequestAlarmDetail) *CreateOpaStrategyNewRequest {
	s.AlarmDetail = v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetClusterId(v string) *CreateOpaStrategyNewRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetClusterName(v string) *CreateOpaStrategyNewRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetDescription(v string) *CreateOpaStrategyNewRequest {
	s.Description = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetImageName(v []*string) *CreateOpaStrategyNewRequest {
	s.ImageName = v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetLabel(v []*string) *CreateOpaStrategyNewRequest {
	s.Label = v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetMaliciousImage(v bool) *CreateOpaStrategyNewRequest {
	s.MaliciousImage = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetRuleAction(v int32) *CreateOpaStrategyNewRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetScopes(v []*CreateOpaStrategyNewRequestScopes) *CreateOpaStrategyNewRequest {
	s.Scopes = v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetStrategyId(v int64) *CreateOpaStrategyNewRequest {
	s.StrategyId = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetStrategyName(v string) *CreateOpaStrategyNewRequest {
	s.StrategyName = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetStrategyTemplateId(v int64) *CreateOpaStrategyNewRequest {
	s.StrategyTemplateId = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetUnScanedImage(v bool) *CreateOpaStrategyNewRequest {
	s.UnScanedImage = &v
	return s
}

func (s *CreateOpaStrategyNewRequest) SetWhiteList(v []*string) *CreateOpaStrategyNewRequest {
	s.WhiteList = v
	return s
}

func (s *CreateOpaStrategyNewRequest) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetail struct {
	// The baseline risks.
	Baseline *CreateOpaStrategyNewRequestAlarmDetailBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	// The configuration of image build risk.
	BuildRisk *CreateOpaStrategyNewRequestAlarmDetailBuildRisk `json:"BuildRisk,omitempty" xml:"BuildRisk,omitempty" type:"Struct"`
	// The malicious sample risks.
	MaliciousFile *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty" type:"Struct"`
	// The configuration of sensitive file.
	SensitiveFile *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile `json:"SensitiveFile,omitempty" xml:"SensitiveFile,omitempty" type:"Struct"`
	// The vulnerability risks.
	Vul *CreateOpaStrategyNewRequestAlarmDetailVul `json:"Vul,omitempty" xml:"Vul,omitempty" type:"Struct"`
}

func (s CreateOpaStrategyNewRequestAlarmDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetail) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) GetBaseline() *CreateOpaStrategyNewRequestAlarmDetailBaseline {
	return s.Baseline
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) GetBuildRisk() *CreateOpaStrategyNewRequestAlarmDetailBuildRisk {
	return s.BuildRisk
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) GetMaliciousFile() *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	return s.MaliciousFile
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) GetSensitiveFile() *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	return s.SensitiveFile
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) GetVul() *CreateOpaStrategyNewRequestAlarmDetailVul {
	return s.Vul
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) SetBaseline(v *CreateOpaStrategyNewRequestAlarmDetailBaseline) *CreateOpaStrategyNewRequestAlarmDetail {
	s.Baseline = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) SetBuildRisk(v *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) *CreateOpaStrategyNewRequestAlarmDetail {
	s.BuildRisk = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) SetMaliciousFile(v *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) *CreateOpaStrategyNewRequestAlarmDetail {
	s.MaliciousFile = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) SetSensitiveFile(v *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) *CreateOpaStrategyNewRequestAlarmDetail {
	s.SensitiveFile = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) SetVul(v *CreateOpaStrategyNewRequestAlarmDetailVul) *CreateOpaStrategyNewRequestAlarmDetail {
	s.Vul = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetail) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailBaseline struct {
	// The baseline check items.
	Item []*CreateOpaStrategyNewRequestAlarmDetailBaselineItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailBaseline) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailBaseline) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaseline) GetItem() []*CreateOpaStrategyNewRequestAlarmDetailBaselineItem {
	return s.Item
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaseline) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaseline) SetItem(v []*CreateOpaStrategyNewRequestAlarmDetailBaselineItem) *CreateOpaStrategyNewRequestAlarmDetailBaseline {
	s.Item = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaseline) SetRiskLevel(v []*string) *CreateOpaStrategyNewRequestAlarmDetailBaseline {
	s.RiskLevel = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaseline) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailBaselineItem struct {
	// The ID of the baseline check item.
	//
	// >  You can call the [GetOpaClusterBaseLineList](https://help.aliyun.com/document_detail/2539883.html) operation to query the ID.
	//
	// example:
	//
	// hc.image.checklist.identify.hc_exploit_couchdb_linux.item
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the baseline check item.
	//
	// >  You can call the [GetOpaClusterBaseLineList](https://help.aliyun.com/document_detail/2539883.html) operation to query the name.
	//
	// example:
	//
	// Unauthorized access to CouchDB configuration risk
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailBaselineItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailBaselineItem) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaselineItem) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaselineItem) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaselineItem) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailBaselineItem {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaselineItem) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailBaselineItem {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBaselineItem) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailBuildRisk struct {
	// The configuration of image build risk.
	Item []*CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailBuildRisk) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailBuildRisk) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) GetItem() []*CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	return s.Item
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) SetItem(v []*CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) *CreateOpaStrategyNewRequestAlarmDetailBuildRisk {
	s.Item = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) SetRiskLevel(v []*string) *CreateOpaStrategyNewRequestAlarmDetailBuildRisk {
	s.RiskLevel = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRisk) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem struct {
	// The ID of the image build risk.
	//
	// >  You can call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the image build risk.
	//
	// >  You can call the [ListImageBuildRiskItem](~~ListImageBuildRiskItem~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailBuildRiskItem) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailMaliciousFile struct {
	// The malicious samples.
	Item []*CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) GetItem() []*CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	return s.Item
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) SetItem(v []*CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	s.Item = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) SetRiskLevel(v []*string) *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile {
	s.RiskLevel = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFile) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem struct {
	// The ID of the malicious sample.
	//
	// >  You can call the [DescribeMatchedMaliciousNames](~~DescribeMatchedMaliciousNames~~) operation to query the ID.
	//
	// example:
	//
	// 3685699
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the malicious sample.
	//
	// >  You can call the [DescribeMatchedMaliciousNames](~~DescribeMatchedMaliciousNames~~) operation to query the name.
	//
	// example:
	//
	// abnormal binary file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailMaliciousFileItem) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailSensitiveFile struct {
	// The configuration of sensitive file.
	Item []*CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) GetItem() []*CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	return s.Item
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) SetItem(v []*CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	s.Item = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) SetRiskLevel(v []*string) *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile {
	s.RiskLevel = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFile) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem struct {
	// The ID of the sensitive files.
	//
	// >  You can call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// key
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the sensitive files.
	//
	// >  You can call the [GetSensitiveDefineRuleConfig](~~GetSensitiveDefineRuleConfig~~) operation to query the ID of the malicious sample.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailSensitiveFileItem) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailVul struct {
	// The vulnerabilities.
	Item []*CreateOpaStrategyNewRequestAlarmDetailVulItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// Risk type of vulnerability.
	RiskClass []*CreateOpaStrategyNewRequestAlarmDetailVulRiskClass `json:"RiskClass,omitempty" xml:"RiskClass,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailVul) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailVul) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) GetItem() []*CreateOpaStrategyNewRequestAlarmDetailVulItem {
	return s.Item
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) GetRiskClass() []*CreateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	return s.RiskClass
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) SetItem(v []*CreateOpaStrategyNewRequestAlarmDetailVulItem) *CreateOpaStrategyNewRequestAlarmDetailVul {
	s.Item = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) SetRiskClass(v []*CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) *CreateOpaStrategyNewRequestAlarmDetailVul {
	s.RiskClass = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) SetRiskLevel(v []*string) *CreateOpaStrategyNewRequestAlarmDetailVul {
	s.RiskLevel = v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVul) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailVulItem struct {
	// The ID of the vulnerability.
	//
	// >  You can call the [DescribeVulListPage](https://help.aliyun.com/document_detail/471928.html) operation to query the ID.
	//
	// example:
	//
	// CVE-2023-36034
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability.
	//
	// >  You can call the [DescribeVulListPage](https://help.aliyun.com/document_detail/471928.html) operation to query the name.
	//
	// example:
	//
	// Microsoft Edge vul
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailVulItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailVulItem) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulItem) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulItem) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulItem) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailVulItem {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulItem) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailVulItem {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulItem) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestAlarmDetailVulRiskClass struct {
	// The ID of the vulnerability types. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// example:
	//
	// cve
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability. Valid values:
	//
	// 	- **system vulnerability**
	//
	// 	- **application vulnerability**
	//
	// example:
	//
	// system vulnerability
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) GetId() *string {
	return s.Id
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) GetName() *string {
	return s.Name
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) SetId(v string) *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	s.Id = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) SetName(v string) *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass {
	s.Name = &v
	return s
}

func (s *CreateOpaStrategyNewRequestAlarmDetailVulRiskClass) Validate() error {
	return dara.Validate(s)
}

type CreateOpaStrategyNewRequestScopes struct {
	// The ID of the cluster node to which the rule is applied.
	//
	// > This parameter is not required when you create the instance.
	//
	// example:
	//
	// ack-p-1
	AckPolicyInstanceId *string `json:"AckPolicyInstanceId,omitempty" xml:"AckPolicyInstanceId,omitempty"`
	// Specifies whether to include all namespaces. Valid values:
	//
	// 	- **1**: includes all namespaces.
	//
	// 	- **0**: does not include all namespaces.
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the cluster that is specified in the rule.
	//
	// >  You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/421736.html) operation to query the cluster ID.
	//
	// example:
	//
	// cc50d***015d2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	//
	// > This parameter is valid only when the AllNamespace parameter is set to 0.
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewRequestScopes) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewRequestScopes) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewRequestScopes) GetAckPolicyInstanceId() *string {
	return s.AckPolicyInstanceId
}

func (s *CreateOpaStrategyNewRequestScopes) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *CreateOpaStrategyNewRequestScopes) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOpaStrategyNewRequestScopes) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *CreateOpaStrategyNewRequestScopes) SetAckPolicyInstanceId(v string) *CreateOpaStrategyNewRequestScopes {
	s.AckPolicyInstanceId = &v
	return s
}

func (s *CreateOpaStrategyNewRequestScopes) SetAllNamespace(v int32) *CreateOpaStrategyNewRequestScopes {
	s.AllNamespace = &v
	return s
}

func (s *CreateOpaStrategyNewRequestScopes) SetClusterId(v string) *CreateOpaStrategyNewRequestScopes {
	s.ClusterId = &v
	return s
}

func (s *CreateOpaStrategyNewRequestScopes) SetNamespaceList(v []*string) *CreateOpaStrategyNewRequestScopes {
	s.NamespaceList = v
	return s
}

func (s *CreateOpaStrategyNewRequestScopes) Validate() error {
	return dara.Validate(s)
}
