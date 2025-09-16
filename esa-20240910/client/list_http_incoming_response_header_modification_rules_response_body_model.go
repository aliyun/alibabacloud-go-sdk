// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingResponseHeaderModificationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetConfigs() []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody
	GetTotalPage() *int32
}

type ListHttpIncomingResponseHeaderModificationRulesResponseBody struct {
	Configs []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
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
	// CF02C6F6-DB59-5438-8C05-3CE42DFCB0AD
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

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetConfigs() []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetConfigs(v []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetPageNumber(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetPageSize(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetRequestId(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetTotalCount(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) SetTotalPage(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs struct {
	// example:
	//
	// 430559776208896
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// rule
	ConfigType                 *string                                                                                         `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ResponseHeaderModification []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
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

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetResponseHeaderModification() []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetConfigId(v int64) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetConfigType(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetResponseHeaderModification(v []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.ResponseHeaderModification = v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetRule(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetRuleEnable(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetRuleName(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetSequence(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}

type ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification struct {
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

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetName(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetOperation(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetType(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetValue(v string) *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
