// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v int32) *DescribeRulesRequest
	GetCategory() *int32
	SetContentCategory(v int32) *DescribeRulesRequest
	GetContentCategory() *int32
	SetCooperationChannel(v string) *DescribeRulesRequest
	GetCooperationChannel() *string
	SetCurrentPage(v int32) *DescribeRulesRequest
	GetCurrentPage() *int32
	SetCustomType(v int32) *DescribeRulesRequest
	GetCustomType() *int32
	SetFeatureType(v int32) *DescribeRulesRequest
	GetFeatureType() *int32
	SetGroupId(v string) *DescribeRulesRequest
	GetGroupId() *string
	SetKeywordCompatible(v bool) *DescribeRulesRequest
	GetKeywordCompatible() *bool
	SetLang(v string) *DescribeRulesRequest
	GetLang() *string
	SetMatchType(v int32) *DescribeRulesRequest
	GetMatchType() *int32
	SetName(v string) *DescribeRulesRequest
	GetName() *string
	SetPageSize(v int32) *DescribeRulesRequest
	GetPageSize() *int32
	SetProductCode(v int32) *DescribeRulesRequest
	GetProductCode() *int32
	SetProductId(v int64) *DescribeRulesRequest
	GetProductId() *int64
	SetRiskLevelId(v int64) *DescribeRulesRequest
	GetRiskLevelId() *int64
	SetRuleType(v int32) *DescribeRulesRequest
	GetRuleType() *int32
	SetSimplify(v bool) *DescribeRulesRequest
	GetSimplify() *bool
	SetStatus(v int32) *DescribeRulesRequest
	GetStatus() *int32
	SetSupportForm(v int32) *DescribeRulesRequest
	GetSupportForm() *int32
	SetWarnLevel(v int32) *DescribeRulesRequest
	GetWarnLevel() *int32
}

type DescribeRulesRequest struct {
	// The type of content in the sensitive data detection rule. Valid values:
	//
	// - **0**: keyword
	//
	// - **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content type. Valid values:
	//
	// - **1**: SQL injection exploits
	//
	// - **2**: SQL injection bypass attempts
	//
	// - **3**: stored procedure abuse
	//
	// - **4**: buffer overflows
	//
	// - **5**: error-based SQL injections
	//
	// example:
	//
	// 1
	ContentCategory *int32 `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The source of the external cooperation request. Valid values:
	//
	// - DAS
	//
	// - YAOCHI
	//
	// example:
	//
	// DAS
	CooperationChannel *string `json:"CooperationChannel,omitempty" xml:"CooperationChannel,omitempty"`
	// The page number of the paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// - **0**: built-in
	//
	// - **1**: custom
	//
	// example:
	//
	// 1
	CustomType *int32 `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The parent group of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether the keyword is compatible with earlier versions. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > In earlier versions, the Category parameter for keywords had a value of 0. In the current version, it has a value of 5. Enable this parameter based on your business needs.
	//
	// example:
	//
	// true
	KeywordCompatible *bool `json:"KeywordCompatible,omitempty" xml:"KeywordCompatible,omitempty"`
	// The language of the request and response messages. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// - 1: rule-based match
	//
	// - 2: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The name of the sensitive data detection rule. Fuzzy match is supported.
	//
	// example:
	//
	// ***Rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset belongs. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADS**
	//
	// - **OTS**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// example:
	//
	// MaxCompute
	ProductCode *int32 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
	// 1
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
	// The type of the sensitive data detection rule. Valid values:
	//
	// - **1**: data detection rule
	//
	// - **2**: audit policy
	//
	// - **3**: anomaly detection rule
	//
	// - **99**: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// Specifies whether to return a simplified version of the rule that contains only the rule name. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	Simplify *bool `json:"Simplify,omitempty" xml:"Simplify,omitempty"`
	// The status. Valid values:
	//
	// - **1**: Normal
	//
	// - **0**: Disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of data asset that the rule supports. Valid values:
	//
	// - **0**: all assets
	//
	// - **1**: structured assets
	//
	// - **2**: unstructured assets
	//
	// > When you query for rules that support structured or unstructured assets, the response also includes rules that support all asset types.
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The risk level.
	//
	// - **1**: Low
	//
	// - **2**: Medium
	//
	// - **3**: High
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) GetCategory() *int32 {
	return s.Category
}

func (s *DescribeRulesRequest) GetContentCategory() *int32 {
	return s.ContentCategory
}

func (s *DescribeRulesRequest) GetCooperationChannel() *string {
	return s.CooperationChannel
}

func (s *DescribeRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRulesRequest) GetCustomType() *int32 {
	return s.CustomType
}

func (s *DescribeRulesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeRulesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeRulesRequest) GetKeywordCompatible() *bool {
	return s.KeywordCompatible
}

func (s *DescribeRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRulesRequest) GetMatchType() *int32 {
	return s.MatchType
}

func (s *DescribeRulesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRulesRequest) GetProductCode() *int32 {
	return s.ProductCode
}

func (s *DescribeRulesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeRulesRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeRulesRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *DescribeRulesRequest) GetSimplify() *bool {
	return s.Simplify
}

func (s *DescribeRulesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeRulesRequest) GetSupportForm() *int32 {
	return s.SupportForm
}

func (s *DescribeRulesRequest) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *DescribeRulesRequest) SetCategory(v int32) *DescribeRulesRequest {
	s.Category = &v
	return s
}

func (s *DescribeRulesRequest) SetContentCategory(v int32) *DescribeRulesRequest {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesRequest) SetCooperationChannel(v string) *DescribeRulesRequest {
	s.CooperationChannel = &v
	return s
}

func (s *DescribeRulesRequest) SetCurrentPage(v int32) *DescribeRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulesRequest) SetCustomType(v int32) *DescribeRulesRequest {
	s.CustomType = &v
	return s
}

func (s *DescribeRulesRequest) SetFeatureType(v int32) *DescribeRulesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeRulesRequest) SetGroupId(v string) *DescribeRulesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesRequest) SetKeywordCompatible(v bool) *DescribeRulesRequest {
	s.KeywordCompatible = &v
	return s
}

func (s *DescribeRulesRequest) SetLang(v string) *DescribeRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRulesRequest) SetMatchType(v int32) *DescribeRulesRequest {
	s.MatchType = &v
	return s
}

func (s *DescribeRulesRequest) SetName(v string) *DescribeRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeRulesRequest) SetPageSize(v int32) *DescribeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesRequest) SetProductCode(v int32) *DescribeRulesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeRulesRequest) SetProductId(v int64) *DescribeRulesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeRulesRequest) SetRiskLevelId(v int64) *DescribeRulesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesRequest) SetRuleType(v int32) *DescribeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRulesRequest) SetSimplify(v bool) *DescribeRulesRequest {
	s.Simplify = &v
	return s
}

func (s *DescribeRulesRequest) SetStatus(v int32) *DescribeRulesRequest {
	s.Status = &v
	return s
}

func (s *DescribeRulesRequest) SetSupportForm(v int32) *DescribeRulesRequest {
	s.SupportForm = &v
	return s
}

func (s *DescribeRulesRequest) SetWarnLevel(v int32) *DescribeRulesRequest {
	s.WarnLevel = &v
	return s
}

func (s *DescribeRulesRequest) Validate() error {
	return dara.Validate(s)
}
