// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingRequestHeaderModificationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetConfigs() []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody
	GetTotalPage() *int32
}

type ListHttpIncomingRequestHeaderModificationRulesResponseBody struct {
	Configs []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 7FB6EBC8-8849-5FC6-890E-3A761A5CD42D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetConfigs() []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetConfigs(v []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetPageNumber(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetPageSize(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetRequestId(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetTotalCount(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) SetTotalPage(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs struct {
	// example:
	//
	// 429422870243328
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// rule
	ConfigType                *string                                                                                       `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestHeaderModification []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetRequestHeaderModification() []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetConfigId(v int64) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetConfigType(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetRequestHeaderModification(v []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.RequestHeaderModification = v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetRule(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetRuleEnable(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetRuleName(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetSequence(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}

type ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification struct {
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetName(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetOperation(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetType(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) SetValue(v string) *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
