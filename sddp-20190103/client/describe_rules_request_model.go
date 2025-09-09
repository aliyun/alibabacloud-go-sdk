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
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
	//
	// example:
	//
	// 1
	ContentCategory *int32 `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The external cooperation channel. Valid values:
	//
	// 	- DAS
	//
	// 	- YAOCHI
	//
	// example:
	//
	// DAS
	CooperationChannel *string `json:"CooperationChannel,omitempty" xml:"CooperationChannel,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: built-in rule
	//
	// 	- **1**: custom rule
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
	// The parent group type of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to allow earlier versions of request parameters to support keywords that are supported in later versions of request parameters. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > To specify keywords as the content type of the sensitive data detection rule, you can set the Category parameter to 0 for earlier versions of request parameters and set the Category parameter to 5 for later versions of request parameters. You can specify the KeywordCompatible parameter based on your business requirements.
	//
	// example:
	//
	// true
	KeywordCompatible *bool `json:"KeywordCompatible,omitempty" xml:"KeywordCompatible,omitempty"`
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
	// 	- 1: rule-based match
	//
	// 	- 2: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The name of the sensitive data detection rule. Fuzzy match is supported.
	//
	// example:
	//
	// \\*\\*\\	- rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *int32 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 1
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
	// Specifies whether to query a simplified rule. The simplified rule contains only the rule name. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Simplify *bool `json:"Simplify,omitempty" xml:"Simplify,omitempty"`
	// The status of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
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
	// The severity level of the alert. Valid values:
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
