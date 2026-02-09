// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentContext(v *RunSearchGenerationRequestAgentContext) *RunSearchGenerationRequest
	GetAgentContext() *RunSearchGenerationRequestAgentContext
	SetChatConfig(v *RunSearchGenerationRequestChatConfig) *RunSearchGenerationRequest
	GetChatConfig() *RunSearchGenerationRequestChatConfig
	SetFileUrl(v string) *RunSearchGenerationRequest
	GetFileUrl() *string
	SetModelId(v string) *RunSearchGenerationRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunSearchGenerationRequest
	GetOriginalSessionId() *string
	SetPrompt(v string) *RunSearchGenerationRequest
	GetPrompt() *string
	SetTaskId(v string) *RunSearchGenerationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunSearchGenerationRequest
	GetWorkspaceId() *string
}

type RunSearchGenerationRequest struct {
	AgentContext *RunSearchGenerationRequestAgentContext `json:"AgentContext,omitempty" xml:"AgentContext,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	ChatConfig *RunSearchGenerationRequestChatConfig `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty" type:"Struct"`
	// example:
	//
	// http://xxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// qwen-max-latest
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// xxx
	OriginalSessionId *string `json:"OriginalSessionId,omitempty" xml:"OriginalSessionId,omitempty"`
	// example:
	//
	// xxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunSearchGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequest) GetAgentContext() *RunSearchGenerationRequestAgentContext {
	return s.AgentContext
}

func (s *RunSearchGenerationRequest) GetChatConfig() *RunSearchGenerationRequestChatConfig {
	return s.ChatConfig
}

func (s *RunSearchGenerationRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunSearchGenerationRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunSearchGenerationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunSearchGenerationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunSearchGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunSearchGenerationRequest) SetAgentContext(v *RunSearchGenerationRequestAgentContext) *RunSearchGenerationRequest {
	s.AgentContext = v
	return s
}

func (s *RunSearchGenerationRequest) SetChatConfig(v *RunSearchGenerationRequestChatConfig) *RunSearchGenerationRequest {
	s.ChatConfig = v
	return s
}

func (s *RunSearchGenerationRequest) SetFileUrl(v string) *RunSearchGenerationRequest {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationRequest) SetModelId(v string) *RunSearchGenerationRequest {
	s.ModelId = &v
	return s
}

func (s *RunSearchGenerationRequest) SetOriginalSessionId(v string) *RunSearchGenerationRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunSearchGenerationRequest) SetPrompt(v string) *RunSearchGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunSearchGenerationRequest) SetTaskId(v string) *RunSearchGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunSearchGenerationRequest) SetWorkspaceId(v string) *RunSearchGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunSearchGenerationRequest) Validate() error {
	if s.AgentContext != nil {
		if err := s.AgentContext.Validate(); err != nil {
			return err
		}
	}
	if s.ChatConfig != nil {
		if err := s.ChatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationRequestAgentContext struct {
	BizContext *RunSearchGenerationRequestAgentContextBizContext `json:"BizContext,omitempty" xml:"BizContext,omitempty" type:"Struct"`
}

func (s RunSearchGenerationRequestAgentContext) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestAgentContext) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestAgentContext) GetBizContext() *RunSearchGenerationRequestAgentContextBizContext {
	return s.BizContext
}

func (s *RunSearchGenerationRequestAgentContext) SetBizContext(v *RunSearchGenerationRequestAgentContextBizContext) *RunSearchGenerationRequestAgentContext {
	s.BizContext = v
	return s
}

