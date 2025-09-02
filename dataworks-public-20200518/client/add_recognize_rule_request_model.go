// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecognizeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *AddRecognizeRuleRequest
	GetAccountName() *string
	SetColExclude(v string) *AddRecognizeRuleRequest
	GetColExclude() *string
	SetColScan(v string) *AddRecognizeRuleRequest
	GetColScan() *string
	SetCommentScan(v string) *AddRecognizeRuleRequest
	GetCommentScan() *string
	SetContentScan(v string) *AddRecognizeRuleRequest
	GetContentScan() *string
	SetHitThreshold(v int32) *AddRecognizeRuleRequest
	GetHitThreshold() *int32
	SetLevel(v string) *AddRecognizeRuleRequest
	GetLevel() *string
	SetLevelName(v string) *AddRecognizeRuleRequest
	GetLevelName() *string
	SetNodeId(v string) *AddRecognizeRuleRequest
	GetNodeId() *string
	SetNodeParent(v string) *AddRecognizeRuleRequest
	GetNodeParent() *string
	SetOperationType(v int32) *AddRecognizeRuleRequest
	GetOperationType() *int32
	SetRecognizeRules(v string) *AddRecognizeRuleRequest
	GetRecognizeRules() *string
	SetRecognizeRulesType(v string) *AddRecognizeRuleRequest
	GetRecognizeRulesType() *string
	SetSensitiveDescription(v string) *AddRecognizeRuleRequest
	GetSensitiveDescription() *string
	SetSensitiveName(v string) *AddRecognizeRuleRequest
	GetSensitiveName() *string
	SetStatus(v int32) *AddRecognizeRuleRequest
	GetStatus() *int32
	SetTemplateId(v string) *AddRecognizeRuleRequest
	GetTemplateId() *string
	SetTenantId(v string) *AddRecognizeRuleRequest
	GetTenantId() *string
}

