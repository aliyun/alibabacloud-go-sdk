// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPulishAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfig(v *CreateAndPulishAgentRequestApplicationConfig) *CreateAndPulishAgentRequest
	GetApplicationConfig() *CreateAndPulishAgentRequestApplicationConfig
	SetInstructions(v string) *CreateAndPulishAgentRequest
	GetInstructions() *string
	SetModelId(v string) *CreateAndPulishAgentRequest
	GetModelId() *string
	SetName(v string) *CreateAndPulishAgentRequest
	GetName() *string
	SetSampleLibrary(v *CreateAndPulishAgentRequestSampleLibrary) *CreateAndPulishAgentRequest
	GetSampleLibrary() *CreateAndPulishAgentRequestSampleLibrary
}

type CreateAndPulishAgentRequest struct {
	ApplicationConfig *CreateAndPulishAgentRequestApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Instructions      *string                                       `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                       `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibrary     *CreateAndPulishAgentRequestSampleLibrary     `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty" type:"Struct"`
}

func (s CreateAndPulishAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequest) GetApplicationConfig() *CreateAndPulishAgentRequestApplicationConfig {
	return s.ApplicationConfig
}

func (s *CreateAndPulishAgentRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *CreateAndPulishAgentRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateAndPulishAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateAndPulishAgentRequest) GetSampleLibrary() *CreateAndPulishAgentRequestSampleLibrary {
	return s.SampleLibrary
}

func (s *CreateAndPulishAgentRequest) SetApplicationConfig(v *CreateAndPulishAgentRequestApplicationConfig) *CreateAndPulishAgentRequest {
	s.ApplicationConfig = v
	return s
}

func (s *CreateAndPulishAgentRequest) SetInstructions(v string) *CreateAndPulishAgentRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAndPulishAgentRequest) SetModelId(v string) *CreateAndPulishAgentRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAndPulishAgentRequest) SetName(v string) *CreateAndPulishAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateAndPulishAgentRequest) SetSampleLibrary(v *CreateAndPulishAgentRequestSampleLibrary) *CreateAndPulishAgentRequest {
	s.SampleLibrary = v
	return s
}

