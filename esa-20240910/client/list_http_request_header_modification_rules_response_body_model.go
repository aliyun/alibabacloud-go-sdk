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
	Configs    []*ListHttpRequestHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                       `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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

type ListHttpRequestHeaderModificationRulesResponseBodyConfigs struct {
	ConfigId                  *int64                                                                                `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType                *string                                                                               `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestHeaderModification []*ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	Rule                      *string                                                                               `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                *string                                                                               `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                  *string                                                                               `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                  *int32                                                                                `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion               *int32                                                                                `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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
	if s.RequestHeaderModification != nil {
		for _, item := range s.RequestHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHttpRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
