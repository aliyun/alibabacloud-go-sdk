// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *InsertInterveneRuleRequest
	GetAgentKey() *string
	SetInterveneRuleConfig(v *InsertInterveneRuleRequestInterveneRuleConfig) *InsertInterveneRuleRequest
	GetInterveneRuleConfig() *InsertInterveneRuleRequestInterveneRuleConfig
}

type InsertInterveneRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey            *string                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfig *InsertInterveneRuleRequestInterveneRuleConfig `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty" type:"Struct"`
}

func (s InsertInterveneRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *InsertInterveneRuleRequest) GetInterveneRuleConfig() *InsertInterveneRuleRequestInterveneRuleConfig {
	return s.InterveneRuleConfig
}

func (s *InsertInterveneRuleRequest) SetAgentKey(v string) *InsertInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleRequest) SetInterveneRuleConfig(v *InsertInterveneRuleRequestInterveneRuleConfig) *InsertInterveneRuleRequest {
	s.InterveneRuleConfig = v
	return s
}

func (s *InsertInterveneRuleRequest) Validate() error {
	if s.InterveneRuleConfig != nil {
		if err := s.InterveneRuleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertInterveneRuleRequestInterveneRuleConfig struct {
	AnswerConfig        []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig        `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig        *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig          `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	InterveneConfigList []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList `json:"InterveneConfigList,omitempty" xml:"InterveneConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// tf-test-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetAnswerConfig() []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	return s.AnswerConfig
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetEffectConfig() *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	return s.EffectConfig
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetInterveneConfigList() []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	return s.InterveneConfigList
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetInterveneType() *int32 {
	return s.InterveneType
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetRuleId() *int64 {
	return s.RuleId
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) GetRuleName() *string {
	return s.RuleName
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetAnswerConfig(v []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.AnswerConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetEffectConfig(v *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.EffectConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneConfigList(v []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneConfigList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneType(v int32) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetNamespaceList(v []*string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.NamespaceList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleId(v int64) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleId = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleName(v string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleName = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) Validate() error {
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
	if s.InterveneConfigList != nil {
		for _, item := range s.InterveneConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig struct {
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

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GetAnswerType() *int32 {
	return s.AnswerType
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GetMessage() *string {
	return s.Message
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetAnswerType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetMessage(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetNamespace(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Namespace = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) Validate() error {
	return dara.Validate(s)
}

type InsertInterveneRuleRequestInterveneRuleConfigEffectConfig struct {
	// example:
	//
	// 0
	EffectType *int32 `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	// example:
	//
	// 2023-03-28 06:04:29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-03-28 06:04:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GetEffectType() *int32 {
	return s.EffectType
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEffectType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EffectType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEndTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EndTime = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetStartTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.StartTime = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) Validate() error {
	return dara.Validate(s)
}

type InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList struct {
	// id
	//
	// example:
	//
	// 37249
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	OperationType *int32  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GetId() *string {
	return s.Id
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GetOperationType() *int32 {
	return s.OperationType
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GetQuery() *string {
	return s.Query
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetId(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Id = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetOperationType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.OperationType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetQuery(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Query = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) Validate() error {
	return dara.Validate(s)
}