func (s *CreateAndPulishAgentRequest) Validate() error {
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

type CreateAndPulishAgentRequestApplicationConfig struct {
	HistoryConfig  *CreateAndPulishAgentRequestApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *CreateAndPulishAgentRequestApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *CreateAndPulishAgentRequestApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *CreateAndPulishAgentRequestApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	SecurityConfig *CreateAndPulishAgentRequestApplicationConfigSecurityConfig `json:"securityConfig,omitempty" xml:"securityConfig,omitempty" type:"Struct"`
	Tools          []*CreateAndPulishAgentRequestApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*CreateAndPulishAgentRequestApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s CreateAndPulishAgentRequestApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetHistoryConfig() *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	return s.HistoryConfig
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetLongTermMemory() *CreateAndPulishAgentRequestApplicationConfigLongTermMemory {
	return s.LongTermMemory
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetParameters() *CreateAndPulishAgentRequestApplicationConfigParameters {
	return s.Parameters
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetRagConfig() *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	return s.RagConfig
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetSecurityConfig() *CreateAndPulishAgentRequestApplicationConfigSecurityConfig {
	return s.SecurityConfig
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetTools() []*CreateAndPulishAgentRequestApplicationConfigTools {
	return s.Tools
}

func (s *CreateAndPulishAgentRequestApplicationConfig) GetWorkFlows() []*CreateAndPulishAgentRequestApplicationConfigWorkFlows {
	return s.WorkFlows
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetHistoryConfig(v *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetLongTermMemory(v *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) *CreateAndPulishAgentRequestApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetParameters(v *CreateAndPulishAgentRequestApplicationConfigParameters) *CreateAndPulishAgentRequestApplicationConfig {
	s.Parameters = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetRagConfig(v *CreateAndPulishAgentRequestApplicationConfigRagConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetSecurityConfig(v *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.SecurityConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetTools(v []*CreateAndPulishAgentRequestApplicationConfigTools) *CreateAndPulishAgentRequestApplicationConfig {
	s.Tools = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetWorkFlows(v []*CreateAndPulishAgentRequestApplicationConfigWorkFlows) *CreateAndPulishAgentRequestApplicationConfig {
	s.WorkFlows = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) Validate() error {
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
	if s.SecurityConfig != nil {
		if err := s.SecurityConfig.Validate(); err != nil {
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

type CreateAndPulishAgentRequestApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GetEnableAdbRecord() *bool {
	return s.EnableAdbRecord
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GetEnableRecord() *bool {
	return s.EnableRecord
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GetRegion() *string {
	return s.Region
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GetStoreCode() *string {
	return s.StoreCode
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetEnableRecord(v bool) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetInstanceId(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetRegion(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetStoreCode(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigLongTermMemory) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) GetEnable() *bool {
	return s.Enable
}

func (s *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) SetEnable(v bool) *CreateAndPulishAgentRequestApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigParameters struct {
	DialogRound    *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	EnableThinking *bool    `json:"enable_thinking,omitempty" xml:"enable_thinking,omitempty"`
	MaxTokens      *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature    *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) GetDialogRound() *int32 {
	return s.DialogRound
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) GetTemperature() *float64 {
	return s.Temperature
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetDialogRound(v int32) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetEnableThinking(v bool) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.EnableThinking = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetMaxTokens(v int32) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetTemperature(v float64) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.Temperature = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigRagConfig struct {
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

func (s CreateAndPulishAgentRequestApplicationConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetAnswerScope() *string {
	return s.AnswerScope
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetFixedReplyDetail() *string {
	return s.FixedReplyDetail
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetRagRejectType() *string {
	return s.RagRejectType
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetRejectFilterPrompt() *string {
	return s.RejectFilterPrompt
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetRejectFilterType() *string {
	return s.RejectFilterType
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetAnswerScope(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.AnswerScope = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetEnableCitation(v bool) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetEnableSearch(v bool) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetEnableWebSearch(v bool) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.EnableWebSearch = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetFixedReplyDetail(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.FixedReplyDetail = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetPromptStrategy(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.PromptStrategy = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetRagRejectType(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.RagRejectType = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetRejectFilterPrompt(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.RejectFilterPrompt = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetRejectFilterType(v string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.RejectFilterType = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetRetrieveMaxLength(v int32) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.RetrieveMaxLength = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetTopK(v int32) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigSecurityConfig struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigSecurityConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigSecurityConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) GetProcessingStrategy() *string {
	return s.ProcessingStrategy
}

func (s *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) SetProcessingStrategy(v string) *CreateAndPulishAgentRequestApplicationConfigSecurityConfig {
	s.ProcessingStrategy = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigTools) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigTools) GetType() *string {
	return s.Type
}

func (s *CreateAndPulishAgentRequestApplicationConfigTools) SetType(v string) *CreateAndPulishAgentRequestApplicationConfigTools {
	s.Type = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigTools) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigWorkFlows) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigWorkFlows) GetType() *string {
	return s.Type
}

func (s *CreateAndPulishAgentRequestApplicationConfigWorkFlows) SetType(v string) *CreateAndPulishAgentRequestApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigWorkFlows) Validate() error {
	return dara.Validate(s)
}

type CreateAndPulishAgentRequestSampleLibrary struct {
	EnableSample        *bool     `json:"enableSample,omitempty" xml:"enableSample,omitempty"`
	SampleLibraryIdList []*string `json:"sampleLibraryIdList,omitempty" xml:"sampleLibraryIdList,omitempty" type:"Repeated"`
	TopK                *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateAndPulishAgentRequestSampleLibrary) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentRequestSampleLibrary) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestSampleLibrary) GetEnableSample() *bool {
	return s.EnableSample
}

func (s *CreateAndPulishAgentRequestSampleLibrary) GetSampleLibraryIdList() []*string {
	return s.SampleLibraryIdList
}

func (s *CreateAndPulishAgentRequestSampleLibrary) GetTopK() *int32 {
	return s.TopK
}

func (s *CreateAndPulishAgentRequestSampleLibrary) SetEnableSample(v bool) *CreateAndPulishAgentRequestSampleLibrary {
	s.EnableSample = &v
	return s
}

func (s *CreateAndPulishAgentRequestSampleLibrary) SetSampleLibraryIdList(v []*string) *CreateAndPulishAgentRequestSampleLibrary {
	s.SampleLibraryIdList = v
	return s
}

func (s *CreateAndPulishAgentRequestSampleLibrary) SetTopK(v int32) *CreateAndPulishAgentRequestSampleLibrary {
	s.TopK = &v
	return s
}

func (s *CreateAndPulishAgentRequestSampleLibrary) Validate() error {
	return dara.Validate(s)
}
