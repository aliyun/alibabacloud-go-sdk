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
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. You can specify a regular expression or keywords that are used to match sensitive fields or text.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
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
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
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
	// This parameter is required.
	//
	// example:
	//
	// rule-tst
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the service to which data in the column of the table belongs. Valid values include **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
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
	// 	- **1**: sensitive data detection rule
	//
	// 	- **2**: audit rule
	//
	// 	- **3**: anomalous event detection rule
	//
	// 	- **99**: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The statistical expression.
	//
	// example:
	//
	// 1
	StatExpress *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	// Specifies whether to enable the sensitive data detection rule. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the data asset. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data asset
	//
	// 	- **2**: unstructured data asset
	//
	// > If you set the parameter to 1 or 2, rules that support all data assets and rules that support the queried data asset type are returned.
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The code of the service to which the sensitive data detection rule is applied. Valid values include **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level of the alert that is triggered. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
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
