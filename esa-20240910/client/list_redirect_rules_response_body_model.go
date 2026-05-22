// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRedirectRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListRedirectRulesResponseBodyConfigs) *ListRedirectRulesResponseBody
	GetConfigs() []*ListRedirectRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListRedirectRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRedirectRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRedirectRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRedirectRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListRedirectRulesResponseBody
	GetTotalPage() *int32
}

type ListRedirectRulesResponseBody struct {
	Configs    []*ListRedirectRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListRedirectRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesResponseBody) GetConfigs() []*ListRedirectRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListRedirectRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRedirectRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRedirectRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRedirectRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRedirectRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListRedirectRulesResponseBody) SetConfigs(v []*ListRedirectRulesResponseBodyConfigs) *ListRedirectRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListRedirectRulesResponseBody) SetPageNumber(v int32) *ListRedirectRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetPageSize(v int32) *ListRedirectRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetRequestId(v string) *ListRedirectRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetTotalCount(v int32) *ListRedirectRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRedirectRulesResponseBody) SetTotalPage(v int32) *ListRedirectRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListRedirectRulesResponseBody) Validate() error {
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

type ListRedirectRulesResponseBodyConfigs struct {
	ConfigId           *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType         *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	Rule               *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable         *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence           *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion        *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	StatusCode         *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TargetUrl          *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRedirectRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRedirectRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRedirectRulesResponseBodyConfigs) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListRedirectRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRedirectRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListRedirectRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRedirectRulesResponseBodyConfigs) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListRedirectRulesResponseBodyConfigs) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *ListRedirectRulesResponseBodyConfigs) GetType() *string {
	return s.Type
}

func (s *ListRedirectRulesResponseBodyConfigs) SetConfigId(v int64) *ListRedirectRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetConfigType(v string) *ListRedirectRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetReserveQueryString(v string) *ListRedirectRulesResponseBodyConfigs {
	s.ReserveQueryString = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRule(v string) *ListRedirectRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRuleEnable(v string) *ListRedirectRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetRuleName(v string) *ListRedirectRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetSequence(v int32) *ListRedirectRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListRedirectRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetStatusCode(v string) *ListRedirectRulesResponseBodyConfigs {
	s.StatusCode = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetTargetUrl(v string) *ListRedirectRulesResponseBodyConfigs {
	s.TargetUrl = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) SetType(v string) *ListRedirectRulesResponseBodyConfigs {
	s.Type = &v
	return s
}

func (s *ListRedirectRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
