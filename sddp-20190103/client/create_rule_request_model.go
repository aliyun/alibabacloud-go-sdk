// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v int32) *CreateRuleRequest
	GetCategory() *int32
	SetContent(v string) *CreateRuleRequest
	GetContent() *string
	SetContentCategory(v int32) *CreateRuleRequest
	GetContentCategory() *int32
	SetDescription(v string) *CreateRuleRequest
	GetDescription() *string
	SetLang(v string) *CreateRuleRequest
	GetLang() *string
	SetMatchType(v int32) *CreateRuleRequest
	GetMatchType() *int32
	SetModelRuleIds(v string) *CreateRuleRequest
	GetModelRuleIds() *string
	SetName(v string) *CreateRuleRequest
	GetName() *string
	SetProductCode(v string) *CreateRuleRequest
	GetProductCode() *string
	SetProductId(v int64) *CreateRuleRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *CreateRuleRequest
	GetRiskLevelId() *int64
	SetRuleType(v int32) *CreateRuleRequest
	GetRuleType() *int32
	SetSourceIp(v string) *CreateRuleRequest
	GetSourceIp() *string
	SetStatExpress(v string) *CreateRuleRequest
	GetStatExpress() *string
	SetStatus(v int32) *CreateRuleRequest
	GetStatus() *int32
	SetSupportForm(v int32) *CreateRuleRequest
	GetSupportForm() *int32
	SetTarget(v string) *CreateRuleRequest
	GetTarget() *string
	SetTemplateRuleIds(v string) *CreateRuleRequest
	GetTemplateRuleIds() *string
	SetWarnLevel(v int32) *CreateRuleRequest
	GetWarnLevel() *int32
}

type CreateRuleRequest struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// - **0**: keyword.
	//
	// - **2**: regular expression.
	//
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. The content can be a regular expression or a keyword that is used to match sensitive data fields or text.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type. Valid values:
	//
	// - **1**: SQL injection exploit.
	//
	// - **2**: SQL injection bypass.
	//
	// - **3**: stored procedure abuse.
	//
	// - **4**: buffer overflow.
	//
	// - **5**: error-based SQL injection.
	//
	// example:
	//
	// 1
	ContentCategory *int32 `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// ID card
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The language of the content in the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
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
	// The IDs of the model rules for sensitive data auditing.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-tst
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the service to which the data asset belongs. Valid values: **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level ID of the sensitive data detection rule. Valid values:
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
	// The type of the rule. Valid values:
	//
	// - **1**: data detection rule.
	//
	// - **2**: audit policy.
	//
	// - **3**: anomaly detection rule.
	//
	// - **99**: custom rule.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 106.11.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The statistical expression.
	//
	// example:
	//
	// {"time":"24","groupby":"1003","having":[{"key":"2001","value":"50"}]}
	StatExpress *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	// The status of the rule. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of data assets that the rule supports. Valid values:
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
	// The code of the service to which the rule applies. Valid values: **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The IDs of the template rules for sensitive data auditing.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level. Valid values:
	//
	// - **1**: Low.
	//
	// - **2**: Medium.
	//
	// - **3**: High.
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) GetCategory() *int32 {
	return s.Category
}

func (s *CreateRuleRequest) GetContent() *string {
	return s.Content
}

func (s *CreateRuleRequest) GetContentCategory() *int32 {
	return s.ContentCategory
}

func (s *CreateRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateRuleRequest) GetMatchType() *int32 {
	return s.MatchType
}

func (s *CreateRuleRequest) GetModelRuleIds() *string {
	return s.ModelRuleIds
}

func (s *CreateRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateRuleRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateRuleRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *CreateRuleRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *CreateRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *CreateRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateRuleRequest) GetStatExpress() *string {
	return s.StatExpress
}

func (s *CreateRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateRuleRequest) GetSupportForm() *int32 {
	return s.SupportForm
}

func (s *CreateRuleRequest) GetTarget() *string {
	return s.Target
}

func (s *CreateRuleRequest) GetTemplateRuleIds() *string {
	return s.TemplateRuleIds
}

func (s *CreateRuleRequest) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *CreateRuleRequest) SetCategory(v int32) *CreateRuleRequest {
	s.Category = &v
	return s
}

func (s *CreateRuleRequest) SetContent(v string) *CreateRuleRequest {
	s.Content = &v
	return s
}

func (s *CreateRuleRequest) SetContentCategory(v int32) *CreateRuleRequest {
	s.ContentCategory = &v
	return s
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleRequest) SetLang(v string) *CreateRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateRuleRequest) SetMatchType(v int32) *CreateRuleRequest {
	s.MatchType = &v
	return s
}

func (s *CreateRuleRequest) SetModelRuleIds(v string) *CreateRuleRequest {
	s.ModelRuleIds = &v
	return s
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateRuleRequest) SetProductCode(v string) *CreateRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateRuleRequest) SetProductId(v int64) *CreateRuleRequest {
	s.ProductId = &v
	return s
}

func (s *CreateRuleRequest) SetRiskLevelId(v int64) *CreateRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleType(v int32) *CreateRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateRuleRequest) SetSourceIp(v string) *CreateRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateRuleRequest) SetStatExpress(v string) *CreateRuleRequest {
	s.StatExpress = &v
	return s
}

func (s *CreateRuleRequest) SetStatus(v int32) *CreateRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateRuleRequest) SetSupportForm(v int32) *CreateRuleRequest {
	s.SupportForm = &v
	return s
}

func (s *CreateRuleRequest) SetTarget(v string) *CreateRuleRequest {
	s.Target = &v
	return s
}

func (s *CreateRuleRequest) SetTemplateRuleIds(v string) *CreateRuleRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *CreateRuleRequest) SetWarnLevel(v int32) *CreateRuleRequest {
	s.WarnLevel = &v
	return s
}

func (s *CreateRuleRequest) Validate() error {
	return dara.Validate(s)
}
