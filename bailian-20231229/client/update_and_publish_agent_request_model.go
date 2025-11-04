// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfig(v *UpdateAndPublishAgentRequestApplicationConfig) *UpdateAndPublishAgentRequest
	GetApplicationConfig() *UpdateAndPublishAgentRequestApplicationConfig
	SetInstructions(v string) *UpdateAndPublishAgentRequest
	GetInstructions() *string
	SetModelId(v string) *UpdateAndPublishAgentRequest
	GetModelId() *string
	SetName(v string) *UpdateAndPublishAgentRequest
	GetName() *string
	SetSampleLibrary(v *UpdateAndPublishAgentRequestSampleLibrary) *UpdateAndPublishAgentRequest
	GetSampleLibrary() *UpdateAndPublishAgentRequestSampleLibrary
}

type UpdateAndPublishAgentRequest struct {
	ApplicationConfig *UpdateAndPublishAgentRequestApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Instructions      *string                                        `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                        `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibrary     *UpdateAndPublishAgentRequestSampleLibrary     `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty" type:"Struct"`
}

func (s UpdateAndPublishAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequest) GetApplicationConfig() *UpdateAndPublishAgentRequestApplicationConfig {
	return s.ApplicationConfig
}

func (s *UpdateAndPublishAgentRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAndPublishAgentRequest) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateAndPublishAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAndPublishAgentRequest) GetSampleLibrary() *UpdateAndPublishAgentRequestSampleLibrary {
	return s.SampleLibrary
}

