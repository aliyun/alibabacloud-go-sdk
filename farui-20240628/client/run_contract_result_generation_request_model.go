// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractResultGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractResultGenerationRequest
	GetAppId() *string
	SetAssistant(v *RunContractResultGenerationRequestAssistant) *RunContractResultGenerationRequest
	GetAssistant() *RunContractResultGenerationRequestAssistant
	SetStream(v bool) *RunContractResultGenerationRequest
	GetStream() *bool
}

type RunContractResultGenerationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                      `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunContractResultGenerationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractResultGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractResultGenerationRequest) GetAssistant() *RunContractResultGenerationRequestAssistant {
	return s.Assistant
}

func (s *RunContractResultGenerationRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunContractResultGenerationRequest) SetAppId(v string) *RunContractResultGenerationRequest {
	s.AppId = &v
	return s
}

func (s *RunContractResultGenerationRequest) SetAssistant(v *RunContractResultGenerationRequestAssistant) *RunContractResultGenerationRequest {
	s.Assistant = v
	return s
}

func (s *RunContractResultGenerationRequest) SetStream(v bool) *RunContractResultGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunContractResultGenerationRequest) Validate() error {
	if s.Assistant != nil {
		if err := s.Assistant.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractResultGenerationRequestAssistant struct {
	MetaData *RunContractResultGenerationRequestAssistantMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// example:
	//
	// contract_examime
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunContractResultGenerationRequestAssistant) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistant) GetMetaData() *RunContractResultGenerationRequestAssistantMetaData {
	return s.MetaData
}

func (s *RunContractResultGenerationRequestAssistant) GetType() *string {
	return s.Type
}

func (s *RunContractResultGenerationRequestAssistant) GetVersion() *string {
	return s.Version
}

func (s *RunContractResultGenerationRequestAssistant) SetMetaData(v *RunContractResultGenerationRequestAssistantMetaData) *RunContractResultGenerationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunContractResultGenerationRequestAssistant) SetType(v string) *RunContractResultGenerationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistant) SetVersion(v string) *RunContractResultGenerationRequestAssistant {
	s.Version = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistant) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractResultGenerationRequestAssistantMetaData struct {
	CustomRuleConfig *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig `json:"customRuleConfig,omitempty" xml:"customRuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// 9a6b1ba60d9944249363ec3cc1529b7b
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// b265b416-ca1f-425d-9340-c968f39624e1
	RuleTaskId *string                                                     `json:"ruleTaskId,omitempty" xml:"ruleTaskId,omitempty"`
	Rules      []*RunContractResultGenerationRequestAssistantMetaDataRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationRequestAssistantMetaData) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaData) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaData) GetCustomRuleConfig() *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig {
	return s.CustomRuleConfig
}

func (s *RunContractResultGenerationRequestAssistantMetaData) GetFileId() *string {
	return s.FileId
}

func (s *RunContractResultGenerationRequestAssistantMetaData) GetPosition() *string {
	return s.Position
}

func (s *RunContractResultGenerationRequestAssistantMetaData) GetRuleTaskId() *string {
	return s.RuleTaskId
}

func (s *RunContractResultGenerationRequestAssistantMetaData) GetRules() []*RunContractResultGenerationRequestAssistantMetaDataRules {
	return s.Rules
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetCustomRuleConfig(v *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) *RunContractResultGenerationRequestAssistantMetaData {
	s.CustomRuleConfig = v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetFileId(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.FileId = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetPosition(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.Position = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetRuleTaskId(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.RuleTaskId = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetRules(v []*RunContractResultGenerationRequestAssistantMetaDataRules) *RunContractResultGenerationRequestAssistantMetaData {
	s.Rules = v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) Validate() error {
	if s.CustomRuleConfig != nil {
		if err := s.CustomRuleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig struct {
	CustomRules []*RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules `json:"customRules,omitempty" xml:"customRules,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) GetCustomRules() []*RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	return s.CustomRules
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) SetCustomRules(v []*RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig {
	s.CustomRules = v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) Validate() error {
	if s.CustomRules != nil {
		for _, item := range s.CustomRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules struct {
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RuleDesc  *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	RuleTitle *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) GetRuleTitle() *string {
	return s.RuleTitle
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRiskLevel(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRuleDesc(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RuleDesc = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRuleTitle(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RuleTitle = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) Validate() error {
	return dara.Validate(s)
}

type RunContractResultGenerationRequestAssistantMetaDataRules struct {
	// example:
	//
	// medium
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 2.1
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataRules) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataRules) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) GetRuleSequence() *string {
	return s.RuleSequence
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) GetRuleTag() *string {
	return s.RuleTag
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) GetRuleTitle() *string {
	return s.RuleTitle
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRiskLevel(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleSequence(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleSequence = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleTag(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleTag = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleTitle(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleTitle = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) Validate() error {
	return dara.Validate(s)
}
