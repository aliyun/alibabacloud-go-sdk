// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRecognizeRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountName(v string) *EditRecognizeRuleRequest
  GetAccountName() *string 
  SetColExclude(v string) *EditRecognizeRuleRequest
  GetColExclude() *string 
  SetColScan(v string) *EditRecognizeRuleRequest
  GetColScan() *string 
  SetCommentScan(v string) *EditRecognizeRuleRequest
  GetCommentScan() *string 
  SetContentScan(v string) *EditRecognizeRuleRequest
  GetContentScan() *string 
  SetHitThreshold(v int32) *EditRecognizeRuleRequest
  GetHitThreshold() *int32 
  SetLevelName(v string) *EditRecognizeRuleRequest
  GetLevelName() *string 
  SetNodeId(v string) *EditRecognizeRuleRequest
  GetNodeId() *string 
  SetNodeParent(v string) *EditRecognizeRuleRequest
  GetNodeParent() *string 
  SetOperationType(v int32) *EditRecognizeRuleRequest
  GetOperationType() *int32 
  SetRecognizeRules(v string) *EditRecognizeRuleRequest
  GetRecognizeRules() *string 
  SetRecognizeRulesType(v string) *EditRecognizeRuleRequest
  GetRecognizeRulesType() *string 
  SetSensitiveDescription(v string) *EditRecognizeRuleRequest
  GetSensitiveDescription() *string 
  SetSensitiveId(v string) *EditRecognizeRuleRequest
  GetSensitiveId() *string 
  SetSensitiveName(v string) *EditRecognizeRuleRequest
  GetSensitiveName() *string 
  SetStatus(v int32) *EditRecognizeRuleRequest
  GetStatus() *int32 
  SetTemplateId(v string) *EditRecognizeRuleRequest
  GetTemplateId() *string 
  SetTenantId(v string) *EditRecognizeRuleRequest
  GetTenantId() *string 
  SetLevel(v string) *EditRecognizeRuleRequest
  GetLevel() *string 
}

type EditRecognizeRuleRequest struct {
  // The Alibaba Cloud account that is used to create the sensitive data identification rule. Enter the username of the Alibaba Cloud account.
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
  // Customer/personal/personal_Natural_Information/personal_basic_profile_information
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
  // The sensitive field ID. You can call the [QuerySensNodeInfo](https://help.aliyun.com/document_detail/2747189.html) operation to query the IDs of all sensitive fields. You can also call the [QueryRecognizeRuleDetail](https://help.aliyun.com/document_detail/2766023.html) operation to query the IDs of specific sensitive fields.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1a928de7-3962-4e07-93ac-e1973baa1024
  SensitiveId *string `json:"SensitiveId,omitempty" xml:"SensitiveId,omitempty"`
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
  // The sensitivity level of the sensitive field. You can select one from all sensitivity levels that are defined in a template as the sensitivity level of the sensitive field, such as level 1 to level 10.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Level *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s EditRecognizeRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EditRecognizeRuleRequest) GoString() string {
  return s.String()
}

func (s *EditRecognizeRuleRequest) GetAccountName() *string  {
  return s.AccountName
}

func (s *EditRecognizeRuleRequest) GetColExclude() *string  {
  return s.ColExclude
}

func (s *EditRecognizeRuleRequest) GetColScan() *string  {
  return s.ColScan
}

func (s *EditRecognizeRuleRequest) GetCommentScan() *string  {
  return s.CommentScan
}

func (s *EditRecognizeRuleRequest) GetContentScan() *string  {
  return s.ContentScan
}

func (s *EditRecognizeRuleRequest) GetHitThreshold() *int32  {
  return s.HitThreshold
}

func (s *EditRecognizeRuleRequest) GetLevelName() *string  {
  return s.LevelName
}

func (s *EditRecognizeRuleRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EditRecognizeRuleRequest) GetNodeParent() *string  {
  return s.NodeParent
}

func (s *EditRecognizeRuleRequest) GetOperationType() *int32  {
  return s.OperationType
}

func (s *EditRecognizeRuleRequest) GetRecognizeRules() *string  {
  return s.RecognizeRules
}

func (s *EditRecognizeRuleRequest) GetRecognizeRulesType() *string  {
  return s.RecognizeRulesType
}

func (s *EditRecognizeRuleRequest) GetSensitiveDescription() *string  {
  return s.SensitiveDescription
}

func (s *EditRecognizeRuleRequest) GetSensitiveId() *string  {
  return s.SensitiveId
}

func (s *EditRecognizeRuleRequest) GetSensitiveName() *string  {
  return s.SensitiveName
}

func (s *EditRecognizeRuleRequest) GetStatus() *int32  {
  return s.Status
}

func (s *EditRecognizeRuleRequest) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *EditRecognizeRuleRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EditRecognizeRuleRequest) GetLevel() *string  {
  return s.Level
}

func (s *EditRecognizeRuleRequest) SetAccountName(v string) *EditRecognizeRuleRequest {
  s.AccountName = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetColExclude(v string) *EditRecognizeRuleRequest {
  s.ColExclude = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetColScan(v string) *EditRecognizeRuleRequest {
  s.ColScan = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetCommentScan(v string) *EditRecognizeRuleRequest {
  s.CommentScan = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetContentScan(v string) *EditRecognizeRuleRequest {
  s.ContentScan = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetHitThreshold(v int32) *EditRecognizeRuleRequest {
  s.HitThreshold = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetLevelName(v string) *EditRecognizeRuleRequest {
  s.LevelName = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetNodeId(v string) *EditRecognizeRuleRequest {
  s.NodeId = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetNodeParent(v string) *EditRecognizeRuleRequest {
  s.NodeParent = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetOperationType(v int32) *EditRecognizeRuleRequest {
  s.OperationType = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetRecognizeRules(v string) *EditRecognizeRuleRequest {
  s.RecognizeRules = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetRecognizeRulesType(v string) *EditRecognizeRuleRequest {
  s.RecognizeRulesType = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetSensitiveDescription(v string) *EditRecognizeRuleRequest {
  s.SensitiveDescription = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetSensitiveId(v string) *EditRecognizeRuleRequest {
  s.SensitiveId = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetSensitiveName(v string) *EditRecognizeRuleRequest {
  s.SensitiveName = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetStatus(v int32) *EditRecognizeRuleRequest {
  s.Status = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetTemplateId(v string) *EditRecognizeRuleRequest {
  s.TemplateId = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetTenantId(v string) *EditRecognizeRuleRequest {
  s.TenantId = &v
  return s
}

func (s *EditRecognizeRuleRequest) SetLevel(v string) *EditRecognizeRuleRequest {
  s.Level = &v
  return s
}

func (s *EditRecognizeRuleRequest) Validate() error {
  return dara.Validate(s)
}

