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
	Configs    []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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

type ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigs struct {
	ConfigId                  *int64                                                                                        `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType                *string                                                                                       `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestHeaderModification []*ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	Rule                      *string                                                                                       `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                *string                                                                                       `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                  *string                                                                                       `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                  *int32                                                                                        `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion               *int32                                                                                        `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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

type ListHttpIncomingRequestHeaderModificationRulesResponseBodyConfigsRequestHeaderModification struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
