// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpResponseHeaderModificationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpResponseHeaderModificationRulesResponseBodyConfigs) *ListHttpResponseHeaderModificationRulesResponseBody
	GetConfigs() []*ListHttpResponseHeaderModificationRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpResponseHeaderModificationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpResponseHeaderModificationRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpResponseHeaderModificationRulesResponseBody
	GetTotalPage() *int32
}

type ListHttpResponseHeaderModificationRulesResponseBody struct {
	Configs    []*ListHttpResponseHeaderModificationRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetConfigs() []*ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetConfigs(v []*ListHttpResponseHeaderModificationRulesResponseBodyConfigs) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetRequestId(v string) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetTotalCount(v int32) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) SetTotalPage(v int32) *ListHttpResponseHeaderModificationRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBody) Validate() error {
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

type ListHttpResponseHeaderModificationRulesResponseBodyConfigs struct {
	ConfigId                   *int64                                                                                  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType                 *string                                                                                 `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ResponseHeaderModification []*ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	Rule                       *string                                                                                 `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                 *string                                                                                 `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                   *string                                                                                 `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                   *int32                                                                                  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion                *int32                                                                                  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetResponseHeaderModification() []*ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetConfigId(v int64) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetConfigType(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetResponseHeaderModification(v []*ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.ResponseHeaderModification = v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetRule(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetRuleEnable(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetRuleName(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetSequence(v int32) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListHttpResponseHeaderModificationRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigs) Validate() error {
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

type ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetName(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetOperation(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetType(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) SetValue(v string) *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponseBodyConfigsResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
