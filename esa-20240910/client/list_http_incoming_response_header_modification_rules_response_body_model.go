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
	// The list of incoming response header modification rules.
	Configs []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The number of the returned page. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 500. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF02C6F6-DB59-5438-8C05-3CE42DFCB0AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of entries per page.
	//
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
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 430559776208896
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The type of the configuration. Specifies whether to check the image used by the instance supports hot migration. Valid values:
	//
	// 	- global: global configuration.
	//
	// 	- rule: rule configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The configurations of modifying response headers. You can add, delete, or modify a response header.
	ResponseHeaderModification []*ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configuration. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example:(http.host eq "video.example.com"): Match the specified request.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configuration. Specifies whether to check the image used by the instance supports hot migration. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The order in which the rule is executed. A smaller value gives priority to the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
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
	if s.ResponseHeaderModification != nil {
		for _, item := range s.ResponseHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHttpIncomingResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification struct {
	// The name of the response header.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action. Specifies whether to check the image used by the instance supports hot migration. Valid values:
	//
	// 	- add: adds a response header.
	//
	// 	- del: deletes a response header.
	//
	// 	- modify: modifies a response header.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The type of the header. Valid values:
	//
	// 	- static
	//
	// 	- dynamic
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the response header.
	//
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