type AddRecognizeRuleRequest struct {
	// The Alibaba Cloud account that is used to create a sensitive data identification rule. Enter the username of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// dsg-uat
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Excludes fields. The system does not identify fields that are assigned with values.
	//
	// 	- The value must be in the ${Project name}.${Table name}.${Field name} or ${Project name}.${Schema name}.${Table name}.${Field name} format.
	//
	// 	- *Wildcards are supported. For example, the asterisk (\\*) in default.table.column1\\	- can be used to match any content following default.table.column1, such as default.table.column10.
	//
	// example:
	//
	// default.qujian.*6
	ColExclude *string `json:"ColExclude,omitempty" xml:"ColExclude,omitempty"`
	// Scans fields. The system identifies only fields that are assigned with values.
	//
	// 	- The value must be in the ${Project name}.${Table name}.${Field name} or ${Project name}.${Schema name}.${Table name}.${Field name} format.
	//
	// 	- *Wildcards are supported. For example, the asterisk (\\*) in default.table.column1\\	- can be used to match any content following default.table.column1, such as default.table.column10.
	//
	// example:
	//
	// default.qujian.*
	ColScan *string `json:"ColScan,omitempty" xml:"ColScan,omitempty"`
	// Scans content. The value is the text of each field comment in your data asset. Fuzzy match is supported.
	//
	// example:
	//
	// test
	CommentScan *string `json:"CommentScan,omitempty" xml:"CommentScan,omitempty"`
	// Identifies content. You can call the [QuerySensNodeInfo](https://help.aliyun.com/document_detail/2747189.html) operation to query the value of the current parameter for a built-in sensitive field.
	//
	// example:
	//
	// {"_clazz":"com.alipay.dsgclient.sdk.dsg.fastscan.engine.cond.NationalityCond"}
	ContentScan *string `json:"ContentScan,omitempty" xml:"ContentScan,omitempty"`
	// The hit ratio threshold. If more than 60%, which is a sample hit ratio threshold, of all sample data records hit the Name Entity Recognition (NER) model, the sensitive field is hit. The value can be an integer from 0 to 100.
	//
	// example:
	//
	// 50
	HitThreshold *int32 `json:"HitThreshold,omitempty" xml:"HitThreshold,omitempty"`
	// The sensitivity level of the sensitive field. You can select one from all sensitivity levels that are defined in a template as the sensitivity level of the sensitive field, such as level 1 to level 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the sensitivity level. You can call the [QueryDefaultTemplate](https://help.aliyun.com/document_detail/2743948.html) operation to obtain the name of the sensitivity level in the related template.
	//
	// example:
	//
	// Confidential
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// The ID of the data category. You can call the [QuerySensClassification](https://help.aliyun.com/document_detail/2746850.html) operation to query the ID of all data categories. Then, you can select a data category to create a sensitive field. Enter the ID of the selected data category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0ce67949-0810-400f-a24a-cc5ffafe1024
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The information about the parent data category of the current data category. You can call the [QuerySensClassification](https://help.aliyun.com/document_detail/2746850.html) operation to obtain the ID of a data category.
	//
	// This parameter is required.
	//
	// example:
	//
	// Customer/personal/personal_Natural _Information/personal basic_profile_information
	NodeParent *string `json:"NodeParent,omitempty" xml:"NodeParent,omitempty"`
	// The type of the arithmetic operation. Valid values:
	//
	// 	- 0: OR
	//
	// 	- 1: AND
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperationType *int32 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The content of the sensitive data identification rule. You can call the [QuerySensNodeInfo](https://help.aliyun.com/document_detail/2747189.html) operation to query the value of the current parameter for a built-in sensitive field.
	//
	// example:
	//
	// {"contentRule":{"_clazz":"com.alipay.dsgclient.sdk.dsg.fastscan.engine.cond.GenderCond"},"_clazz":"com.alipay.dsg.dal.model.RuleContent"}
	RecognizeRules *string `json:"RecognizeRules,omitempty" xml:"RecognizeRules,omitempty"`
	// The type of the sensitive data identification rule. Valid values:
	//
	// 	- 1: regular expression
	//
	// 	- 2: built-in rule
	//
	// 	- 3: sample library
	//
	// 	- 4: self-generated data identification model
	//
	// example:
	//
	// 1
	RecognizeRulesType *string `json:"RecognizeRulesType,omitempty" xml:"RecognizeRulesType,omitempty"`
	// The description of the sensitive field. Enter a string that is less than 128 characters in length.
	//
	// example:
	//
	// This is a sensitive field that identifies the name.
	SensitiveDescription *string `json:"SensitiveDescription,omitempty" xml:"SensitiveDescription,omitempty"`
	// The name of the custom sensitive field. Enter a string that is less than 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	SensitiveName *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
	// The status of the sensitive field. Valid values:
	//
	// 	- 0: draft
	//
	// 	- 1: effective
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template ID. You can call the [QueryDefaultTemplate](https://help.aliyun.com/document_detail/2743948.html) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1970541-6cf5-4d23-b101-d5b66f6e1024
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AddRecognizeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecognizeRuleRequest) GoString() string {
	return s.String()
}

func (s *AddRecognizeRuleRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *AddRecognizeRuleRequest) GetColExclude() *string {
	return s.ColExclude
}

func (s *AddRecognizeRuleRequest) GetColScan() *string {
	return s.ColScan
}

func (s *AddRecognizeRuleRequest) GetCommentScan() *string {
	return s.CommentScan
}

func (s *AddRecognizeRuleRequest) GetContentScan() *string {
	return s.ContentScan
}

func (s *AddRecognizeRuleRequest) GetHitThreshold() *int32 {
	return s.HitThreshold
}

func (s *AddRecognizeRuleRequest) GetLevel() *string {
	return s.Level
}

func (s *AddRecognizeRuleRequest) GetLevelName() *string {
	return s.LevelName
}

func (s *AddRecognizeRuleRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AddRecognizeRuleRequest) GetNodeParent() *string {
	return s.NodeParent
}

func (s *AddRecognizeRuleRequest) GetOperationType() *int32 {
	return s.OperationType
}

func (s *AddRecognizeRuleRequest) GetRecognizeRules() *string {
	return s.RecognizeRules
}

func (s *AddRecognizeRuleRequest) GetRecognizeRulesType() *string {
	return s.RecognizeRulesType
}

func (s *AddRecognizeRuleRequest) GetSensitiveDescription() *string {
	return s.SensitiveDescription
}

func (s *AddRecognizeRuleRequest) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *AddRecognizeRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *AddRecognizeRuleRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddRecognizeRuleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddRecognizeRuleRequest) SetAccountName(v string) *AddRecognizeRuleRequest {
	s.AccountName = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetColExclude(v string) *AddRecognizeRuleRequest {
	s.ColExclude = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetColScan(v string) *AddRecognizeRuleRequest {
	s.ColScan = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetCommentScan(v string) *AddRecognizeRuleRequest {
	s.CommentScan = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetContentScan(v string) *AddRecognizeRuleRequest {
	s.ContentScan = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetHitThreshold(v int32) *AddRecognizeRuleRequest {
	s.HitThreshold = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetLevel(v string) *AddRecognizeRuleRequest {
	s.Level = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetLevelName(v string) *AddRecognizeRuleRequest {
	s.LevelName = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetNodeId(v string) *AddRecognizeRuleRequest {
	s.NodeId = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetNodeParent(v string) *AddRecognizeRuleRequest {
	s.NodeParent = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetOperationType(v int32) *AddRecognizeRuleRequest {
	s.OperationType = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetRecognizeRules(v string) *AddRecognizeRuleRequest {
	s.RecognizeRules = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetRecognizeRulesType(v string) *AddRecognizeRuleRequest {
	s.RecognizeRulesType = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetSensitiveDescription(v string) *AddRecognizeRuleRequest {
	s.SensitiveDescription = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetSensitiveName(v string) *AddRecognizeRuleRequest {
	s.SensitiveName = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetStatus(v int32) *AddRecognizeRuleRequest {
	s.Status = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetTemplateId(v string) *AddRecognizeRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *AddRecognizeRuleRequest) SetTenantId(v string) *AddRecognizeRuleRequest {
	s.TenantId = &v
	return s
}

func (s *AddRecognizeRuleRequest) Validate() error {
	return dara.Validate(s)
}
