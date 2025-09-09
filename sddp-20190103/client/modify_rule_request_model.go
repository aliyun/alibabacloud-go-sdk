// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v int32) *ModifyRuleRequest
	GetCategory() *int32
	SetContent(v string) *ModifyRuleRequest
	GetContent() *string
	SetId(v int64) *ModifyRuleRequest
	GetId() *int64
	SetLang(v string) *ModifyRuleRequest
	GetLang() *string
	SetMatchType(v int32) *ModifyRuleRequest
	GetMatchType() *int32
	SetModelRuleIds(v string) *ModifyRuleRequest
	GetModelRuleIds() *string
	SetName(v string) *ModifyRuleRequest
	GetName() *string
	SetProductCode(v string) *ModifyRuleRequest
	GetProductCode() *string
	SetProductId(v int64) *ModifyRuleRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *ModifyRuleRequest
	GetRiskLevelId() *int64
	SetRuleType(v int32) *ModifyRuleRequest
	GetRuleType() *int32
	SetSupportForm(v int32) *ModifyRuleRequest
	GetSupportForm() *int32
	SetTemplateRuleIds(v string) *ModifyRuleRequest
	GetTemplateRuleIds() *string
	SetWarnLevel(v int32) *ModifyRuleRequest
	GetWarnLevel() *int32
}

type ModifyRuleRequest struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **2**: regular expression
	//
	// 	- **3**: algorithm
	//
	// 	- **5**: keyword
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. You can specify a regular expression, an algorithm, or keywords that are used to match sensitive fields or text.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// 	- **1**: rule-based match
	//
	// 	- **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The IDs of the models for sensitive data audit.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// esw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service to which the sensitive data detection rule is applied. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: data detection rule
	//
	// 	- **2**: audit rule
	//
	// 	- **3**: anomalous event detection rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The data assets supported by the sensitive data detection rule. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data assets
	//
	// 	- **2**: unstructured data assets
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level of the alert that is triggered by the sensitive data detection rule. Valid values:
	//
	// 	- **1**: low level
	//
	// 	- **2**: medium level
	//
	// 	- **3**: high level
	//
	// example:
	//
	// 1
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s ModifyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequest) GetCategory() *int32 {
	return s.Category
}

func (s *ModifyRuleRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyRuleRequest) GetMatchType() *int32 {
	return s.MatchType
}

func (s *ModifyRuleRequest) GetModelRuleIds() *string {
	return s.ModelRuleIds
}

func (s *ModifyRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyRuleRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ModifyRuleRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ModifyRuleRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *ModifyRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ModifyRuleRequest) GetSupportForm() *int32 {
	return s.SupportForm
}

func (s *ModifyRuleRequest) GetTemplateRuleIds() *string {
	return s.TemplateRuleIds
}

func (s *ModifyRuleRequest) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *ModifyRuleRequest) SetCategory(v int32) *ModifyRuleRequest {
	s.Category = &v
	return s
}

func (s *ModifyRuleRequest) SetContent(v string) *ModifyRuleRequest {
	s.Content = &v
	return s
}

func (s *ModifyRuleRequest) SetId(v int64) *ModifyRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleRequest) SetLang(v string) *ModifyRuleRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleRequest) SetMatchType(v int32) *ModifyRuleRequest {
	s.MatchType = &v
	return s
}

func (s *ModifyRuleRequest) SetModelRuleIds(v string) *ModifyRuleRequest {
	s.ModelRuleIds = &v
	return s
}

func (s *ModifyRuleRequest) SetName(v string) *ModifyRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyRuleRequest) SetProductCode(v string) *ModifyRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyRuleRequest) SetProductId(v int64) *ModifyRuleRequest {
	s.ProductId = &v
	return s
}

func (s *ModifyRuleRequest) SetRiskLevelId(v int64) *ModifyRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ModifyRuleRequest) SetRuleType(v int32) *ModifyRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyRuleRequest) SetSupportForm(v int32) *ModifyRuleRequest {
	s.SupportForm = &v
	return s
}

func (s *ModifyRuleRequest) SetTemplateRuleIds(v string) *ModifyRuleRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *ModifyRuleRequest) SetWarnLevel(v int32) *ModifyRuleRequest {
	s.WarnLevel = &v
	return s
}

func (s *ModifyRuleRequest) Validate() error {
	return dara.Validate(s)
}
