// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamWatermarkRulesResponseBody
	GetRequestId() *string
	SetRuleInfoList(v *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) *DescribeLiveStreamWatermarkRulesResponseBody
	GetRuleInfoList() *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList
	SetTotal(v int32) *DescribeLiveStreamWatermarkRulesResponseBody
	GetTotal() *int32
}

type DescribeLiveStreamWatermarkRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b9676b3
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfoList *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList `json:"RuleInfoList,omitempty" xml:"RuleInfoList,omitempty" type:"Struct"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetRuleInfoList() *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList {
	return s.RuleInfoList
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetRequestId(v string) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetRuleInfoList(v *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.RuleInfoList = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) SetTotal(v int32) *DescribeLiveStreamWatermarkRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBody) Validate() error {
	if s.RuleInfoList != nil {
		if err := s.RuleInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList struct {
	RuleInfo []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) GetRuleInfo() []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	return s.RuleInfo
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) SetRuleInfo(v []*DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList {
	s.RuleInfo = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoList) Validate() error {
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

type DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo struct {
	App         *string `json:"App,omitempty" xml:"App,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Stream      *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetApp() *string {
	return s.App
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetName() *string {
	return s.Name
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetApp(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.App = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetDescription(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetDomain(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Domain = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetName(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Name = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetRuleId(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.RuleId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetStream(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.Stream = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) SetTemplateId(v string) *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo {
	s.TemplateId = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponseBodyRuleInfoListRuleInfo) Validate() error {
	return dara.Validate(s)
}
