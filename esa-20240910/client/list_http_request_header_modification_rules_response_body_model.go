// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpRequestHeaderModificationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs) *ListHttpRequestHeaderModificationRulesResponseBody
	GetConfigs() []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpRequestHeaderModificationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpRequestHeaderModificationRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpRequestHeaderModificationRulesResponseBody
	GetTotalPage() *int32
}

type ListHttpRequestHeaderModificationRulesResponseBody struct {
	// List of HTTP request header modification configurations.
	Configs []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default **500**, with a range of **1~500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetConfigs() []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetConfigs(v []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetRequestId(v string) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetTotalCount(v int32) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) SetTotalPage(v int32) *ListHttpRequestHeaderModificationRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHttpRequestHeaderModificationRulesResponseBodyConfigs struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration;
	//
	// - rule: Rule configuration;
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Modify request headers, supporting add, delete, and modify operations.
	RequestHeaderModification []*ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the configuration, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetRequestHeaderModification() []*ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetConfigId(v int64) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetConfigType(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetRequestHeaderModification(v []*ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.RequestHeaderModification = v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetRule(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetRuleEnable(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetRuleName(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetSequence(v int32) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListHttpRequestHeaderModificationRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}

type ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification struct {
	// The name of the request header.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. The value range is as follows:
	//
	// - add: Add.
	//
	// - del: Delete.
	//
	// - modify: Modify.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetName(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetOperation(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetType(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetValue(v string) *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
