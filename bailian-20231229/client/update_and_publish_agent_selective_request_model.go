// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentSelectiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfig(v *UpdateAndPublishAgentSelectiveRequestApplicationConfig) *UpdateAndPublishAgentSelectiveRequest
	GetApplicationConfig() *UpdateAndPublishAgentSelectiveRequestApplicationConfig
	SetInstructions(v string) *UpdateAndPublishAgentSelectiveRequest
	GetInstructions() *string
	SetModelId(v string) *UpdateAndPublishAgentSelectiveRequest
	GetModelId() *string
	SetName(v string) *UpdateAndPublishAgentSelectiveRequest
	GetName() *string
	SetSampleLibrary(v *UpdateAndPublishAgentSelectiveRequestSampleLibrary) *UpdateAndPublishAgentSelectiveRequest
	GetSampleLibrary() *UpdateAndPublishAgentSelectiveRequestSampleLibrary
}

type UpdateAndPublishAgentSelectiveRequest struct {
	ApplicationConfig *UpdateAndPublishAgentSelectiveRequestApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Instructions      *string                                                 `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                                 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibrary     *UpdateAndPublishAgentSelectiveRequestSampleLibrary     `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty" type:"Struct"`
}

func (s UpdateAndPublishAgentSelectiveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequest) GetApplicationConfig() *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	return s.ApplicationConfig
}

func (s *UpdateAndPublishAgentSelectiveRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAndPublishAgentSelectiveRequest) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateAndPublishAgentSelectiveRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAndPublishAgentSelectiveRequest) GetSampleLibrary() *UpdateAndPublishAgentSelectiveRequestSampleLibrary {
	return s.SampleLibrary
}

func (s *UpdateAndPublishAgentSelectiveRequest) SetApplicationConfig(v *UpdateAndPublishAgentSelectiveRequestApplicationConfig) *UpdateAndPublishAgentSelectiveRequest {
	s.ApplicationConfig = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequest) SetInstructions(v string) *UpdateAndPublishAgentSelectiveRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequest) SetModelId(v string) *UpdateAndPublishAgentSelectiveRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequest) SetName(v string) *UpdateAndPublishAgentSelectiveRequest {
	s.Name = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequest) SetSampleLibrary(v *UpdateAndPublishAgentSelectiveRequestSampleLibrary) *UpdateAndPublishAgentSelectiveRequest {
	s.SampleLibrary = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequest) Validate() error {
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

type UpdateAndPublishAgentSelectiveRequestApplicationConfig struct {
	HistoryConfig  *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*UpdateAndPublishAgentSelectiveRequestApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetHistoryConfig() *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	return s.HistoryConfig
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetLongTermMemory() *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory {
	return s.LongTermMemory
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetParameters() *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters {
	return s.Parameters
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetRagConfig() *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	return s.RagConfig
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetSecurity() *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity {
	return s.Security
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetTools() []*UpdateAndPublishAgentSelectiveRequestApplicationConfigTools {
	return s.Tools
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) GetWorkFlows() []*UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows {
	return s.WorkFlows
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetHistoryConfig(v *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetLongTermMemory(v *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetParameters(v *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.Parameters = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetRagConfig(v *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetSecurity(v *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.Security = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetTools(v []*UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.Tools = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) SetWorkFlows(v []*UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) *UpdateAndPublishAgentSelectiveRequestApplicationConfig {
	s.WorkFlows = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfig) Validate() error {
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

type UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GetEnableAdbRecord() *bool {
	return s.EnableAdbRecord
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GetEnableRecord() *bool {
	return s.EnableRecord
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GetRegion() *string {
	return s.Region
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) GetStoreCode() *string {
	return s.StoreCode
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) SetEnableRecord(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) SetInstanceId(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) SetRegion(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) SetStoreCode(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) SetEnable(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigLongTermMemory) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters struct {
	DialogRound    *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	EnableThinking *bool    `json:"enable_thinking,omitempty" xml:"enable_thinking,omitempty"`
	MaxTokens      *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature    *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) GetDialogRound() *int32 {
	return s.DialogRound
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) GetTemperature() *float64 {
	return s.Temperature
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) SetDialogRound(v int32) *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) SetEnableThinking(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters {
	s.EnableThinking = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) SetMaxTokens(v int32) *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) SetTemperature(v float64) *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters {
	s.Temperature = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig struct {
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

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetAnswerScope() *string {
	return s.AnswerScope
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetFixedReplyDetail() *string {
	return s.FixedReplyDetail
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetRagRejectType() *string {
	return s.RagRejectType
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetRejectFilterPrompt() *string {
	return s.RejectFilterPrompt
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetRejectFilterType() *string {
	return s.RejectFilterType
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetAnswerScope(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.AnswerScope = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetEnableCitation(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetEnableSearch(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetEnableWebSearch(v bool) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.EnableWebSearch = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetFixedReplyDetail(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.FixedReplyDetail = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetPromptStrategy(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.PromptStrategy = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetRagRejectType(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.RagRejectType = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetRejectFilterPrompt(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.RejectFilterPrompt = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetRejectFilterType(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.RejectFilterType = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetRetrieveMaxLength(v int32) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.RetrieveMaxLength = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) SetTopK(v int32) *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) GetProcessingStrategy() *string {
	return s.ProcessingStrategy
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) SetProcessingStrategy(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigSecurity) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) GetType() *string {
	return s.Type
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) SetType(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigTools {
	s.Type = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigTools) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) GetType() *string {
	return s.Type
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) SetType(v string) *UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestApplicationConfigWorkFlows) Validate() error {
	return dara.Validate(s)
}

type UpdateAndPublishAgentSelectiveRequestSampleLibrary struct {
	EnableSample        *bool     `json:"enableSample,omitempty" xml:"enableSample,omitempty"`
	SampleLibraryIdList []*string `json:"sampleLibraryIdList,omitempty" xml:"sampleLibraryIdList,omitempty" type:"Repeated"`
	TopK                *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveRequestSampleLibrary) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveRequestSampleLibrary) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) GetEnableSample() *bool {
	return s.EnableSample
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) GetSampleLibraryIdList() []*string {
	return s.SampleLibraryIdList
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) SetEnableSample(v bool) *UpdateAndPublishAgentSelectiveRequestSampleLibrary {
	s.EnableSample = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) SetSampleLibraryIdList(v []*string) *UpdateAndPublishAgentSelectiveRequestSampleLibrary {
	s.SampleLibraryIdList = v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) SetTopK(v int32) *UpdateAndPublishAgentSelectiveRequestSampleLibrary {
	s.TopK = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveRequestSampleLibrary) Validate() error {
	return dara.Validate(s)
}
