// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInterveneRulesResponseBody
	GetCode() *string
	SetData(v *ListInterveneRulesResponseBodyData) *ListInterveneRulesResponseBody
	GetData() *ListInterveneRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListInterveneRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInterveneRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInterveneRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInterveneRulesResponseBody
	GetSuccess() *bool
}

type ListInterveneRulesResponseBody struct {
	// example:
	//
	// 0
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInterveneRulesResponseBody) GetData() *ListInterveneRulesResponseBodyData {
	return s.Data
}

func (s *ListInterveneRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInterveneRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInterveneRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterveneRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInterveneRulesResponseBody) SetCode(v string) *ListInterveneRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetData(v *ListInterveneRulesResponseBodyData) *ListInterveneRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneRulesResponseBody) SetHttpStatusCode(v int32) *ListInterveneRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetMessage(v string) *ListInterveneRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetRequestId(v string) *ListInterveneRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetSuccess(v bool) *ListInterveneRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInterveneRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInterveneRulesResponseBodyData struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Count             *int64                                                 `json:"Count,omitempty" xml:"Count,omitempty"`
	InterveneRuleList []*ListInterveneRulesResponseBodyDataInterveneRuleList `json:"InterveneRuleList,omitempty" xml:"InterveneRuleList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListInterveneRulesResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *ListInterveneRulesResponseBodyData) GetInterveneRuleList() []*ListInterveneRulesResponseBodyDataInterveneRuleList {
	return s.InterveneRuleList
}

func (s *ListInterveneRulesResponseBodyData) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneRulesResponseBodyData) SetCode(v int32) *ListInterveneRulesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetCount(v int64) *ListInterveneRulesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetInterveneRuleList(v []*ListInterveneRulesResponseBodyDataInterveneRuleList) *ListInterveneRulesResponseBodyData {
	s.InterveneRuleList = v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageIndex(v int32) *ListInterveneRulesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageSize(v int32) *ListInterveneRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) Validate() error {
	if s.InterveneRuleList != nil {
		for _, item := range s.InterveneRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterveneRulesResponseBodyDataInterveneRuleList struct {
	AnswerConfig []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-06-05 15:17:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2023-04-03 02:42:01
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// mr-iuo9pi9w555phfbb
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// ruletest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetAnswerConfig() []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	return s.AnswerConfig
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetEffectTime() *string {
	return s.EffectTime
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetInterveneType() *int32 {
	return s.InterveneType
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetAnswerConfig(v []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.AnswerConfig = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetCreateTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.CreateTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetEffectTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.EffectTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetInterveneType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.InterveneType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetNamespaceList(v []*string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.NamespaceList = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleId(v int64) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleId = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleName(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleName = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) Validate() error {
	if s.AnswerConfig != nil {
		for _, item := range s.AnswerConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig struct {
	// example:
	//
	// 0
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GetAnswerType() *int32 {
	return s.AnswerType
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GetMessage() *string {
	return s.Message
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetAnswerType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetMessage(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetNamespace(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Namespace = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) Validate() error {
	return dara.Validate(s)
}