func (s *UpdateAndPublishAgentRequest) SetApplicationConfig(v *UpdateAndPublishAgentRequestApplicationConfig) *UpdateAndPublishAgentRequest {
	s.ApplicationConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetInstructions(v string) *UpdateAndPublishAgentRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetModelId(v string) *UpdateAndPublishAgentRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetName(v string) *UpdateAndPublishAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetSampleLibrary(v *UpdateAndPublishAgentRequestSampleLibrary) *UpdateAndPublishAgentRequest {
	s.SampleLibrary = v
	return s
}

func (s *UpdateAndPublishAgentRequest) Validate() error {
	if s.ApplicationConfig != nil {
		if err := s.ApplicationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SampleLibrary != nil {
		if err := s.SampleLibrary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAndPublishAgentRequestApplicationConfig struct {
	HistoryConfig  *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *UpdateAndPublishAgentRequestApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *UpdateAndPublishAgentRequestApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *UpdateAndPublishAgentRequestApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*UpdateAndPublishAgentRequestApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*UpdateAndPublishAgentRequestApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s UpdateAndPublishAgentRequestApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetHistoryConfig() *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	return s.HistoryConfig
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetLongTermMemory() *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory {
	return s.LongTermMemory
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetParameters() *UpdateAndPublishAgentRequestApplicationConfigParameters {
	return s.Parameters
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetRagConfig() *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	return s.RagConfig
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetSecurity() *UpdateAndPublishAgentRequestApplicationConfigSecurity {
	return s.Security
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetTools() []*UpdateAndPublishAgentRequestApplicationConfigTools {
	return s.Tools
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) GetWorkFlows() []*UpdateAndPublishAgentRequestApplicationConfigWorkFlows {
	return s.WorkFlows
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetHistoryConfig(v *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) *UpdateAndPublishAgentRequestApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetLongTermMemory(v *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) *UpdateAndPublishAgentRequestApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetParameters(v *UpdateAndPublishAgentRequestApplicationConfigParameters) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Parameters = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetRagConfig(v *UpdateAndPublishAgentRequestApplicationConfigRagConfig) *UpdateAndPublishAgentRequestApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetSecurity(v *UpdateAndPublishAgentRequestApplicationConfigSecurity) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Security = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetTools(v []*UpdateAndPublishAgentRequestApplicationConfigTools) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Tools = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetWorkFlows(v []*UpdateAndPublishAgentRequestApplicationConfigWorkFlows) *UpdateAndPublishAgentRequestApplicationConfig {
	s.WorkFlows = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) Validate() error {
	if s.HistoryConfig != nil {
		if err := s.HistoryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LongTermMemory != nil {
		if err := s.LongTermMemory.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Security != nil {
		if err := s.Security.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkFlows != nil {
		for _, item := range s.WorkFlows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAndPublishAgentRequestApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GetEnableAdbRecord() *bool {
	return s.EnableAdbRecord
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GetEnableRecord() *bool {
	return s.EnableRecord
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GetRegion() *string {
	return s.Region
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GetStoreCode() *string {
	return s.StoreCode
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetEnableRecord(v bool) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetInstanceId(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetRegion(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetStoreCode(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) SetEnable(v bool) *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigParameters struct {
	DialogRound    *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	EnableThinking *bool    `json:"enable_thinking,omitempty" xml:"enable_thinking,omitempty"`
	MaxTokens      *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature    *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) GetDialogRound() *int32 {
	return s.DialogRound
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) GetTemperature() *float64 {
	return s.Temperature
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetDialogRound(v int32) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetEnableThinking(v bool) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.EnableThinking = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetMaxTokens(v int32) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetTemperature(v float64) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.Temperature = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigRagConfig struct {
	AnswerScope           *string   `json:"answerScope,omitempty" xml:"answerScope,omitempty"`
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	EnableWebSearch       *bool     `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	FixedReplyDetail      *string   `json:"fixedReplyDetail,omitempty" xml:"fixedReplyDetail,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	PromptStrategy        *string   `json:"promptStrategy,omitempty" xml:"promptStrategy,omitempty"`
	RagRejectType         *string   `json:"ragRejectType,omitempty" xml:"ragRejectType,omitempty"`
	RejectFilterPrompt    *string   `json:"rejectFilterPrompt,omitempty" xml:"rejectFilterPrompt,omitempty"`
	RejectFilterType      *string   `json:"rejectFilterType,omitempty" xml:"rejectFilterType,omitempty"`
	RetrieveMaxLength     *int32    `json:"retrieveMaxLength,omitempty" xml:"retrieveMaxLength,omitempty"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetAnswerScope() *string {
	return s.AnswerScope
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetFixedReplyDetail() *string {
	return s.FixedReplyDetail
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetRagRejectType() *string {
	return s.RagRejectType
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetRejectFilterPrompt() *string {
	return s.RejectFilterPrompt
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetRejectFilterType() *string {
	return s.RejectFilterType
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetAnswerScope(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.AnswerScope = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetEnableCitation(v bool) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetEnableSearch(v bool) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetEnableWebSearch(v bool) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.EnableWebSearch = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetFixedReplyDetail(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.FixedReplyDetail = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetPromptStrategy(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.PromptStrategy = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetRagRejectType(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.RagRejectType = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetRejectFilterPrompt(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.RejectFilterPrompt = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetRejectFilterType(v string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.RejectFilterType = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetRetrieveMaxLength(v int32) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.RetrieveMaxLength = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetTopK(v int32) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigSecurity) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigSecurity) GetProcessingStrategy() *string {
	return s.ProcessingStrategy
}

func (s *UpdateAndPublishAgentRequestApplicationConfigSecurity) SetProcessingStrategy(v string) *UpdateAndPublishAgentRequestApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigSecurity) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigTools) GetType() *string {
	return s.Type
}

func (s *UpdateAndPublishAgentRequestApplicationConfigTools) SetType(v string) *UpdateAndPublishAgentRequestApplicationConfigTools {
	s.Type = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigTools) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigWorkFlows) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigWorkFlows) GetType() *string {
	return s.Type
}

func (s *UpdateAndPublishAgentRequestApplicationConfigWorkFlows) SetType(v string) *UpdateAndPublishAgentRequestApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigWorkFlows) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentRequestSampleLibrary struct {
	EnableSample        *bool     `json:"enableSample,omitempty" xml:"enableSample,omitempty"`
	SampleLibraryIdList []*string `json:"sampleLibraryIdList,omitempty" xml:"sampleLibraryIdList,omitempty" type:"Repeated"`
	TopK                *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateAndPublishAgentRequestSampleLibrary) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentRequestSampleLibrary) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) GetEnableSample() *bool {
	return s.EnableSample
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) GetSampleLibraryIdList() []*string {
	return s.SampleLibraryIdList
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) SetEnableSample(v bool) *UpdateAndPublishAgentRequestSampleLibrary {
	s.EnableSample = &v
	return s
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) SetSampleLibraryIdList(v []*string) *UpdateAndPublishAgentRequestSampleLibrary {
	s.SampleLibraryIdList = v
	return s
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) SetTopK(v int32) *UpdateAndPublishAgentRequestSampleLibrary {
	s.TopK = &v
	return s
}

func (s *UpdateAndPublishAgentRequestSampleLibrary) Validate() error {
	return dara.Validate(s)
}
