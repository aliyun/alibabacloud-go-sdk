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
	// The type of the content in the sensitive data detection rule. Valid values:
	//
	// - **2**: regular expression.
	//
	// - **3**: algorithm.
	//
	// - **5**: keyword.
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. The rule can be a regular expression, an algorithm, or a keyword, and matches fields or text that contain sensitive data.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The unique ID of the sensitive data detection rule.
	//
	// You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the request and response. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Simplified Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// - **1**: rule-based match.
	//
	// - **2**: dictionary-based match.
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// A collection of model IDs for sensitive data auditing.
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
	// The type of the product resource that contains the sensitive data detection rule. Valid values:
	//
	// - **MaxCompute**.
	//
	// - **OSS**.
	//
	// - **ADS**.
	//
	// - **OTS**.
	//
	// - **RDS**.
	//
	// - **SELF_DB**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the product that contains the sensitive data detection rule. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: ADS.
	//
	// - **4**: OTS.
	//
	// - **5**: RDS.
	//
	// - **6**: SELF_DB.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the risk level for the sensitive data detection rule. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// - **1**: data detection rule.
	//
	// - **2**: audit policy.
	//
	// - **3**: abnormal event rule.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The type of data asset that the rule supports. Valid values:
	//
	// - **0**: all assets.
	//
	// - **1**: structured assets.
	//
	// - **2**: unstructured assets.
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// A collection of template IDs for sensitive data auditing.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level of the sensitive data detection rule. Valid values:
	//
	// - **1**: low.
	//
	// - **2**: medium.
	//
	// - **3**: high.
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
