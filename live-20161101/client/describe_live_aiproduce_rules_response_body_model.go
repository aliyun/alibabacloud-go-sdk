// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIProduceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveAIProduceRulesResponseBody
	GetRequestId() *string
	SetRuleInfoList(v *DescribeLiveAIProduceRulesResponseBodyRuleInfoList) *DescribeLiveAIProduceRulesResponseBody
	GetRuleInfoList() *DescribeLiveAIProduceRulesResponseBodyRuleInfoList
}

type DescribeLiveAIProduceRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b96****
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfoList *DescribeLiveAIProduceRulesResponseBodyRuleInfoList `json:"RuleInfoList,omitempty" xml:"RuleInfoList,omitempty" type:"Struct"`
}

func (s DescribeLiveAIProduceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIProduceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIProduceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveAIProduceRulesResponseBody) GetRuleInfoList() *DescribeLiveAIProduceRulesResponseBodyRuleInfoList {
	return s.RuleInfoList
}

func (s *DescribeLiveAIProduceRulesResponseBody) SetRequestId(v string) *DescribeLiveAIProduceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBody) SetRuleInfoList(v *DescribeLiveAIProduceRulesResponseBodyRuleInfoList) *DescribeLiveAIProduceRulesResponseBody {
	s.RuleInfoList = v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBody) Validate() error {
	if s.RuleInfoList != nil {
		if err := s.RuleInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveAIProduceRulesResponseBodyRuleInfoList struct {
	RuleInfo []*DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveAIProduceRulesResponseBodyRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIProduceRulesResponseBodyRuleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoList) GetRuleInfo() []*DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	return s.RuleInfo
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoList) SetRuleInfo(v []*DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) *DescribeLiveAIProduceRulesResponseBodyRuleInfoList {
	s.RuleInfo = v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoList) Validate() error {
	if s.RuleInfo != nil {
		for _, item := range s.RuleInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo struct {
	App           *string `json:"App,omitempty" xml:"App,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	IsLazy        *bool   `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	LiveTemplate  *string `json:"LiveTemplate,omitempty" xml:"LiveTemplate,omitempty"`
	RulesId       *string `json:"RulesId,omitempty" xml:"RulesId,omitempty"`
	StudioName    *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	SubtitleName  *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	SuffixName    *string `json:"SuffixName,omitempty" xml:"SuffixName,omitempty"`
}

func (s DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetApp() *string {
	return s.App
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetIsLazy() *bool {
	return s.IsLazy
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetLiveTemplate() *string {
	return s.LiveTemplate
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetRulesId() *string {
	return s.RulesId
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetStudioName() *string {
	return s.StudioName
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) GetSuffixName() *string {
	return s.SuffixName
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetApp(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.App = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetDescription(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.Description = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetDomain(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.Domain = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetGmtModifyTime(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.GmtModifyTime = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetIsLazy(v bool) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.IsLazy = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetLiveTemplate(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.LiveTemplate = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetRulesId(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.RulesId = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetStudioName(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.StudioName = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetSubtitleName(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.SubtitleName = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) SetSuffixName(v string) *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo {
	s.SuffixName = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponseBodyRuleInfoListRuleInfo) Validate() error {
	return dara.Validate(s)
}
