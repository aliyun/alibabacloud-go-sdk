// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInterveneRuleDetailResponseBody
	GetCode() *string
	SetData(v *GetInterveneRuleDetailResponseBodyData) *GetInterveneRuleDetailResponseBody
	GetData() *GetInterveneRuleDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetInterveneRuleDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInterveneRuleDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInterveneRuleDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInterveneRuleDetailResponseBody
	GetSuccess() *bool
}

type GetInterveneRuleDetailResponseBody struct {
	// example:
	//
	// 0
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInterveneRuleDetailResponseBody) GetData() *GetInterveneRuleDetailResponseBodyData {
	return s.Data
}

func (s *GetInterveneRuleDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInterveneRuleDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterveneRuleDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInterveneRuleDetailResponseBody) SetCode(v string) *GetInterveneRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetData(v *GetInterveneRuleDetailResponseBodyData) *GetInterveneRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetHttpStatusCode(v int32) *GetInterveneRuleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetMessage(v string) *GetInterveneRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetRequestId(v string) *GetInterveneRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetSuccess(v bool) *GetInterveneRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneRuleDetailResponseBodyData struct {
	Code                *int32                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	InterveneRuleDetail *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail `json:"InterveneRuleDetail,omitempty" xml:"InterveneRuleDetail,omitempty" type:"Struct"`
}

func (s GetInterveneRuleDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *GetInterveneRuleDetailResponseBodyData) GetInterveneRuleDetail() *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	return s.InterveneRuleDetail
}

func (s *GetInterveneRuleDetailResponseBodyData) SetCode(v int32) *GetInterveneRuleDetailResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyData) SetInterveneRuleDetail(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) *GetInterveneRuleDetailResponseBodyData {
	s.InterveneRuleDetail = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyData) Validate() error {
	if s.InterveneRuleDetail != nil {
		if err := s.InterveneRuleDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail struct {
	AnswerConfig []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig   `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// 100418
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetAnswerConfig() []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	return s.AnswerConfig
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetEffectConfig() *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	return s.EffectConfig
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetInterveneType() *int32 {
	return s.InterveneType
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GetRuleName() *string {
	return s.RuleName
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetAnswerConfig(v []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.AnswerConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetEffectConfig(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.EffectConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetInterveneType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.InterveneType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetNamespaceList(v []*string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.NamespaceList = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleId(v int64) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleName(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleName = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) Validate() error {
	if s.AnswerConfig != nil {
		for _, item := range s.AnswerConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EffectConfig != nil {
		if err := s.EffectConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig struct {
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

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GetAnswerType() *int32 {
	return s.AnswerType
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetAnswerType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetMessage(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetNamespace(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Namespace = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) Validate() error {
	return dara.Validate(s)
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig struct {
	// example:
	//
	// 0
	EffectType *int32 `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	// example:
	//
	// 2023-11-25 14:21:15
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-11-25 14:21:15
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GetEffectType() *int32 {
	return s.EffectType
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEffectType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EffectType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEndTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EndTime = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetStartTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.StartTime = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) Validate() error {
	return dara.Validate(s)
}
