// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomResponseCodeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListCustomResponseCodeRulesResponseBodyConfigs) *ListCustomResponseCodeRulesResponseBody
	GetConfigs() []*ListCustomResponseCodeRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListCustomResponseCodeRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomResponseCodeRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomResponseCodeRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomResponseCodeRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListCustomResponseCodeRulesResponseBody
	GetTotalPage() *int32
}

type ListCustomResponseCodeRulesResponseBody struct {
	Configs    []*ListCustomResponseCodeRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCustomResponseCodeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesResponseBody) GetConfigs() []*ListCustomResponseCodeRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListCustomResponseCodeRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomResponseCodeRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomResponseCodeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomResponseCodeRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomResponseCodeRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListCustomResponseCodeRulesResponseBody) SetConfigs(v []*ListCustomResponseCodeRulesResponseBodyConfigs) *ListCustomResponseCodeRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetPageNumber(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetPageSize(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetRequestId(v string) *ListCustomResponseCodeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetTotalCount(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) SetTotalPage(v int32) *ListCustomResponseCodeRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBody) Validate() error {
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

type ListCustomResponseCodeRulesResponseBodyConfigs struct {
	ConfigId    *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageId      *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ReturnCode  *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable  *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence    *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCustomResponseCodeRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetPageId() *string {
	return s.PageId
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetConfigId(v int64) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetConfigType(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetPageId(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.PageId = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetReturnCode(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.ReturnCode = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRule(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRuleEnable(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetRuleName(v string) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetSequence(v int32) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListCustomResponseCodeRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