func (s *RunSearchGenerationRequestAgentContext) Validate() error {
	if s.BizContext != nil {
		if err := s.BizContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationRequestAgentContextBizContext struct {
	AskUser                  *string                                                                   `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	AskUserKeywords          []*string                                                                 `json:"AskUserKeywords,omitempty" xml:"AskUserKeywords,omitempty" type:"Repeated"`
	CurrentStep              *string                                                                   `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	MultimodalMediaSelection *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection `json:"MultimodalMediaSelection,omitempty" xml:"MultimodalMediaSelection,omitempty" type:"Struct"`
	NextStep                 *string                                                                   `json:"NextStep,omitempty" xml:"NextStep,omitempty"`
	SkipCurrentSupplement    *bool                                                                     `json:"SkipCurrentSupplement,omitempty" xml:"SkipCurrentSupplement,omitempty"`
	SupplementDataType       *string                                                                   `json:"SupplementDataType,omitempty" xml:"SupplementDataType,omitempty"`
	SupplementEnable         *bool                                                                     `json:"SupplementEnable,omitempty" xml:"SupplementEnable,omitempty"`
	UserBack                 *string                                                                   `json:"UserBack,omitempty" xml:"UserBack,omitempty"`
	UserBackKeywords         []*string                                                                 `json:"UserBackKeywords,omitempty" xml:"UserBackKeywords,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationRequestAgentContextBizContext) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestAgentContextBizContext) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetAskUser() *string {
	return s.AskUser
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetAskUserKeywords() []*string {
	return s.AskUserKeywords
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetMultimodalMediaSelection() *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	return s.MultimodalMediaSelection
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetNextStep() *string {
	return s.NextStep
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetSkipCurrentSupplement() *bool {
	return s.SkipCurrentSupplement
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetSupplementDataType() *string {
	return s.SupplementDataType
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetSupplementEnable() *bool {
	return s.SupplementEnable
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetUserBack() *string {
	return s.UserBack
}

func (s *RunSearchGenerationRequestAgentContextBizContext) GetUserBackKeywords() []*string {
	return s.UserBackKeywords
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetAskUser(v string) *RunSearchGenerationRequestAgentContextBizContext {
	s.AskUser = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetAskUserKeywords(v []*string) *RunSearchGenerationRequestAgentContextBizContext {
	s.AskUserKeywords = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetCurrentStep(v string) *RunSearchGenerationRequestAgentContextBizContext {
	s.CurrentStep = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetMultimodalMediaSelection(v *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) *RunSearchGenerationRequestAgentContextBizContext {
	s.MultimodalMediaSelection = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetNextStep(v string) *RunSearchGenerationRequestAgentContextBizContext {
	s.NextStep = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetSkipCurrentSupplement(v bool) *RunSearchGenerationRequestAgentContextBizContext {
	s.SkipCurrentSupplement = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetSupplementDataType(v string) *RunSearchGenerationRequestAgentContextBizContext {
	s.SupplementDataType = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetSupplementEnable(v bool) *RunSearchGenerationRequestAgentContextBizContext {
	s.SupplementEnable = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetUserBack(v string) *RunSearchGenerationRequestAgentContextBizContext {
	s.UserBack = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) SetUserBackKeywords(v []*string) *RunSearchGenerationRequestAgentContextBizContext {
	s.UserBackKeywords = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContext) Validate() error {
	if s.MultimodalMediaSelection != nil {
		if err := s.MultimodalMediaSelection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection struct {
	// example:
	//
	// 原始会话唯一标识：搜索结果取这个会话中的全量，目前仅媒资搜索场景需要
	OriginalSessionId *string `json:"OriginalSessionId,omitempty" xml:"OriginalSessionId,omitempty"`
	// example:
	//
	// TextGenerate
	SearchModel *string `json:"SearchModel,omitempty" xml:"SearchModel,omitempty"`
	// example:
	//
	// 分类1
	SearchModelDataValue *string `json:"SearchModelDataValue,omitempty" xml:"SearchModelDataValue,omitempty"`
	// example:
	//
	// all
	SelectionType *string `json:"SelectionType,omitempty" xml:"SelectionType,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId        *string                                                                                   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TextSearchResult *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult `json:"TextSearchResult,omitempty" xml:"TextSearchResult,omitempty" type:"Struct"`
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetSearchModel() *string {
	return s.SearchModel
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetSearchModelDataValue() *string {
	return s.SearchModelDataValue
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetSelectionType() *string {
	return s.SelectionType
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetSessionId() *string {
	return s.SessionId
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) GetTextSearchResult() *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult {
	return s.TextSearchResult
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetOriginalSessionId(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.OriginalSessionId = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetSearchModel(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.SearchModel = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetSearchModelDataValue(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.SearchModelDataValue = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetSelectionType(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.SelectionType = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetSessionId(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.SessionId = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) SetTextSearchResult(v *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection {
	s.TextSearchResult = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelection) Validate() error {
	if s.TextSearchResult != nil {
		if err := s.TextSearchResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult struct {
	SearchResult []*RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) GetSearchResult() []*RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	return s.SearchResult
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) SetSearchResult(v []*RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult {
	s.SearchResult = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult struct {
	Chunks []*string `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Repeated"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xxx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-11-25 14:25:59
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 互联网搜索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// SystemSearch
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// example:
	//
	// 新华社
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetChunks() []*string {
	return s.Chunks
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetContent() *string {
	return s.Content
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetScore() *float32 {
	return s.Score
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetSearchSource() *string {
	return s.SearchSource
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetSource() *string {
	return s.Source
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetTitle() *string {
	return s.Title
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) GetUrl() *string {
	return s.Url
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetChunks(v []*string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Chunks = v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetContent(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Content = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetDocId(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.DocId = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetDocUuid(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.DocUuid = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetPubTime(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.PubTime = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetScore(v float32) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Score = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetSearchSource(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.SearchSource = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetSearchSourceName(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetSearchSourceType(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.SearchSourceType = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetSource(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Source = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetSummary(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Summary = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetTitle(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Title = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) SetUrl(v string) *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult {
	s.Url = &v
	return s
}

func (s *RunSearchGenerationRequestAgentContextBizContextMultimodalMediaSelectionTextSearchResultSearchResult) Validate() error {
	return dara.Validate(s)
}

type RunSearchGenerationRequestChatConfig struct {
	EnableThinking         *bool     `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
	ExcludeGenerateOptions []*string `json:"ExcludeGenerateOptions,omitempty" xml:"ExcludeGenerateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// concise
	GenerateLevel *string `json:"GenerateLevel,omitempty" xml:"GenerateLevel,omitempty"`
	// example:
	//
	// copilotPrecise
	GenerateTechnology *string `json:"GenerateTechnology,omitempty" xml:"GenerateTechnology,omitempty"`
	// example:
	//
	// # 角色 你是一个专业的文章检索和问答机器人，擅长文章检索和回答用户问题。  # 任务目标 请你根据检索到的相关文章，回答或表述用户问题“{query}”。  # 任务限制 - 如果用户问题中提到具体日期，请考虑知识日期做筛选。 - 生成内容结构条理。 - 生成内容尽量精简。 - 控制在30字以内 - 不要使用其他数据，不要杜撰。 - 如果不能回答用户问题，请输出对应语言的拒识文案:   - 中文：\\"根据已知信息无法回答。\\"   - 英文：\\"Unable to answer based on the known information.\\"  # 输入数据 ## 检索到的相关文章 {content}
	ModelCustomPromptTemplate *string `json:"ModelCustomPromptTemplate,omitempty" xml:"ModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// # 角色 你是一个专业的文章检索和问答机器人，擅长文章检索和回答用户问题。   # 任务目标 请你根据检索到的相关文章和图片，回答或表述用户问题“{query}”。  # 任务限制  - 如果用户问题中提到具体日期，请考虑知识日期做筛选。  - 生成内容结构条理。  - 生成内容尽量精简。  - 控制在30字以内。 - 如果图片内容可以回答，可以忽略文章内容。 - 不要使用其他数据，不要杜撰。  - 如果不能回答用户问题，请输出对应语言的拒识文案:    	- 中文：\\"根据已知信息无法回答。\\"    	- 英文：\\"Unable to answer based on the known information.\\"    # 输入数据  ## 检索到的相关文章  {content}
	ModelCustomVlPromptTemplate *string                                          `json:"ModelCustomVlPromptTemplate,omitempty" xml:"ModelCustomVlPromptTemplate,omitempty"`
	SearchModels                []*string                                        `json:"SearchModels,omitempty" xml:"SearchModels,omitempty" type:"Repeated"`
	SearchParam                 *RunSearchGenerationRequestChatConfigSearchParam `json:"SearchParam,omitempty" xml:"SearchParam,omitempty" type:"Struct"`
}

func (s RunSearchGenerationRequestChatConfig) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestChatConfig) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestChatConfig) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *RunSearchGenerationRequestChatConfig) GetExcludeGenerateOptions() []*string {
	return s.ExcludeGenerateOptions
}

func (s *RunSearchGenerationRequestChatConfig) GetGenerateLevel() *string {
	return s.GenerateLevel
}

func (s *RunSearchGenerationRequestChatConfig) GetGenerateTechnology() *string {
	return s.GenerateTechnology
}

func (s *RunSearchGenerationRequestChatConfig) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunSearchGenerationRequestChatConfig) GetModelCustomVlPromptTemplate() *string {
	return s.ModelCustomVlPromptTemplate
}

func (s *RunSearchGenerationRequestChatConfig) GetSearchModels() []*string {
	return s.SearchModels
}

func (s *RunSearchGenerationRequestChatConfig) GetSearchParam() *RunSearchGenerationRequestChatConfigSearchParam {
	return s.SearchParam
}

func (s *RunSearchGenerationRequestChatConfig) SetEnableThinking(v bool) *RunSearchGenerationRequestChatConfig {
	s.EnableThinking = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetExcludeGenerateOptions(v []*string) *RunSearchGenerationRequestChatConfig {
	s.ExcludeGenerateOptions = v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetGenerateLevel(v string) *RunSearchGenerationRequestChatConfig {
	s.GenerateLevel = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetGenerateTechnology(v string) *RunSearchGenerationRequestChatConfig {
	s.GenerateTechnology = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetModelCustomPromptTemplate(v string) *RunSearchGenerationRequestChatConfig {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetModelCustomVlPromptTemplate(v string) *RunSearchGenerationRequestChatConfig {
	s.ModelCustomVlPromptTemplate = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetSearchModels(v []*string) *RunSearchGenerationRequestChatConfig {
	s.SearchModels = v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) SetSearchParam(v *RunSearchGenerationRequestChatConfigSearchParam) *RunSearchGenerationRequestChatConfig {
	s.SearchParam = v
	return s
}

func (s *RunSearchGenerationRequestChatConfig) Validate() error {
	if s.SearchParam != nil {
		if err := s.SearchParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchGenerationRequestChatConfigSearchParam struct {
	CategoryUuids   []*string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty" type:"Repeated"`
	CreateTimeEnd   *int64    `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *int64    `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	DocIds          []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	DocUuids        []*string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty" type:"Repeated"`
	// example:
	//
	// 1725983999999
	EndTime               *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Extend1               *string   `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2               *string   `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3               *string   `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	MultimodalSearchTypes []*string `json:"MultimodalSearchTypes,omitempty" xml:"MultimodalSearchTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 0.6
	SearchAudioMinScore *float64 `json:"SearchAudioMinScore,omitempty" xml:"SearchAudioMinScore,omitempty"`
	// example:
	//
	// 0.6
	SearchImageMinScore *float64                                                        `json:"SearchImageMinScore,omitempty" xml:"SearchImageMinScore,omitempty"`
	SearchSources       []*RunSearchGenerationRequestChatConfigSearchParamSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// example:
	//
	// 0.6
	SearchTextMinScore *float64 `json:"SearchTextMinScore,omitempty" xml:"SearchTextMinScore,omitempty"`
	// example:
	//
	// 0.6
	SearchVideoMinScore *float64 `json:"SearchVideoMinScore,omitempty" xml:"SearchVideoMinScore,omitempty"`
	// example:
	//
	// 1725983999999
	StartTime *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RunSearchGenerationRequestChatConfigSearchParam) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestChatConfigSearchParam) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetCategoryUuids() []*string {
	return s.CategoryUuids
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetDocIds() []*string {
	return s.DocIds
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetDocUuids() []*string {
	return s.DocUuids
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetExtend1() *string {
	return s.Extend1
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetExtend2() *string {
	return s.Extend2
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetExtend3() *string {
	return s.Extend3
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetMultimodalSearchTypes() []*string {
	return s.MultimodalSearchTypes
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetSearchAudioMinScore() *float64 {
	return s.SearchAudioMinScore
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetSearchImageMinScore() *float64 {
	return s.SearchImageMinScore
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetSearchSources() []*RunSearchGenerationRequestChatConfigSearchParamSearchSources {
	return s.SearchSources
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetSearchTextMinScore() *float64 {
	return s.SearchTextMinScore
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetSearchVideoMinScore() *float64 {
	return s.SearchVideoMinScore
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) GetTags() []*string {
	return s.Tags
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetCategoryUuids(v []*string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.CategoryUuids = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetCreateTimeEnd(v int64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.CreateTimeEnd = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetCreateTimeStart(v int64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.CreateTimeStart = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetDocIds(v []*string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.DocIds = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetDocUuids(v []*string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.DocUuids = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetEndTime(v int64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.EndTime = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetExtend1(v string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.Extend1 = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetExtend2(v string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.Extend2 = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetExtend3(v string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.Extend3 = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetMultimodalSearchTypes(v []*string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.MultimodalSearchTypes = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetSearchAudioMinScore(v float64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.SearchAudioMinScore = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetSearchImageMinScore(v float64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.SearchImageMinScore = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetSearchSources(v []*RunSearchGenerationRequestChatConfigSearchParamSearchSources) *RunSearchGenerationRequestChatConfigSearchParam {
	s.SearchSources = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetSearchTextMinScore(v float64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.SearchTextMinScore = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetSearchVideoMinScore(v float64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.SearchVideoMinScore = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetStartTime(v int64) *RunSearchGenerationRequestChatConfigSearchParam {
	s.StartTime = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) SetTags(v []*string) *RunSearchGenerationRequestChatConfigSearchParam {
	s.Tags = v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParam) Validate() error {
	if s.SearchSources != nil {
		for _, item := range s.SearchSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSearchGenerationRequestChatConfigSearchParamSearchSources struct {
	// example:
	//
	// SystemSearch
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// QuarkCommonNews
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s RunSearchGenerationRequestChatConfigSearchParamSearchSources) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationRequestChatConfigSearchParamSearchSources) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationRequestChatConfigSearchParamSearchSources) GetCode() *string {
	return s.Code
}

func (s *RunSearchGenerationRequestChatConfigSearchParamSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RunSearchGenerationRequestChatConfigSearchParamSearchSources) SetCode(v string) *RunSearchGenerationRequestChatConfigSearchParamSearchSources {
	s.Code = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParamSearchSources) SetDatasetName(v string) *RunSearchGenerationRequestChatConfigSearchParamSearchSources {
	s.DatasetName = &v
	return s
}

func (s *RunSearchGenerationRequestChatConfigSearchParamSearchSources) Validate() error {
	return dara.Validate(s)
}
